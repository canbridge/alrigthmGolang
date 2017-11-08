package skip

import (
	"log"
	"math/rand"
)

type Skip struct {
	head  *node
	lv    int
	maxlv int
}

type node struct {
	v       int
	forward []*node
}

func (skip *Skip) get_lv() int {
	lv := 1
	for rand.Float64() < 0.5 && lv < skip.maxlv {
		lv++
	}
	return lv
}

func NewSkip(maxlv int) *Skip {
	update := make([]*node, maxlv)

	for i := 0; i <= maxlv; i++ {
		update[0] = new(node)
	}
	head := &node{forward: update}
	skip := &Skip{head: head, maxlv: maxlv}
	return skip
}

func (skip *Skip) Insert(v int) {
	update := make([]*node, skip.maxlv)
	c := skip.head

	//	log.Println("lv, c", c, c.forward)

	//	// 插入第1个
	//	log.Println(c)

	//	node := new(node)
	//	node.v = v

	//	//	insert_lv := skip.get_lv()
	//	insert_lv := 3

	//	for i := 0; i <= insert_lv; i++ {
	//		c.forward[i] = node
	//	}
	//	log.Println(insert_lv)

	for lv := skip.maxlv - 1; lv > 0; lv-- {
		log.Println(11111, c.v, c.forward == nil)
		for c.forward[lv].v < v {
			c = c.forward[lv]
			log.Println(2121211)
		}
		update[lv] = c
	}

	log.Println(9871, c, update)
}
