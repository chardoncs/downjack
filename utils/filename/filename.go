package filename

import "strings"

func GetFilePrefix(filename string) string {
	before, _, _ := strings.Cut(filename, ".")
	return before
}
