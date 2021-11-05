package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"rename/conf"
	"rename/tools"
)

func main() {
	var c conf.Conf
	log.Info(c)
	//配置文件路径
	ymlPathUrl := ""
	if len(os.Args) != 0 {
		ymlPathUrl = os.Args[1]
	}
	c.GetConf(ymlPathUrl)

	//替换文
	tools.ResetName(c.Path, c.FileExtensions, c.PrefixName)
}
