# error wrap

## new feature

新特性相关函数有：

1. fmt.Errorf
2. errors.Unwrap
3. errors.Is
4. errors.As

### fmt.Errorf

fmt.Errorf 提供了新的动词 `w`，使用该动词，并传入一个 `error`，则返回一个包含该`error`的新的`error`.
[演示代码](./unwrap_test.go#L12)
