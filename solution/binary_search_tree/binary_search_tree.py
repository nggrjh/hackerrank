class Node:
    def __init__(self, data):
        self.right = self.left = None
        self.data = data


class Tree:
    def insert(self, root, data):
        if root is None:
            return Node(data)
        else:
            if data <= root.data:
                cur = self.insert(root.left, data)
                root.left = cur
            else:
                cur = self.insert(root.right, data)
                root.right = cur
        return root

    def get_height(self, root):
        return self.count_nodes(root) - 1

    def count_nodes(self, root):
        if root is None:
            return 0

        left = self.count_nodes(root.left)
        right = self.count_nodes(root.right)

        if left > right:
            return left + 1
        else:
            return right + 1

    def level_order(self, root):
        order = [root.data]
        self.get_children(root, order=order)
        return order

    def get_children(self, root, queue=[], order=[]):
        if root is None:
            return

        for child in [c for c in (root.left, root.right) if c is not None]:
            queue.append(child)
            order.append(child.data)

        if len(queue):
            self.get_children(queue.pop(0), queue, order)
