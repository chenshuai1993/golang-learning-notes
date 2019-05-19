package main

import "oldboy/day5/日志作业/myloger"

var logger myloger.Logger

func main() {
	logger = myloger.NewLogger("info", "./", "xxx.log")

	//关闭时候close文件
	defer logger.Close()

	for {
		sb := "hello world"
		logger.Debug("%s 是一个好人", sb)
		logger.Info("%s 是一个好人", sb)
		logger.Fatal("%s 我是一个好人", sb)
	}

}
