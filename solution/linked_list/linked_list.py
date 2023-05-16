class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class LinkedList:
    @staticmethod
    def display(head):
        order = []
        current = head
        while current:
            order.append(str(current.data))
            current = current.next
        return ' '.join(order)

    @staticmethod
    def insert(head, data):
        node = Node(data)

        if head is None:
            return node

        current = head
        while current:
            if current.next is None:
                current.next = node
                break

            current = current.next

        return head
