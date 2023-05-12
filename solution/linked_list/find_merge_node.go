package linked_list

func findMergeNode(head1, head2 *linkedListNode) int32 {
	currentHead1 := head1.next
	currentHead2 := head2.next

	var point int32
	for {
		if currentHead1 == nil {
			break
		}

		if currentHead2 == nil {
			currentHead1 = currentHead1.next
			currentHead2 = head2.next
			continue
		}

		if currentHead1.data != currentHead2.data {
			currentHead2 = currentHead2.next
			continue
		}

		if currentHead1.IsEqual(currentHead2) {
			point = currentHead1.data
			break
		}
	}

	return point
}
