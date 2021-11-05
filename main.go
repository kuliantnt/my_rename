package main

import (
	log "github.com/sirupsen/logrus"
	"rename/conf"
	"rename/tools"
)


func main() {
	var c conf.Conf
	log.Info(c)
	c.GetConf()
	//替换文
	tools.ResetName(c.Path, c.FileExtensions, c.PrefixName)
}
