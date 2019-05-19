package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 解析日志库的配置文件

// Config 是一个日志配置项
type Config struct {
	filePath string `conf:"file_path"`
	fileName string `conf:"file_name"`
	maxSize  int64  `conf:"max_size"`
}

// 从conf文件中读取内容赋值给结构体指针
func parseConf(fileName string, c interface{}) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	// 1. 打开文件
	// 2. 读取内容
	// 3. 一行一行读取内容 根据tag找结构体里面的对应字段
	// 4. 找到了要赋值

	var c = &Config{} // 用来存储读取的数据
	filename := "src/oldboy/day5/反射/xxx.conf"
	//err := parseConf(filename, c)
	parseConf(filename, c)
	/*if err != nil {
		fmt.Println(err)
	}*/

	//fmt.Println(c) // 解析之后的数据

}
