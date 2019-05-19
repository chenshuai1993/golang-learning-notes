## 字符串和字节数组

字符串和字节数组是紧密相关的。我们可以轻松地在他们之间转换：

```go
stra := "the spice must flow"
byts := []byte(stra)
strb := string(byts)
```

实际上，这种转换方式在各种类型之间是通用的。一些函数显示地需要一个  `int32` 或者  `int64` 或者它们的无符号部分。你可能发现你必须这样做：

```go
int64(count)
```

然而，当它涉及到字节和字符串时，这可能是你经常做的事情。一定记着当你使用 `[]byte(X)` 或者 `string(X)` 时，你实际上创建了数据的副本。这是必要的，因为字符串是不可变的。

那些由 Unicode 码点 `runes` 构成的字符串，如果你获取字符串的长度，你可能不能得到你期望的。下面的结果是3：

    fmt.Println(len("椒"))

如果你用 `range` 迭代一个字符串，你将得到 runes，而不是字节。当然，当你将字符串转换为 `[]byte` 类型时，你将得到正确的数据。