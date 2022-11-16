## 数组

### 稀疏数组 sparsearray

#### 一个数组中大部分元素相同

1.记录几行几列，有多少个不同的值

2.把不同的值记录在一个小规模数组中

row    col     val

![截屏2022-11-16 09.49.17](/Users/admin/Library/Application Support/typora-user-images/sparearray.png)

type ValNode struct {

​    row int

​    col int

​    val int

}

## 队列 Queue

### 数组实现队列

1.maxsize

2.front  