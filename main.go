package main

import (
	"flag"
	"rename/conf"
	"rename/tools"

	log "github.com/sirupsen/logrus"
)

var cliPath = flag.String("path", "/data/sh/conf.yml", "Input Your yaml path")

func main() {
	flag.Parse()
	var c conf.Conf
	//配置文件路径
	c.GetConf(*cliPath)
	//打印配置文件内容
	log.Debug(c)

	//替换文文件名
	tools.ResetName(c.Path, c.FileExtensions, c.PrefixName)
}
