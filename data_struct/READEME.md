# data struct

## 对比

### map slice array struct 可比较性

- array:如果元素类型是可比较的,数组就是可比较的.比较结果是两边元素的值是否完全相同
- slice:和数组不同 slice 无法做比较,因此不能用`==`来测试两个 slice 是否拥有相同的元素,slice 唯一允许的比较操作和是 nil 作比较
- map:和 slice 一样,map 不可比较,唯一合法的比较就是和 nil 做比较. 为了判断两个 map 是否拥有相同的键值对,必须使用循环
- struct: 如果所有的 field 是可比较的,那么这个`struct`就是可比较的

### slice channel map 的 make

#### slice make

slice := make([]string[,len[,cap]])
slice 使用 make 参数是`type`紧接着两个可选参数分别表示 len 和 cap 分别用于设置 slice 的 len 和 cap
len 如果不提供默认为 0, cap 如果不提供默认为 len,注意 len<=cap 否则会报错

#### map make

map := make(map[string]int[,size])
map 的 size 用于控制创建 map 时 预先创建 size 个元素的内存空间

#### channel make

channel := make(chan struct {}[,buff_size])
buff_size 用于控制创建的 channel 缓存的大小,如果省略或者为零的话这就是无缓存 chan,否则就是有缓存 chan

###
channel := make(chan struct{},)

## map

## 引用类型

map 是引用类型

```go
	graph := make(map[string]map[string]bool)
	AddEdge := func(from, to string) {
		var edges map[string]bool = graph[from]
		if edges == nil {
			edges = make(map[string]bool)
		}
		graph[from] = edges
		edges[to] = true
	}
//
```

上面的代码不是很优雅,因为如果`edges`不等于`nil`的话,graph[from]就是该结构的引用,就没必要再次调用`graph[from] = edges`
可以修改成如下代码,这里只有当`edges == nil`才会使用`make`创建一个新的 map,然后将其放入 graph 中,`edges !=nil`的时候直接填入职.
```go
	AddEdge := func(from, to string) {
		edges := graph[from]
		if edges == nil {
			edges = make(map[string]bool)
			graph[from] = edges
		}
		edges[to] = true
	}
```

### range delete

python 中的 dict 在在遍历的时候不能执行删除操作否则会触发异常

```py
a = {"foo": 1, "bar": 2}

print(f'before:{a}')

for k, v in a.items():
    del a[k]

for k in a.keys():
    del a[k]


for k in list(a.keys()):
    del a[k]


print(f'after:{a}')

```

但是 golang 就没有这个限制