package ex17

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	phone := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
	}

	for _, digit := range strings.Split(digits, "") {
		chars, isValidDigit := phone[digit]
		if !isValidDigit {
			panic(fmt.Sprintf("invalid digit: %v", digit))
		}

		return chars
	}

	return []string{}
}
