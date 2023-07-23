package solution

import "testing"

func Test_angryProfessor(t *testing.T) {
	t.Parallel()
	type args struct {
		attendance int32
		students []int32
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"case_1": {
			args: args{
				attendance: 3,
				students: []int32{-1, -3, 4, 2},
			},
			want: "YES",
		},
		"case_2": {
			args: args{
				attendance: 2,
				students: []int32{0, -1, 2, 1},
			},
			want: "NO",
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := angryProfessor(tt.args.attendance, tt.args.students); got != tt.want {
				t.Errorf("angryProfessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
