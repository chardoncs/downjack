package gitignore

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/chardoncs/downjack/internal/cli"
	"github.com/chardoncs/downjack/internal/cli/ask"
	"github.com/chardoncs/downjack/internal/cli/fuzzy"
	lib "github.com/chardoncs/downjack/internal/gitignore"
	"github.com/chardoncs/downjack/internal/gitignore/search"
	"github.com/chardoncs/downjack/internal/utils"
)

var (
	overwrite bool
	dir       string
	title     string
	noTitle   bool
	listing   bool
)

var aliases = []string{"g", "git", "i", "ignore"}

var GitignoreCmd = &cobra.Command{
	Use:     "gitignore [name]",
	Aliases: aliases,
	Short: fmt.Sprintf(
		"Create or append a `.gitignore` file in the project (aliases: %s)",
		strings.Join(aliases, "/"),
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string

		if len(args) < 1 {
			files, err := utils.ListFilenames(lib.Root, lib.DirPrefix)
			if err != nil {
				return err
			}

			names := make([]string, len(files))
			for i, filename := range files {
				names[i], _ = strings.CutSuffix(filename, ".gitignore")
			}

			selected, err := fuzzy.Find("Find a gitignore template", names)
			if err != nil {
				return err
			}

			if selected == "" {
				cli.Warnf("Nothing is selected")
				return nil
			}

			name = selected + ".gitignore"
		} else {
			name = args[0]
		}

		result, err := search.SearchEmbed(name)
		if err != nil {
			return err
		}

		if len(result.Filenames) < 1 {
			return utils.NotFoundError("gitignore", name)
		}

		var content string
		var filename string

		if result.IsExact {
			filename = result.Filenames[0]
			cli.Infof("Found exact template: %s", filename)

			content = result.ExactContent
		} else {
			cli.Infof("Found template(s):")
			cli.PrintItems(result.Filenames)

			fmt.Println()

			input, err := ask.AskInt("Choose template", len(result.Filenames))
			if err != nil {
				return err
			}

			filename = result.Filenames[input-1]

			cli.Infof("Selected %s", filename)

			content, err = utils.ReadEmbedToString(
				&lib.Root,
				filepath.Join(lib.DirPrefix, filename),
			)
			if err != nil {
				return err
			}
		}

		targetFile := filepath.Join(dir, ".gitignore")

		if overwrite {
			cli.Warnf("%s will be overwritten with template `%s`", targetFile, filename)

			confirmed := ask.AskConfirm("Do you want to proceed?")
			if !confirmed {
				cli.Infof("Aborted")
				return nil
			}
		} else {
			cli.Infof(
				"Template `%s` will be appended into %s",
				filename,
				targetFile,
			)
		}

		var resultTitle string
		if !noTitle {
			if title == "" {
				resultTitle = utils.GetFilePrefix(filename)
			} else {
				resultTitle = title
			}
		}

		cli.InfoProgressf("Writing into .gitignore")

		if err := lib.SaveTo(dir, content, lib.SaveToOptions{
			Overwrite: overwrite,
			Title:     resultTitle,
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
		&listing,
		"list", "l",
		false,
		"list or search all available gitignore templates",
	)
}
