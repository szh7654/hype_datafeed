package dataretriver

import (
	"fmt"
	"os"
	"time"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

const (
	MainnetAPIURL = "https://api.hyperliquid.xyz"
	//TestnetAPIURL = "https://api.hyperliquid-testnet.xyz"
	LocalAPIURL = "http://localhost:3001"
)

var (
	httpClient = fasthttp.Client{
		MaxConnsPerHost:          20,
		MaxIdleConnDuration:      30 * time.Minute,
		ReadBufferSize:           4096,
		WriteBufferSize:          4096,
		NoDefaultUserAgentHeader: true,
	}
)

func post(url string, payload []byte) ([]byte, *fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()

	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.SetBody(payload)

	err := httpClient.Do(req, resp)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		log.Error().Msgf("Non-OK status code: %d, Body preview: %s", resp.StatusCode(), string(resp.Body())[:min(100, len(resp.Body()))])
		return nil, resp, fmt.Errorf("received non-OK status code: %d", resp.StatusCode())
	}

	return resp.Body(), resp, nil

}

func FetchUserState(address string) (*UserState, error) {
	payload := []byte(fmt.Sprintf(`{"type": "clearinghouseState", "user": "%s"}`, address))
	body, resp, err := post(LocalAPIURL+"/info", payload)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user state: %w", err)
	}

	var result UserState
	if err := sonic.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user state: %w", err)
	}
	return &result, nil
}

func FetchMetaAndAssetCtxs() (*MetaAndAssetCtxs, error) {
	payload := []byte(`{"type": "metaAndAssetCtxs"}`)
	body, resp, err := post(MainnetAPIURL+"/info", payload)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch meta and asset contexts: %w", err)
	}
	var res MetaAndAssetCtxs
	err = sonic.Unmarshal(body, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal meta and asset contexts: %w", err)
	}
	return &res, nil
}

func FetchL4Snapshot() *L4SnapShot {
	log.Info().Msg("fetching l4 snapshot")
	os.Remove("/root/l4Snapshots.json")
	payload := []byte(`{"type": "fileSnapshot",
     "request": {"type": "l4Snapshots", "includeUsers": true, "includeTriggerOrders": true}, 
     "outPath": "/root/l4Snapshots.json", 
     "includeHeightInOutput": true
     }`)
	body, resp, err := post(LocalAPIURL+"/info", payload)
	defer fasthttp.ReleaseResponse(resp)
	if err != nil {
		panic(err)
	}
	log.Info().Msgf("fileSnapshot response Body: %s", string(body))

	f, err := os.Open("/root/l4Snapshots.json")
	if err != nil {
		panic(err)
	}

	var res L4SnapShot
	decoder := sonic.ConfigDefault.NewDecoder(f)
	var l4SnapShot L4SnapShot
	err = decoder.Decode(&l4SnapShot)
	if err != nil {
		panic(err)
	}
	log.Info().Msgf("l4 snapshot block number: %d", l4SnapShot.BlockNumber)
	return &res
}
