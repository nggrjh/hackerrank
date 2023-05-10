package solution

import "testing"

func Test_twoArrays(t *testing.T) {
	t.Parallel()
	type args struct {
		k int32
		A []int32
		B []int32
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"case_0": {
			args: args{
				k: 1,
				A: []int32{0, 1},
				B: []int32{0, 2},
			},
			want: "YES",
		},
		"case_1.1": {
			args: args{
				k: 10,
				A: []int32{2, 1, 3},
				B: []int32{7, 9, 8},
			},
			want: "YES",
		},
		"case_1.2": {
			args: args{
				k: 5,
				A: []int32{1, 2, 2, 1},
				B: []int32{3, 3, 3, 4},
			},
			want: "NO",
		},
		"case_2.1": {
			args: args{
				k: 91,
				A: []int32{18, 73, 55, 59, 37, 13, 1, 33},
				B: []int32{81, 11, 77, 49, 65, 26, 29, 49},
			},
			want: "NO",
		},
		"case_2.2": {
			args: args{
				k: 94,
				A: []int32{84, 2, 50, 51, 19, 58, 12, 90, 81, 68, 54, 73, 81, 31, 79, 85, 39, 2},
				B: []int32{53, 102, 40, 17, 33, 92, 18, 79, 66, 23, 84, 25, 38, 43, 27, 55, 8, 19},
			},
			want: "YES",
		},
		"case_2.3": {
			args: args{
				k: 88,
				A: []int32{66, 66, 32, 75, 77, 34, 23, 35},
				B: []int32{61, 17, 52, 20, 48, 11, 50, 5},
			},
			want: "NO",
		},
		"case_2.4": {
			args: args{
				k: 45,
				A: []int32{11, 16, 43, 5, 25, 22, 19, 32, 10, 26, 2, 42, 16, 1},
				B: []int32{33, 1, 1, 20, 26, 7, 17, 39, 23, 34, 10, 11, 38, 20},
			},
			want: "NO",
		},
		"case_2.5": {
			args: args{
				k: 59,
				A: []int32{15, 16, 44, 58, 5, 10, 16, 9, 4, 20, 24},
				B: []int32{37, 45, 41, 46, 8, 23, 59, 57, 51, 44, 59},
			},
			want: "YES",
		},
		"case_2.6": {
			args: args{
				k: 32,
				A: []int32{28, 14, 24, 25, 2, 20, 1, 26},
				B: []int32{4, 3, 1, 11, 25, 22, 2, 4},
			},
			want: "NO",
		},
		"case_2.7": {
			args: args{
				k: 57,
				A: []int32{1, 7, 42, 26, 45, 14},
				B: []int32{37, 49, 42, 26, 4, 11},
			},
			want: "NO",
		},
		"case_2.8": {
			args: args{
				k: 88,
				A: []int32{25, 60, 49, 4},
				B: []int32{65, 46, 85, 34},
			},
			want: "YES",
		},
		"case_2.9": {
			args: args{
				k: 9,
				A: []int32{2, 1, 1, 4, 1, 7, 3, 4, 3, 7, 3, 1, 3, 5, 6, 7},
				B: []int32{1, 1, 1, 1, 4, 1, 2, 1, 7, 1, 1, 1, 1, 1, 1, 2},
			},
			want: "NO",
		},
		"case_2.10": {
			args: args{
				k: 70,
				A: []int32{40},
				B: []int32{38},
			},
			want: "YES",
		},
		"case_10": {
			args: args{
				k: 10,
				A: []int32{7, 6, 8, 4, 2},
				B: []int32{5, 2, 6, 3, 1},
			},
			want: "NO",
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := twoArrays(tt.args.k, tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("twoArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
