package hack

import (
	_ "embed"
	"fmt"
	"gopkg.in/yaml.v3"
)

type Admin struct {
	Address string `yaml:"address"`
}

type Email struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	From     string `yaml:"from"`
	FromName string `yaml:"fromName"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	TimeZone string `yaml:"timezone"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Mongodb struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	DbName string `yaml:"dbname"`
}

type AliYun struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	EndPoint        string `yaml:"endPoint"`
}

type Config struct {
	Mysql   Mysql   `yaml:"mysql"`
	Email   Email   `yaml:"email"`
	Redis   Redis   `yaml:"redis"`
	Mongodb Mongodb `yaml:"mongodb"`
	AliYun  AliYun  `yaml:"aliyun"`
	Admin   Admin   `yaml:"admin"`
}

//go:embed config.yaml
var data []byte

var config *Config

func InitConfig() error {
	if data == nil {
		return fmt.Errorf("读取YAML配置文件失败: %w")
	}
	var newConfig Config
	if err := yaml.Unmarshal(data, &newConfig); err != nil {
		return fmt.Errorf("获取YAML配置文件失败: %w", err)
	}
	config = &newConfig
	return nil
}

func GetConfig() (*Config, error) {
	if config == nil {
		if err := InitConfig(); err != nil {
			return nil, err
		}
	}
	return config, nil
}
