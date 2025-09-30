package gitignore

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const GitIgnoreFileName = ".gitignore"

type SaveToOptions struct {
	Overwrite		bool
	Title		string
}

func SaveTo(dir string, content string, opts SaveToOptions) error {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil
	}

	path := filepath.Join(dir, GitIgnoreFileName)

	fp, err := os.OpenFile(path, buildFileFlag(opts.Overwrite), 0666)
	if err != nil {
		return err
	}

	defer fp.Close()

	if !opts.Overwrite {
		empty, err := isFileEmpty(fp)
		if err != nil {
			return err
		}

		if !empty {
			// Add a blank line after the existing content
			if _, err := fp.WriteString("\n\n"); err != nil {
				return err
			}
		}
	}

	// Write title
	if title := opts.Title; title != "" {
		if _, err := fmt.Fprintf(fp, "#-- %s --#\n\n", title); err != nil {
			return err
		}
	}

	// Write new content
	if _, err := fp.WriteString(content); err != nil {
		return err
	}

	return nil
}

func buildFileFlag(overwrite bool) int {
	flag := os.O_RDWR | os.O_CREATE

	if overwrite {
		flag |= os.O_TRUNC
	} else {
		flag |= os.O_APPEND
	}

	return flag
}

func isFileEmpty(fp *os.File) (bool, error) {
	size, err := fp.Seek(0, io.SeekCurrent)
	if err != nil {
		return true, err
	}

	return size == 0, nil
}

