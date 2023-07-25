package solution

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/nggrjh/hackerrank/common"
)

func Test_getMax(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
	}{
		"case_0": {},
		"case_4": {},
	}
	for name := range tests {
		nm := name
		t.Run(nm, func(t *testing.T) {
			t.Parallel()

			inputs := common.ReadLines(fmt.Sprintf("../case/max_element/%s_input.txt", nm))

			var operations []string
			for i, input := range inputs {
				if i == 0 {
					continue
				}
				operations = append(operations, input)
			}

			outputs := common.ReadLines(fmt.Sprintf("../case/max_element/%s_output.txt", nm))

			var want []int32
			for _, output := range outputs {
				o, _ := strconv.ParseInt(output, 10, 32)

				want = append(want, int32(o))
			}

			s := time.Now()
			defer fmt.Printf("%s\t: %s\n", nm, time.Since(s).String())
			if got := getMax(operations); !reflect.DeepEqual(got, want) {
				t.Errorf("getMax() = %v, want %v", got, want)
			}
		})
	}
}
