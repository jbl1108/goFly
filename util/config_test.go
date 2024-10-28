package util

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {

}

func shutdown() {

}

func TestRedis(t *testing.T) {
	conf := NewConfig()
	redis := conf.RedisDBAddr()
	if redis == "" {
		t.FailNow()
	}
}
