from unittest import TestCase

from solution.linked_list.linked_list import LinkedList


class TestSolution(TestCase):
    def test_display(self):
        ll = LinkedList()

        head = None
        for i in [2, 3, 4, 1]:
            head = ll.insert(head, i)

        got = ll.display(head)
        want = "2 3 4 1"
        self.assertEqual(want, got)
