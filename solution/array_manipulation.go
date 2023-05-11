package solution

import (
	"sort"
)

// https://www.hackerrank.com/challenges/crush/problem
func arrayManipulation(n int32, queries [][]int32) int64 {
	sort.SliceStable(queries, func(i, j int) bool { return queries[i][2] > queries[j][2] })

	incrA := make([]int64, n)
	incrB := make([]int64, n+1)
	for i := 0; i < len(queries); i++ {
		a := queries[i][0]
		b := queries[i][1]
		k := int64(queries[i][2])

		incrA[a-1] += k
		incrB[b] += k
	}

	/*
		Explanation:
			Carry over previous `sum` and keep adding it with first array,
			assuming it has the same integer combination as the previous one.
			It keeps going until it gets reduced by the second array.
	*/
	var max, sum int64
	for i := 0; i < int(n); i++ {
		newNumberAdded := incrA[i]
		oldNumberReduces := incrB[i]

		sum = sum + newNumberAdded - oldNumberReduces
		if sum > max {
			max = sum
		}
	}

	return max
}
