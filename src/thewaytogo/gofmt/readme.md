### go fmt
大多数 Go 程序遵循相同的格式化规则，换句话说，一个 tab 键用于缩进，左括号和他们的声明语句在同一行。
我知道，你可能有自己的风格，并且想坚持它。这也是我长期以来所做的事情，但我很高兴我最终放弃了。一个大原因是 go fmt 命令。它易于使用而且具有权威性（所以就没有人争论无意义的偏好）。
当你在一个项目内的时候，你可以运用格式化规则到这个项目及其所有子目录：

```go
go fmt ./...
```


> 试一试，它不仅缩进你的代码，也对齐了声明的字段和按字母书序导入。