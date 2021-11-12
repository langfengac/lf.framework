package nlog

import (
	"fmt"
	"github.com/langfengac/lf.framework/appconfig"
	"github.com/langfengac/lf.framework/framework"
	"time"
)

func NewInfoNLog(dirPath string, isPrint bool) *NLog {
	return NewNLogForHours(dirPath, "info.txt", "【INFO】", isPrint)
}

func (log NLog) Info(text ...interface{}) {
	str := fmt.Sprintln(text...)

	f := log.OsFile
	defer f.Close()

	t := time.Now().Format(lf.TimeF) + " "
	c := fmt.Sprintf("%s %s %s \r\n", log.FirstTag, t, str)
	f.WriteString(c)

	if log.IsPrint {
		fmt.Print(c)
	}
}

func Info(text ...interface{}) {
	c := appconfig.NewInitAppConfig()
	dirPath := c.ReadString("nlog", "path", "")

	log := NewInfoNLog(dirPath, true)
	log.Info(text...)
}
func InfoDir(dir string, text ...interface{}) {
	c := appconfig.NewInitAppConfig()
	dirPath := c.ReadString("nlog", "path", "")
	dirPath += "/" + dir

	log := NewInfoNLog(dirPath, true)
	log.Info(text...)
}
