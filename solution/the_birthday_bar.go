package solution

// https://www.hackerrank.com/challenges/one-month-preparation-kit-the-birthday-bar/problem
func birthday(s []int32, d, m int32) int32 {
	ln := len(s)
	mx := ln - int(m)

	var tmp [][]int32
	if mx < 1 {
		tmp = append(tmp, s)
	} else {
		for i := 0; i <= mx; i++ {
			var t []int32
			for j := i; j < i+int(m); j++ {
				t = append(t, s[j])
			}
			tmp = append(tmp, t)
		}
	}

	count := int32(0)
	for i := 0; i < len(tmp); i++ {
		var sum int32
		for j := 0; j < len(tmp[i]); j++ {
			sum += tmp[i][j]
		}
		if sum == d {
			count += 1
		}
	}

	return count
}
