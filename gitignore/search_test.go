package gitignore

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

func TestSearchWords1(t *testing.T) {
	keyword := "go"
	names := []string{
		"Go.gitignore",
		"Rust.gitignore",
		"Go.AllowList.gitignore",
		"LittleGo.something",
		"GZip",
		"NotGoHere",
		"GoHere",
		"ogz",
	}

	expected := []string{
		"Go.gitignore",
		"Go.AllowList.gitignore",
		"LittleGo.something",
		"NotGoHere",
		"GoHere",
	}
	actual := searchWords(keyword, names)

	if len(expected) != len(actual) {
		t.Errorf("got %s, expect %s", actual, expected)
	}

	for i, item := range expected {
		if item != actual[i] {
			t.Errorf("got %s, expect %s", actual, expected)
		}
	}
}
