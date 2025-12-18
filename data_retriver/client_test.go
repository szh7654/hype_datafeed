package dataretriver

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestClient_UserState(t *testing.T) {
	// 创建客户端实例
	client := newClient(MainnetAPIURL, true)

	// 调用 UserState 方法
	result, err := client.UserState("0x0448e125d7c83a35a211722a12627311c3e6a657")
	if err != nil {
		t.Fatalf("UserState returned error: %v", err)
	}
	log.Info().Any("res", result).Send()
}
