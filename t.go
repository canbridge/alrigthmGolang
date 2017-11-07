package main

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
	lv      int
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

	for i := 0; i < maxlv; i++ {
		update[0] = new(node)
	}
	head := &node{forward: update, lv: maxlv}
	skip := &Skip{head: head, maxlv: maxlv}
	return skip
}

func (skip *Skip) Insert(v int) {
	c := skip.head
	lv := c.lv
	log.Println(c)
	for c.v < v && lv > 0 {
		if c.forward[lv].forward == nil {
			lv = lv - 1
			c = c.forward(lv)

		}
		c = c.forward(lv)
	}
}

func main() {
	skip := NewSkip(10)
	log.Println(skip)

	t := make([]int, 1, 10)
	log.Println(t)
}
