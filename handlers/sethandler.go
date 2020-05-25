package handlers

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
	"strings"
)

type SetHandler struct {
	commandSlots map[string]int
	redisClient  redis.Conn
}

// 支持的命令列表
var setSupportCommands map[string]int = map[string]int{
	"SCARD":       1,
	"SRANDMEMBER": 1,
	"SMEMBERS":    1,
}

func NewSetHandler(conn redis.Conn) (*SetHandler, error) {
	sh := &SetHandler{
		commandSlots: setSupportCommands,
		redisClient:  conn,
	}
	return sh, nil
}

func (sh *SetHandler) Hand(command string, doCommand libredis.DoCommand, params ...interface{}) (ret interface{}, err error) {
	command = strings.ToUpper(command)
	if _, hit := sh.commandSlots[command]; !hit {
		return nil, errors.New(fmt.Sprintf("unsupport command for %s", command))
	}

	return doCommand(sh.redisClient, command, params...)
}

func (sh *SetHandler) GetSupportCommands() (commands map[string]int, err error) {

	return sh.commandSlots, nil
}
