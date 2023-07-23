package solution

import "strings"

// FIXME: https://www.hackerrank.com/challenges/append-and-delete/problem
func appendAndDelete(initial, desired string, operations int32) string {
	if initial == desired {
		return "Yes"
	}

	if len(initial)+len(desired) == int(operations) {
		return "Yes"
	}

	ln := len(initial)
	word := initial
	for i := 0; i < ln; i++ {
		if operations < 1 {
			return "No"
		}

		word = initial[:len(initial)-i-1]
		operations--

		if strings.HasPrefix(desired, word) {
			break
		}
	}

	if len(desired)-len(word) <= int(operations) {
		return "Yes"
	}

	return "No"
}
