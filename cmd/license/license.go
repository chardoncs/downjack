package license

import (
	"fmt"
	"strings"

	lib "github.com/chardoncs/downjack/internal/licenses"
	"github.com/chardoncs/downjack/utils"
	"github.com/spf13/cobra"
)

var (
	overwrite		bool
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

		return nil
	},
}

func init() {
	f := LicenseCmd.PersistentFlags()

	f.BoolVarP(
		&overwrite,
		"overwrite", "o",
		false,
		"overwrite the `LICENSE` file instead of ignoring or creating a new license file",
	)
}
