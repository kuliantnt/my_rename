package tools

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//ResetName 重设文件名
func ResetName(folder string, fileExtensions string, prefixName string) {
	log.Infof("后缀名: %s, 文件名： %s, 文件夹目录: %s", fileExtensions, prefixName, folder)
	var cstZone = time.FixedZone("CST", 8*3600)
	files, _ := ioutil.ReadDir(folder)
	var count int
	for _, file := range files {
		//首先不是目录，且后缀名为fileExtensions
		var flag = !file.IsDir() && strings.HasSuffix(file.Name(), fileExtensions)
		//改文件名
		if flag {
			fileName := file.Name()
			editTime := file.ModTime().In(cstZone)
			editTimeStr := fmt.Sprintf("%04d%02d%02d_%02d%02d",
				editTime.Year(), editTime.Month(), editTime.Day(), editTime.Hour(), editTime.Month())
			newFileName := prefixName + editTimeStr
			log.Infof("reName %s -> %s", fileName, newFileName)

			err := os.Rename(folder+fileName, folder+newFileName)
			if err != nil {
				fmt.Println("reName Error", err)
				continue
			} else {
				count++
			}
		}
	}
	log.Infof("更改了%d个文件", count)
}
