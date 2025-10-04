package gitignore

import (
	"fmt"

	"github.com/spf13/cobra"

	lib "github.com/chardoncs/downjack/gitignore"
	"github.com/chardoncs/downjack/internal/cli"
)

var (
	overwrite			bool

	dir					string

	title				string

	noTitle				bool

	yes					bool
)

var GitignoreCmd = &cobra.Command{
	Use: "gitignore [flags [values]] <name>",
	Aliases: []string{ "g", "git", "i", "ignore" },
	Short: "Create or append a `.gitignore` file in the project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("the name of the ignore type is required")
		}

		name := args[0]

		result, err := lib.SearchEmbed(name)
		if err != nil {
			return err
		}

		if len(result.Filenames) < 1 {
			return fmt.Errorf("No gitignore template named like \"%s\" found.", name)
		}

		var content string
		var filename string

		if result.IsExact {
			filename = result.Filenames[0]
			cli.Info("Found exact template: %s", filename)

			content = result.ExactContent
		} else {
			cli.Info("Found template(s):")
			cli.PrintItems(result.Filenames)

			if yes {
				filename = result.Filenames[0]
			} else {
				fmt.Println()

				input, err := cli.AskInt("Choose template", len(result.Filenames))
				if err != nil {
					return err
				}

				filename = result.Filenames[input - 1]
			}

			cli.Info("Selected %s", filename)

			var err error
			content, err = lib.ReadToString(filename)
			if err != nil {
				return err
			}
		}

		var resultTitle string
		if !noTitle {
			if title == "" {
				resultTitle = lib.GetFilePrefix(filename)
			} else {
				resultTitle = title
			}
		}

		cli.InfoProgress("Writing into .gitignore")

		if err := lib.SaveTo(dir, content, lib.SaveToOptions{
			Overwrite: overwrite,
			Title: resultTitle,
		}); err != nil {
			fmt.Println()
			return err
		}

		cli.Done()

		return nil
	},
}

func init() {
	f := GitignoreCmd.PersistentFlags()

	f.BoolVarP(
		&overwrite,
		"overwrite", "o",
		false,
		"overwrite the existing ignore file instead of appending the snippet in it",
	)

	f.StringVarP(
		&dir,
		"dir", "d",
		".",
		"specify the directory where the ignore file resides",
	)

	f.StringVarP(
		&title,
		"title", "t",
		"",
		"specify a custom title for the snippet",
	)

	f.BoolVar(
		&noTitle,
		"no-title",
		false,
		"do not add title for the snippet",
	)

	f.BoolVarP(
		&yes,
		"yes", "y",
		false,
		"skip confirmations",
	)
}
