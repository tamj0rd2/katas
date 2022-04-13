package ex17

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		digits   string
		expected []string
	}{
		//{
		//	digits:   "23",
		//	expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		//},
		{
			digits:   "",
			expected: []string{},
		},
		{
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			digits:   "3",
			expected: []string{"d", "e", "f"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.digits, func(t *testing.T) {
			actual := letterCombinations(tc.digits)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
