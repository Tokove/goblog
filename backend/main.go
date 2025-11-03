package main

import "goblog/config"

func main() {
	config.InitConfig()
	config.InitDB()
	config.InitRedis()
}
