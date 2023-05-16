from unittest import TestCase

from solution.linked_list.linked_list import LinkedList


class TestLinkedList(TestCase):
    def test_display(self):
        ll = LinkedList()

        head = None
        for i in [2, 3, 4, 1]:
            head = ll.insert(head, i)

        got = ll.display(head)
        want = "2 3 4 1"
        self.assertEqual(want, got)

    def test_remove_duplicates(self):
        ll = LinkedList()

        head = None
        for i in [1, 2, 2, 3, 3, 4]:
            head = ll.insert(head, i)

        head = ll.remove_duplicates(head)

        got = ll.display(head)
        want = "1 2 3 4"
        self.assertEqual(want, got)
