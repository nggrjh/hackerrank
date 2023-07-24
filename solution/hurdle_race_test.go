package solution

import "testing"

func Test_hurdleRace(t *testing.T) {
	t.Parallel()
	type args struct {
		jump    int32
		heights []int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_1": {
			args: args{
				jump:    4,
				heights: []int32{1, 6, 3, 5, 2},
			},
			want: 2,
		},
		"case_2": {
			args: args{
				jump:    7,
				heights: []int32{2, 5, 4, 5, 2},
			},
			want: 0,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := hurdleRace(tt.args.jump, tt.args.heights); got != tt.want {
				t.Errorf("hurdleRace() = %v, want %v", got, tt.want)
			}
		})
	}
}
