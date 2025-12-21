package license

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.chardoncs.dev/downjack/internal/cli"
	"go.chardoncs.dev/downjack/internal/cli/ask"
	"go.chardoncs.dev/downjack/internal/cli/fuzzy"
	lib "go.chardoncs.dev/downjack/internal/licenses"
	"go.chardoncs.dev/downjack/internal/utils"
	"github.com/spf13/cobra"
)

var (
	dir   string
	force bool
)

var aliases = []string{"l"}

var LicenseCmd = &cobra.Command{
	Use:     "license",
	Aliases: aliases,
	Short:   fmt.Sprintf("Add an open source license (aliases: %s)", strings.Join(aliases, "/")),
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		var err error

		var selected *lib.MatchedItem

		if len(args) < 1 {
			name, err = findFiles("")
			if err != nil {
				return err
			}

			if name == "" {
				cli.Warnf("Nothing is selected")
				return nil
			}

			selected = &lib.MatchedItem{
				Id:       lib.GetLicenseId(name),
				Filename: name,
			}
		} else if len(args) > 1 {
			return fmt.Errorf("expect 1 argument, found %d", len(args))
		} else {
			name = args[0]

			// Try find matched files
			result, err := lib.SearchEmbed(name)
			if err != nil {
				return err
			}

			if len(result.Items) == 1 {
				selected = &result.Items[0]
				cli.Infof("Found exact license: %s", selected.Filename)
			} else {
				cli.Infof("Not sure which one")
				name, err = findFiles(name)
				if err != nil {
					return err
				}
				if name == "" {
					cli.Infof("No file selected")
					return nil
				}

				selected = &lib.MatchedItem{
					Id:       lib.GetLicenseId(name),
					Filename: name,
				}
			}
		}

		targetFile := "LICENSE"
		extName := utils.GetFormatExtName(selected.Filename)

		isPlainText := extName == "txt"
		if !isPlainText {
			targetFile = fmt.Sprintf("%s.%s", targetFile, extName)
		}

		target := filepath.Join(dir, targetFile)

		exists, err := licenseFileExists(target)
		if err != nil {
			return err
		}

		if exists {
			cli.Warnf("%s already exists", targetFile)

			if force {
				cli.Warnf("%s will be overwritten", targetFile)
			} else {
				var candidateFilename string
				if isPlainText {
					candidateFilename = fmt.Sprintf("LICENSE-%s", selected.Id)
				} else {
					candidateFilename = fmt.Sprintf("LICENSE-%s.%s", selected.Id, extName)
				}

				input := strings.ToLower(strings.TrimSpace(ask.Askf(`What do you want to do?
- [a]dd a new file named %s
- [o]verwrite the existing %s
- [N]o action
: `,
					candidateFilename, targetFile)))

				if input == "" {
					input = " "
				}

				switch input[0] {
				case 'a':
					target = filepath.Join(dir, candidateFilename)
				case 'o':
					cli.Warnf("%s will be overwritten", targetFile)
				default:
					cli.Infof("Aborted")
					return nil
				}
			}
		}

		cli.InfoProgressf("Writing license into `%s`", target)

		if err := lib.WriteLicense(*selected, target); err != nil {
			return err
		}

		cli.Done()

		return nil
	},
}

func init() {
	f := LicenseCmd.PersistentFlags()

	f.BoolVarP(
		&force,
		"force", "f",
		false,
		"overwrite the `LICENSE` file without confirmation",
	)

	f.StringVarP(
		&dir,
		"dir", "d",
		".",
		"specify the directory where the license file resides",
	)
}

func licenseFileExists(target string) (bool, error) {
	stat, err := os.Stat(target)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	if stat.IsDir() {
		return false, nil
	}

	return true, nil
}

func findFiles(initialInput string) (string, error) {
	files, err := utils.ListFilenames(lib.Root, lib.DirPrefix)
	if err != nil {
		return "", err
	}

	selected, err := fuzzy.Find("Find a license template", files, initialInput)
	if err != nil {
		return "", err
	}

	return selected, nil
}
