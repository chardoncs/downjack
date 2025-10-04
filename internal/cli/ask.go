package cli

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func AskConfirm(prompt string) (confirmed bool) {
	fmt.Printf("%s [Y/n]: ", prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result := strings.TrimSpace(scanner.Text())
		confirmed = result == "" || strings.HasPrefix(strings.ToLower(result), "y")
	}

	return
}

func AskInt(prompt string, numRange ...int) (int, error) {
	var low, high int = 1, 0
	numLen := len(numRange)

	invalid := math.MinInt

	if numLen == 1 {
		high = numRange[0]
	} else if numLen == 2 {
		low, high = numRange[0], numRange[1]
	} else if numLen > 2 {
		return invalid, fmt.Errorf("Too many arguments, expect 1~3, got %d", numLen + 1)
	}

	hasRange := low <= high
	if hasRange {
		invalid = min(0, low - 1)
	}

	fmt.Printf("%s", prompt)
	if hasRange {
		if low == high {
			fmt.Printf(" [%d]", low)
		} else {
			fmt.Printf(" [%d-%d]", low, high)
		}
	}
	fmt.Print(": ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return invalid, err
		}

		if hasRange && (num < low || num > high) {
			return num, fmt.Errorf("The input number is out of range")
		}

		return num, nil
	}

	// If scanner is unavailable
	return invalid, fmt.Errorf("No content read")
}
