package stack_queue

import (
	dsa "algorithms/stack/src"
)

type StackQueue struct {
	stackPush dsa.ArrayStack
	stackPop  dsa.ArrayStack
}

func (queue *StackQueue) Enqueue(v interface{}) {
	queue.stackPush.Push(v)
	if queue.stackPop.Size() == 0 {
		for {
			if queue.stackPush.Size() == 0 {
				break
			} else {
				queue.stackPop.Push(queue.stackPush.Pop())
			}
		}
	}
}

func (queue *StackQueue) Dequeue() interface{} {
	if queue.stackPop.Size() == 0 {
		for {
			if queue.stackPush.Size() == 0 {
				break
			} else {
				queue.stackPop.Push(queue.stackPush.Pop())
			}
		}
	}
	return queue.stackPop.Pop()
}

func (queue *StackQueue) Peek() interface{} {
	if queue.stackPop.Size() == 0 {
		for {
			if queue.stackPush.Size() == 0 {
				break
			} else {
				queue.stackPop.Push(queue.stackPush.Pop())
			}
		}
	}
	return queue.stackPop.Top()
}

func (queue *StackQueue) Size() int {
	return queue.stackPop.Size() + queue.stackPush.Size()
}
