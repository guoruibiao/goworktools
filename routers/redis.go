package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/goworktools/controllers"
)

func RedisRouterInit(router *gin.Engine) {
	redisRouter := router.Group("/redis")
	{
		// string
		redisRouter.GET("/strlen", controllers.Strlen)
		redisRouter.GET("/get", controllers.Get)

		// list
		redisRouter.GET("/llen", controllers.Llen)
		redisRouter.GET("/lrange", controllers.Lrange)

		// set
		redisRouter.GET("/scard", controllers.Scard)
		redisRouter.GET("/srandmember", controllers.Srandmember)
		redisRouter.GET("/smembers", controllers.Smembers)

		// hash
		redisRouter.GET("/hlen", controllers.Hlen)
		redisRouter.GET("/hget", controllers.Hget)
		redisRouter.GET("/hmget", controllers.Hmget)
		redisRouter.GET("/hgetall", controllers.Hgetall)

		// zset
		redisRouter.GET("/zcard", controllers.Zcard)
		redisRouter.GET("/zscore", controllers.Zscore)
		redisRouter.GET("/zrange", controllers.Zrange)
		redisRouter.GET("/zrevrange", controllers.Zrevrange)

		// BNS List
		redisRouter.GET("/bnslist", controllers.BNSList)
		redisRouter.GET("/getinstances", controllers.GetInstanceByBNS)
	}
}
