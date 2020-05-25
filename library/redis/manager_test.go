package redis

import "testing"

func TestGetRedisConn(t *testing.T) {
	conn, err := GetRedisConn()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(conn)
	}
}
