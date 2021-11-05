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
	Path           string `yaml:"path"`
	FileExtensions string `yaml:"fileExtensions"`
	PrefixName     string `yaml:"prefixName"`
}

//GetConf 获取配置文件
func (conf *Conf) GetConf() *Conf {

	yamlFileURL :=  "conf.yaml"
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	yamlFileURL = path.Join(exPath, yamlFileURL)
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile(yamlFileURL)
	if err != nil {
		log.Error(err.Error())
	}
	//读取yaml文件
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Error(err.Error())
	}
	return conf
}
