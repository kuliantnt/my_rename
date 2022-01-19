package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

//ResetName 重设文件名
func ResetName(folder string, fileExtensions string, prefixName string) {
	log.Infof("后缀名: %s, 文件名: %s, 文件夹目录: %s", fileExtensions, prefixName, folder)
	var cstZone = time.FixedZone("CST", 8*3600)
	files, _ := ioutil.ReadDir(folder)
	var count int
	for _, file := range files {
		//判断不是目录
		dirFlag := !file.IsDir()
		//是否改名, 如果改名了就不执行
		matchFlag, _ := regexp.MatchString(prefixName+"\\d{8}_\\d{4}"+fileExtensions, file.Name())
		matchFlag = !matchFlag
		//是否为prefix开头，后缀名为fileExtensions
		doFlag := strings.HasPrefix(file.Name(), prefixName) && strings.HasSuffix(file.Name(), fileExtensions)
		//如果为否，则不改文件名
		flag := dirFlag && matchFlag && doFlag
		//改文件名
		if flag {
			fileName := file.Name()
			editTime := file.ModTime().In(cstZone)
			editTimeStr := fmt.Sprintf("%04d%02d%02d_%02d%02d",
				editTime.Year(), editTime.Month(), editTime.Day(), editTime.Hour(), editTime.Month())
			newFileName := prefixName + editTimeStr + fileExtensions
			log.Infof("ReName %s -> %s", fileName, newFileName)

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
