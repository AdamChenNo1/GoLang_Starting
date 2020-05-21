# 反转单向和双向链表
## 要求：
分别实现反转单向链表和反转双向链表的函数。
如果链表长度为N ，时间复杂度要求为O (N )，额外空间复杂度要求为O (1)。
## 原理：
- 对于单链表，若K等于链表长度L，则删除首节点，K为小于L的非负整数，则删除第L-K+1个节点，见"src\algorithms\linked_list\LinkedList\src\SGLinkedList.go" func (list *SGLinkedList) Reverse(k int)  
- 对于双向链表，见"src\algorithms\linked_list\LinkedList\src\DLLinkedList.go" func (list *DLLinkedList) Reverse(k int)