package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var MyConfigInfo *MysqlConfig

type  MysqlConfig struct {
	App struct {
		Name string `yaml:"name"`
	}
	Mysql struct {
		Host string `yaml:"host"`
		Port int32 `yaml:"port"`
		DbName string `yaml:"dbname"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
	}
	Cache struct {
		Enable bool `yaml:"enable"`
		List []string `yaml:"list,flow"`
	}
}

func ReadConfigFromYaml(config string)  {
	myConf  := new(MysqlConfig)
	yamlFile, err := ioutil.ReadFile(config)
	log.Println("yamlFile:", string(yamlFile))
	if err != nil {
		log.Printf("yamlFile.Get err %v", err)
	}
	err = yaml.Unmarshal(yamlFile, myConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v when to struct", err)
	}
	log.Println("conf", myConf)
	log.Printf("mysqlHost:%s, port:%d, cacheList:%s", myConf.Mysql.Host, myConf.Mysql.Port, myConf.Cache.List)

	MyConfigInfo = myConf
	//configMap := make(map[string]interface{})
	//err = yaml.Unmarshal(yamlFile, configMap)

	if err != nil {
		log.Fatalf("Unmarshal: %v when to map", err)
	}

	fmt.Printf("---遍历结束---\n")

}

func init() {

	ReadConfigFromYaml("./config/mysql.yaml")
	fmt.Println("user:"+MyConfigInfo.Mysql.User)


}