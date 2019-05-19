package myloger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志

//文件日志结构体

//FileLogger 文件日志的结构体  字段 字段类型
type FileLogger struct {
	level    Level
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
	maxSize  int64
}

//文件日志结构体的构造函数
func NewLogger(levelStr, fileName, filePath string) *FileLogger {
	level := parseLogLevel(levelStr)
	fl := &FileLogger{
		level:    level,
		fileName: fileName,
		filePath: filePath,
		maxSize:  10 * 1024 * 1024,
	}
	fl.initFile() //根据上面的文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应的字段

	return fl
}

//初始化 将制定的日志文件打开， 赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	//打开文件
	fileobj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败", logName, err))
	}
	f.file = fileobj

	//打开错误日志文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开错误日志文件%s失败", logName, err))
	}
	f.errFile = errFileObj
}

//方法debug
//func (f *FileLogger)Debug(format string, args ...interface{}) {
//	//fmt.Fprintf fmt.Eprintf fmt.Sprintf
//	//f.file.Write()
//	msg := fmt.Sprintf(format, args...)  //得到要记录的日志
//
//	nowStr := time.Now().Format("2006-01-02 15:04:05.000") //
//	fileName, line, funcName := getCallerInfo(2) //skip
//	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, funcName, "debug", msg)
//	//日志格式:[时间][文件:行号][函数名][日志级别][日志信息]
//	fmt.Fprintln(f.file, logMsg)  		//利用fmt包将msg字符串写入f.file
//}

//方法debug
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.Log(DebugLevel, format, args)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.Log(InfoLevel, format, args)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.Log(FatalLevel, format, args)
}

//日志
func (f *FileLogger) Log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}

	msg := fmt.Sprintf(format, args...)                    //得到要记录的日志
	nowStr := time.Now().Format("2006-01-02 15:04:05.000") //z
	fileName, line, funcName := getCallerInfo(2)           //skip
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s][%s]", nowStr, fileName, line, funcName, logLevelStr, msg)

	//切分日志
	if f.checkSplit(f.file) {
		f.file = f.splitLogFile(f.file)
	}

	//日志格式:[时间][文件:行号][函数名][日志级别][日志信息]
	fmt.Fprintln(f.file, logMsg) //利用fmt包将msg字符串写入f.file

	//如果是 error 或者是 fatal日志记录到 f.errfile
	if level >= ErrorLevel {
		if f.checkSplit(f.file) {
			f.file = f.splitLogFile(f.file)
		}
		fmt.Fprintln(f.errFile, logMsg) //利用fmt包将msg字符串写入f.file
	}
}

//封装一个切分日志文件的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	//切分文件
	fileName := file.Name()

	// 1. 原来的文件关闭
	file.Close()

	//2 备份原来的文件
	backname := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
	os.Rename(fileName, backname)

	//
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开错误日志文件失败"))
	}

	return fileObj
}

//检查是否要拆分
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileinfo, _ := file.Stat()
	return fileinfo.Size() >= f.maxSize
}

//close
func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}
