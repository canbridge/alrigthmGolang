/*用数组来实现Stack, 当然也可以用链表来实现Stack,
最好的肯定是C语言自己实现的stack
注意:TOP指的是栈顶, 栈顶是最新一个PUSH进去的那个元素, 而不是
最老进去的元素(栈底)

注意： 所有的出栈都需要检验是否为Empty, 省略了
*/
package stack

import (
	"log"
)

type Stack []int //非线程安全, 如果想做成安全的, 需要加mutex

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Empty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Len() int {
	return len(*s)
}

// 栈的排序, 这个好像有点难度,想了好半天才写出来
func (s *Stack) sortStack() {
	s2 := new(Stack)
	for !s.Empty() {
		tmp := s.Top()
		s.Pop()

		// 确保S2的 栈底的元素为最大的， 有点类似数组的选择排序的感觉
		for !s2.Empty() && s2.Top() < tmp {
			s.Push(s2.Top())
			s2.Pop()
		}
		s2.Push(tmp)
	}
	for !s2.Empty() {
		s.Push(s2.Top())
		s2.Pop()
	}
}

// 利用栈的POP来遍历下, 方便查看内容
func (s *Stack) Iteral() {
	for !s.Empty() {
		log.Println(s.Pop())
	}
}

/*
min stack, 在o(1)时间内
Design a stack that supports push, pop, top, and retrieving the minimum
element in constant time.
原理:添加一个新栈, 栈顶元素为原始栈的最小值
每当PUSH的值小于等于新栈栈顶时候则往新栈里面PUSH该元素;
POP时候将POP的值同新栈栈顶元素进行比较, 小于或等于则新栈POP
*/
type MinStack struct {
	s1 *Stack
	s2 *Stack // 存储最小值
}

func NewMS() *MinStack {
	s1 := new(Stack)
	s2 := new(Stack)
	ms := &MinStack{s1: s1, s2: s2}
	return ms

}

func (ms *MinStack) Push(v int) {
	if ms.s1.Empty() {
		ms.s2.Push(v)
	} else {
		top := ms.s2.Top()
		if v <= top {
			ms.s2.Push(v)
		}
	}
	ms.s1.Push(v)
}

func (ms *MinStack) Pop() int {
	top1 := ms.s1.Top()
	top2 := ms.s2.Top()
	//实际上要校验是否有值,此处省略
	if top1 <= top2 {
		ms.s2.Pop()
	}

	return ms.s1.Pop()
}

func (ms *MinStack) GetMin() int {
	return ms.s2.Top()
}

// 用两个栈实现队列
type Queue struct {
	s1 *Stack
	s2 *Stack
}

func NewQ() *Queue {
	s1 := new(Stack)
	s2 := new(Stack)
	ms := &Queue{s1: s1, s2: s2}
	return ms

}

// 进队列
func (q *Queue) in(v int) {
	for !q.s2.Empty() {
		q.s1.Push(q.s2.Pop())
	}
	q.s1.Push(v)
}

// 出队列
func (q *Queue) out() int {
	for !q.s1.Empty() {
		q.s2.Push(q.s1.Pop())
	}
	return q.s2.Pop()
}

//func main() {
//	q := NewQ()
//	q.in(1)
//	q.in(2)
//	q.in(3)
//	log.Println(q.out())
//	log.Println(q.out())
//	q.in(4)
//	log.Println(q.out())
//	log.Println(q.out())
//	log.Println(q.out())
//	log.Println(q.out())

//	ms := new(Stack)
//	ms.Push(9)
//	ms.Push(19)
//	ms.Push(11)
//	ms.Push(8)
//	ms.Push(1)

//	ms.sortStack()
//	ms.Iteral()

//	log.Println("MIN", ms.GetMin())
//	ms.Pop()
//	log.Println("MIN", ms.GetMin())
//	ms.Pop()
//	log.Println("MIN", ms.GetMin())
//	ms.Pop()
//	log.Println("MIN", ms.GetMin())
//	ms.Pop()
//	log.Println("MIN", ms.GetMin())

//}
