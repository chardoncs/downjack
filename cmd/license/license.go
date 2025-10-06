package license

import "github.com/spf13/cobra"

var (
	overwrite		bool
)

var LicenseCmd = &cobra.Command{
	Use: "license [flags] <name> [extra-names...]",
	Aliases: []string{ "l" },
	Short: "Add an open source license (alias: l)",
	RunE: func(cmd *cobra.Command, args []string) error {
		names := args

		return nil
	},
}

func init() {
	f := LicenseCmd.PersistentFlags()

	f.BoolVarP(
		&overwrite,
		"overwrite", "o",
		false,
		"overwrite the `LICENSE` file instead of creating a new license file",
	)
}
