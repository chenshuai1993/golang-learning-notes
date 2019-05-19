package main

import (
	"fmt"
	"time"
)

func main() {

	//time1()

	//time2()

	//当前时间
	//timenow()

	//当前时间戳
	//timestampDemo()

	//时间戳 转 时间
	//timestamp2time(time.Now().Unix())

	//定时器
	//tick()

	//时间差
	//sub()

	//时间差1
	sub1()

	//十分钟之后
	//tenmin()

}

func time1() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分
	fmt.Println(now.Format("2006-01-02 15:04.000"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

//时间戳2
func time2() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

//当前时间
func timenow() {
	now := time.Now()
	//fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//时间戳
func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

}

//时间戳 to 时间
func timestamp2time(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	//fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

//定时器
func tick() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println(i) //每秒都会执行的任务
	}
}

//时间差
func sub() {
	temp := 0
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		for j := 0; j < 10000; j++ {
			temp++
		}
	}
	fmt.Println(temp)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

}

//程序毫秒时间差
func sub1() {
	s := time.Now().UnixNano() / 1000

	for i := 1; i < 10000000; i++ {
	}

	d := time.Now().UnixNano() / 1000

	fmt.Println("毫秒时间差:", int(d-s), "ms")
}

//十分钟之后
func tenmin() {
	now := time.Now()
	m, _ := time.ParseDuration("10m")
	m1 := now.Add(m)
	fmt.Println(m1)

}
