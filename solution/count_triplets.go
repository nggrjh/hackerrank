package solution

// FIXME: https://www.hackerrank.com/challenges/count-triplets-1/problem
func countTriplets(arr []int64, r int64) int64 {
	ln := len(arr)

	var count int64
	for i := 0; i < ln-2; i++ {
		for j := i + 1; j < ln-1; j++ {
			if r*arr[i] != arr[j] {
				continue
			}

			for k := j + 1; k < ln; k++ {
				if r*arr[j] != arr[k] {
					continue
				}

				count++
			}
		}
	}

	return count
}

//func countTriplets(arr []int64, r int64) int64 {
//	ln := len(arr)
//
//	if r == 1 {
//		return int64(common.Combinations(float64(ln), 3))
//	}
//
//	wg := &sync.WaitGroup{}
//	tmp := make([]bool, (ln*3)-3)
//
//	for i := 0; i < ln-2; i++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//
//			for j := i + 1; j < ln-1; j++ {
//				if r*arr[i] != arr[j] {
//					continue
//				}
//
//				wg.Add(1)
//				go func(j int) {
//					defer wg.Done()
//
//					for k := j + 1; k < ln; k++ {
//						if r*arr[j] != arr[k] {
//							continue
//						}
//
//						tmp[i+j+k] = true
//					}
//				}(j)
//			}
//		}(i)
//	}
//
//	wg.Wait()
//
//	var count int64
//	for _, v := range tmp {
//		if v {
//			count++
//		}
//	}
//	return count
//}
