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
	dir			string
	force		bool
)

var aliases = []string{ "l" }

var LicenseCmd = &cobra.Command{
	Use: "license [flags] <name>",
	Aliases: aliases,
	Short: fmt.Sprintf("Add an open source license (aliases: %s)", strings.Join(aliases, "/")),
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
			cli.Info("Found exact license: %s", selected.Id)
		} else {
			cli.Info("Found license(s):")

			names := make([]string, len(result.Items))
			for i, item := range result.Items {
				names[i] = item.Filename
			}
			cli.PrintItems(names)

			num, err := cli.AskInt("Choose license", len(names))
			if err != nil {
				return err
			}

			selected = &result.Items[num - 1]
		}

		// Future-proof check
		if selected == nil {
			cli.Info("Nothing is selected")
			return nil
		}

		target := filepath.Join(dir, "LICENSE")

		exists, err := licenseFileExists(target)
		if err != nil {
			return err
		}

		if exists {
			cli.Warn("LICENSE already exists")

			if force {
				cli.Warn("LICENSE will be overwritten")
			} else {
				candidateFilename := fmt.Sprintf("LICENSE-%s", selected.Id)
				input := strings.ToLower(strings.TrimSpace(cli.Ask(`What do you want to do?
- [a]dd a new file named %s
- [o]verwrite the existing LICENSE
- [N]o action
: `,
				candidateFilename)))

				if input == "" {
					input = " "
				}

				switch input[0] {
				case 'a':
					target = filepath.Join(dir, candidateFilename)
				case 'o':
					cli.Warn("LICENSE will be overwritten")
				default:
					cli.Info("Aborted")
					return nil
				}
			}
		}

		cli.InfoProgress("Writing license into `%s`", target)

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
