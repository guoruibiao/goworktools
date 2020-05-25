package handlers

import (
	"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
)

func TestSetHandler_GetSupportCommands(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewSetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	commands, err := sh.GetSupportCommands()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(commands)
}

func TestSetHandler_Hand(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewSetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"set"}
	ret, err := sh.Hand("scard", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret)
}

func TestSetHandler_Hand2(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewSetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"set"}
	ret, err := sh.Hand("srandmember", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(ret.([]byte)))
}

func TestSetHandler_Hand3(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewSetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"set", 3}
	ret, err := sh.Hand("srandmember", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range ret.([]interface{}) {
		t.Log(string(item.([]byte)))
	}
}

func TestSetHandler_Hand4(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewSetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"set"}
	ret, err := sh.Hand("smembers", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}
	for _, item := range ret.([]interface{}) {
		t.Log(string(item.([]byte)))
	}
}

func TestNewSetHandler(t *testing.T) {
	type args struct {
		conn redis.Conn
	}
	conn, _ := libredis.GetRedisConn()
	sh, _ := NewSetHandler(conn)
	tests := []struct {
		name    string
		args    args
		want    *SetHandler
		wantErr bool
	}{
		{
			"GET",
			args{conn},
			sh,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSetHandler(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSetHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
