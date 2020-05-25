package handlers

import (
	"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
)

func TestZsetHandler_GetSupportCommands(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewZsetHandler(conn)
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

func TestZsetHandler_Hand(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewZsetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"zset"}
	ret, err := redis.Int(hh.Hand("zcard", libredis.DoCommandImpl, params...))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret)
}

func TestZsetHandler_Hand2(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewZsetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"zset", "a"}
	ret, err := redis.Float64(hh.Hand("zscore", libredis.DoCommandImpl, params...))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret)
}

func TestZsetHandler_Hand3(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewZsetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"zset", 0, -1, "withscores"}
	ret, err := redis.StringMap(hh.Hand("zrange", libredis.DoCommandImpl, params...))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret)
}

func TestZsetHandler_Hand4(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	hh, err := NewZsetHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"zset", 0, -1, "withscores"}
	ret, err := redis.StringMap(hh.Hand("zrevrange", libredis.DoCommandImpl, params...))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ret)
}

func TestNewZsetHandler(t *testing.T) {
	type args struct {
		conn redis.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    *ZsetHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewZsetHandler(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewZsetHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewZsetHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
