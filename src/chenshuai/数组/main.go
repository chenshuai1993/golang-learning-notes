package main

func main() {

	//
	array()
}

func array() {

	const (
		name1 int = iota
		name2
		name3
	)
	var names [3]string = [3]string{name1: "hello", name2: "world", name3: "!"}
	var name [3]string = [3]string{name1: "hello", name2: "world", name3: "!"}

	println(names == name)

	for i, v := range names {
		println(i, v)
	}
}
