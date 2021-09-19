package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/schollz/progressbar/v2"
)

type entry struct {
	commit    string
	timestamp int
	data      string
}

func DumpHistory(directory string, command string, output io.Writer, limit int, after string, before string) error {
	var err error

	if command == "" {
		flag.PrintDefaults()
		return fmt.Errorf("No command given\n")
	}

	if directory != "" {
		err = os.Chdir(directory)
		if err != nil {
			return fmt.Errorf("Invalid directory: %s\n", directory)
		}
	}

	// make sure we leave a clean state afterwards
	branch, err := execute("git", "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return err
	}

	// switch back to original branch
	defer execute("git", "checkout", branch)

	commits, err := getCommits(limit, after, before)
	if err != nil {
		return err
	}

	bar := progressbar.NewOptions(len(commits), progressbar.OptionSetWriter(os.Stderr))

	result := make([]entry, 0, len(commits))

	for _, commit := range commits {
		bar.Add(1)

		entity, err := evaluate(commit, command)
		if err != nil {
			log.Print(err)
			continue
		}
		result = append(result, entity)
	}

	return writeCsv(result, output)
}

func writeCsv(logs []entry, file io.Writer) error {
	writer := csv.NewWriter(file)
	err := writer.Write([]string{
		"time",
		"result",
		"commit",
	})
	if err != nil {
		return err
	}

	for _, row := range logs {
		err := writer.Write([]string{
			time.Unix(int64(row.timestamp), 0).Format(time.RFC3339),
			row.data,
			row.commit,
		})
		if err != nil {
			return err
		}
	}

	writer.Flush()

	return nil
}
