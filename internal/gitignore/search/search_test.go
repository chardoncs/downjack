package search

import "testing"

func TestFilenameConstructionPureName(t *testing.T) {
	keyword := "go"

	expected := "Go.gitignore"
	actual, err := constructPlausibleFilename(keyword)
	if err != nil {
		t.Error(err)
	}

	if expected != actual {
		t.Errorf("got %s, expect %s", actual, expected)
	}
}

func TestFilenameConstructionFilename(t *testing.T) {
	keyword := "go.allowlist.gitignore"

	expected := "Go.Allowlist.gitignore"
	actual, err := constructPlausibleFilename(keyword)
	if err != nil {
		t.Error(err)
	}

	if expected != actual {
		t.Errorf("got %s, expect %s", actual, expected)
	}
}
