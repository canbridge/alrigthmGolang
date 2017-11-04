// 链表相关的一些算法
package list

import (
	"log"
)

type Node struct {
	next *Node
	v    int
}

// 将数组添加到空链表
func add_slice(root *Node, slice []int) {
	var prev *Node

	for i, v := range slice {
		var c *Node

		if i == 0 {
			root.v = v
			c = root
		} else {
			c = &Node{nil, v}
			prev.next = c
		}
		prev = c
	}

}

// 链表的遍历
func iteral(root *Node) {
	c := root
	for {
		log.Println(c.v)
		if c.next == nil {
			break
		}
		c = c.next

	}
}

// 返回倒数第K个元素
func last_k(head *Node, k int) {
	slow := head
	fast := head
	diff := 0
	for {
		log.Println("fast pointer", fast.v, fast.next)

		if diff < k {
			diff++
		}

		if fast.next == nil {

			log.Println("diff", diff)
			if diff < k {
				log.Println("no")
			} else {
				log.Println("last k:", slow.v)
			}
			break
		}
		fast = fast.next

		if diff == k {
			slow = slow.next
		}

	}

}

// 链表的反转
func reverse(head *Node) *Node {
	p1 := head
	p2 := p1.next
	p1.next = nil
	for {
		if p2 == nil {
			return p1
		}

		p3 := p2.next
		p2.next = p1

		p1 = p2
		p2 = p3

	}
}

// 链表的归并排序
func listMergeSort(head *Node) *Node {
	if head.next == nil || head == nil {
		return head
	}

	/* Split the nodes of the given list into front and back halves,
	   and return the two lists using the references parameters.
	   If the length is odd, the extra node shold go in the front list.
	   Uses the fast/slow pointer strategy. */
	/*最终结果是将一个链表切成两个: head_one, head_two,然后对head_one, head_two分别递归,
	且head_two的长度小于等于2, 比如 a > b > c > d > e > f 会被
	切成 a > b > c > d 和 e > f 两个链表
	*/
	head_one := head
	head_two := head.next
	for (head_two != nil) && head_two.next != nil {
		head = head.next
		head_two = head.next.next
	}
	head_two = head.next
	head.next = nil

	return mergelist(listMergeSort(head_one), listMergeSort(head_two))
}

// 两个排序好链表merge
func mergelist(l1, l2 *Node) *Node {
	var c *Node
	if l2 == nil {
		c = l1
		return c
	}

	if l1 == nil {
		c = l2
		return c
	}

	if l1.v <= l2.v {
		c = l1
		c.next = mergelist(l1.next, l2)
	} else {
		c = l2
		c.next = mergelist(l1, l2.next)
	}
	return c
}

// 合并排序好的N个链表
func mergeSortLists(lists []*Node) *Node {
	length := len(lists)

	// 递归退出条件
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}
	head_one := mergeSortLists(lists[length/2:])
	head_two := mergeSortLists(lists[:length/2])
	return mergelist(head_one, head_two)
}
