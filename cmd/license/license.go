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

			ids := make([]string, len(result.Items))
			for i, item := range result.Items {
				ids[i] = item.Id
			}
			cli.PrintItems(ids)

			num, err := cli.AskInt("Choose license", len(ids))
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

		var target string = filepath.Join(dir, "LICENSE")
		if !force {
			stat, err := os.Stat(target)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					goto WriteIntoTarget
				}

				return err
			}

			if stat.IsDir() {
				goto WriteIntoTarget
			}

			cli.Warn("`LICENSE` already exists")
			// TODO
		}

WriteIntoTarget:

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
