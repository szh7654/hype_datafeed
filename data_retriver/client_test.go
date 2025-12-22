package dataretriver

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestClient_FetchUserState(t *testing.T) {
	baseURL = MainnetAPIURL
	result, err := FetchUserState("0x0448e125d7c83a35a211722a12627311c3e6a657")
	if err != nil {
		t.Fatalf("UserState returned error: %v", err)
	}
	log.Info().Any("res", result).Send()
}

func TestClient_FetchAsset(t *testing.T) {
	baseURL = MainnetAPIURL
	result, err := FetchMetaAndAssetCtxs()
	if err != nil {
		t.Fatalf("UserState returned error: %v", err)
	}
	log.Info().Any("res", result).Send()
}
