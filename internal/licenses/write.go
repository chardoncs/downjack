package licenses

import (
	html_tmpl "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	text_tmpl "text/template"

	"go.chardoncs.dev/downjack/internal/utils"
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
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0o664,
	)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	if isGoTemplate(item.Filename) {
		tmpl, err := parseTemplate(item, string(b))
		if err != nil {
			return err
		}

		if err := executeTemplate(tmpl, targetFile); err != nil {
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

func parseTemplate(item MatchedItem, content string) (templateExecutor, error) {
	if utils.GetFormatExtName(item.Filename) == "html" {
		// HTML template
		return html_tmpl.New(item.Id).Parse(content)
	}

	// Text template
	return text_tmpl.New(item.Id).Parse(content)
}

func executeTemplate(tmpl templateExecutor, writer io.Writer) error {
	data, err := getLicenseInfo()
	if err != nil {
		return err
	}

	if err := tmpl.Execute(writer, data); err != nil {
		return err
	}

	return nil
}
