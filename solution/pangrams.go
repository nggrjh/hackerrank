package solution

import (
	"strings"
)

// https://www.hackerrank.com/challenges/one-month-preparation-kit-pangrams/problem
func pangrams(s string) string {
	s = strings.ReplaceAll(strings.ToLower(s), " ", "")

	m := make(map[string]bool)
	for i, r := range s {
		if r >= 'a' || r <= 'z' {
			m[s[i:i+1]] = true
		}
	}

	if len(m) < 26 {
		return "not pangram"
	}

	return "pangram"
}
