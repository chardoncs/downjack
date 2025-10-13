package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// A replacement of Cobra's `CheckErr()` with color print.
func CheckErr(msg any) {
	if msg != nil {
		fmt.Fprintln(color.Error, color.RedString("Error:"), msg)
		os.Exit(1)
	}
}
