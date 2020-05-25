package redis

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

//
type RedisManager struct {
	pool []redis.Conn
}

var (
	host string = "localhost"
	port int    = 6379
)

func Init() {

}

func GetRedisConn() (conn redis.Conn, err error) {
	address := host + ":" + strconv.Itoa(port)
	return redis.Dial("tcp", address)
}

func GetRedisConnByAddress(address string) (conn redis.Conn, err error) {
	return redis.Dial("tcp", address)
}

func ReturnRedisConn(conn redis.Conn) (err error) {

	return
}

func ReleaseRedisConn(conn redis.Conn) {
	if conn != nil{
		conn.Close()
	}
}

// 执行器
type DoCommand func(conn redis.Conn, commandname string, params ...interface{}) (result interface{}, err error)

// 执行器实现
func DoCommandImpl(conn redis.Conn, commandname string, params ...interface{}) (result interface{}, err error) {
	return conn.Do(commandname, params...)
}
