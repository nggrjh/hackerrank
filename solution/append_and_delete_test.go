package solution

import "testing"

func Test_appendAndDelete(t *testing.T) {
	t.Parallel()
	type args struct {
		initial    string
		desired    string
		operations int32
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"case_0.0": {
			args: args{
				initial:    "abc",
				desired:    "def",
				operations: 6,
			},
			want: "Yes",
		},
		"case_0.1": {
			args: args{
				initial:    "hackerhappy",
				desired:    "hackerrank",
				operations: 9,
			},
			want: "Yes",
		},
		"case_1": {
			args: args{
				initial:    "aba",
				desired:    "aba",
				operations: 7,
			},
			want: "Yes",
		},
		"case_2": {
			args: args{
				initial:    "aaaaaaaaaa",
				desired:    "aaaaa",
				operations: 7,
			},
			want: "Yes",
		},
		"case_3": {
			args: args{
				initial:    "zzzzz",
				desired:    "zzzzzzz",
				operations: 4,
			},
			want: "Yes",
		},
		"case_4": {
			args: args{
				initial:    "qwerasdf",
				desired:    "qwerbsdf",
				operations: 6,
			},
			want: "No",
		},
		"case_5": {
			args: args{
				initial:    "y",
				desired:    "yu",
				operations: 2,
			},
			want: "No",
		},
		"case_10": {
			args: args{
				initial:    "abcd",
				desired:    "abcdert",
				operations: 10,
			},
			want: "No",
		},
		"case_13": {
			args: args{
				initial:    "ashley",
				desired:    "ash",
				operations: 2,
			},
			want: "No",
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := appendAndDelete(tt.args.initial, tt.args.desired, tt.args.operations); got != tt.want {
				t.Errorf("appendAndDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
