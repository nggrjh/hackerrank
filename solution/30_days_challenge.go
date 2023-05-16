package solution

import (
	"math"
)

func solveOperators(mealCost float64, tipPercent int32, taxPercent int32) int32 {
	return int32(math.Round(mealCost + mealCost*(float64(tipPercent)/100) + mealCost*(float64(taxPercent)/100)))
}

// https://www.hackerrank.com/challenges/30-2d-arrays/problem
func maxHourGlass(arr [][]int32) int32 {
	last := len(arr) - 2

	max := int32(math.MinInt32)
	for i := 0; i < last; i++ {
		for j := 0; j < last; j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
