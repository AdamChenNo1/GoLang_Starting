package stack

import "sync"

type ArrayStack struct {
	array []interface{}
	size  int
	sync.Mutex
}

//入栈
func (stack *ArrayStack) Push(v interface{}) {
	stack.Lock()
	defer stack.Unlock()

	stack.array = append(stack.array, v)

	stack.size++
}

//出栈
func (stack *ArrayStack) Pop() interface{} {
	stack.Lock()
	defer stack.Unlock()

	//栈已空
	if stack.size == 0 {
		return nil
	}

	//栈顶元素
	v := stack.array[stack.size-1]

	//动态改变底层slice的容量
	var newArray []interface{}
	length := len(stack.array) - 1
	if length < 1/2*cap(stack.array) {
		newArray = stack.array[0:length:length]
	} else {
		newArray = stack.array[0:length]
	}
	stack.array = newArray

	stack.size--
	return v
}

//取栈顶元素
func (stack *ArrayStack) Top() interface{} {
	stack.Lock()
	defer stack.Unlock()

	//栈已空
	if stack.size == 0 {
		return nil
	}

	//栈顶元素
	v := stack.array[stack.size-1]

	return v
}

//获取栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

//判断栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
