# 删除链表的中间节点和a/b处的节点
## 要求：
给定链表的头节点head、整数a和b，实现删除位于a/b处节点的函数。
## 原理：
- 对于单链表，见"src\algorithms\linked_list\LinkedList\src\SGLinkedList.go" func (list *SGLinkedList) DeleteElementAt(a,b int)  
- 对于双向链表，见"src\algorithms\linked_list\LinkedList\src\DLLinkedList.go" func (list *DLLinkedList) DeleteElementAt(a,b int)