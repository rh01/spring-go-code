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
        def encode(node):
            if root is None:
                res.append("#")
                return
            else:
                res.append(str(node.val))
            res.append()
            encode(node.left)
            encode(node.right)
        encode(root)
        return "_".join(res)

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        ds = data.split("_")
        l = len(ds)
        def decode(level):
            if level >= l:
                return 
            
            if ds[level] == "#":
                return None 
            
            node = TreeNode(ds[level])
            node.left(level+1)
            node.right(level+1)         
            return node
        return decode(0)
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))