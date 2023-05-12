package linked_list

import (
	"reflect"
	"testing"
)

func Test_linkedListNode_GetQueue(t *testing.T) {
	t.Parallel()
	type args struct {
		nodes []int32
	}
	tests := map[string]struct {
		args args
		want []int32
	}{
		"case_1": {
			args: args{
				nodes: []int32{16, 13},
			},
			want: []int32{16, 13},
		},
		"case_2": {
			args: args{
				nodes: []int32{141, 302, 164, 530, 474},
			},
			want: []int32{141, 302, 164, 530, 474},
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			list := &linkedList{}
			for _, n := range tt.args.nodes {
				list.queue(n)
			}

			if got := list.head.GetQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkedListNode_GetStack(t *testing.T) {
	t.Parallel()
	type args struct {
		nodes []int32
	}
	tests := map[string]struct {
		args args
		want []int32
	}{
		"case_1": {
			args: args{
				nodes: []int32{16, 13},
			},
			want: []int32{13, 16},
		},
		"case_2": {
			args: args{
				nodes: []int32{141, 302, 164, 530, 474},
			},
			want: []int32{474, 530, 164, 302, 141},
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			list := &linkedList{}
			for _, n := range tt.args.nodes {
				list.queue(n)
			}

			if got := list.tail.GetStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkedListNode_IsEqual(t *testing.T) {
	t.Parallel()
	type args struct {
		nodes1 []int32
		nodes2 []int32
	}
	tests := map[string]struct {
		args args
		want bool
	}{
		"case_1": {
			args: args{
				nodes1: []int32{1, 2, 3},
				nodes2: []int32{1, 2, 3},
			},
			want: true,
		},
		"case_2": {
			args: args{
				nodes1: []int32{1, 2, 3},
				nodes2: []int32{1, 3},
			},
			want: false,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			l1 := &linkedList{}
			for i := range tt.args.nodes1 {
				l1.queue(tt.args.nodes1[i])
			}

			l2 := &linkedList{}
			for i := range tt.args.nodes2 {
				l2.queue(tt.args.nodes2[i])
			}

			if got := l1.head.IsEqual(l2.head); got != tt.want {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkedListNode_MergeSort(t *testing.T) {
	t.Parallel()
	type args struct {
		nodes1 []int32
		nodes2 []int32
	}
	tests := map[string]struct {
		args args
		want []int32
	}{
		"case_1": {
			args: args{
				nodes1: []int32{1, 2, 3},
				nodes2: []int32{1, 2, 3},
			},
			want: []int32{1, 1, 2, 2, 3, 3},
		},
		"case_2": {
			args: args{
				nodes1: []int32{1, 2, 3},
				nodes2: []int32{1, 3},
			},
			want: []int32{1, 1, 2, 3, 3},
		},
		"case_3": {
			args: args{
				nodes1: []int32{1, 2, 3},
				nodes2: []int32{1, 3},
			},
			want: []int32{1, 1, 2, 3, 3},
		},
		"case_4": {
			args: args{
				nodes1: []int32{1, 3, 7},
				nodes2: []int32{1, 2},
			},
			want: []int32{1, 1, 2, 3, 7},
		},
		"case_5": {
			args: args{
				nodes1: []int32{1, 3, 7},
				nodes2: []int32{3, 4},
			},
			want: []int32{1, 3, 3, 4, 7},
		},
		"case_6": {
			args: args{
				nodes1: []int32{1, 2, 3},
				nodes2: []int32{3, 4},
			},
			want: []int32{1, 2, 3, 3, 4},
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			l1 := &linkedList{}
			for i := range tt.args.nodes1 {
				l1.queue(tt.args.nodes1[i])
			}

			l2 := &linkedList{}
			for i := range tt.args.nodes2 {
				l2.queue(tt.args.nodes2[i])
			}

			l1.head.MergeSort(l2.head)
			if got := l1.head.GetQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkedList_sort(t *testing.T) {
	t.Parallel()
	type args struct {
		nodes []int32
	}
	tests := map[string]struct {
		args args
		want []int32
	}{
		"case_1": {
			args: args{
				nodes: []int32{16, 13},
			},
			want: []int32{13, 16},
		},
		"case_2": {
			args: args{
				nodes: []int32{141, 302, 164, 530, 474},
			},
			want: []int32{141, 164, 302, 474, 530},
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			list := &linkedList{}
			for _, n := range tt.args.nodes {
				list.queue(n)
			}

			list.sort()
			if got := list.head.GetQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
