package main44

import (
	"fmt"
	"io/ioutil"
	_ "encoding/json"
	yaml "gopkg.in/yaml.v2"
)

//profile variables
type conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
	Devs        []string       `yaml:"devs"`
	FilterTypes map[string]int `yaml:"filter_types"`
}
func main() {



}

func (c *conf)getConf() *conf  {

	yamlFile,err :=ioutil.ReadFile("config.yaml")
	 if err != nil {

	 }
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c

}
