package myloger

import (
	"path"
	"runtime"
)

//存放一些公用的函数
func getCallerInfo(skip int) (fileName string, line int, funcName string) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}

	//从 filename 中（全路径中）剥离文件名
	path.Base(fileName) // x/y/xx.log

	//根据pc拿到函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)

	return
}
