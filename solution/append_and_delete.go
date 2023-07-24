package solution

import "strings"

// https://www.hackerrank.com/challenges/append-and-delete/problem
func appendAndDelete(initial, desired string, operations int32) string {
	if len(initial)+len(desired) <= int(operations) {
		return "Yes"
	}

	ln := len(initial)
	for i := 0; i < ln; i++ {
		if desired == initial && operations%2 == 0 {
			return "Yes"
		}

		initial = initial[:len(initial)-1]
		operations--

		if strings.HasPrefix(desired, initial) {
			desired = desired[len(initial):]
			break
		}
	}

	if len(desired) == int(operations) {
		return "Yes"
	} else if len(desired) < int(operations) && operations%2 == 0 {
		return "Yes"
	}

	return "No"
}
