package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/schollz/progressbar/v2"
)

type entry struct {
	commit    string
	timestamp int
	data      string
}

var directory string
var command string
var branch string
var outputFile string
var limit int
var after string
var before string

func main() {
	flag.StringVar(&directory, "directory", "", "git workspace (by default current directory)")
	flag.StringVar(&command, "command", "", "command to execute")
	flag.StringVar(&branch, "branch", "master", "git branch to check")
	flag.StringVar(&outputFile, "output", "out.csv", "output csv file")
	flag.StringVar(&after, "after", "", "optional begin date of history search")
	flag.StringVar(&before, "before", "", "optional end date of history search")
	flag.IntVar(&limit, "limit", 500, "max number of commits to check")
	flag.Parse()

	// if no parameter given, use history of current history
	var err error
	if directory == "" {
		directory, err = os.Getwd()
		checkError(err)
	}

	if stat, err := os.Stat(directory); err != nil || !stat.IsDir() {
		fmt.Printf("Invalid directory: %s\n", directory)
		flag.PrintDefaults()
		os.Exit(1)
	}

	commits, err := getCommits(branch, limit, after, before)
	checkError(err)

	bar := progressbar.New(len(commits))
	result := make([]entry, 0, len(commits))

	for _, commit := range commits {
		bar.Add(1)
		result = append(result, evaluate(commit))
	}

	// checkout to initial version
	checkout(branch)

	out, err := os.Create(outputFile)
	checkError(err)
	defer out.Close()

	err = writeCsv(result, out)
	checkError(err)
	fmt.Printf("\nWrote file %s\n", outputFile)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func evaluate(commit entry) entry {
	checkout(commit.commit)

	stdout, _ := execute("sh", "-c", command)

	commit.data = strings.TrimSpace(stdout)

	return commit
}

func execute(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = directory
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Command '%s' failed: %s: %s", cmd.String(), err, string(out))
		os.Exit(1)
	}

	return string(out), err
}
