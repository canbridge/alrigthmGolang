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
	log.Println("update,11111", update)
	node := &node{forward: update}
	skip := &SkipList{level: 0, size: size, head: node}
	return skip
}

func (list *SkipList) Insert(key int) {
	// update 用来保存每一行最后一个检索到节点的列表
	update := make([]*node, list.size+1)
	current := list.head
	for i := list.size; i >= 0; i-- {
		// 最终current 保存小于key的最大值的node
		log.Println(1111, "current", current.forward[0])
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	log.Println("current, update", current.forward[0], update)

	current = current.forward[0]
	log.Println("current, if nil", current)
	if current == nil {
		// 这种情况实际上是没有插入数据插入数据

	} else if current.key == key {
		// 需要插入的数据已经存在于列表中了

	} else if current.key > key {

	}

}

func main() {
	log.Println("21")
	list := InitSkip(3)
	log.Println(list.head.forward[2])
	list.Insert(4)
	log.Println(list.head.forward[0])
}
