# 根据前序和中序遍历重建二叉树

from binarytree import Node, tree

# 待调试
def BuildBinaryTree(pre_list, in_list, pre_start, pre_end, in_start, in_end):    
    # 构建新的二叉树节点
    root = Node(pre_list[pre_start])

    # 前序已经是最后节点
    if pre_start == pre_end:
        return root

    # 中序遍历中查找根节点位置（不考虑异常情况）
    index = in_start
    while index < in_end and pre_list[index] != root.value:
        index += 1
    
    # 构建左子树
    if index > in_start:
        # pre_start+index-in_start 中序遍历当前位置
        root.left = BuildBinaryTree(pre_list, in_list, pre_start+1, 
        pre_start+index-in_start, in_start, index-1)
    if index < in_end:
        root.right = BuildBinaryTree(pre_list, in_list, pre_start+index-in_start+1,
         pre_end, index+1, in_end)
    return root

def BuildTree(pre_list, in_list):
    if pre_list == None or in_list == None:
        return None
    return BuildBinaryTree(pre_list, in_list, 0, len(pre_list)-1, 0, len(in_list)-1)

pre_list = [1, 2, 4, 7, 3, 5, 6, 8]
in_list = [4, 7, 2, 1, 5, 3, 8, 6]
print(BuildTree(pre_list, in_list))
