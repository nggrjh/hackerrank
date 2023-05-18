package common

import "testing"

func TestFactorial(t *testing.T) {
	t.Parallel()
	type args struct {
		n float64
	}
	tests := map[string]struct {
		args args
		want float64
	}{
		"case_1": {
			args: args{
				n: 4,
			},
			want: 24,
		},
		"case_2": {
			args: args{
				n: 10,
			},
			want: 3628800,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinations(t *testing.T) {
	t.Parallel()
	type args struct {
		n float64
		r float64
	}
	tests := map[string]struct {
		args args
		want float64
	}{
		"case_1": {
			args: args{
				n: 100,
				r: 3,
			},
			want: 161700,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Combinations(tt.args.n, tt.args.r); int64(got) != int64(tt.want) {
				t.Errorf("Combinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
