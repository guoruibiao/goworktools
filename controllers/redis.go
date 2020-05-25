package controllers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/goworktools/handlers"
	"github.com/guoruibiao/goworktools/library"
	libredis "github.com/guoruibiao/goworktools/library/redis"
	"net/http"
	"strconv"
	"strings"
)

var (
	stringHandler *handlers.StringHandler
	listHandler   *handlers.ListHandler
	hashHandler   *handlers.HashHandler
	setHandler    *handlers.SetHandler
	zsetHandler   *handlers.ZsetHandler
	redisConn     redis.Conn
)

type WebResponse struct {
	Data interface{} `json:"data"`
	Err  error       `json:"err"`
}

func init() {
	redisConn, err := libredis.GetRedisConn()
	if err != nil {
		panic(err)
		return
	}
	stringHandler, _ = handlers.NewStringHandler(redisConn)
	listHandler, _ = handlers.NewListHandler(redisConn)
	hashHandler, _ = handlers.NewHashHandler(redisConn)
	setHandler, _ = handlers.NewSetHandler(redisConn)
	zsetHandler, _ = handlers.NewZsetHandler(redisConn)
}

func prepareConnection(cmd, redisAddress string) (err error){
	conn, err := libredis.GetRedisConnByAddress(redisAddress)
	if err != nil {
		return
	}
	switch cmd[0] {
	case 's':
		// 区分来自 set 或者 string
		if stringHandler == nil {
			stringHandler, _ = handlers.NewStringHandler(conn)
		}
		if setHandler == nil {
			setHandler, _ = handlers.NewSetHandler(conn)
		}
	case 'l':
		if listHandler == nil {
			listHandler, _ = handlers.NewListHandler(conn)
		}
	case 'h':
		if hashHandler == nil {
			hashHandler, _ = handlers.NewHashHandler(conn)
		}
	case 'z':
		if zsetHandler == nil {
			zsetHandler, _ = handlers.NewZsetHandler(conn)
		}
	}
	return nil
}

func prepare(key, address string) (err error) {
	if address == "" {
		err = fmt.Errorf("No redis address specified ")
		return
	}

	err = prepareConnection(key, address)
	if err != nil {
		err = fmt.Errorf("prepare for redis connection failed with %s", err.Error())
		return
	}
	return nil
}


