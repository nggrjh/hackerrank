package solution

import "testing"

func Test_solveOperators(t *testing.T) {
	t.Parallel()
	type args struct {
		mealCost   float64
		tipPercent int32
		taxPercent int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0": {
			args: args{
				mealCost:   12,
				tipPercent: 20,
				taxPercent: 8,
			},
			want: 15,
		},
		"case_1": {
			args: args{
				mealCost:   15.5,
				tipPercent: 15,
				taxPercent: 10,
			},
			want: 19,
		},
		"case_2": {
			args: args{
				mealCost:   10.25,
				tipPercent: 17,
				taxPercent: 5,
			},
			want: 13,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			if got := solveOperators(tt.args.mealCost, tt.args.tipPercent, tt.args.taxPercent); got != tt.want {
				t.Errorf("solveOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxHourGlass(t *testing.T) {
	t.Parallel()
	type args struct {
		arr [][]int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case": {
			args: args{
				arr: [][]int32{
					{1, 1, 1, 0, 0, 0},
					{0, 1, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0},
				},
			},
			want: 7,
		},
		"case1": {
			args: args{
				arr: [][]int32{
					{1, 1, 1, 0, 0, 0},
					{0, 1, 0, 0, 0, 0},
					{1, 1, 1, 0, 0, 0},
					{0, 0, 2, 4, 4, 0},
					{0, 0, 0, 2, 0, 0},
					{0, 0, 1, 2, 4, 0},
				},
			},
			want: 19,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := maxHourGlass(tt.args.arr); got != tt.want {
				t.Errorf("maxHourGlass() = %v, want %v", got, tt.want)
			}
		})
	}
}
