from unittest import TestCase

from solution.binary_search_tree.binary_search_tree import Tree


class TestTreeGetHeight(TestCase):
    def test_case_1(self):
        tree = Tree()
        root = None
        for i in [3, 5, 2, 1, 4, 6, 7]:
            root = tree.insert(root, i)

        got = tree.get_height(root)
        want = 3
        self.assertEqual(want, got)


class TestTreeLevelOrder(TestCase):
    def test_case_1(self):
        tree = Tree()
        root = None
        for i in [3, 5, 4, 7, 2, 1]:
            root = tree.insert(root, i)

        got = tree.level_order(root)
        want = [3, 2, 5, 1, 4, 7]
        self.assertEqual(want, got)

    def test_case_2(self):
        tree = Tree()
        root = None
        for i in [20, 50, 35, 44, 9, 15, 62, 11, 13]:
            root = tree.insert(root, i)

        got = tree.level_order(root)
        want = [20, 9, 50, 15, 35, 62, 11, 44, 13]
        self.assertEqual(want, got)
