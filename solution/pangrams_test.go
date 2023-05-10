package solution

import "testing"

func Test_pangrams(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"case_1": {
			args: args{
				s: "The quick brown fox jumps over the lazy dog",
			},
			want: "pangram",
		},
		"case_2": {
			args: args{
				s: "We promptly judged antique ivory buckles for the next prize",
			},
			want: "pangram",
		},
		"case_3": {
			args: args{
				s: "We promptly judged antique ivory buckles for the prize",
			},
			want: "not pangram",
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := pangrams(tt.args.s); got != tt.want {
				t.Errorf("pangrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
