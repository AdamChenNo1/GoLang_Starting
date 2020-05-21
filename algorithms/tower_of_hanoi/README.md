# 用栈来求解汉诺塔问题
## 要求：
限制不能从最左侧的塔直接移动到最右侧，也不能从最右侧直接移动到最左侧，而是必须经过中间。当塔有N层的时候，打印最优移动过程和最优移动总步数。
## 原理：
递归的方法:(见recurisuvely.go)  
- 外部调用函数为Hannoi(n int, from int, to int)，其中n为汉诺塔层数，from和to分别为起始点和终点，均为int型，以-1、0、1分别代表左中右  
- Hannoi函数中先进行输入值判断,然后调用move(n int, from int, to int) int函数，这是主要的函数，它实现了将1-n层汉诺塔移动的过程，并记录了移动的总步数
- move函数依次调用moveBottle(n int, from int, to int) (steps int,begin int) 函数按从底至顶的顺序将每个底层移动至目的地
- moveBottle函数中分情况处理：
  - 如果n=1,则调用moveN(n int, from int, to int) int函数将该层移动至目的地,并返回移动的步数和上层此时的位置0
  - 否则，判断中转地的值，
    - 若中转地值不为0，则先调用move函数将n-1层移动至中转地，然后调用moveN(n int, from int, to int) int函数将第n层移动至目的地,并返回移动的步数和上层此时的位置
    - 若，中转地值为0，则则先调用move函数将n-1层移动至目的地，然后调用moveN函数将第n层移动至中转地，接着调用move函数将n-1层移动至起始地，最后调用moveN函数将第n层移动至目的地,并返回移动的步数和上层此时的位置
- moveN函数的功能为将第n层移动至目的地，并记录移动的步数，如果起始地或者目的地为0，则调用1次printStep打印移动的步骤，否则调用2次printStep
- printStep函数输出"Move n from x to y"

栈的方法：(见stackful.go)
- 外部调用函数为Hannoi_v2(n int, from int, to int)，其中n为汉诺塔层数，from和to分别为起始点和终点，均为int型，以-1、0、1分别代表左中右 
- Hannoi函数中先进行输入值判断,并根据from和to的值生成相应的打印列表maps,然后调用move_v2(maps map[int]string, n int, from int, to int) int函数，这是主要的函数，它实现了将1-n层汉诺塔移动的过程，并记录了移动的总步数 
- move_v2函数首先调用InitStack(n int, from int) map[int]*dsa.ArrayStack 函数初始化栈；然后分情况讨论：
  - 如果n=1，则按起始点或终点是否为0，分别调用1次和2次printStepWithMap(maps map[int]string, n int, from int, to int)
  - 否则，当目的栈的长度不为n时，持续循环
    - 若LS为空，则判断MS的长度的奇偶性：
      - 若为偶数，则MS顶部元素出栈，进入LS栈，打印元素值移动，步数加1
      - 若为奇数：
        - 如果RS不为空并且MS顶部元素值大于RS顶部元素值，则RS顶部元素出栈，进入MS，打印元素值移动，步数加1
        - 否则，MS顶部元素值出栈进入RS栈，打印元素值移动，步数加1
    - 否则如果MS为空：
      - 如果RS为空，或者RS顶部元素值小于LS顶部元素值，则LS顶部元素出栈，进入RS栈，打印2次元素值移动，步数加2
      - 否则，LS顶部元素出栈，进入MS栈，打印元素值移动，步数加1
    - 否则，如果RS为空，判断MS长度奇偶性：
      -  若为偶数：
         -  如果LS顶部元素值大于MS顶部元素值，则MS顶部元素出栈，进入LS，打印元素值移动，步数加1
         -  否则，LS顶部元素出栈，进入MS，打印元素值移动，步数加1
      - 若为奇数，则MS顶部元素出栈，进入RS，打印元素值移动，步数加1
    - 若，LS、MS、RS均不空，判断MS长度奇偶性：
      - 若为奇数：
        - 如果MS顶部元素值小于RS顶部元素值，则MS顶部元素出栈，进入RS，打印元素值移动，步数加1
        - 否则，RS顶部元素值出栈，进入MS，打印元素值移动，步数加1
      - 若为偶数：
        - 如果LS顶部元素值小于MS顶部元素值，则LS顶部元素出栈，进入MS，打印元素值移动，步数加1
        - 否则，MS顶部元素出栈，进入LS，打印元素值移动，步数加1
- InitStack函数中首先将1-n按从大到小的顺序依次进栈，然后分情况处理：
  - 如果from=0,则调用返回{-1: LS, 0: stack, 1: RS}
  - 否则，返回{-1: stack, 0: LS , 1: RS}
  - 其中除stack外均为空栈
- printStepWithMap函数按照maps输出"Move n from x to y"
