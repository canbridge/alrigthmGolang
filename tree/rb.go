package tree

/*红黑书, 是一种平衡二叉树， 只不过是不严格的平衡二叉树， 去掉了自平衡二叉树左右子树高度叉小于等于1的限制， \
换成了从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点, 这也是变相的高度差
1.节点是红色或黑色。
2.根节点是黑色。
3.每个叶子节点都是黑色的空节点（NIL节点）。
4 每个红色节点的两个子节点都是黑色。(从每个叶子到根的所有路径上不能有两个连续的红色节点)
5.从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点。
特性4和特性5 保证了从根到任意一个叶子节点的最长路径不会超过最短路径的两倍(2n-1)：举个栗子，
最短路径 b >b >nil, 最长路径b>r>b>r>nil
*/
