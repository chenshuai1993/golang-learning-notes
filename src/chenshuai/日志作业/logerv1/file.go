package logerv1

import (
	"fmt"
	"os"
	"path"
)

//往文件里面写日志

// FileLogger 文件日志的结构体
type FileLogger struct {
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
}

//文件日志结构体的构造函数
func NewFileLogger(fileName, filePath string) *FileLogger {

	fl := FileLogger{
		fileName: fileName,
		filePath: filePath,
	}
	fl.initFile()
	return fl
}

//将制定的日志文件打开，并赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)

	//打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件 %s 失败, %v", logName, err))
	}

	f.file = fileObj

	//打开错误日志文件
	//errLogName =
	/*fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件 %s 失败, %v", logName, err))
	}*/
}

//func
func (f *FileLogger) Debug(format string, args ...interface{}) {
	fmt.Printf("%s")
}
