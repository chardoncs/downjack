package ask

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Askf(prompt string, a ...any) (answer string) {
	fmt.Printf(prompt, a...)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer = strings.TrimSpace(scanner.Text())
	}

	return answer
}

func AskConfirm(prompt string) (confirmed bool) {
	fmt.Printf("%s [Y/n]: ", prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result := strings.TrimSpace(scanner.Text())
		confirmed = result == "" || strings.HasPrefix(strings.ToLower(result), "y")
	}

	return confirmed
}

func AskInt(prompt string, numRange ...int) (int, error) {
	low, high := 1, 0
	numLen := len(numRange)

	invalid := math.MinInt

	if numLen == 1 {
		high = numRange[0]
	} else if numLen == 2 {
		low, high = numRange[0], numRange[1]
	} else if numLen > 2 {
		return invalid, fmt.Errorf("too many arguments, expect 1~3, got %d", numLen+1)
	}

	hasRange := low <= high
	if hasRange {
		invalid = min(0, low-1)
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
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			return invalid, fmt.Errorf("no input")
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			return invalid, err
		}

		if hasRange && (num < low || num > high) {
			return num, fmt.Errorf("input number out of range")
		}

		return num, nil
	}

	// If scanner is unavailable
	return invalid, fmt.Errorf("no content read")
}
