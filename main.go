package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"rename/conf"
	"rename/tools"
)

func main() {
	var c conf.Conf
	//配置文件路径
	ymlPathUrl := ""
	if len(os.Args) >= 2 {
		ymlPathUrl = os.Args[1]
	}
	c.GetConf(ymlPathUrl)
	//打印配置文件内容
	log.Info(c)

	//替换文文件名
	tools.ResetName(c.Path, c.FileExtensions, c.PrefixName)
}
