# 构造数组的MaxTree
## 要求：
定义二叉树节点如下：  
type Node struct {  
&emsp;value int  
&emsp;left *Node   
&emsp;right *Node   
}  
一个数组的MaxTree定义如下。  
●　数组必须没有重复元素。  
●　MaxTree是一棵二叉树，数组的每一个值对应一个二叉树节点。  
●　包括MaxTree树在内且在其中的每一棵子树上，值最大的节点都是树的头。  
给定一个没有重复元素的数组arr，写出生成这个数组的MaxTree的函数，要求如果数组长度为N ，则时间复杂度为O (N )、额外空间复杂度为O (N )。
## 原理：
关键在于利用双端队列来实现窗口最大值的更新。首先生成双端队列l，l中存放数组arr中的下标。

假设遍历到arr[i]，qmax的放入规则为：

1．如果qmax为空，直接把下标i放进qmax，放入过程结束。

2．如果qmax不为空，取出当前qmax队尾存放的下标，假设为j。

1）如果arr[j]>arr[i]，直接把下标i放进qmax的队尾，放入过程结束。

2）如果arr[j]<=arr[i]，把j从qmax中弹出，继续qmax的放入规则。

假设遍历到arr[i]，qmax的弹出规则为：

如果qmax队头的下标等于i-w，说明当前qmax队头的下标已过期，弹出当前对头的下标即可。

对每个i，当i >= window-1时，l队首的元素值即为窗口最大值的序号，将对应元素加入最大值数组末尾即可