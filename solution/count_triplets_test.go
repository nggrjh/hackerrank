package solution

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/nggrjh/hackerrank/common"
)

func Test_countTriplets(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		want int64
	}{
		"case_0": {want: 2},
		"case_1": {want: 6},
		"case_2": {want: 161700},
		//"case_3":  {want: 166661666700000},
		//"case_10": {want: 1339347780085},
		"case_12": {want: 4},
	}
	for name, test := range tests {
		tt := test
		nm := name
		t.Run(nm, func(t *testing.T) {
			t.Parallel()

			lines := common.ReadLines(fmt.Sprintf("../case/count_triplets/%s_input.txt", nm))
			queries := strings.Split(lines[0], " ")

			r, _ := strconv.ParseInt(queries[1], 10, 64)

			ints := strings.Split(lines[1], " ")

			var arr []int64
			for _, s := range ints {
				a, _ := strconv.ParseInt(s, 10, 64)
				arr = append(arr, a)
			}

			s := time.Now()
			defer fmt.Printf("%s\t: %s\n", nm, time.Since(s).String())
			if got := countTriplets(arr, r); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
