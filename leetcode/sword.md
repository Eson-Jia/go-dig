# sword

## 2.4.1 排序

`请实现一个排序算法,要求时间效率为 O(n)`,当我看到这段话的时间,我的心想:搞错了吧,哪有时间效率为 O(n) 的排序算法.但是在读完整个上下文之后,发现还真有. 因为这里有个特殊限制:`这些数字是人的年龄`
,也就是说,被排序的数字在一个很小的范围内,对于这种情况,位图(bitmap)的时间效率是 O(n). 所以首先认真读完题,遇到无法理解的问题不要怀疑别人出错题了,实在不理解就跟面试官沟通清楚.

有点搞笑,前面一本正经对之前的自己进行说教,然后给出了位图这个解决方案.事实却是又一次没有认真思考,题目要求是`实现一个排序算法`.这些排序的数字是公司员工的年龄,
有个特点就是可能会重复,但是位图只能对不重复的数字进行排序.因为位图中使用一`bit`来表示一个数字是否存在,但不能用来表示这个数字存在多少个. 所以正确方式是使用数组这个结构,下标代表年龄,对应的值代表相同年龄的员工有多少位.

## 面试题9: 斐波那契数列

遇到一个题目,没有思路的时候要尝试使用最朴素的方式解题,这样能更深入的理解题目,为更好的解法带来灵感.

## 移位

### 原码 反码 补码 补数 模

- 原码:
- 反码:原码除符号位以外全部取反
- 补码:反码+1
- 补数:对一个n位整数

参考 <<编码 隐匿在计算机软硬件背后的语言>>13. 如何实现减法

