package solution

import (
	"reflect"
	"testing"
)

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
		"case_0": {
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
		"case_1": {
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
		"case_3": {
			args: args{
				arr: [][]int32{
					{0, -4, -6, 0, -7, -6},
					{-1, -2, -6, -8, -3, -1},
					{-8, -4, -2, -8, -8, -6},
					{-3, -1, -2, -5, -7, -4},
					{-3, -5, -3, -6, -6, -6},
					{-3, -6, 0, -8, -6, -7},
				},
			},
			want: -19,
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

func Test_countSwaps(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := map[string]struct {
		args  args
		want  []int32
		want1 int32
	}{
		"case": {
			args: args{
				arr: []int32{4, 3, 1, 2},
			},
			want:  []int32{1, 2, 3, 4},
			want1: 5,
		},
		"case1": {
			args: args{
				arr: []int32{1, 2, 3, 4},
			},
			want:  []int32{1, 2, 3, 4},
			want1: 0,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			got, got1 := countSwaps(tt.args.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSwaps() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countSwaps() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
