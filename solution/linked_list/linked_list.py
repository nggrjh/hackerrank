class Node:
    def __init__(self, data):
        self.data = data
        self.next = None


class LinkedList:
    @staticmethod
    def get(head):
        order = []
        current = head
        while current:
            order.append(current.data)
            current = current.next
        return order

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

    @staticmethod
    def remove_duplicates(head):
        t = list()
        s = LinkedList()

        current = head
        head = None
        while current:
            if current.data not in t:
                head = s.insert(head, current.data)
                t.append(current.data)

            current = current.next
        return head

    def merge(self, head1, head2):
        arr = self.get(head1) + self.get(head2)
        arr.sort()

        head = None
        for i in arr:
            head = self.insert(head, i)

        return head
