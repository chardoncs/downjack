package gitignore

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	lib "github.com/chardoncs/downjack/gitignore"
)

const repo = lib.Repo

var (
	overwrite			bool

	allowNet			bool
	forceNet			bool

	dir					string

	GitignoreCmd = &cobra.Command{
		Use: "gitignore [flags] <name>",
		Aliases: []string{ "g", "git", "i", "ignore" },
		Short: "Create or append a `.gitignore` file in the project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("the name of the ignore type is required")
			}

			name := args[0]
			var content		string
			var err			error

			var useNet		bool = forceNet

			if !forceNet {
				// Try getting the embedded
				content, err = lib.FetchEmbedded(name)
				if err != nil {
					if errors.Is(err, os.ErrNotExist) && allowNet {
						useNet = true
					} else {
						return err
					}
				}
			}

			if useNet {
				content, err = lib.FetchRepo(name)
				if err != nil {
					return err
				}
			}

			if err := lib.SaveTo(dir, content, overwrite); err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	f := GitignoreCmd.PersistentFlags()

	f.BoolVarP(
		&overwrite,
		"overwrite", "o",
		false,
		"overwrite the existing ignore file instead of appending in it",
	)

	f.BoolVarP(
		&allowNet,
		"net", "n",
		false,
		fmt.Sprintf("allow finding files from the online repo (namely `%s`) when there are no matched embedded files", repo),
	)

	f.BoolVarP(
		&forceNet,
		"force-net", "N",
		false,
		fmt.Sprintf("always find files from the online repo (namely `%s`) and ignore the embedded ones", repo),
	)

	f.StringVarP(
		&dir,
		"dir", "d",
		".",
		"specify the directory where the ignore file resides",
	)
}
