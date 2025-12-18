package dataretriver

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestBlockFill_UnmarshalJSON(t *testing.T) {
	var blockFill BlockFill
	err := blockFill.UnmarshalJSON([]byte(blockFillSample))
	if err != nil {
		t.Error(err)
	}
	log.Info().Any("blockfill", blockFill).Send()
}

var blockFillSample = `
{
  "local_time": "2025-10-27T06:30:38.191000595",
  "block_time": "2025-10-27T06:30:38.019609515",
  "block_number": 776534347,
  "events": [
    [
      "0xd7e678a2ad378286d79e278bbccb6438f1a8a064",
      {
        "coin": "HYPE",
        "px": "48.334",
        "sz": "1.87",
        "side": "B",
        "time": 1761546638019,
        "startPosition": "334.68",
        "dir": "Open Long",
        "closedPnl": "0.0",
        "hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "oid": 213637501487,
        "crossed": true,
        "fee": "0.040673",
        "tid": 990908538720549,
        "feeToken": "USDC",
        "twapId": 1305284
      }
    ],
    [
      "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
      {
        "coin": "HYPE",
        "px": "48.334",
        "sz": "1.87",
        "side": "A",
        "time": 1761546638019,
        "startPosition": "-700.92",
        "dir": "Open Short",
        "closedPnl": "0.0",
        "hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "oid": 213637501029,
        "crossed": false,
        "fee": "-0.000903",
        "tid": 990908538720549,
        "cloid": "0x9152106fa754fb298f52649ec28a812a",
        "feeToken": "USDC",
        "twapId": null
      }
    ]
  ]
}
`
