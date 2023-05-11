package solution

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/nggrjh/hackerrank/common"
)

func Test_arrayManipulation(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		want int64
	}{
		"case_0": {want: 200},
		"case_1": {want: 882},
		"case_4": {want: 7542539201},
		"case_6": {want: 7515267971},
		//"case_7": {want: 2497169732},
		//"case_12": {want: 2517519438},
		"case_14": {want: 10},
		"case_15": {want: 31},
	}
	for name, test := range tests {
		tt := test
		nm := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var n int32
			var queries [][]int32

			lines := common.ReadLines(fmt.Sprintf("../input/array_manipulation/%s.txt", nm))
			for i, line := range lines {
				texts := strings.Split(line, " ")

				var values []int32
				for _, text := range texts {
					v, _ := strconv.ParseInt(text, 10, 32)
					values = append(values, int32(v))
				}

				if i == 0 {
					n = values[0]
					continue
				}

				queries = append(queries, values)
			}

			s := time.Now()
			defer fmt.Printf("%s\t: %s\n", nm, time.Since(s).String())
			if got := arrayManipulation(n, queries); got != tt.want {
				t.Errorf("arrayManipulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
