package solution

import (
	"sort"
)

// FIXME: https://www.hackerrank.com/challenges/one-month-preparation-kit-two-arrays/problem
func twoArrays(k int32, A, B []int32) string {
	if len(A) > 1 {
		return evaluate(k, A, B)
	}

	if A[0]+B[0] < k {
		return "NO"
	} else {
		return "YES"
	}
}

func evaluate(k int32, A, B []int32) string {
	sort.SliceStable(A, func(i, j int) bool { return A[j] < A[i] }) // desc
	sort.SliceStable(B, func(i, j int) bool { return B[i] < B[j] }) // asc

	var setB, setD []int32

	countB := make(map[int32]int32)
	countD := make(map[int32]int32)

	for i := 0; i < len(A); i++ {
		diff := k - A[i]

		countD[diff]++
		if l := len(setD); l < 1 || setD[l-1] < diff {
			setD = append(setD, diff)
		}

		countB[B[i]]++
		if l := len(setB); l < 1 || setB[l-1] < diff {
			setB = append(setB, B[i])
		}
	}

	current := 0
	for _, d := range setD {
		for i := current; i < len(setB); i++ {
			b := setB[i]
			need := countD[d]
			have := countB[b]

			if d > b {
				return "NO"
			}

			if need > have {
				continue
			}

			countD[d] -= need
			countB[b] -= need

			if countB[b] < 1 {
				current += 1
			}

			break
		}
	}

	return "YES"
}
