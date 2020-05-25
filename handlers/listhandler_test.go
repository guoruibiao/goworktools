package handlers

import (
	"reflect"
	"testing"

	"github.com/garyburd/redigo/redis"
	libredis "github.com/guoruibiao/goworktools/library/redis"
)

func TestListHandler_GetSupportCommands(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	lh, err := NewListHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	commands, err := lh.GetSupportCommands()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(commands)
}

func TestListHandler_Hand(t *testing.T) {
	conn, err := libredis.GetRedisConn()
	if err != nil {
		t.Error(err)
		return
	}

	lh, err := NewListHandler(conn)
	if err != nil {
		t.Error(err)
		return
	}

	params := []interface{}{"list", 0, -1}
	list, err := lh.Hand("lrange", libredis.DoCommandImpl, params...)
	if err != nil {
		t.Error(err)
		return
	}

	for item := range list.([]interface{}) {
		t.Log(item)
	}
}

func TestNewListHandler(t *testing.T) {
	type args struct {
		conn redis.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    *ListHandler
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewListHandler(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewListHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
