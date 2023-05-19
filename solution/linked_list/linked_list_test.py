from unittest import TestCase

from solution.linked_list.linked_list import LinkedList


class TestLinkedListDisplay(TestCase):
    def test_case_1(self):
        ll = LinkedList()

        head = None
        for i in [2, 3, 4, 1]:
            head = ll.insert(head, i)

        got = ll.display(head)
        want = "2 3 4 1"
        self.assertEqual(want, got)


class TestLinkedListRemoveDuplicates(TestCase):
    def test_case_1(self):
        ll = LinkedList()

        head = None
        for i in [1, 2, 2, 3, 3, 4]:
            head = ll.insert(head, i)

        head = ll.remove_duplicates(head)

        got = ll.display(head)
        want = "1 2 3 4"
        self.assertEqual(want, got)


class TestLinkedListMerge(TestCase):
    def test_case_0(self):
        ll = LinkedList()

        head1 = None
        for i in [1, 2, 3]:
            head1 = ll.insert(head1, i)

        head2 = None
        for i in [3, 4]:
            head2 = ll.insert(head2, i)

        got = ll.display(ll.merge(head1, head2))
        want = "1 2 3 3 4"
        self.assertEqual(want, got)

    def test_case_1(self):
        ll = LinkedList()

        head1 = None
        for i in [4, 5, 6]:
            head1 = ll.insert(head1, i)

        head2 = None
        for i in [1, 2, 10]:
            head2 = ll.insert(head2, i)

        got = ll.display(ll.merge(head1, head2))
        want = "1 2 4 5 6 10"
        self.assertEqual(want, got)
