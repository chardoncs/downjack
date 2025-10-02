package gitignore

import "testing"

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
