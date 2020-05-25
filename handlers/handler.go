package handlers

import (
	libredis "github.com/guoruibiao/goworktools/library/redis"
)

type Handler interface {
	Hand(command string, doCommand libredis.DoCommand, params ...interface{}) (result interface{}, err error)
}
