package handlers

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
	"strings"
)

type StringHandler struct {
	commandSlots map[string]int
	redisClient  redis.Conn
}

// 支持的命令列表
var stringSupportCommands map[string]int = map[string]int{
	"GET":    1,
	"STRLEN": 1,
	//"SET": 1,
}

// 第一版只支持一个 key 的 读操作处理吧
func NewStringHandler(conn redis.Conn) (*StringHandler, error) {
	sh := &StringHandler{
		commandSlots: stringSupportCommands,
		redisClient:  conn,
	}
	return sh, nil
}

func (sh *StringHandler) Hand(command string, doCommand libredis.DoCommand, params ...interface{}) (ret interface{}, err error) {
	command = strings.ToUpper(command)
	if _, hit := sh.commandSlots[command]; !hit {
		return nil, errors.New(fmt.Sprintf("unsupport command for %s", command))
	}

	return doCommand(sh.redisClient, command, params...)
}

func (sh *StringHandler) GetStringSupportCommands() (commands map[string]int, err error) {

	return sh.commandSlots, nil
}
