# 反转单向和双向链表
## 要求：
给定一个链表，以及两个整数from和to，在链表上把第from个节点到第to个节点这一部分进行反转。
如果链表长度为N ，时间复杂度要求为O (N )，额外空间复杂度要求为O (1)。
## 原理：
- 对于单链表，见"src\algorithms\linked_list\LinkedList\src\SGLinkedList.go" func (list *SGLinkedList) ReversePart(from,to int)  
- 对于双向链表，见"src\algorithms\linked_list\LinkedList\src\DLLinkedList.go" func (list *DLLinkedList) ReversePart(from,to int)