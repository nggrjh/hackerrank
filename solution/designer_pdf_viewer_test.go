package solution

import "testing"

func Test_designerPdfViewer(t *testing.T) {
	t.Parallel()
	type args struct {
		heights []int32
		word string
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_1": {
			args: args{
				heights: []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
				word: "abc",
			},
			want: 9,
		},
		"case_2": {
			args: args{
				heights: []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
				word: "zaba",
			},
			want: 28,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := designerPdfViewer(tt.args.heights, tt.args.word); got != tt.want {
				t.Errorf("designerPdfViewer() = %v, want %v", got, tt.want)
			}
		})
	}
}
