package solution

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (s *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if s.head == nil {
		s.head = node
	} else {
		s.tail.next = node
	}

	s.tail = node
}

func printLinkedList(head *SinglyLinkedListNode) {
	for _, v := range getLinkedList(head) {
		fmt.Println(v)
	}
}

func getLinkedList(head *SinglyLinkedListNode) []int32 {
	var values []int32
	for {
		if head == nil {
			break
		}

		values = append(values, head.data)
		head = head.next
	}
	return values
}
