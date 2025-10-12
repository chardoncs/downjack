package ext

import (
	"regexp"
	"sync"
)

var (
	pattern		*regexp.Regexp
	once		sync.Once
)

func GetRecognizedExtPattern() *regexp.Regexp {
	once.Do(func() {
		pattern = regexp.MustCompile(`\.(txt|md|html?)$`)
	})

	return pattern
}
