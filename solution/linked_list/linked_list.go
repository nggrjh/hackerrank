package linked_list

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
		}

		if head1.data != head2.data {
			return false
		}

		head1 = head1.next
		head2 = head2.next
	}
}

type linkedList struct {
	head *linkedListNode
	tail *linkedListNode
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
