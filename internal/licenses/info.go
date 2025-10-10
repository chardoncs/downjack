package licenses

import (
	"time"

	"github.com/chardoncs/downjack/utils"
)

type LicenseInfo struct {
	Year		int
	Name		string
}

func newLicenseInfo() (*LicenseInfo, error) {
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

