# defer

## defer命令的拆解

理解这些坑的关键是这条语句：

```go
return xxx
```

上面这条语句经过编译之后，变成了三条指令：

```go
返回值 = xxx
调用`defer`函数
return 空
```

## 拆解例子

### 1

```go
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
```

拆解为：

```go
func f() (r int) {
	t := 5
	r = t
	func() {
		t = t + 5
	}()
	return
}
```

### 2

```go
func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
```

拆解为：

```go
func f() (r int) {
	r = 1
	func(r int) {
		r = r + 5
	}(r)
	return
}
```

### 3

```go
func modifyReturnValue(in int) (r int) {
	r = in
	defer func() {
		fmt.Println("r in defer bgein", r)
		r++
	}()
	return in * 2
}
```

拆解为：

```go
func modifyReturnValue(in int) (r int) {
	r = in
	r = in * 2
	func() {
		r++
	}()
	return
}
```
