package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var directory string
	var command string
	var outputFile string
	var limit int
	var after string
	var before string

	flag.StringVar(&directory, "directory", "", "git workspace (by default current directory)")
	flag.StringVar(&command, "command", "", "command to execute")
	flag.StringVar(&outputFile, "output", "out.csv", "output csv file")
	flag.StringVar(&after, "after", "", "optional begin date of history search")
	flag.StringVar(&before, "before", "", "optional end date of history search")
	flag.IntVar(&limit, "limit", 500, "max number of commits to check")
	flag.Parse()

	// if no parameter given, use history of current history
	var err error

	out, err := os.Create(outputFile)
	checkError(err)
	defer out.Close()

	err = DumpHistory(directory, command, out, limit, after, before)
	checkError(err)

	fmt.Printf("\nWrote file %s\n", outputFile)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func evaluate(commit entry, command string) entry {
	execute("git", "checkout", commit.commit)

	stdout, _ := execute("sh", "-c", command)

	data := strings.SplitN(stdout, "\n", 2)
	commit.data = data[0]

	return commit
}

func execute(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Command '%s' failed: %s: %s", cmd.String(), err, string(out))
		os.Exit(1)
	}

	return strings.TrimSpace(string(out)), err
}
