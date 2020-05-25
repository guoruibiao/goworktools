package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/goworktools/models"
	"github.com/guoruibiao/goworktools/routers"
	"log"
)

func main() {
	router := gin.Default()
	// 读取配置文件到内存
	var globalConfig models.App
	if _, err := toml.DecodeFile("./conf/app.toml", &globalConfig); err != nil {
		log.Fatal(err)
	}

	// 静态文件路由设置
	router.Static("/static/", "./web/static")
	router.Static("/templates/", "./web/templates")

	// 动态接口路由设置 !!! 注意路由冲突问题, 优先级：静态优先，动态兜底。
	routers.RedisRouterInit(router)

	// run
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
