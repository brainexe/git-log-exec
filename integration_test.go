package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const GitDir = "git_directory"

func TestRemoteRepo(t *testing.T) {
	t.Run("slack-bot-1", func(t *testing.T) {
		before := ""
		after := "2019-09-04T17:14:43+02:00"
		limit := 20

		test(t, "../tests/test1.csv", "ls * | wc -l", "https://github.com/innogames/slack-bot.git", limit, before, after)
	})
}

func test(t *testing.T, expectedFile string, command string, gitUrl string, limit int, after string, before string) {
	os.RemoveAll(GitDir)
	cmd := exec.Command("git", "clone", gitUrl, GitDir)
	out, err := cmd.CombinedOutput()
	assert.NoError(t, err, string(out))
	defer os.RemoveAll(GitDir)

	var actual bytes.Buffer
	err = DumpHistory(GitDir, command, bufio.NewWriter(&actual), limit, after, before)
	assert.NoError(t, err)

	if expectedFile != "" {
		expected, err := ioutil.ReadFile(expectedFile)
		assert.NoError(t, err, expected)

		assert.Equal(t, string(expected), actual.String())
	}
}
