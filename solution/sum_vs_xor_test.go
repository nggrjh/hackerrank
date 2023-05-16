package solution

import (
	"fmt"
	"testing"
	"time"
)

func Test_sumXor(t *testing.T) {
	t.Parallel()
	type args struct {
		n int64
	}
	tests := map[string]struct {
		args args
		want int64
	}{
		"case_0": {
			args: args{
				n: 5,
			},
			want: 2,
		},
		"case_1": {
			args: args{
				n: 0,
			},
			want: 1,
		},
		//"case_8": {
		//	args: args{
		//		n: 3434444444333,
		//	},
		//	want: 262144,
		//},
		//"case_10": {
		//	args: args{
		//		n: 1099511627776,
		//	},
		//	want: 1099511627776,
		//},
		"case_11": {
			args: args{
				n: 10,
			},
			want: 4,
		},
	}
	for name, test := range tests {
		tt := test
		nm := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			s := time.Now()
			defer fmt.Printf("%s\t: %s\n", nm, time.Since(s).String())
			if got := sumXor(tt.args.n); got != tt.want {
				t.Errorf("sumXor() = %v, want %v", got, tt.want)
			}

		})
	}
}
