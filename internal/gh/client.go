package gh

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func IssueList() ([]string, error) {
	out, err := run("issue", "list", "--limit", "10")
	if err != nil {
		return nil, err
	}
	out = strings.TrimSpace(out)
	if out == "" {
		return []string{"No open issues found."}, nil
	}
	return strings.Split(out, "\n"), nil
}

func CreatePR(title string, body string) (string, error) {
	if strings.TrimSpace(title) == "" {
		return "", fmt.Errorf("PR title cannot be empty")
	}
	args := []string{"pr", "create", "--title", strings.TrimSpace(title), "--body", strings.TrimSpace(body)}
	return run(args...)
}

func run(args ...string) (string, error) {
	cmd := exec.Command("gh", args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	out := strings.TrimSpace(stdout.String())
	errOut := strings.TrimSpace(stderr.String())
	if err != nil {
		if errOut != "" {
			return strings.TrimSpace(out + "\n" + errOut), fmt.Errorf("%s", errOut)
		}
		return out, err
	}
	if errOut != "" {
		return strings.TrimSpace(out + "\n" + errOut), nil
	}
	return out, nil
}
