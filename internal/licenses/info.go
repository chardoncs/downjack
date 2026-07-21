package licenses

import (
	"os/exec"
	"time"

	"go.chardoncs.dev/downjack/internal/utils"
)

type LicenseInfo struct {
	Year int
	Name string
}

func getLicenseInfo() (*LicenseInfo, error) {
	name, err := utils.GetUserName()
	if err != nil {
		if _, ok := err.(*exec.Error); ok {
			goto SetNamePlaceholder
		}
		return nil, err
	}

SetNamePlaceholder:
	if name == "" {
		name = "<name>"
	}

	return &LicenseInfo{
		Year: time.Now().Year(),
		Name: name,
	}, nil
}
