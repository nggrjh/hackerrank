package solution

import "testing"

func Test_birthday(t *testing.T) {
	t.Parallel()
	type args struct {
		s []int32
		d int32
		m int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0.1": {
			args: args{
				s: []int32{2, 2, 1, 3, 2},
				d: 4,
				m: 2,
			},
			want: 2,
		},
		"case_0.2": {
			args: args{
				s: []int32{1, 2, 1, 3, 2},
				d: 3,
				m: 2,
			},
			want: 2,
		},
		"case_1": {
			args: args{
				s: []int32{1, 1, 1, 1, 1, 1},
				d: 3,
				m: 2,
			},
			want: 0,
		},
		"case_2": {
			args: args{
				s: []int32{4},
				d: 4,
				m: 1,
			},
			want: 1,
		},
		"case_10": {
			args: args{
				s: []int32{3, 5, 4, 1, 2, 5, 3, 4, 3, 2, 1, 1, 2, 4, 2, 3, 4, 5, 3, 1, 2, 5, 4, 5, 4, 1, 1, 5, 3, 1, 4, 5, 2, 3, 2, 5, 2, 5, 2, 2, 1, 5, 3, 2, 5, 1, 2, 4, 3, 1, 5, 1, 3, 3, 5},
				d: 18,
				m: 6,
			},
			want: 10,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := birthday(tt.args.s, tt.args.d, tt.args.m); got != tt.want {
				t.Errorf("birthday() = %v, want %v", got, tt.want)
			}
		})
	}
}
