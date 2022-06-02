package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Mysql struct {
	Host string `json:"host"`
	Port int32 `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Swagger struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	BasePath string `json:"base_path"`
	Version string `json:"version"`
	Host string `json:"host"`
}
type ConfigData struct {
	DebugMode string `json:"debug_mode"`
	Addr string `json:"addr"`
	ReadTimeout uint `json:"read_timeout"`
	WriteTimeout uint `json:"write_timeout"`
	MaxHeaderBytes uint `json:"max_header_bytes"`
	Secret string `json:"secret"`
	Mysql Mysql `json:"mysql"`
	Swagger Swagger `json:"swagger"`
}



var Config  ConfigData
// ReadConfigFile 读取配置文件
func ReadConfigFile(jsonFile string) {
	jByte, err := ioutil.ReadFile(jsonFile)
	if err != nil  {
		log.Fatal("read config file fail:", err.Error())
	} else {
		err = json.Unmarshal(jByte,&Config)
		if err != nil {
			log.Fatal("json unmarshal fail:",err.Error())
		}
	}
}
