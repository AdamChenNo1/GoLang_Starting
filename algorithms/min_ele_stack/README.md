# 设计一个有getMin功能的栈  
## 题目：
实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作。
## 要求：
1. pop、push、getMin操作的时间复杂度都是O (1)。  
2. 设计的栈类型可以使用现成的栈结构。
## 解答：
使用两个栈，一个栈用来保存当前栈中的元素，其功能和一个正常的栈没有区别，这个栈记为stackData；另一个栈用于保存每一步的最小值，这个栈记为stackMin。  
具体的实现方式有两种：
- 方案一
  - 压入数据规则  
    假设当前数据为newNum，先将其压入stackData。  
    然后判断stackMin是否为空：
    - 如果为空，则newNum也压入stackMin
    - 如果不为空，则比较newNum和stackMin的栈顶元素：
      - 如果newNum更小，则newNum也压入stackMin
      - 如果相等或者stackMin的栈顶元素更小，则stackMin不压入任何内容：
  - 弹出数据规则  
    先在stackData中弹出栈顶元素，记为value。
    然后比较当前stackMin的栈顶元素和value哪一个更小：
      - 当value等于stackMin的栈顶元素时，stackMin弹出栈顶元素
      - 当value大于stackMin的栈顶元素时，stackMin不弹出栈顶元素；返回value  
    <strong>Note</strong>:  stackMin中存在的元素是从栈底到栈顶逐渐变小的，stackMin栈顶的元素既是stackMin栈的最小值，也是当前stackData栈的最小值。  
    所以不会出现value比stackMin的栈顶元素更小的情况，value只可能大于或等于stackMin的栈顶元素。
  - 查询当前栈中最小值  
    stackMin的栈顶元素始终是当前stackData中的最小值
   