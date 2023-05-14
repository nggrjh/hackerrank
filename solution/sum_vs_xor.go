package solution

import (
	"sync"
)

// FIXME: https://www.hackerrank.com/challenges/one-month-preparation-kit-sum-vs-xor
func sumXor(n int64) int64 {
	result := make([]bool, n+1)

	wg := sync.WaitGroup{}
	for i := int64(0); i <= n; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			if n+i == n^i {
				result[i] = true
			}
		}(i)

	}

	wg.Wait()

	var count int64
	for _, r := range result {
		if r {
			count += 1
		}
	}
	return count
}
