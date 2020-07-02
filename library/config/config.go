package config

import (
	"activities/common"

	"github.com/spf13/viper"
)

// config 配置
var config *Config

// Config 配置文件结构
type Config struct {
	Server  ServerConfig  `json:"server"`
	Storage StorageConfig `json:"storage"`
	Nats    NatsConfig    `json:"nats"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

// StorageConfig 存储配置
type StorageConfig struct {
	SQL struct {
		DBs map[string]string `json:"dbs"`
	} `json:"sql"`
	Mgo struct {
		URI      string `json:"uri"`
		DataBase string `json:"database"`
	} `json:"mgo"`
	Rds struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
	} `json:"rds"`
}

// NatsConfig 消息中间件集群配置
type NatsConfig struct {
	Client  string `json:"client"`  // 客户ID
	Cluster string `json:"cluster"` // 集群ID
	URLs    string `json:"urls"`    // 集群地址
}

func init() {
	config = new(Config)

	vpr := viper.New()
	vpr.SetConfigName(common.ConfigFileName.String())
	vpr.SetConfigType(common.ConfigFileType.String())
	vpr.AddConfigPath(common.ConfigFilePath.String())
	if err := vpr.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := vpr.Unmarshal(&config); err != nil {
		panic(err)
	}
}

// GetServerConf 获取服务器配置
func GetServerConf() ServerConfig {
	return config.Server
}

// GetStorageConf 获取存储配置
func GetStorageConf() StorageConfig {
	return config.Storage
}

// GetNatsConf 获取Nats配置
func GetNatsConf() NatsConfig {
	return config.Nats
}
