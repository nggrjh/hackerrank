package solution

import (
	"math"
)

func solveOperators(mealCost float64, tipPercent int32, taxPercent int32) int32 {
	return int32(math.Round(mealCost + mealCost*(float64(tipPercent)/100) + mealCost*(float64(taxPercent)/100)))
}

// FIXME: https://www.hackerrank.com/challenges/30-2d-arrays/problem
func maxHourGlass(arr [][]int32) int32 {
	return 0
}
