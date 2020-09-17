package ch2

// stack的定义：LIFO的对象的容器。支持pop(o) 和 push()两种方法的容器
// 基于数组实现

type IntStack []int

func (h *IntStack) push(o interface{}) {
	*h = append(*h, o.(int))
}

func (h *IntStack) pop() interface{},error {
	old = *h;
	l := len(old)
	if l < 0 {
		return nil,new error{}
	} 
	*h = old[0,l-1]
}
