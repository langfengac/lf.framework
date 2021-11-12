package nlog

import (
	"fmt"
	"os"
	"time"
)

type NLog struct {
	DirPath  string
	FileName string
	FirstTag string
	OsFile   *os.File
	IsPrint  bool
}

func NewNLog(dirPath, fileName, firstTag string, isPrint bool) *NLog {
	log := new(NLog)
	log.DirPath = dirPath
	log.FileName = fileName
	log.FirstTag = firstTag
	log.IsPrint = isPrint

	//创建日志目录
	dirName := dirPath + "/" + time.Now().Format("2006/01/02")
	os.MkdirAll(dirName, os.ModePerm)

	filename := dirName + "/" + log.FileName
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.OsFile = f
	return log
}

func NewNLogForHours(dirPath, fileName, firstTag string, isPrint bool) *NLog {
	log := new(NLog)
	log.DirPath = dirPath
	log.FileName = fmt.Sprintf("%v_%v", time.Now().Hour(), fileName)
	log.FirstTag = firstTag
	log.IsPrint = isPrint

	//创建日志目录
	dirName := dirPath + "/" + time.Now().Format("2006/01/02")
	os.MkdirAll(dirName, os.ModePerm)

	filename := dirName + "/" + log.FileName
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.OsFile = f
	return log
}
