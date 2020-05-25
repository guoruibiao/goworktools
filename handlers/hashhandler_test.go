package handlers

import (
	"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
)

func TestHashHandler_GetSupportCommands(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewHashHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	commands, err := hh.GetSupportCommands()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(commands)
}

func TestHashHandler_Hand(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewHashHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"hash"}
	ret, err := hh.Hand("hlen", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret.(int64))
}

func TestHashHandler_Hand2(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewHashHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"hash", "name"}
	ret, err := hh.Hand("hget", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(ret.([]uint8)))
}

func TestHashHandler_Hand3(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewHashHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"hash", "name", "age"}
	ret, err := hh.Hand("hmget", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range ret.([]interface{}) {
		t.Log(string(item.([]byte)))
	}
}

func TestHashHandler_Hand4(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewHashHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"hash"}
	ret, err := hh.Hand("hgetall", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range ret.([]interface{}) {
		t.Log(string(item.([]byte)))
	}
}

func TestNewHashHandler(t *testing.T) {
	type args struct {
		conn redis.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    *HashHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHashHandler(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHashHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
