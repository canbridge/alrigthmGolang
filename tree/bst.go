package tree

/* 外面套一层Tree 是必要的， 方便用面向对象的方式来写结构体方法,否则需要用return 的方式来返回值，
而不是用修改指针所指向的 值的方法来修改/改变值， 尤其是在Tree.insert（v） 的时候。 如T果不嵌套一层Tree，直接用TreeNode
插入的话， 需要盘点TreeNode是否为空， 代码写起来很丑， 同理， 链表的实现也可以用这种方式： List， ListNode的方式
*/
type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	v           int
	left, right *TreeNode
}

func (tree *Tree) Find(v int) *TreeNode {
	return tree.root.find(v)
}

func (root *TreeNode) find(v int) *TreeNode {
	for root != nil {
		if root.v == v {
			return root
		} else if root.v > v {
			return root.left.find(v)
		} else if root.v < v {
			return root.right.find(v)
		}
	}
	return nil
}

func (root *TreeNode) insert(v int) {

	if root.v < v {
		if root.right == nil {
			root.right = &TreeNode{v: v}
		} else {
			root.right.insert(v)
		}
	} else {
		if root.left == nil {
			root.left = &TreeNode{v: v}
		} else {
			root.left.insert(v)
		}
	}

}

// 这里Struct 需要是一个已经实例化了的具体对象， 跟OO里面的实例方法有点类似
func (tree *Tree) Insert(v int) {
	root := tree.root
	if root == nil {
		tree.root = &TreeNode{v: v}
	} else if v < root.v {
		root.insert(v)
	} else if v > root.v {
		root.insert(v)
	}
}

//func main() {
//	tree := new(Tree)
//	tree.Insert(21)
//	tree.Insert(11)
//	tree.Insert(33)
//	log.Println(111, tree.root)
//	log.Println(111, tree.root.left)
//	log.Println(111, tree.root.right)
//	log.Println(2222, tree.Find(21))
//}
