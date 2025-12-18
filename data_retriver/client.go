package dataretriver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	MainnetAPIURL = "https://api.hyperliquid.xyz"
	TestnetAPIURL = "https://api.hyperliquid-testnet.xyz"
	LocalAPIURL   = "http://localhost:3001"

	// httpErrorStatusCode is the minimum status code considered an error
	httpErrorStatusCode = 400
)

type Client struct {
	debug      bool
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string, debug bool) *Client {
	return newClient(baseURL, debug)
}

func newClient(baseURL string, debug bool) *Client {
	if baseURL == "" {
		baseURL = MainnetAPIURL
	}
	transport := &http.Transport{
		MaxIdleConns:        10,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     365 * 24 * time.Hour,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second, // 请求超时
	}

	cli := &Client{
		debug:      debug,
		baseURL:    baseURL,
		httpClient: client,
	}
	return cli
}

func (c *Client) post(path string, payload any) ([]byte, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := c.baseURL + path
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.debug {
		log.Debug().Any("body", payload).Msg("HTTP request")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body := make([]byte, 0)
	if resp.Body != nil {
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
	}
	if c.debug {
		log.Debug().Any("body", string(body)).Msg("HTTP response")
	}

	if resp.StatusCode >= httpErrorStatusCode {
		return nil, fmt.Errorf("status %d, body %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (c *Client) UserState(address string) (*UserState, error) {
	payload := map[string]any{
		"type": "clearinghouseState",
		"user": address,
	}

	resp, err := c.post("/info", payload)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user state: %w", err)
	}

	var result UserState
	if err := result.UnmarshalJSON(resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user state: %w", err)
	}
	return &result, nil
}

func (c *Client) MetaAndAssetCtxs() (*MetaAndAssetCtxs, error) {
	resp, err := c.post("/info", map[string]any{
		"type": "metaAndAssetCtxs",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch meta and asset contexts: %w", err)
	}

	var result []any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal meta and asset contexts: %w", err)
	}

	if len(result) < 2 {
		return nil, fmt.Errorf("expected at least 2 elements in response, got %d", len(result))
	}

	metaBytes, err := json.Marshal(result[0])
	if err != nil {
		return nil, fmt.Errorf("failed to marshal meta data: %w", err)
	}

	meta, err := parseMetaResponse(metaBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse meta: %w", err)
	}

	ctxsBytes, err := json.Marshal(result[1])
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ctxs data: %w", err)
	}

	var ctxs []AssetCtx
	if err := json.Unmarshal(ctxsBytes, &ctxs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ctxs: %w", err)
	}

	metaAndAssetCtxs := &MetaAndAssetCtxs{
		Meta: *meta,
		Ctxs: ctxs,
	}

	return metaAndAssetCtxs, nil
}

func parseMetaResponse(resp []byte) (*Meta, error) {
	var meta map[string]json.RawMessage
	if err := json.Unmarshal(resp, &meta); err != nil {
		return nil, fmt.Errorf("failed to unmarshal meta response: %w", err)
	}

	var universe []AssetInfo
	if err := json.Unmarshal(meta["universe"], &universe); err != nil {
		return nil, fmt.Errorf("failed to unmarshal universe: %w", err)
	}

	var marginTables [][]any
	if err := json.Unmarshal(meta["marginTables"], &marginTables); err != nil {
		return nil, fmt.Errorf("failed to unmarshal margin tables: %w", err)
	}

	marginTablesResult := make([]MarginTable, len(marginTables))
	for i, marginTable := range marginTables {
		id := marginTable[0].(float64)
		tableBytes, err := json.Marshal(marginTable[1])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal margin table data: %w", err)
		}

		var marginTableData map[string]any
		if err := json.Unmarshal(tableBytes, &marginTableData); err != nil {
			return nil, fmt.Errorf("failed to unmarshal margin table data: %w", err)
		}

		marginTiersBytes, err := json.Marshal(marginTableData["marginTiers"])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal margin tiers: %w", err)
		}

		var marginTiers []MarginTier
		if err := json.Unmarshal(marginTiersBytes, &marginTiers); err != nil {
			return nil, fmt.Errorf("failed to unmarshal margin tiers: %w", err)
		}

		marginTablesResult[i] = MarginTable{
			ID:          int(id),
			Description: marginTableData["description"].(string),
			MarginTiers: marginTiers,
		}
	}

	return &Meta{
		Universe:     universe,
		MarginTables: marginTablesResult,
	}, nil
}
