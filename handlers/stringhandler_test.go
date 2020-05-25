package handlers

import (
	"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
)

func TestStringHandler_Hand(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewStringHandler(conn)
	if err != nil {
		t.Error(err)
	} else {
		params := []interface{}{"name", "Tiger"}
		ret, err := sh.Hand("set", libredis.DoCommandImpl, params...)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(ret)
		}
	}
}

func TestStringHandler_Hand2(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewStringHandler(conn)
	if err != nil {
		t.Error(err)
	} else {
		params := []interface{}{"age"}
		ret, err := redis.String(sh.Hand("GET", libredis.DoCommandImpl, params...))
		if err != nil {
			t.Error(err)
		} else {
			t.Log(ret)
		}
	}
}

func TestStringHandler_Hand3(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewStringHandler(conn)
	if err != nil {
		t.Error(err)
	} else {
		params := []interface{}{"name"}
		ret, err := redis.String(sh.Hand("SETBIT", libredis.DoCommandImpl, params...))
		if err != nil {
			t.Log(err)
		} else {
			t.Error(ret)
		}
	}
}

func TestStringHandler_Hand4(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	sh, err := NewStringHandler(conn)
	if err != nil {
		t.Error(err)
	} else {
		params := []interface{}{"name"}
		ret, err := redis.Int64(sh.Hand("strlen", libredis.DoCommandImpl, params...))
		if err != nil {
			t.Error(err)
		} else {
			t.Log(ret)
		}
	}
}

func TestNewStringHandler(t *testing.T) {
	type args struct {
		conn redis.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    *StringHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStringHandler(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStringHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringHandler_GetStringSupportCommands(t *testing.T) {
	type fields struct {
		commandSlots map[string]int
		redisClient  redis.Conn
	}
	tests := []struct {
		name         string
		fields       fields
		wantCommands map[string]int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &StringHandler{
				commandSlots: tt.fields.commandSlots,
				redisClient:  tt.fields.redisClient,
			}
			gotCommands, err := sh.GetStringSupportCommands()
			if (err != nil) != tt.wantErr {
				t.Errorf("StringHandler.GetStringSupportCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCommands, tt.wantCommands) {
				t.Errorf("StringHandler.GetStringSupportCommands() = %v, want %v", gotCommands, tt.wantCommands)
			}
		})
	}
}
