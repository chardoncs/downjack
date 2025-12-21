package licenses

import (
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
		return nil, err
	}
	if name == "" {
		name = "<name>"
	}

	return &LicenseInfo{
		Year: time.Now().Year(),
		Name: name,
	}, nil
}
