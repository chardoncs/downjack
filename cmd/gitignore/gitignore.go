package gitignore

import (
	"fmt"

	"github.com/spf13/cobra"

	lib "github.com/chardoncs/downjack/gitignore"
)

const repo = lib.Repo

var (
	overwrite			bool

	allowNet			bool
	forceNet			bool

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

		//name := args[0]

		//var info		*lib.ReadInfo
		//var err			error

		//var useNet		bool = forceNet

		//if !forceNet {
		//	// Try getting the embedded
		//	info, err = lib.FetchEmbedded(name)
		//	if err != nil {
		//		if errors.Is(err, os.ErrNotExist) && allowNet {
		//			useNet = true
		//		} else {
		//			return err
		//		}
		//	}
		//}

		//if useNet {
		//	info, err = lib.FetchRepo(name)
		//	if err != nil {
		//		return err
		//	}
		//}

		//if info == nil {
		//	return nil
		//}

		//var title	string = title
		//if noTitle {
		//	title = ""
		//} else if title == "" {
		//	// E.g.,
		//	// 1. Go.gitignore -> Go
		//	// 2. Go.AllowList.gitignore -> Go
		//	title, _, _ = strings.Cut(info.Filename, ".")
		//}

		//if err := lib.SaveTo(dir, info.Content, lib.SaveToOptions{
		//	Overwrite: overwrite,
		//	Title: title,
		//}); err != nil {
		//	return err
		//}

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
