package licenses

import (
	html_tmpl "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	text_tmpl "text/template"

	"github.com/chardoncs/downjack/utils"
)

type templateExecutor interface {
	Execute(io.Writer, any) error
}

func WriteLicense(item MatchedItem, target string) error {
	b, err := Root.ReadFile(filepath.Join(DirPrefix, item.Filename))
	if err != nil {
		return err
	}

	targetFile, err := os.OpenFile(
		target,
		os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
		0664,
	)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	if isGoTemplate(item.Filename) {
		var tmpl templateExecutor
		var err error

		if utils.GetFormatExtName(item.Filename) == "html" {
			// HTML template
			tmpl, err = html_tmpl.New(item.Id).Parse(string(b))
		} else {
			// Text template
			tmpl, err = text_tmpl.New(item.Id).Parse(string(b))
		}
		if err != nil {
			return err
		}

		data, err := getLicenseInfo()
		if err != nil {
			return err
		}

		if err := tmpl.Execute(targetFile, data); err != nil {
			return err
		}
	} else {
		if _, err := targetFile.Write(b); err != nil {
			return err
		}
	}

	return nil
}

func isGoTemplate(filename string) bool {
	return strings.HasSuffix(filename, ".tmpl")
}
