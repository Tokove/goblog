package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}

	Redis struct {
		Addr     string
		DB       int
		Password string
	}
}

func InitConfig() {
	AppConfig = &Config{}

	// 设置读取文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 读取文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	// 解析读取文件
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	log.Println("Configure successfully")
}
