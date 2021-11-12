package nlog

import (
	"fmt"
	"github.com/langfengac/lf.framework/appconfig"
	lf "github.com/langfengac/lf.framework/framework"
	"time"
)

func NewErrorNLog(dirPath string, isPrint bool) *NLog {
	return NewNLogForHours(dirPath, "error.txt", "【ERROR】", isPrint)
}

func (log NLog) Error(err error, remark string) {
	f := log.OsFile
	defer f.Close()

	t := time.Now().Format(lf.TimeF) + " "

	c := fmt.Sprintf("%s %s %s | %s \r\n", log.FirstTag, t, remark, err.Error())
	f.WriteString(c)

	if log.IsPrint {
		fmt.Print(c)
	}
}

func Error(err error, remark string) {
	c := appconfig.NewInitAppConfig()
	dirPath := c.ReadString("nlog", "path", "")

	log := NewErrorNLog(dirPath, true)
	log.Error(err, remark)
}
func ErrorDir(dir string, err error, remark string) {
	c := appconfig.NewInitAppConfig()
	dirPath := c.ReadString("nlog", "path", "")
	dirPath += "/" + dir

	log := NewErrorNLog(dirPath, true)
	log.Error(err, remark)
}
