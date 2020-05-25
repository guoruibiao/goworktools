package handlers

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
	"strings"
)

type ListHandler struct {
	commandSlots map[string]int
	redisClient  redis.Conn
}

// 支持的命令列表
var listSupportCommands map[string]int = map[string]int{
	"LLEN":   1,
	"LRANGE": 1,
}

// 第一版只支持一个 key 的 读操作处理吧
func NewListHandler(conn redis.Conn) (*ListHandler, error) {
	lh := &ListHandler{
		commandSlots: listSupportCommands,
		redisClient:  conn,
	}
	return lh, nil
}

func (lh *ListHandler) Hand(command string, doCommand libredis.DoCommand, params ...interface{}) (ret interface{}, err error) {
	command = strings.ToUpper(command)
	if _, hit := lh.commandSlots[command]; !hit {
		return nil, errors.New(fmt.Sprintf("unsupport command for %s", command))
	}

	return doCommand(lh.redisClient, command, params...)
}

func (lh *ListHandler) GetSupportCommands() (commands map[string]int, err error) {

	return lh.commandSlots, nil
}
