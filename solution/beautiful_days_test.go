package solution

import "testing"

func Test_beautifulDays(t *testing.T) {
	t.Parallel()
	type args struct {
		i int32
		j int32
		k int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case": {
			args: args{
				i: 20,
				j: 23,
				k: 6,
			},
			want: 2,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := beautifulDays(tt.args.i, tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("beautifulDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
