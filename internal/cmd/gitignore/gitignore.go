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
		var err error

		if len(args) < 1 {
			name, err = findFiles("")
			if err != nil {
				return err
			}

			if name == "" {
				cli.Warnf("Nothing is selected")
				return nil
			}
		} else if len(args) > 1 {
			return fmt.Errorf("Expect 1 argument, found %d", len(args))
		} else {
			name = args[0]
		}

		filename := addSuffix(name)

		// Try search
		result, err := search.MatchFiles(name)
		if err != nil {
			return err
		}

		var content string

		if len(result.Filenames) == 1 {
			filename = result.Filenames[0]
			cli.Infof("Found exact template: %s", filename)
		} else {
			cli.Infof("No exact template found")

			name, err = findFiles(name)
			if err != nil {
				return err
			}

			if name == "" {
				cli.Infof("No file selected")
				return nil
			}

			filename = addSuffix(name)
		}

		b, err := lib.Root.ReadFile(filepath.Join(lib.DirPrefix, filename))
		if err != nil {
			return err
		}

		content = string(b)

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

		cli.InfoProgressf("Writing into .gitignore")

		if err := lib.SaveTo(dir, content, lib.SaveToOptions{
			Overwrite: overwrite,
			Title:     name,
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

func findFiles(initialInput string) (string, error) {
	files, err := utils.ListFilenames(lib.Root, lib.DirPrefix)
	if err != nil {
		return "", err
	}

	names := make([]string, len(files))
	for i, filename := range files {
		names[i], _ = strings.CutSuffix(filename, ".gitignore")
	}

	selected, err := fuzzy.Find("Find a gitignore template", names, initialInput)
	if err != nil {
		return "", err
	}

	return selected, nil
}

func addSuffix(name string) string {
	if strings.HasSuffix(name, ".gitignore") {
		return name
	}
	return name + ".gitignore"
}
