Go work tools.

自用工具合集。


## redis
```shell
➜  goworktools git:(master) go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /templates/*filepath      --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] HEAD   /templates/*filepath      --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
[GIN-debug] GET    /redis/strlen             --> github.com/guoruibiao/goworktools/controllers.Strlen (3 handlers)
[GIN-debug] GET    /redis/get                --> github.com/guoruibiao/goworktools/controllers.Get (3 handlers)
[GIN-debug] GET    /redis/llen               --> github.com/guoruibiao/goworktools/controllers.Llen (3 handlers)
[GIN-debug] GET    /redis/lrange             --> github.com/guoruibiao/goworktools/controllers.Lrange (3 handlers)
[GIN-debug] GET    /redis/scard              --> github.com/guoruibiao/goworktools/controllers.Scard (3 handlers)
[GIN-debug] GET    /redis/srandmember        --> github.com/guoruibiao/goworktools/controllers.Srandmember (3 handlers)
[GIN-debug] GET    /redis/smembers           --> github.com/guoruibiao/goworktools/controllers.Smembers (3 handlers)
[GIN-debug] GET    /redis/hlen               --> github.com/guoruibiao/goworktools/controllers.Hlen (3 handlers)
[GIN-debug] GET    /redis/hget               --> github.com/guoruibiao/goworktools/controllers.Hget (3 handlers)
[GIN-debug] GET    /redis/hmget              --> github.com/guoruibiao/goworktools/controllers.Hmget (3 handlers)
[GIN-debug] GET    /redis/hgetall            --> github.com/guoruibiao/goworktools/controllers.Hgetall (3 handlers)
[GIN-debug] GET    /redis/zcard              --> github.com/guoruibiao/goworktools/controllers.Zcard (3 handlers)
[GIN-debug] GET    /redis/zscore             --> github.com/guoruibiao/goworktools/controllers.Zscore (3 handlers)
[GIN-debug] GET    /redis/zrange             --> github.com/guoruibiao/goworktools/controllers.Zrange (3 handlers)
[GIN-debug] GET    /redis/zrevrange          --> github.com/guoruibiao/goworktools/controllers.Zrevrange (3 handlers)
[GIN-debug] GET    /redis/bnslist            --> github.com/guoruibiao/goworktools/controllers.BNSList (3 handlers)
[GIN-debug] GET    /redis/getinstances       --> github.com/guoruibiao/goworktools/controllers.GetInstanceByBNS (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

>>>>>>> 0df1d13652174eeaa8004955a8b1cc0076c94e7b
