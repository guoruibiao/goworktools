package handlers

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
	"strings"
)

type ZsetHandler struct {
	commandSlots map[string]int
	redisClient  redis.Conn
}

// 支持的命令列表
var zsetSupportCommands map[string]int = map[string]int{
	"ZSCORE":    1,
	"ZRANGE":    1,
	"ZREVRANGE": 1,
	"ZCARD":     1,
}

func NewZsetHandler(conn redis.Conn) (*ZsetHandler, error) {
	sh := &ZsetHandler{
		commandSlots: zsetSupportCommands,
		redisClient:  conn,
	}
	return sh, nil
}

func (zh *ZsetHandler) Hand(command string, doCommand libredis.DoCommand, params ...interface{}) (ret interface{}, err error) {
	command = strings.ToUpper(command)
	if _, hit := zh.commandSlots[command]; !hit {
		return nil, errors.New(fmt.Sprintf("unsupport command for %s", command))
	}

	return doCommand(zh.redisClient, command, params...)
}

func (zh *ZsetHandler) GetSupportCommands() (commands map[string]int, err error) {

	return zh.commandSlots, nil
}
