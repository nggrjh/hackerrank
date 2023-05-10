package solution

import "testing"

func Test_activityNotifications(t *testing.T) {
	t.Parallel()
	type args struct {
		expenditure []int32
		days        int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0": {
			args: args{
				expenditure: []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
				days:        5,
			},
			want: 2,
		},
		"case_6": {
			args: args{
				expenditure: []int32{1, 2, 3, 4, 4},
				days:        4,
			},
			want: 0,
		},
		"case_7": {
			args: args{
				expenditure: []int32{10, 20, 30, 40, 50},
				days:        3,
			},
			want: 1,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := activityNotifications(tt.args.expenditure, tt.args.days); got != tt.want {
				t.Errorf("activityNotifications() = %v, want %v", got, tt.want)
			}
		})
	}
}
