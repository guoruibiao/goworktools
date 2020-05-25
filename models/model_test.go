package models

import (
	"github.com/BurntSushi/toml"
	"log"
	"testing"
)

func TestConfRead(t *testing.T) {
	var globalConfig App
	if _, err := toml.DecodeFile("../conf/app.toml", &globalConfig); err != nil {
		log.Fatal(err)
	}
	t.Log("server.port", globalConfig.Server.Port)
	t.Log("redis.test.host", globalConfig.RedisTest)
}