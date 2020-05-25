package handlers

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
	"strings"
)

type HashHandler struct {
	commandSlots map[string]int
	redisClient  redis.Conn
}

// 支持的命令列表
var hashSupportCommands map[string]int = map[string]int{
	"HLEN":    1,
	"HGET":    1,
	"HMGET":   1,
	"HGETALL": 1,
}

func NewHashHandler(conn redis.Conn) (*HashHandler, error) {
	sh := &HashHandler{
		commandSlots: hashSupportCommands,
		redisClient:  conn,
	}
	return sh, nil
}

func (hh *HashHandler) Hand(command string, doCommand libredis.DoCommand, params ...interface{}) (ret interface{}, err error) {
	command = strings.ToUpper(command)
	if _, hit := hh.commandSlots[command]; !hit {
		return nil, errors.New(fmt.Sprintf("unsupport command for %s", command))
	}

	return doCommand(hh.redisClient, command, params...)
}

func (hh *HashHandler) GetSupportCommands() (commands map[string]int, err error) {

	return hh.commandSlots, nil
}
