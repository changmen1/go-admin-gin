//读取配置文件
//@author: zxl

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// config 总配置文件
type config struct {
	System system `yaml:"system"`
	Logger logger `yaml:"logger"`
	Mysql  mysql  `yaml:"mysql"`
	Redis  redis  `yaml:"redis"`
}

// 系统配置
type system struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

// logger 日志
type logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

// mysql配置
type mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxOpen  int    `yaml:"max_open"`
}

// redis配置
type redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

var Config *config

// init 初始化配置
func init() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return
	}
	yaml.Unmarshal(yamlFile, &Config)
}
