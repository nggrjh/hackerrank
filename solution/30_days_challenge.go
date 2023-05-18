package solution

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
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

// https://www.hackerrank.com/challenges/30-sorting/problem
func countSwaps(arr []int32) ([]int32, int32) {
	n := len(arr)

	var numberOfSwaps int32
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				numberOfSwaps++
			}
		}

		if numberOfSwaps < 1 {
			break
		}
	}

	return arr, numberOfSwaps
}

// https://www.hackerrank.com/challenges/30-nested-logic/problem
func calculateFine(returnedDate, dueDate string) int32 {
	date := func(s string) time.Time {
		a := strings.Split(s, " ")
		d, _ := strconv.ParseInt(a[0], 10, 64)
		m, _ := strconv.ParseInt(a[1], 10, 64)
		y, _ := strconv.ParseInt(a[2], 10, 64)
		return time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC)
	}

	returned := date(returnedDate)
	due := date(dueDate)

	if returned.After(due) {
		if returned.Year() != due.Year() {
			return 10000
		}

		if returned.Month() == due.Month() {
			return 15 * int32(returned.Day()-due.Day())
		}

		return 500 * int32(returned.Month()-due.Month())
	}

	return 0
}

// https://www.hackerrank.com/challenges/30-running-time-and-complexity/
func checkPrime(arr []int64) []string {
	st := make([]string, len(arr))
	for i, v := range arr {
		st[i] = "Prime"

		if !big.NewInt(v).ProbablyPrime(0) {
			st[i] = "Not prime"
		}
	}
	return st
}

// https://www.hackerrank.com/challenges/30-bitwise-and/problem
func bitwiseAnd(N, K int32) int32 {
	var max int32
	for i := 1; i <= int(N); i++ {
		for j := i + 1; j <= int(N); j++ {
			if x := int32(i & j); x < K && x > max {
				max = x
			}
		}
	}
	return max
}
