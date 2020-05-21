# 链表
## 结构：
单链表节点  
type Node struct {  
&emsp;Value interface{}  
&emsp;Next *Node  
}

双向链表节点  
type DLNode struct {  
&emsp;Value interface{}  
&emsp;Next *DLNode  
&emsp;Prev *DLNode  
}

