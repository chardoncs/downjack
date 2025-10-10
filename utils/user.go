package utils

import "os/exec"

func GetUserName() (string, error) {
	// Get from Git
	cmd := exec.Command("git", "config", "user.name")
	if err := cmd.Run(); err != nil {
		return "", err
	}

	b, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(b), nil
}
