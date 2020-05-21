# 在单链表和双链表中删除倒数第 K 个节点
## 要求：
分别实现两个函数，一个可以删除单链表中倒数第K个节点，另一个可以删除双链表中倒数第K个节点。
## 原理：
- 对于单链表，若K等于链表长度L，则删除首节点，K为小于L的非负整数，则删除第L-K+1个节点，见"src\algorithms\linked_list\LinkedList\src\SGLinkedList.go" func (list *SGLinkedList) DeleteReverseK(k int)  
- 对于双向链表，见"src\algorithms\linked_list\LinkedList\src\DLLinkedList.go" func (list *DLLinkedList) DeleteReverseK(k int)
  - 若K等于链表长度L，则删除首节点，K介于L和L/2之间，则从首节点开始遍历，删除第L-K+1个节点
  - 若K等于1，则删除末尾节点，若K介于1和L/2之间，则从尾节点开始遍历，删除第K个节点