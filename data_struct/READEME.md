# data struct

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