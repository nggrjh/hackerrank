package solution

import (
	"math"
	"sync"
)

func beautifulDays(start int32, end int32, divisor int32) int32 {
	ln := end - start + 1
	beauties := make([]bool, ln)

	wg := sync.WaitGroup{}
	for i := int32(0); i < ln; i++ {
		wg.Add(1)
		go func(i, n int32) {
			defer wg.Done()
			if diff := math.Abs(float64(n) - float64(reverseInt(n))); int32(diff)%divisor == 0 {
				beauties[i] = true
			}
		}(i, i+start)
	}

	wg.Wait()

	var count int32
	for _, b := range beauties {
		if b {
			count++
		}
	}
	return count
}

func reverseInt(i int32) int32 {
	var numbers []int32
	for i > 0 {
		numbers = append([]int32{i % 10}, numbers...)
		i = i / 10
	}

	var r int32
	for idx, number := range numbers {
		r += (number * int32(math.Pow(10, float64(idx))))
	}

	return r
}
