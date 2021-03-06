package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCommits(limit int, after string, before string) ([]entry, error) {
	args := []string{
		"log",
		"--reverse",
		"--pretty=\"%h %ct\"",
	}
	if after != "" {
		args = append(args, fmt.Sprintf("--after='%s'", after))
	}
	if before != "" {
		args = append(args, fmt.Sprintf("--before='%s'", before))
	}

	stdout, err := execute("git", args...)

	commits := make([]entry, 0)
	if err != nil {
		return commits, err
	}

	lines := strings.Split(stdout, "\n")

	stepSize := (len(lines) / limit) - 1
	if stepSize == 0 {
		stepSize = 1
	}

	for i, line := range lines {
		if line == "" {
			continue
		}
		if i%stepSize != 0 && i != len(lines)-1 { // include every X once + the last one for sure
			continue
		}
		parts := strings.Split(strings.ReplaceAll(line, `"`, ""), " ")
		timestamp, _ := strconv.Atoi(parts[1])
		commits = append(commits, entry{
			commit:    parts[0],
			timestamp: timestamp,
		})
	}

	fmt.Fprintf(os.Stderr, "Commits: %d, Step size: %d (%d commits to check)\n", len(lines), stepSize, len(commits))

	return commits, nil
}
