package nlog

import (
	"github.com/langfengac/lf.framework/appconfig"
	"os"
)

func NewGinLogFile() *os.File {
	c := appconfig.NewInitAppConfig()
	dirPath := c.ReadString("nlog", "path", "")
	log := NewNLogForHours(dirPath, "gin.txt", "", false)
	return log.OsFile
}
