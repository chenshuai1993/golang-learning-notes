
## 带初始化的 if

Go 对 `if` 语句做了稍微修改，支持在条件语句被求值之前先进行初始化：

```go
if x := 10; count > x {
  ...
}
```

这是一个比较蠢的例子，更现实的是，你可能会像下面这样做：

```go
if err := process(); err != nil {
  return err
}
```

有意思的是，虽然 `err` 不能在 `if` 语句之外使用，但他可以在任何  `else if`  或者  `else` 之内使用。