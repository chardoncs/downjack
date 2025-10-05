package gitignore

import (
	"path/filepath"

	utils_io "github.com/chardoncs/downjack/utils/io"
	"github.com/chardoncs/downjack/utils/search"
)

func SearchEmbed(name string) (*search.SearchResult, error) {
	return search.SearchEmbed(name, &root, dirPrefix, "gitignore")
}

func ReadToString(filename string) (string, error) {
	return utils_io.ReadEmbedToString(
		&root,
		filepath.Join(dirPrefix, filename),
	)
}

func SaveTo(dir string, content string, opts utils_io.SaveToOptions) error {
	return utils_io.SaveTo(filepath.Join(dir, ".gitignore"), content, opts)
}
