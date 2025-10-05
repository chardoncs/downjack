package gitignore

import (
	"bufio"
	"fmt"
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

	file, err := os.OpenFile(path, buildFileFlag(opts.Overwrite), 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	if !opts.Overwrite {
		empty, err := isFileEmpty(file)
		if err != nil {
			return err
		}

		if !empty {
			// Add a blank line after the existing content
			if _, err := writer.WriteString("\n"); err != nil {
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

func isFileEmpty(fp *os.File) (bool, error) {
	stat, err := fp.Stat()
	if err != nil {
		return true, err
	}

	return stat.Size() == 0, nil
}
