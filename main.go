package main

import (
	"flag"
	"fmt"
	"io"
	"log"
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
	var stdout bool

	flag.StringVar(&directory, "directory", "", "git workspace (by default current directory)")
	flag.StringVar(&command, "command", "", "command to execute")
	flag.StringVar(&outputFile, "output", "out.csv", "output csv file")
	flag.StringVar(&after, "after", "", "optional begin date of history search")
	flag.StringVar(&before, "before", "", "optional end date of history search")
	flag.IntVar(&limit, "limit", 500, "max number of commits to check")
	flag.BoolVar(&stdout, "stdout", false, "write csv content to stdout")
	flag.Parse()

	// if no parameter given, use history of current history
	var err error
	var out io.Writer

	if stdout {
		out = os.Stdout
	} else {
		out, err := os.Create(outputFile)
		checkError(err)
		defer out.Close()
	}

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

func evaluate(commit entry, command string) (entry, error) {
	if _, err := execute("git", "reset", "--hard"); err != nil {
		return commit, err
	}
	if _, err := execute("git", "checkout", commit.commit); err != nil {
		return commit, err
	}

	stdout, _ := execute("sh", "-c", command)

	data := strings.SplitN(stdout, "\n", 2)
	commit.data = data[0]

	return commit, nil
}

func execute(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Command '%s' failed: %s: %s", cmd.String(), err, string(out))
	}

	return strings.TrimSpace(string(out)), err
}
