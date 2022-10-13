package config

import (
	"attendance_service/src/entities"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var ConfPath = flag.String("c", "./conf/local.yaml", "配置文件")

var conf entities.ConfigStruct

//初始化命令行参数以及配置文件
func init() {
	flag.Parse()
	f, err := os.Open(*ConfPath)
	if err != nil {
		panic(fmt.Errorf("open config file:%q failed,%v", *ConfPath, err))
	}
	confContent, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("read config file:%q failed,%v", *ConfPath, err))
	}
	err = yaml.Unmarshal(confContent, &conf)
	if err != nil {
		panic(fmt.Errorf("parse config content:%q failed,%v", string(confContent), err))
	}
}

func GetConfig() entities.ConfigStruct {
	return conf
}
