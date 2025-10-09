package utils

import "fmt"

func ArgsError(expect, actual int) error {
	var pluralExpect string
	if expect != 1 {
		pluralExpect = "s"
	}

	return fmt.Errorf("expect %d argument%s, found %d", expect, pluralExpect, actual)
}
