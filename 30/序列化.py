# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        res = []

        def in_order(node):
            if node is None:
                res.append("#")
                return
            else:
                res.append(str(node.val))
            in_order(node.left)
            in_order(node.right)

        in_order(root)
        return "_".join(res)

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        datas = data.split("_")
        l = len(data)

        def in_order(data, level):
            if level >= l:
                return   

            if data[level] == "#":
                return None
            node = TreeNode(data[level])
            level+=1 
            node.left = in_order(data, level)
            level+=1
            node.right = in_order(data, level)
            return node 
        return in_order(datas, 0)
        
        



# Your Codec object will be instantiated and called as such:
codec = Codec()
root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.right.left = TreeNode(4)
root.right.right = TreeNode(5)
print codec.serialize(root)
# codec.deserialize(codec.serialize(root))