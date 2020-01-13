package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getCommits(branch string, limit int, after string, before string) ([]entry, error) {
	checkout(branch)

	args := []string{
		"log",
		"--pretty=\"%h %ct\"",
	}
	if after != "" {
		args = append(args, fmt.Sprintf("--from='%s'", after))
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
		if i%stepSize != 0 {
			continue
		}
		parts := strings.Split(strings.ReplaceAll(line, "\"", ""), " ")
		timestamp, _ := strconv.Atoi(parts[1])
		commits = append(commits, entry{
			commit:    parts[0],
			timestamp: timestamp,
		})
	}

	fmt.Printf("Commits: %d, Step size: %d (%d commits to check)\n", len(lines), stepSize, len(commits))

	return commits, nil
}

func checkout(branch string) error {
	_, err := execute("git", "checkout", branch)

	return err
}
