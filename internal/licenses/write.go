package licenses

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

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
		tmpl, err := template.New(item.Id).Parse(string(b))
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