func Strlen(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	length, err := redis.Int64(stringHandler.Hand("STRLEN", libredis.DoCommandImpl, params...))
	resp.Data = length
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Get(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	ret, err := redis.String(stringHandler.Hand("GET", libredis.DoCommandImpl, params...))
	resp.Data = ret
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Llen(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	length, err := redis.Int64(listHandler.Hand("LLEN", libredis.DoCommandImpl, params...))
	resp.Data = length
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Lrange(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key, 0, -1}
	list, err := redis.Strings(listHandler.Hand("LRANGE", libredis.DoCommandImpl, params...))
	resp.Data = list
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Scard(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	length, err := redis.Int(setHandler.Hand("SCARD", libredis.DoCommandImpl, params...))
	resp.Data = length
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Srandmember(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	number, err := strconv.Atoi(ctx.DefaultQuery("number", "1"))
	if err != nil {
		number = 1
	}
	params := []interface{}{key, number}
	list, err := redis.Strings(setHandler.Hand("SRANDMEMBER", libredis.DoCommandImpl, params...))
	resp.Data = list
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Smembers(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	list, err := redis.Strings(setHandler.Hand("SMEMBERS", libredis.DoCommandImpl, params...))
	resp.Data = list
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Hlen(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	length, err := redis.Int(hashHandler.Hand("HLEN", libredis.DoCommandImpl, params...))
	resp.Data = length
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Hget(ctx *gin.Context) {
	key := ctx.Query("key")
	member := ctx.DefaultQuery("member", "")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key, member}
	value, err := redis.String(hashHandler.Hand("HGET", libredis.DoCommandImpl, params...))
	resp.Data = value
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Hmget(ctx *gin.Context) {
	key := ctx.Query("key")
	members := strings.Split(ctx.DefaultQuery("members", ""), ",")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	for _, member := range members {
		params = append(params, member)
	}
	list, err := redis.Strings(hashHandler.Hand("HMGET", libredis.DoCommandImpl, params...))
	mapper := make(map[string]string, len(members))
	for index, item := range list {
		mapper[members[index]] = item
	}
	resp.Data = mapper
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Hgetall(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	list, err := redis.Strings(hashHandler.Hand("HGETALL", libredis.DoCommandImpl, params...))
	resp.Data = list
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Zcard(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key}
	length, err := redis.Int(zsetHandler.Hand("ZCARD", libredis.DoCommandImpl, params...))
	resp.Data = length
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Zscore(ctx *gin.Context) {
	key := ctx.Query("key")
	member := ctx.DefaultQuery("member", "")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key, member}
	score, err := redis.Int(zsetHandler.Hand("ZSCORE", libredis.DoCommandImpl, params...))
	resp.Data = score
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Zrange(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key, 0, -1, "WITHSCORES"}
	values, err := redis.Values(zsetHandler.Hand("ZRANGE", libredis.DoCommandImpl, params...))

	jsondata := "{"
	for index := 0; index < len(values); index += 2 {
		k := string(values[index].([]byte))
		v, _ := strconv.ParseFloat(string(values[index+1].([]byte)), 10)
		jsondata += fmt.Sprintf("\"%s\": %f", k, v) + ","
	}
	jsondata = strings.TrimRight(jsondata, ",")
	jsondata += "}"

	/*
		var data map[string]float64
		json.Unmarshal([]byte(jsondata), &data)
	*/
	resp.Data = jsondata
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func Zrevrange(ctx *gin.Context) {
	key := ctx.Query("key")
	address := ctx.DefaultQuery("address", "")
	resp := &WebResponse{
		Data: nil,
		Err:  nil,
	}

	if err := prepare(key, address); err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	params := []interface{}{key, 0, -1, "WITHSCORES"}
	values, err := redis.Values(zsetHandler.Hand("ZREVRANGE", libredis.DoCommandImpl, params...))
	/*
		mapper := make(map[string]float64)
		for index:=0; index<len(values)-1; index+=2 {
			mapper[string(values[index].([]byte))], _ = strconv.ParseFloat(string(values[index+1].([]byte)), 10)
		}
		fmt.Println(mapper)
	*/
	jsondata := "{"
	for index := 0; index < len(values); index += 2 {
		k := string(values[index].([]byte))
		v, _ := strconv.ParseFloat(string(values[index+1].([]byte)), 10)
		jsondata += fmt.Sprintf("\"%s\": %f", k, v) + ","
	}
	jsondata = strings.TrimRight(jsondata, ",")
	jsondata += "}"

	resp.Data = jsondata
	resp.Err = err
	ctx.JSON(http.StatusOK, resp)
}

func BNSList(ctx *gin.Context) {
	bnsList := []string{
		"localhost",
		// todo 写到配置文件中！
		"group.bdrp-bdrp-common-test-v308-2-proxy.redis.all",
		"下面是线上，小心选择,本项不可选择",
		"group.bdrp-rmb-haokan-proxy.MAP.all",
	}
	resp := &WebResponse{
		Data: bnsList,
		Err:  nil,
	}
	ctx.JSON(http.StatusOK, resp)
}

func GetInstanceByBNS(ctx *gin.Context) {
	bns := ctx.DefaultQuery("bns", "")
	resp := &WebResponse{
		Data: nil,
		Err:  fmt.Errorf("No bns specified "),
	}
	if bns == "" {
		ctx.JSON(http.StatusOK, resp)
		return
	}

	bnsList, err := library.GetInstanceByBNS(bns)
	if err != nil {
		resp.Err = err
		ctx.JSON(http.StatusOK, resp)
		return
	}

	resp.Data = bnsList
	ctx.JSON(http.StatusOK, resp)
}