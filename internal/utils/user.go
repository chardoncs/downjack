package utils

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

func GetUserName() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get from Git
	cmd := exec.CommandContext(ctx, "git", "config", "user.name")

	b, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}
