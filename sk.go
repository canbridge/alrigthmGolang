package main

import (
	"log"
)

type SkipList struct {
	level int
	size  int
	head  *node
}

type node struct {
	key     int
	forward []*node
}

func InitSkip(size int) *SkipList {
	update := make([]*node, size+1)
	//	for i := 0; i <= size; i++ {
	//		update[i] = new(node)
	//	}
	node := &node{forward: update}
	skip := &SkipList{level: 0, size: size, head: node}
	return skip
}

func (list *SkipList) getRandLevel() int {
	return 3
}

func (list *SkipList) update() {
	insertLv := list.getRandLevel()
	insertForwards := make([]*node, insertLv+1)
	insertNode := &node{key: key, forward: insertForwards}
	for i := insertLv; i >= 0; i-- {
		insertNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = insertNode
	}
	if insertLv > list.level {
		list.level = insertLv
	}

}

func (list *SkipList) Insert(key int) {
	// update 用来保存每一行最后一个检索到节点的列表
	update := make([]*node, list.size+1)
	current := list.head
	for i := list.size; i >= 0; i-- {
		// 最终current 最下面一层链表的小于key的最大值的node
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]

	if current == nil {
		// 这种情况实际上是在之前没有插入数据
		insertLv := list.getRandLevel()
		insertForwards := make([]*node, insertLv+1)
		insertNode := &node{key: key, forward: insertForwards}
		for i := insertLv; i >= 0; i-- {
			insertNode.forward[i] = update[i].forward[i]
			update[i].forward[i] = insertNode
		}
		if insertLv > list.level {
			list.level = insertLv
		}
	} else if current.key > key {
		// 这种情况实际上是在之前没有插入数据
		insertLv := list.getRandLevel()
		insertForwards := make([]*node, insertLv+1)
		insertNode := &node{key: key, forward: insertForwards}
		for i := insertLv; i >= 0; i-- {
			insertNode.forward[i] = update[i].forward[i]
			update[i].forward[i] = insertNode
		}

		if insertLv > list.level {
			list.level = insertLv
		}
	} else if current.key == key {
		// 需要插入的数据已经存在于列表中了， 不需要插入
	}

}

func (list *SkipList) iteral() {
	head := list.head
	c := head.forward[0]
	for c != nil {
		log.Println(c.key)
		c = c.forward[0]
	}
	log.Println("head", head)
	next := head.forward[0]
	log.Println("next", next.forward[0])
}

func main() {

	list := InitSkip(3)
	for _, v := range []int{1, 98, 89, 87, 87, 65, 11} {
		list.Insert(v)
	}

	list.iteral()
}
