package dataretriver

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestClient_UserState(t *testing.T) {
	// 创建客户端实例
	client := newClient(MainnetAPIURL, true)

	// 调用 UserState 方法
	result, err := client.UserState("0xfffff7d3762c2c6fe16ff0dbcdf0032b0549bb06")
	if err != nil {
		t.Fatalf("UserState returned error: %v", err)
	}
	log.Info().Any("res", result).Send()
}
