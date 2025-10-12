package utils

import (
	"os/exec"
	"strings"
)

func GetUserName() (string, error) {
	// Get from Git
	cmd := exec.Command("git", "config", "user.name")

	b, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
