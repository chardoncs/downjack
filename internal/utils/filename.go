package utils

import (
	"strings"

	"github.com/chardoncs/downjack/internal/licenses/regex/ext"
)

func GetFormatExtName(filename string) string {
	cutFilename, _ := strings.CutSuffix(filename, ".tmpl")
	result, _ := strings.CutPrefix(
		ext.GetRecognizedExtPattern().FindString(cutFilename),
		".",
	)
	return result
}
