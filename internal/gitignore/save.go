package gitignore

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chardoncs/downjack/utils/fs"
)

type SaveToOptions struct {
	Overwrite		bool
	Title			string
}

func SaveTo(dir string, content string, opts SaveToOptions) error {
	targetPath := filepath.Join(dir, ".gitignore")
	content = strings.TrimSpace(content)
	if content == "" {
		return nil
	}

	file, err := os.OpenFile(targetPath, buildFileFlag(opts.Overwrite), 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	if !opts.Overwrite {
		empty, err := fs.IsFileEmpty(file)
		if err != nil {
			return err
		}

		if !empty {
			// Add a blank line after the existing content
			if _, err := writer.WriteString("\n\n"); err != nil {
				return err
			}
		}
	}

	// Write title
	if title := opts.Title; title != "" {
		if _, err := fmt.Fprintf(writer, "#-- %s --#\n\n", title); err != nil {
			return err
		}
	}

	// Write new content
	if _, err := writer.WriteString(content); err != nil {
		return err
	}

	if err := writer.Flush(); err != nil {
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
