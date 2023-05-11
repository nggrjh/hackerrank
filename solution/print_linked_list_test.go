package solution

import (
	"reflect"
	"testing"
)

func Test_getLinkedList(t *testing.T) {
	type args struct {
		nodes []int32
	}
	tests := map[string]struct {
		args args
		want []int32
	}{
		"case": {
			args: args{
				nodes: []int32{16, 13},
			},
			want: []int32{16, 13},
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			list := &SinglyLinkedList{}
			for _, n := range tt.args.nodes {
				list.insertNodeIntoSinglyLinkedList(n)
			}

			if got := getLinkedList(list.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
