package linked_list

import (
	"sort"
)

type linkedListNode struct {
	data int32
	next *linkedListNode
	prev *linkedListNode
}

func (n *linkedListNode) GetQueue() []int32 {
	var nodes []int32
	head := n
	for {
		if head == nil {
			break
		}

		nodes = append(nodes, head.data)
		head = head.next
	}
	return nodes
}

func (n *linkedListNode) GetStack() []int32 {
	var nodes []int32
	tail := n
	for {
		if tail == nil {
			break
		}

		nodes = append(nodes, tail.data)
		tail = tail.prev
	}
	return nodes
}

func (n *linkedListNode) IsEqual(head *linkedListNode) bool {
	head1 := n
	head2 := head

	for {
		if head1 == nil && head2 == nil {
			return true
		} else if head1 == nil || head2 == nil {
			return false
		}

		if head1.data != head2.data {
			return false
		}

		head1 = head1.next
		head2 = head2.next
	}
}

func (n *linkedListNode) MergeSort(head *linkedListNode) {
	head1 := n
	head2 := head

	for {
		if head2 == nil {
			break
		}

		if head1.data <= head2.data {
			if head1.next == nil {
				head1.next = head2
				head2.prev = head1

				head2 = head2.next
			}

			head1 = head1.next
			continue
		}

		next := head2.next

		if head1.prev != nil {
			head1.prev.next = head2
		} else {
			*n = *head2
			break
		}

		head2.prev = head1.prev
		head1.prev = head2
		head2.next = head1

		head2 = next
	}
}

type linkedList struct {
	head *linkedListNode
	tail *linkedListNode
}

func (ll *linkedList) sort() {
	nodes := ll.head.GetQueue()
	sort.SliceStable(nodes, func(i, j int) bool { return nodes[i] < nodes[j] })

	*ll = linkedList{}
	for _, n := range nodes {
		ll.queue(n)
	}
}

func (ll *linkedList) queue(data int32) {
	node := &linkedListNode{data: data, prev: ll.tail}

	if ll.head == nil {
		ll.head = node
	} else {
		ll.tail.next = node
	}

	ll.tail = node
}

func (ll *linkedList) stack(data int32) {
	node := &linkedListNode{data: data, next: ll.tail}

	if ll.tail == nil {
		ll.tail = node
	} else {
		ll.head.prev = node
	}

	ll.head = node
}
