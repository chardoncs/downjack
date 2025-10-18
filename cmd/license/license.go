package license

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chardoncs/downjack/internal/cli"
	lib "github.com/chardoncs/downjack/internal/licenses"
	"github.com/chardoncs/downjack/utils"
	"github.com/spf13/cobra"
)

var (
	dir   string
	force bool
)

var aliases = []string{"l"}

var LicenseCmd = &cobra.Command{
	Use:     "license <name>",
	Aliases: aliases,
	Short:   fmt.Sprintf("Add an open source license (aliases: %s)", strings.Join(aliases, "/")),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return utils.ArgsError(1, 0)
		}

		name := args[0]

		result, err := lib.SearchEmbed(name)
		if err != nil {
			return err
		}

		if len(result.Items) < 1 {
			return utils.NotFoundError("license", name)
		}

		var selected *lib.MatchedItem

		if result.Exact {
			selected = &result.Items[0]
			cli.Infof("Found exact license: %s", selected.Id)
		} else {
			cli.Infof("Found license(s):")

			names := make([]string, len(result.Items))
			for i, item := range result.Items {
				names[i] = item.Filename
			}
			cli.PrintItems(names)

			num, err := cli.AskInt("Choose license", len(names))
			if err != nil {
				return err
			}

			selected = &result.Items[num-1]
		}

		// Future-proof check
		if selected == nil {
			cli.Infof("Nothing is selected")
			return nil
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

				input := strings.ToLower(strings.TrimSpace(cli.Askf(`What do you want to do?
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
		"specify the directory where the ignore file resides",
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
