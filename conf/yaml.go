package conf

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//Conf 配置文件结构
type Conf struct {
	//目录
	Path string `yaml:"path"`
	//文件扩展名
	FileExtensions string `yaml:"fileExtensions"`
	//前缀名
	PrefixName string `yaml:"prefixName"`
}

//GetConf 获取配置文件
func (conf *Conf) GetConf(yamlFileURL string) *Conf {
	//获取当前执行文件目录
	ex, err := os.Executable()
	//如果不输入参数，读取当前目录下配置文件
	if yamlFileURL == "" {
		yamlFileURL = "conf.yaml"
		//从当前目录读取
		exPath := filepath.Dir(ex)
		yamlFileURL = path.Join(exPath, yamlFileURL)
		log.Debug("配置文件: %s", yamlFileURL)
	}
	log.Debugf("GetConf: 配置文件地址为: %s", yamlFileURL)
	if err != nil {
		log.Error("获取yaml配置文件出错: ", err.Error())
		panic(err)
	}
	yamlFile, err := ioutil.ReadFile(yamlFileURL)
	if err != nil {
		log.Error("读取文件异常: ", err.Error())
	}
	//读取yaml文件
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Error("读取yaml异常: ", err.Error())
	}
	return conf
}
