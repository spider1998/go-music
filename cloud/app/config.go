package app

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug bool `json:"debug" default:"true"`
	//数据库配置
	Mysql string `json:"mysql" default:"root:123456@tcp(localhost:3306)/cloud?charset=utf8mb4&parseTime=true"` // mysql DSN
	//日志配置
	ConfPath string `json:"conf_path" default:"."` //日志文件路径
}

//加载配置
func LoadConfig() (Config, error) {
	godotenv.Load()
	var config Config
	err := envconfig.Process("", &config)
	return config, err
}
