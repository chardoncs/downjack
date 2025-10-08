package license

import (
	"fmt"
	"strings"

	lib "github.com/chardoncs/downjack/internal/licenses"
	"github.com/spf13/cobra"
)

var (
	overwrite		bool
)

var aliases = []string{ "l" }

var LicenseCmd = &cobra.Command{
	Use: "license [flags] <name> [extra-names...]",
	Aliases: aliases,
	Short: fmt.Sprintf("Add one or more open source licenses (aliases: %s)", strings.Join(aliases, "/")),
	RunE: func(cmd *cobra.Command, args []string) error {

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
