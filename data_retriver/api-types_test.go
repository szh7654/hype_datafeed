package dataretriver

import (
	"testing"
  "os"
	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
)

func TestBlockFill_UnmarshalJSON(t *testing.T) {
	var blockFill BlockFill
	err := sonic.UnmarshalString(blockFillSample, &blockFill)
	if err != nil {
		t.Error(err)
	}
	log.Info().Any("blockfill", blockFill).Send()
}

func TestBlockOrderStatus_UnmarshalJSON(t *testing.T) {
	var blockOrderStatus BlockOrderStatus
	err := sonic.UnmarshalString(blockOrderStatusSample, &blockOrderStatus)
	if err != nil {
		t.Error(err)
	}
	log.Info().Any("blockorderstatus", blockOrderStatus).Send()
}

func TestBlockOrderBookDiff_UnmarshalJSON(t *testing.T) {
	var blockOrderBookDiff BlockOrderBookDiff
	err := sonic.UnmarshalString(blockOrderBookDiffSample, &blockOrderBookDiff)
	if err != nil {
		t.Error(err)
	}
	log.Info().Any("blockorderbookdiff", blockOrderBookDiff).Send()
}

func TestL4Snapshot_UnmarshalJSON(t *testing.T) {
  f, err := os.Open("/home/szh/nt/quant/hyperliquid_indexer/sample_data/l4Snapshots.json")
  if err != nil {
		t.Error(err)
	}
  decoder := sonic.ConfigDefault.NewDecoder(f)
  var l4SnapShot L4SnapShot
  err = decoder.Decode(&l4SnapShot)
  if err!= nil {
		t.Error(err)
	}
	log.Info().Msgf("l4snapshot, block_number: %d", l4SnapShot.BlockNumber)
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

var blockOrderStatusSample = `
{
  "local_time": "2025-10-28T00:00:00.040980047",
  "block_time": "2025-10-27T23:49:22.031892297",
  "block_number": 777324092,
  "events": [
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": "0xdad084335ca62a91dc4a042e55063c0106009c18f7a949637e992f861baa047c",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "W",
        "side": "A",
        "limitPx": "0.07326",
        "sz": "274397.5",
        "oid": 214381798682,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "274397.5",
        "tif": "Alo",
        "cloid": "0x0000019a28138ef006de0fed67d8e7d1"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xae7a0f9f663bb54bfe95378c7e3de7dd6e28d6bc",
      "hash": null,
      "builder": null,
      "status": "filled",
      "order": {
        "coin": "ENA",
        "side": "B",
        "limitPx": "0.5022",
        "sz": "0.0",
        "oid": 214381787192,
        "timestamp": 1761608960628,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1048.0",
        "tif": "Alo",
        "cloid": "0x0014a600639e580000019a281389f01e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xbd37cad03bd7c7de84ef4ffc87bb52079bf3653d",
      "hash": null,
      "builder": null,
      "status": "filled",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.456",
        "sz": "0.0",
        "oid": 214381773966,
        "timestamp": 1761608959620,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.4",
        "tif": "Alo",
        "cloid": "0x2f71eff5cca248cdb135ee7b44bd9560"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": "0x8da3face76376d0d8f1d042e55063c01070012b4113a8bdf316ca621353b46f8",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "PENGU",
        "side": "A",
        "limitPx": "0.021455",
        "sz": "981443.0",
        "oid": 214381798683,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "981443.0",
        "tif": "Alo",
        "cloid": "0x0000019a28138ef202de0fed67d8e7d4"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x716d7ae407194c19592036097eac71109709a19b",
      "hash": null,
      "builder": null,
      "status": "perpMarginRejected",
      "order": {
        "coin": "RENDER",
        "side": "A",
        "limitPx": "2.572",
        "sz": "2591.8",
        "oid": 214381798685,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2591.8",
        "tif": "Alo",
        "cloid": "0x00000000000000001bb0b872bf0c2386"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa1e03bcb1a36da135ca61b86dd1d6da9a9357489",
      "hash": "0xa61e5e9fc2eb3480a798042e55063c010a0076855dee535249e709f281ef0e6b",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MNT",
        "side": "A",
        "limitPx": "1.6644",
        "sz": "222.6",
        "oid": 214381798686,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "222.6",
        "tif": "Alo",
        "cloid": "0x1f65366cc5b7889084feb3c563fa262b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd0254dd9c10498b256b6ec3bd8de37a9ba4329de",
      "hash": "0x0bc54bd5f60db9780d3f042e55063c010c0063bb9100d84aaf8df728b5019362",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "OM",
        "side": "A",
        "limitPx": "0.1135",
        "sz": "29290.9",
        "oid": 214381798687,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "29290.9",
        "tif": "Alo",
        "cloid": "0x000000000000000018727e046e7fe0da"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0xbe98c2710f9efbf3c012042e55063c010d00da56aa921ac562616dc3ce92d5de",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "46.439",
        "sz": "65.18",
        "oid": 214381798688,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "65.18",
        "tif": "Alo",
        "cloid": "0x00000000000008115117500967300337"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
      "hash": "0x716c390c29203e6f72e5042e55063c010e0050f1c4235d411534e45ee824185a",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.49",
        "sz": "198.9",
        "oid": 214381798689,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "198.9",
        "tif": "Alo",
        "cloid": "0x5f611758498c4ee4ac7bda0d718fa10d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
      "hash": "0x200f1f36c57f51c72188042e55063c010e01011c60727099c3d7ca8984732bb1",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.502",
        "sz": "298.26",
        "oid": 214381798690,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "298.26",
        "tif": "Alo",
        "cloid": "0xb7e32003ba6d4e68a83f7c5e30e1f2d6"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
      "hash": "0xceb2056161ce651ed02b042e55063c010e010246fcc183f0727ab0b420c23f09",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.514",
        "sz": "447.3",
        "oid": 214381798691,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "447.3",
        "tif": "Alo",
        "cloid": "0x2265285825654974acbb03820854e8e8"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
      "hash": "0x243fafa742b180eb25b9042e55063c010f00c78cddb49fbdc8085afa01b55ad5",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.495",
        "sz": "375.99",
        "oid": 214381798692,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "375.99",
        "tif": "Alo",
        "cloid": "0xfde62fd5fdf64817800e209fdf2e9aee"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
      "hash": "0xd2e295d1df009442d45c042e55063c010f0101b77a03b31476ab41249e046e2d",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.506",
        "sz": "973.89",
        "oid": 214381798693,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "973.89",
        "tif": "Alo",
        "cloid": "0x07ddcba79a7f4689845084bbc1426f24"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
      "hash": "0x81857bfc7b5fa79a82ff042e55063c010f0102e21652c66c254e274f3a538185",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.518",
        "sz": "973.64",
        "oid": 214381798694,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "973.64",
        "tif": "Alo",
        "cloid": "0xfdc499c96e234753914b8a2a3265dd7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
      "hash": "0x3028622717aebaf231a2042e55063c010f01030cb2a1d9c4d3f10d79d6a294dc",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.576",
        "sz": "972.43",
        "oid": 214381798695,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "972.43",
        "tif": "Alo",
        "cloid": "0xa13b4a08d4574166b40363148633b633"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x44a709e254ac9f54bbd37a1d6388da5ddcb57c7a",
      "hash": null,
      "builder": null,
      "status": "perpMarginRejected",
      "order": {
        "coin": "RUNE",
        "side": "B",
        "limitPx": "0.86423",
        "sz": "17356.5",
        "oid": 214381798697,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "17356.5",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x57dd78cd36e76e2011e8f6dc25cabbaba994494b",
      "hash": "0x89e69cdd75d405e28b60042e55063c011100b4c310d724b42daf483034d7dfcd",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.497",
        "sz": "279.16",
        "oid": 214381798698,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "279.16",
        "tif": "Alo",
        "cloid": "0x00000000000000154952300078047725"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x3cba13788f65485e3e33042e55063c0112002b5e2a686730e082becb4e692248",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114218.0",
        "sz": "0.04203",
        "oid": 214381798699,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.04203",
        "tif": "Alo",
        "cloid": "0x00000000000012948887332941127932"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x91d11931c89ab5a88e1fcaf6fa137ab1c9b04179",
      "hash": "0xef8d8a13a8f68ad9f107042e55063c011300a1f943f9a9ab9356356667fa64c4",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "ME",
        "side": "B",
        "limitPx": "0.509",
        "sz": "11747.4",
        "oid": 214381798700,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "11747.4",
        "tif": "Alo",
        "cloid": "0x90910063475798293651761592790828"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x91d11931c89ab5a88e1fcaf6fa137ab1c9b04179",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ME",
        "side": "B",
        "limitPx": "0.508",
        "sz": "11747.4",
        "oid": 214381793385,
        "timestamp": 1761608961342,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "11747.4",
        "tif": "Alo",
        "cloid": "0x90910063475798293651761592790828"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.74",
        "sz": "74.63",
        "oid": 214381798701,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393db00003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.73",
        "sz": "74.63",
        "oid": 214381798702,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393dc00003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.72",
        "sz": "74.63",
        "oid": 214381798703,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393dd00003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "74.63",
        "oid": 214381798704,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393de00003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "74.63",
        "oid": 214381798705,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393df00003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.69",
        "sz": "74.63",
        "oid": 214381798706,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393e000003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xdbfa2a585b5ed6115950a3eaead82df47918304d",
      "hash": "0x0807ede4f5aa524d0981042e55063c01160005ca90ad711fabd09937b4ae2c37",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "LDO",
        "side": "A",
        "limitPx": "0.9415",
        "sz": "1095.2",
        "oid": 214381798707,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1095.2",
        "tif": "Alo",
        "cloid": "0xf594e967aacd1f4d4d9f036b9edae90e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.27613",
        "oid": 214381798708,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d970"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.27613",
        "oid": 214381798709,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d971"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.27613",
        "oid": 214381798710,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d972"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.27613",
        "oid": 214381798711,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d973"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.27613",
        "oid": 214381798712,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d974"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.27613",
        "oid": 214381798713,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d975"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.27613",
        "oid": 214381798714,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d976"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.27613",
        "oid": 214381798715,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d977"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.27613",
        "oid": 214381798716,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d978"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.27613",
        "oid": 214381798717,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c66d979"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
      "hash": "0x6daedb1b28ccd7446f28042e55063c011800f300c3cff6161177866de7c0b12f",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "LINEA",
        "side": "A",
        "limitPx": "0.014594",
        "sz": "68535.0",
        "oid": 214381798718,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "68535.0",
        "tif": "Alo",
        "cloid": "0x6edd9f291d6c42f073c1024dface5008"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
      "hash": "0x1c51c145c51bea9c1dcb042e55063c011801012b601f096ec01a6c98841fc486",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "LINEA",
        "side": "B",
        "limitPx": "0.014584",
        "sz": "68535.0",
        "oid": 214381798719,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "68535.0",
        "tif": "Alo",
        "cloid": "0x88b9fce2bd0ddc534aee5b18cc7bd548"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x4fb31d5ffaa08be2128cb14485584d76eebef173",
      "hash": "0x208251b6425e19c021fc042e55063c011900699bdd513892c44afd090151f3aa",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.473",
        "sz": "85.67",
        "oid": 214381798720,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "85.67",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000115002ba0"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x4fb31d5ffaa08be2128cb14485584d76eebef173",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.472",
        "sz": "43.04",
        "oid": 214381798721,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "43.04",
        "tif": "Alo",
        "cloid": "0x80000000000000000000000114fea500"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x4334ff1dd9fa46667b718a6f8a489c3d0079b44b",
      "hash": "0x86293eec75709eb787a2042e55063c011b0056d21073bd8929f1ea3f347478a2",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "SNX",
        "side": "A",
        "limitPx": "1.1601",
        "sz": "70.7",
        "oid": 214381798722,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "70.7",
        "tif": "Alo",
        "cloid": "0x7d02a010e1be47f69d33a25ff323760f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6ac569a098a850a6a3c37547bc5d1d85cba80f33",
      "hash": "0x38fcb5878f01e1333a76042e55063c011c00cd6d2a050005dcc560da4e05bb1d",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "WLFI",
        "side": "B",
        "limitPx": "0.14742",
        "sz": "4886.0",
        "oid": 214381798723,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "4886.0",
        "tif": "Alo",
        "cloid": "0x20251028000000000000000000227021"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.27613",
        "oid": 214381798724,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717f4"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.27613",
        "oid": 214381798725,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717f5"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.27613",
        "oid": 214381798726,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717f6"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.27613",
        "oid": 214381798727,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717f7"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.27613",
        "oid": 214381798728,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717f8"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.27613",
        "oid": 214381798729,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717f9"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.27613",
        "oid": 214381798730,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717fa"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.27613",
        "oid": 214381798731,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717fb"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.27613",
        "oid": 214381798732,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717fc"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.27613",
        "oid": 214381798733,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c6717fd"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.5",
        "sz": "12.1226",
        "oid": 214381798734,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1226",
        "tif": "Alo",
        "cloid": "0x370a57e5c391f49136333414a72a1004"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.6",
        "sz": "12.1223",
        "oid": 214381798735,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1223",
        "tif": "Alo",
        "cloid": "0xdea93eac17184553e2307bbeff022e03"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "12.122",
        "oid": 214381798736,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.122",
        "tif": "Alo",
        "cloid": "0x4e130bd9ffff5b32bf3c145eb3833302"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.8",
        "sz": "12.1217",
        "oid": 214381798737,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1217",
        "tif": "Alo",
        "cloid": "0x8d847d63dc493b11d7c6482df797f501"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.9",
        "sz": "12.1215",
        "oid": 214381798738,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1215",
        "tif": "Alo",
        "cloid": "0x790c534a9eb7109ce55c1dae85d9c900"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa880d6cc607a05ea617307ab3b0d335e8d8424ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "6.4327",
        "oid": 214381798739,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "6.4327",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000004653d0f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.74",
        "sz": "74.63",
        "oid": 214381798740,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393e100003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.73",
        "sz": "74.63",
        "oid": 214381798741,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393e200003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.72",
        "sz": "74.63",
        "oid": 214381798742,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393e300003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "74.63",
        "oid": 214381798743,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393e400003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "74.63",
        "oid": 214381798744,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393e500003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.69",
        "sz": "74.63",
        "oid": 214381798745,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393e600003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1918fbc5ab2205e7bc391d250fcfff4e830ce053",
      "hash": "0xb71e068f0ed82d9db897042e55063c0121001e74a9db4c6f5ae6b1e1cddc0788",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "KAS",
        "side": "A",
        "limitPx": "0.057415",
        "sz": "20000.0",
        "oid": 214381798746,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "20000.0",
        "tif": "Alo",
        "cloid": "0x20251028000000000000000000072996"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
      "hash": "0x69f17d2a286970196b6b042e55063c012200950fc36c8eeb0dba287ce76d4a04",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "NIL",
        "side": "B",
        "limitPx": "0.2935",
        "sz": "5108.0",
        "oid": 214381798747,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "5108.0",
        "tif": "Alo",
        "cloid": "0x00000000000000001cc1cdd8d5c279d5"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.27613",
        "oid": 214381798748,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c67567a"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.27613",
        "oid": 214381798749,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c67567b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.27613",
        "oid": 214381798750,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c67567c"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.27613",
        "oid": 214381798751,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c67567d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.27613",
        "oid": 214381798752,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c67567e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.27613",
        "oid": 214381798753,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c67567f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.27613",
        "oid": 214381798754,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c675680"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.27613",
        "oid": 214381798755,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c675681"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.27613",
        "oid": 214381798756,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c675682"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.27613",
        "oid": 214381798757,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c675683"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
      "hash": "0xcf986a605b8bf510d112042e55063c0124008245f68f13e2736115b31a8fcefb",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "NIL",
        "side": "B",
        "limitPx": "0.29373",
        "sz": "10247.0",
        "oid": 214381798758,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "10247.0",
        "tif": "Alo",
        "cloid": "0x00000000000000001cc1cdd8d5c279d7"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
      "hash": "0x826be0fb751d378c83e5042e55063c012500f8e11010565e26348c4e34111177",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "ENA",
        "side": "B",
        "limitPx": "0.50209",
        "sz": "1991.0",
        "oid": 214381798759,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1991.0",
        "tif": "Alo",
        "cloid": "0x1b79effffc443f246bbd77bbf4376b3f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
      "hash": "0x353f57968eae7a0836b9042e55063c0126006f7c29a198dad90802e94da253f2",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "46.406",
        "sz": "23.1",
        "oid": 214381798760,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "23.1",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x621c5551678189b9a6c94d929924c225ff1d63ab",
      "hash": "0xe812ce31a83fbc83e98c042e55063c012700e6174332db558bdb79846733966e",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.492",
        "sz": "217.33",
        "oid": 214381798761,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "217.33",
        "tif": "Alo",
        "cloid": "0x00000000000002390096147980677137"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0xb360a89e0e74c672b4da042e55063c012b00c083a977e544572953f0cd78a05d",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.488",
        "sz": "65.22",
        "oid": 214381798762,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "65.22",
        "tif": "Alo",
        "cloid": "0x00000000000008115117500967300339"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": "0x66341f39280608ee67ad042e55063c012c00371ec30927c009fcca8be709e2d9",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "FARTCOIN",
        "side": "B",
        "limitPx": "0.4026",
        "sz": "31731.1",
        "oid": 214381798763,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "31731.1",
        "tif": "Alo",
        "cloid": "0x19a1a764953891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "FARTCOIN",
        "side": "B",
        "limitPx": "0.40206",
        "sz": "31731.1",
        "oid": 214381796748,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "31731.1",
        "tif": "Alo",
        "cloid": "0x19a1a764953891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": "0x66341f39280608ee67ad042e55063c012c00371ec30927c009fcca8be709e2d9",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "FARTCOIN",
        "side": "B",
        "limitPx": "0.40114",
        "sz": "92193.1",
        "oid": 214381798764,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "92193.1",
        "tif": "Alo",
        "cloid": "0x19a1a764951891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "FARTCOIN",
        "side": "B",
        "limitPx": "0.4009",
        "sz": "92193.1",
        "oid": 214381796749,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "92193.1",
        "tif": "Alo",
        "cloid": "0x19a1a764951891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": "0x66341f39280608ee67ad042e55063c012c00371ec30927c009fcca8be709e2d9",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "FARTCOIN",
        "side": "A",
        "limitPx": "0.40473",
        "sz": "31731.1",
        "oid": 214381798765,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "31731.1",
        "tif": "Alo",
        "cloid": "0x33434ec92c2891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "FARTCOIN",
        "side": "A",
        "limitPx": "0.40488",
        "sz": "31731.1",
        "oid": 214381638516,
        "timestamp": 1761608950444,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "31731.1",
        "tif": "Alo",
        "cloid": "0x33434ec92c2891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": "0x66341f39280608ee67ad042e55063c012c00371ec30927c009fcca8be709e2d9",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "FARTCOIN",
        "side": "A",
        "limitPx": "0.4059",
        "sz": "93842.8",
        "oid": 214381798766,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "93842.8",
        "tif": "Alo",
        "cloid": "0x33434ec98e0891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "FARTCOIN",
        "side": "A",
        "limitPx": "0.40592",
        "sz": "93842.8",
        "oid": 214381786529,
        "timestamp": 1761608960628,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "93842.8",
        "tif": "Alo",
        "cloid": "0x33434ec98e0891200000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
      "hash": "0xcbdb0c6f5b288de5cd54042e55063c012e002454f62bacb76fa3b7c21a2c67d0",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798767,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x0000000000000000e0f5526e6237dbd1"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6a0b567be46b51ad02db4ebdf90d057ecb02f864",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "LINK",
        "side": "A",
        "limitPx": "18.223",
        "sz": "84.7",
        "oid": 214381798768,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "84.7",
        "tif": "Alo",
        "cloid": "0x20251028000000000000000000402569"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.5",
        "sz": "12.1226",
        "oid": 214381798769,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1226",
        "tif": "Alo",
        "cloid": "0xc266b9904559c400ea979581bda6e904"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.6",
        "sz": "12.1223",
        "oid": 214381798770,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1223",
        "tif": "Alo",
        "cloid": "0xf0779e6a79e50d77eb1d23697b5e6903"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "12.122",
        "oid": 214381798771,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.122",
        "tif": "Alo",
        "cloid": "0xaaa5e67f161d6b6bb49438bfc0f56802"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.8",
        "sz": "12.1217",
        "oid": 214381798772,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1217",
        "tif": "Alo",
        "cloid": "0xe2008765835ab8a0eb23ed1203d25101"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba5f4f92f396e6de20d3dae5eefa65f4ecca806a",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.9",
        "sz": "12.1215",
        "oid": 214381798773,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.1215",
        "tif": "Alo",
        "cloid": "0xfcc1485c887e810a26ff6703cd516300"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.5",
        "sz": "1.9892",
        "oid": 214381798774,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9892",
        "tif": "Alo",
        "cloid": "0xecd862f799a27efa466302b7a3ad7b0a"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.6",
        "sz": "1.9893",
        "oid": 214381798775,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9893",
        "tif": "Alo",
        "cloid": "0x3cecd3a49b9503f8b4c36dfb117c7109"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "1.9893",
        "oid": 214381798776,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9893",
        "tif": "Alo",
        "cloid": "0x82a6c4b2fef3e283697777880227d008"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.8",
        "sz": "1.9894",
        "oid": 214381798777,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9894",
        "tif": "Alo",
        "cloid": "0xb85a5f7894c85ebcd212f43a1adbf607"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.9",
        "sz": "1.9894",
        "oid": 214381798778,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9894",
        "tif": "Alo",
        "cloid": "0x01307a0d9a2abcd0108b9d0eeb9a5f06"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.0",
        "sz": "1.9895",
        "oid": 214381798779,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9895",
        "tif": "Alo",
        "cloid": "0x30f74a9a7c31973ba8fda820b9946905"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.1",
        "sz": "1.9895",
        "oid": 214381798780,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9895",
        "tif": "Alo",
        "cloid": "0x77c080b563e23578daae9acc2c984a04"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.2",
        "sz": "1.9896",
        "oid": 214381798781,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9896",
        "tif": "Alo",
        "cloid": "0x0ee38d32cb70c9bb1db79f9e07a27203"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.3",
        "sz": "1.9896",
        "oid": 214381798782,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9896",
        "tif": "Alo",
        "cloid": "0x0e1b96f846a0fcb080fde5b6062d7902"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.4",
        "sz": "1.9897",
        "oid": 214381798783,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9897",
        "tif": "Alo",
        "cloid": "0x554c5475932292184167b9252f95fa01"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7839e2f2c375dd2935193f2736167514efff9916",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.5",
        "sz": "1.9897",
        "oid": 214381798784,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.9897",
        "tif": "Alo",
        "cloid": "0x37b19fdba8d68c0119973b8bbcee4f00"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x20faffd0045455e86b36b827d68f8f99f9c929b4",
      "hash": "0x49fc5d76dafeda504b76042e55063c013300755c75f1f922edc508c999f2b43a",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.473",
        "sz": "4.44",
        "oid": 214381798785,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "4.44",
        "tif": "Alo",
        "cloid": "0x874fb12c53934f5c8ad09a59a00fe1ab"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.27613",
        "oid": 214381798786,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c681209"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.27613",
        "oid": 214381798787,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c68120a"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.27613",
        "oid": 214381798788,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c68120b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.27613",
        "oid": 214381798789,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c68120c"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.27613",
        "oid": 214381798790,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c68120d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.27613",
        "oid": 214381798791,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c68120e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.27613",
        "oid": 214381798792,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c68120f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.27613",
        "oid": 214381798793,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c681210"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.27613",
        "oid": 214381798794,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c681211"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.27613",
        "oid": 214381798795,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c681212"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
      "hash": "0xafa34aad0e115f47b11d042e55063c0135006292a9147e19536bf5ffcd153932",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MELANIA",
        "side": "B",
        "limitPx": "0.13035",
        "sz": "25070.1",
        "oid": 214381798796,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "25070.1",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
      "hash": "0x5e4630d7aa60729f5fbf042e55063c01350101bd45639171020edc2a69644c8a",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MELANIA",
        "side": "A",
        "limitPx": "0.13464",
        "sz": "102936.8",
        "oid": 214381798797,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "102936.8",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x6276c14827a2a1c363f0042e55063c013600d92dc2a5c095063f6c9ae6a67bae",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@142",
        "side": "A",
        "limitPx": "114197.0",
        "sz": "0.65926",
        "oid": 214381798798,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.65926",
        "tif": "Alo",
        "cloid": "0x00000000000003729455888846495049"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x154a37e34133e43f16c3042e55063c0137004fc8dc370311b912e3360037be29",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "2Z",
        "side": "A",
        "limitPx": "0.22977",
        "sz": "15472.0",
        "oid": 214381798799,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "15472.0",
        "tif": "Alo",
        "cloid": "0x00000000000010688361787313138457"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0xc81dae7e5ac526bac997042e55063c013800c663f5c8458c6be659d119c900a5",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "YGG",
        "side": "B",
        "limitPx": "0.13615",
        "sz": "14705.0",
        "oid": 214381798800,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "14705.0",
        "tif": "Alo",
        "cloid": "0x00000000000010698213411120389837"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x7af12519745669367c6a042e55063c0139003cff0f5988081eb9d06c335a4321",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "ATOM",
        "side": "A",
        "limitPx": "3.179",
        "sz": "944.03",
        "oid": 214381798801,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "944.03",
        "tif": "Alo",
        "cloid": "0x00000000000010708627985259303803"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x2dc49bb48de7abb22f3e042e55063c013a00b39a28eaca84d18d47074ceb859c",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "2Z",
        "side": "A",
        "limitPx": "0.22917",
        "sz": "4357.0",
        "oid": 214381798802,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "4357.0",
        "tif": "Alo",
        "cloid": "0x00000000000010688361787313138456"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xbb828515677729c4db2f790a52cf14b1c659e19f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SUI",
        "side": "B",
        "limitPx": "2.6064",
        "sz": "709.3",
        "oid": 214381798803,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "709.3",
        "tif": "Alo",
        "cloid": "0x2c47768ae297640516de2b7cfa5d4e04"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xbb828515677729c4db2f790a52cf14b1c659e19f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SUI",
        "side": "B",
        "limitPx": "2.6065",
        "sz": "709.3",
        "oid": 214381798804,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "709.3",
        "tif": "Alo",
        "cloid": "0x5e61964fe61e20c0edc4e3f71d34b403"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xbb828515677729c4db2f790a52cf14b1c659e19f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SUI",
        "side": "B",
        "limitPx": "2.6066",
        "sz": "709.2",
        "oid": 214381798805,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "709.2",
        "tif": "Alo",
        "cloid": "0xc22bd5e095dceb0fbd79694ae46ad802"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xbb828515677729c4db2f790a52cf14b1c659e19f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SUI",
        "side": "B",
        "limitPx": "2.6067",
        "sz": "709.2",
        "oid": 214381798806,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "709.2",
        "tif": "Alo",
        "cloid": "0x588b15c1a33a56ac7549c27dc76b3001"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xbb828515677729c4db2f790a52cf14b1c659e19f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SUI",
        "side": "B",
        "limitPx": "2.6068",
        "sz": "709.2",
        "oid": 214381798807,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "709.2",
        "tif": "Alo",
        "cloid": "0xf6044c26174d638ebc87faeec7563100"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": "0x936b88eac10a30a994e5042e55063c013c00a0d05c0d4f7b3734343d800e0a94",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MANTA",
        "side": "A",
        "limitPx": "0.11475",
        "sz": "185823.6",
        "oid": 214381798808,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "185823.6",
        "tif": "Alo",
        "cloid": "0x0000019a28138f6004de0fed67d8e7d5"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x953996a8b3e3c9b2ecc13664950910d013c060a5",
      "hash": "0x463eff85da9b732547b8042e55063c013d00176b759e91f7ea07aad8999f4d0f",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "SUI",
        "side": "A",
        "limitPx": "2.608",
        "sz": "1103.6",
        "oid": 214381798809,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1103.6",
        "tif": "Alo",
        "cloid": "0x3f45a7792cc1cd87fa3a1cfba2ccb2af"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa880d6cc607a05ea617307ab3b0d335e8d8424ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.30551",
        "oid": 214381798810,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.30551",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000004653d13"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x44a709e254ac9f54bbd37a1d6388da5ddcb57c7a",
      "hash": null,
      "builder": null,
      "status": "perpMarginRejected",
      "order": {
        "coin": "kPEPE",
        "side": "B",
        "limitPx": "0.006926",
        "sz": "8000915.0",
        "oid": 214381798812,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "8000915.0",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x953996a8b3e3c9b2ecc13664950910d013c060a5",
      "hash": "0x5eb96357274f3a986033042e55063c0140007b3cc242596a02820ea9e6431483",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "NEAR",
        "side": "A",
        "limitPx": "2.3328",
        "sz": "1296.2",
        "oid": 214381798813,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1296.2",
        "tif": "Alo",
        "cloid": "0xa915ce55edc729f8f7d67643af853d88"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0d58586a54a1b2070677f898acebb8317e24e1c2",
      "hash": "0x118cd9f240d07d141306042e55063c014100f1d7dbd39be6b5558544ffd456fe",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "SYRUP",
        "side": "A",
        "limitPx": "0.38527",
        "sz": "4000.0",
        "oid": 214381798814,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "4000.0",
        "tif": "Alo",
        "cloid": "0x20251028000000000000000000075578"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0xc460508d5a61bf8fc5da042e55063c0142006872f564de616828fbe01965997a",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "LINEA",
        "side": "A",
        "limitPx": "0.014595",
        "sz": "140032.0",
        "oid": 214381798815,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "140032.0",
        "tif": "Alo",
        "cloid": "0x00000000000010688080311972452543"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x44a709e254ac9f54bbd37a1d6388da5ddcb57c7a",
      "hash": null,
      "builder": null,
      "status": "perpMarginRejected",
      "order": {
        "coin": "INJ",
        "side": "A",
        "limitPx": "8.9383",
        "sz": "1678.2",
        "oid": 214381798817,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1678.2",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x2a073dc38d8444872b80042e55063c01440055a928876359cdcfe9164c881e71",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "GRASS",
        "side": "A",
        "limitPx": "0.4165",
        "sz": "4919.0",
        "oid": 214381798818,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "4919.0",
        "tif": "Alo",
        "cloid": "0x00000000000010690895061725903586"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
      "hash": "0xdcdab45ea7158702de54042e55063c014500cc444218a5d480a35fb1661960ed",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "WIF",
        "side": "A",
        "limitPx": "0.55251",
        "sz": "1493.0",
        "oid": 214381798819,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1493.0",
        "tif": "Alo",
        "cloid": "0xf3ce430381f99e5390f49a2c81585637"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa880d6cc607a05ea617307ab3b0d335e8d8424ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.30551",
        "oid": 214381798820,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.30551",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000004653d12"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x716d7ae407194c19592036097eac71109709a19b",
      "hash": null,
      "builder": null,
      "status": "perpMarginRejected",
      "order": {
        "coin": "kBONK",
        "side": "B",
        "limitPx": "0.014357",
        "sz": "2321725.0",
        "oid": 214381798822,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2321725.0",
        "tif": "Alo",
        "cloid": "0x00000000000000001d25a904890e2b7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x37c1e493c8a452ab3f5a6d01fda87795b91e7851",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.74",
        "sz": "49.35",
        "oid": 214381798823,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "49.35",
        "tif": "Alo",
        "cloid": "0x000000000051f06f000008dfc9881c25"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x37c1e493c8a452ab3f5a6d01fda87795b91e7851",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.73",
        "sz": "49.35",
        "oid": 214381798824,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "49.35",
        "tif": "Alo",
        "cloid": "0x000000000051f070000008dfc9881c25"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x37c1e493c8a452ab3f5a6d01fda87795b91e7851",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.72",
        "sz": "49.35",
        "oid": 214381798825,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "49.35",
        "tif": "Alo",
        "cloid": "0x000000000051f071000008dfc9881c25"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x37c1e493c8a452ab3f5a6d01fda87795b91e7851",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "49.35",
        "oid": 214381798826,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "49.35",
        "tif": "Alo",
        "cloid": "0x000000000051f072000008dfc9881c25"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x37c1e493c8a452ab3f5a6d01fda87795b91e7851",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "49.35",
        "oid": 214381798827,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "49.35",
        "tif": "Alo",
        "cloid": "0x000000000051f073000008dfc9881c25"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x37c1e493c8a452ab3f5a6d01fda87795b91e7851",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.69",
        "sz": "49.35",
        "oid": 214381798828,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "49.35",
        "tif": "Alo",
        "cloid": "0x000000000051f074000008dfc9881c25"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0xa8288ecb0d5a90f1a9a2042e55063c014900a6b0a85dafc34bf13a1dcc5e6adc",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "AVAX",
        "side": "A",
        "limitPx": "20.31",
        "sz": "146.92",
        "oid": 214381798829,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "146.92",
        "tif": "Alo",
        "cloid": "0x00000000000003727204089027654940"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.27613",
        "oid": 214381798830,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f14"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.27613",
        "oid": 214381798831,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f15"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.27613",
        "oid": 214381798832,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f16"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.27613",
        "oid": 214381798833,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f17"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.27613",
        "oid": 214381798834,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f18"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.27613",
        "oid": 214381798835,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f19"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.27613",
        "oid": 214381798836,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f1a"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.27613",
        "oid": 214381798837,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f1b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.27613",
        "oid": 214381798838,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f1c"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.27613",
        "oid": 214381798839,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c688f1d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x0dcf7c01407d15e90f49042e55063c014b0093e6db7034bbb1982753ff70efd3",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "SAND",
        "side": "B",
        "limitPx": "0.21459",
        "sz": "2226.0",
        "oid": 214381798840,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2226.0",
        "tif": "Alo",
        "cloid": "0x00000000000003742403737752656483"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.74",
        "sz": "46.71",
        "oid": 214381798841,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24660000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.73",
        "sz": "46.71",
        "oid": 214381798842,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24670000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.72",
        "sz": "46.71",
        "oid": 214381798843,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24680000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "46.71",
        "oid": 214381798844,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24690000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "46.71",
        "oid": 214381798845,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f246a0000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.69",
        "sz": "46.71",
        "oid": 214381798846,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f246b0000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.4",
        "sz": "1.2",
        "oid": 214381798847,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7ac80000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.3",
        "sz": "1.2",
        "oid": 214381798848,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7ac90000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.2",
        "sz": "1.2",
        "oid": 214381798849,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7aca0000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.1",
        "sz": "1.2",
        "oid": 214381798850,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7acb0000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.0",
        "sz": "1.2",
        "oid": 214381798851,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7acc0000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.9",
        "sz": "1.2",
        "oid": 214381798852,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7acd0000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.8",
        "sz": "1.2",
        "oid": 214381798853,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7ace0000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "1.2",
        "oid": 214381798854,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7acf0000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.6",
        "sz": "1.2",
        "oid": 214381798855,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7ad00000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x0dd5db9c3748486e747c6a123727f472668cf6ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.5",
        "sz": "1.2",
        "oid": 214381798856,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x0000000003ab7ad10000100adbbd1c8f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x82b409e6c2e617929cf22bd7e4b08564c207abc1",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.9",
        "sz": "1.2",
        "oid": 214381798857,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x00000000029e3de6000037f860561c21"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x82b409e6c2e617929cf22bd7e4b08564c207abc1",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.8",
        "sz": "1.2",
        "oid": 214381798858,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x00000000029e3de7000037f860561c21"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x82b409e6c2e617929cf22bd7e4b08564c207abc1",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "1.2",
        "oid": 214381798859,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x00000000029e3de8000037f860561c21"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x82b409e6c2e617929cf22bd7e4b08564c207abc1",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.6",
        "sz": "1.2",
        "oid": 214381798860,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x00000000029e3de9000037f860561c21"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x82b409e6c2e617929cf22bd7e4b08564c207abc1",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.5",
        "sz": "1.2",
        "oid": 214381798861,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.2",
        "tif": "Alo",
        "cloid": "0x00000000029e3dea000037f860561c21"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.74",
        "sz": "46.71",
        "oid": 214381798862,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f246c0000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.73",
        "sz": "46.71",
        "oid": 214381798863,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f246d0000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.72",
        "sz": "46.71",
        "oid": 214381798864,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f246e0000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "46.71",
        "oid": 214381798865,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f246f0000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "46.71",
        "oid": 214381798866,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24700000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.69",
        "sz": "46.71",
        "oid": 214381798867,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24710000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.5",
        "sz": "1.7348",
        "oid": 214381798868,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90147000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.4",
        "sz": "1.7348",
        "oid": 214381798869,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90148000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.3",
        "sz": "1.7348",
        "oid": 214381798870,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90149000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.2",
        "sz": "1.7348",
        "oid": 214381798871,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9014a000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.1",
        "sz": "1.7348",
        "oid": 214381798872,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9014b000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.0",
        "sz": "1.7348",
        "oid": 214381798873,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9014c000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.9",
        "sz": "1.7348",
        "oid": 214381798874,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9014d000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.8",
        "sz": "1.7348",
        "oid": 214381798875,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9014e000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "1.7348",
        "oid": 214381798876,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9014f000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.6",
        "sz": "1.7348",
        "oid": 214381798877,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90150000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.5",
        "sz": "1.7348",
        "oid": 214381798878,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90151000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": "0x3ec443a3d9d4a4cf403d042e55063c0151005b8974d7c3a1e28ceef698d87eb9",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@151",
        "side": "B",
        "limitPx": "4122.7",
        "sz": "0.738",
        "oid": 214381798879,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.738",
        "tif": "Alo",
        "cloid": "0x00000000000010717353709537080396"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114164.0",
        "sz": "0.02996",
        "oid": 214381798880,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89ba0000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114165.0",
        "sz": "0.02996",
        "oid": 214381798881,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89bb0000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.02996",
        "oid": 214381798882,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89bc0000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.02996",
        "oid": 214381798883,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89bd0000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.02996",
        "oid": 214381798884,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89be0000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.02996",
        "oid": 214381798885,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89bf0000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.02996",
        "oid": 214381798886,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89c00000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.02996",
        "oid": 214381798887,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89c10000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114172.0",
        "sz": "0.02996",
        "oid": 214381798888,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89c20000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.02996",
        "oid": 214381798889,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89c30000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.02996",
        "oid": 214381798890,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89c40000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.02996",
        "oid": 214381798891,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89c50000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xaf2ac71f62f341e5823d6985492409e92c940447",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.02996",
        "oid": 214381798892,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02996",
        "tif": "Alo",
        "cloid": "0x00000000002c89c60000279e68b81c22"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa880d6cc607a05ea617307ab3b0d335e8d8424ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "36.27",
        "oid": 214381798893,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "36.27",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000004653d16"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
      "hash": "0x573ea77526886c4258b8042e55063c015400bf5ac18b8b14fb0752c7e58c462c",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BANANA",
        "side": "A",
        "limitPx": "13.726",
        "sz": "74.0",
        "oid": 214381798894,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.0",
        "tif": "Alo",
        "cloid": "0x00000000000000007d066f289a799edf"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa880d6cc607a05ea617307ab3b0d335e8d8424ee",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "36.27",
        "oid": 214381798895,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "36.27",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000004653d17"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7b7f72a28fe109fa703eeed7984f2a8a68fedee2",
      "hash": "0xbce594ab59aaf139be5f042e55063c015600ac90f4ae100b60ae3ffe18aecb24",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.595",
        "sz": "142.6",
        "oid": 214381798896,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "142.6",
        "tif": "Alo",
        "cloid": "0x00000000a68f13289a01000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.74",
        "sz": "46.71",
        "oid": 214381798897,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24720000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.73",
        "sz": "46.71",
        "oid": 214381798898,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24730000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.72",
        "sz": "46.71",
        "oid": 214381798899,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24740000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "46.71",
        "oid": 214381798900,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24750000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "46.71",
        "oid": 214381798901,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24760000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b3482482899b4890ec5a2093cba2a558d0c14f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.69",
        "sz": "46.71",
        "oid": 214381798902,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "46.71",
        "tif": "Alo",
        "cloid": "0x00000000017f24770000264f4b151c65"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x230633a6e555acac1a865d09db2197c864e58927",
      "hash": "0x228c81e18ccd76312406042e55063c01580099c727c09503c6552d344bc1501b",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MET",
        "side": "A",
        "limitPx": "0.43661",
        "sz": "131.0",
        "oid": 214381798903,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "131.0",
        "tif": "Alo",
        "cloid": "0x20251028000000000000000000308840"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7b7f72a28fe109fa703eeed7984f2a8a68fedee2",
      "hash": "0xd55ff87ca65eb8acd6d9042e55063c01590010624151d77e7928a3cf65529297",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "46.206",
        "sz": "66.84",
        "oid": 214381798904,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "66.84",
        "tif": "Alo",
        "cloid": "0x00000000a38f13289a01000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa289ee1e56c0d5d041db762e6123e78af0f7d9ad",
      "hash": "0x88336f17bfeffb2889ad042e55063c015a0086fd5ae319fa2bfc1a6a7ee3d513",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "AVNT",
        "side": "A",
        "limitPx": "0.69177",
        "sz": "3618.0",
        "oid": 214381798905,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "3618.0",
        "tif": "Alo",
        "cloid": "0x00000000000000009c5ad588c5064723"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xff4cd3826ecee12acd4329aada4a2d3419fc463c",
      "hash": "0x3b06e5b2d9713da43c80042e55063c015b00fd9874745c76decf91059875178e",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "114170.0",
        "sz": "0.03459",
        "oid": 214381798906,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.03459",
        "tif": "Alo",
        "cloid": "0x00000000000002390092134500891887"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.74",
        "sz": "74.63",
        "oid": 214381798907,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393f300003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.73",
        "sz": "74.63",
        "oid": 214381798908,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393f400003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.72",
        "sz": "74.63",
        "oid": 214381798909,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393f500003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.71",
        "sz": "74.63",
        "oid": 214381798910,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393f600003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.7",
        "sz": "74.63",
        "oid": 214381798911,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393f700003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf060e33cfa0727a224a184622606213e8f4e6a6f",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "SOL",
        "side": "B",
        "limitPx": "198.69",
        "sz": "74.63",
        "oid": 214381798912,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "74.63",
        "tif": "Alo",
        "cloid": "0x00000000005393f800003761dcc11c7e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x91d11931c89ab5a88e1fcaf6fa137ab1c9b04179",
      "hash": "0xa0add2e90c93c29ba227042e55063c015d00eacea796e16d44767e3bcb979c86",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "ME",
        "side": "A",
        "limitPx": "0.515",
        "sz": "11747.4",
        "oid": 214381798913,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "11747.4",
        "tif": "Alo",
        "cloid": "0x90911063475798293651761592790829"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x91d11931c89ab5a88e1fcaf6fa137ab1c9b04179",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ME",
        "side": "A",
        "limitPx": "0.514",
        "sz": "11747.4",
        "oid": 214381797692,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "11747.4",
        "tif": "Alo",
        "cloid": "0x90911063475798293651761592790829"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": "0x538149842625051754fb042e55063c015e006169c12823e9f749f4d6e528df01",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MELANIA",
        "side": "B",
        "limitPx": "0.13056",
        "sz": "209117.3",
        "oid": 214381798914,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "209117.3",
        "tif": "Alo",
        "cloid": "0x0000019a28138fad04de0fed67d8e7e0"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": "0x0654c01f3fb6479307ce042e55063c015f00d804dab96665aa1d6b71feba217d",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "WIF",
        "side": "A",
        "limitPx": "0.55232",
        "sz": "445.0",
        "oid": 214381798915,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "445.0",
        "tif": "Alo",
        "cloid": "0x5160254e34e03f5946a1761228f6329f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": "0xb92836ba59478a0ebaa1042e55063c0160004e9ff44aa8e05cf0e20d184b63f9",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@142",
        "side": "B",
        "limitPx": "114124.0",
        "sz": "0.02629",
        "oid": 214381798916,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02629",
        "tif": "Alo",
        "cloid": "0x10416be0a8c519987dfaf0c0efe2e5d3"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x162cc7c861ebd0c06b3d72319201150482518185",
      "hash": "0x6bfbad5572d8cc8a6d75042e55063c016100c53b0ddbeb5c0fc458a831dca675",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798917,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x00000000000000000b56a216233de9dd"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": "0x1ecf23f08c6a0f062048042e55063c0162003bd6276d2dd8c297cf434b6de8f0",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@151",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "12.122",
        "oid": 214381798918,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "12.122",
        "tif": "Alo",
        "cloid": "0x122335fa72a824a2007ccad54f0252f3"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x21fe273e04336c952fa0de439ad00529a9dd7909",
      "hash": "0xd1a29a8ba5fb5181d31c042e55063c016300b27140fe7053756b45de64ff2b6c",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "ME",
        "side": "B",
        "limitPx": "0.508",
        "sz": "13203.8",
        "oid": 214381798919,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": true,
        "orderType": "Limit",
        "origSz": "13203.8",
        "tif": "Alo",
        "cloid": "0x71720133819897882521761608579254"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x21fe273e04336c952fa0de439ad00529a9dd7909",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ME",
        "side": "B",
        "limitPx": "0.509",
        "sz": "13203.8",
        "oid": 214381797586,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": true,
        "orderType": "Limit",
        "origSz": "13203.8",
        "tif": "Alo",
        "cloid": "0x71720133819897882521761608579254"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": "0x84761126bf8c93fd85ef042e55063c016400290c5a8fb2cf283ebc797e806de8",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MOODENG",
        "side": "B",
        "limitPx": "0.11845",
        "sz": "167588.0",
        "oid": 214381798920,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "167588.0",
        "tif": "Alo",
        "cloid": "0x0000019a28138fad06de0fed67d8e7dd"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
      "hash": "0x374987c1d91dd67938c3042e55063c0165009fa77410f54bdb1233149811b063",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798921,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x0000000000000000e0f5526e6237dc40"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
      "hash": "0xea1cfe5cf2af18f4eb96042e55063c01660016428da237c68de5a9afb1a2f2df",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MELANIA",
        "side": "A",
        "limitPx": "0.13135",
        "sz": "62833.8",
        "oid": 214381798922,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "62833.8",
        "tif": "Alo",
        "cloid": "0x0000000000000000e0f5526e6237dc3d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": "0x9cf074f80c305b709e6a042e55063c0167008cdda7337a4240b9204acb34355b",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@188",
        "side": "B",
        "limitPx": "0.0047744",
        "sz": "142056.0",
        "oid": 214381798923,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "142056.0",
        "tif": "Alo",
        "cloid": "0x114240e05bd70e7f077237d41d75e5c5"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
      "hash": "0x4fc3eb9325c19dec513d042e55063c0168000378c0c4bcbef38c96e5e4c577d6",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "ARB",
        "side": "A",
        "limitPx": "0.33117",
        "sz": "79872.2",
        "oid": 214381798924,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "79872.2",
        "tif": "Alo",
        "cloid": "0x000000000000000060fe1e47cbf977f8"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
      "hash": "0x0297622e3f52e0680411042e55063c0169007a13da55ff3aa6600d80fe56ba52",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798925,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x0000000000000000ec509f7b0f456bba"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc6ac58a7a63339898aeda32499a8238a46d88e84",
      "hash": "0xb56ad8c958e422e3b6e4042e55063c016a00f0aef3e741b55933841c17e7fcce",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798926,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x0000000000000000a01b604f9b5672b4"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
      "hash": "0x683e4f647275655f69b8042e55063c016b00674a0d7884310c06fab731793f4a",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798927,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x000000000000000060fe1e47cbf977fb"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x72cb918356c4f6d3f1b2e532928110ba6995f139",
      "hash": "0x1b11c5ff8c06a7db1c8b042e55063c016c00dde52709c6adbeda71524b0a81c5",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@156",
        "side": "B",
        "limitPx": "198.7",
        "sz": "75.473",
        "oid": 214381798928,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "75.473",
        "tif": "Alo",
        "cloid": "0x0000000000000000187170386aee5238"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
      "hash": "0xcde53c9aa597ea56cf5e042e55063c016d005480409b092871ade7ed649bc441",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798929,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x00000000000000001cc1cdd8d5c27a51"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": "0x80b8b335bf292cd28232042e55063c016e00cb1b5a2c4ba424815e887e2d06bd",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "ALT",
        "side": "B",
        "limitPx": "0.019569",
        "sz": "1093199.0",
        "oid": 214381798930,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1093199.0",
        "tif": "Alo",
        "cloid": "0x0000019a28138fa804de0fed67d8e7db"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.27613",
        "oid": 214381798931,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa0"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.27613",
        "oid": 214381798932,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa1"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.27613",
        "oid": 214381798933,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa2"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.27613",
        "oid": 214381798934,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa3"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.27613",
        "oid": 214381798935,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa4"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.27613",
        "oid": 214381798936,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa5"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.27613",
        "oid": 214381798937,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa6"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.27613",
        "oid": 214381798938,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa7"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.27613",
        "oid": 214381798939,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa8"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x2cb562384765bea1612ceb60007e315f904fbe86",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.27613",
        "oid": 214381798940,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.27613",
        "tif": "Alo",
        "cloid": "0x00000000000000000006422c8c694aa9"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xe3b6e3443c8f2080704e7421bad9340f13950acb",
      "hash": "0xe65fa06bf24bb1c9e7d9042e55063c017000b8518d4ed09b8a284bbeb14f8bb4",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798941,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x00000000000000002bca2ffe47853d9d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": "0x993317070bdcf4459aac042e55063c0171002eeca6d013173cfbc259cad0ce30",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "@142",
        "side": "B",
        "limitPx": "114169.0",
        "sz": "0.43794",
        "oid": 214381798942,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.43794",
        "tif": "Alo",
        "cloid": "0x1210b26bdab854609be1422a1d2725bc"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa289ee1e56c0d5d041db762e6123e78af0f7d9ad",
      "hash": "0x4c068da2256e36c14d80042e55063c017200a587c0615593efcf38f4e46210ab",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "BTC",
        "side": "B",
        "limitPx": "102759.0",
        "sz": "0.00194",
        "oid": 214381798943,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.00194",
        "tif": "Alo",
        "cloid": "0x00000000000000009c5ad588c506472c"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114164.0",
        "sz": "0.1125",
        "oid": 214381798944,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f989000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114165.0",
        "sz": "0.1125",
        "oid": 214381798945,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f98a000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114166.0",
        "sz": "0.1125",
        "oid": 214381798946,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f98b000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114167.0",
        "sz": "0.1125",
        "oid": 214381798947,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f98c000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114168.0",
        "sz": "0.1125",
        "oid": 214381798948,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f98d000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114169.0",
        "sz": "0.1125",
        "oid": 214381798949,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f98e000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114170.0",
        "sz": "0.1125",
        "oid": 214381798950,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f98f000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114171.0",
        "sz": "0.1125",
        "oid": 214381798951,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f990000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114172.0",
        "sz": "0.1125",
        "oid": 214381798952,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f991000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114173.0",
        "sz": "0.1125",
        "oid": 214381798953,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f992000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114174.0",
        "sz": "0.1125",
        "oid": 214381798954,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f993000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114175.0",
        "sz": "0.1125",
        "oid": 214381798955,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f994000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd47c1e51b759cb3754bb045ae57bc05aadcf58c2",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "BTC",
        "side": "A",
        "limitPx": "114176.0",
        "sz": "0.1125",
        "oid": 214381798956,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.1125",
        "tif": "Alo",
        "cloid": "0x000000000ae2f995000033f682d51c7b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.5",
        "sz": "1.7348",
        "oid": 214381798957,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90168000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.4",
        "sz": "1.7348",
        "oid": 214381798958,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90169000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.3",
        "sz": "1.7348",
        "oid": 214381798959,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9016a000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.2",
        "sz": "1.7348",
        "oid": 214381798960,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9016b000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.1",
        "sz": "1.7348",
        "oid": 214381798961,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9016c000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4125.0",
        "sz": "1.7348",
        "oid": 214381798962,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9016d000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.9",
        "sz": "1.7348",
        "oid": 214381798963,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9016e000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.8",
        "sz": "1.7348",
        "oid": 214381798964,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b9016f000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.7",
        "sz": "1.7348",
        "oid": 214381798965,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90170000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.6",
        "sz": "1.7348",
        "oid": 214381798966,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90171000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6be2564feecd30ed141696dd4cbbb27f2abf570c",
      "hash": null,
      "builder": null,
      "status": "badAloPxRejected",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.5",
        "sz": "1.7348",
        "oid": 214381798967,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1.7348",
        "tif": "Alo",
        "cloid": "0x0000000008b90172000033bb22161c6e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x716d7ae407194c19592036097eac71109709a19b",
      "hash": null,
      "builder": null,
      "status": "perpMarginRejected",
      "order": {
        "coin": "PAXG",
        "side": "B",
        "limitPx": "3949.3",
        "sz": "5.064",
        "oid": 214381798969,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "5.064",
        "tif": "Alo",
        "cloid": "0x00000000000000000ac2d2f70fd9f018"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": "0x1754680e8ba340b018ce042e55063c0176007ff426a65f82bb1d13614aa71a9a",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MELANIA",
        "side": "A",
        "limitPx": "0.13212",
        "sz": "205204.0",
        "oid": 214381798970,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "205204.0",
        "tif": "Alo",
        "cloid": "0x0000019a28138fb304de0fed67d8e7e5"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "@151",
        "side": "A",
        "limitPx": "4126.3",
        "sz": "36.3573",
        "oid": 214381790440,
        "timestamp": 1761608961027,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "36.3573",
        "tif": "Alo",
        "cloid": "0x10715d2a96dab564ee278c04c898640e"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc3f449389734c3f54a162fd30d0495350db61bb4",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "45.281",
        "sz": "5.26",
        "oid": 214381393441,
        "timestamp": 1761608931572,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "5.26",
        "tif": "Gtc",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ENA",
        "side": "A",
        "limitPx": "0.50258",
        "sz": "2201.0",
        "oid": 214381715337,
        "timestamp": 1761608954907,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2201.0",
        "tif": "Alo",
        "cloid": "0x6700d44bf4f1bc13ba82ec789e43ee73"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "KAS",
        "side": "B",
        "limitPx": "0.057335",
        "sz": "377514.0",
        "oid": 214381737671,
        "timestamp": 1761608956606,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "377514.0",
        "tif": "Alo",
        "cloid": "0x0000019a281379fc01de0fed67d8e57a"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "@107",
        "side": "B",
        "limitPx": "46.433",
        "sz": "62.13",
        "oid": 214381786873,
        "timestamp": 1761608960628,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "62.13",
        "tif": "Alo",
        "cloid": "0x1000a8daa4866a95acd50b26f925cb6d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "AAVE",
        "side": "B",
        "limitPx": "235.1",
        "sz": "108.33",
        "oid": 214381759471,
        "timestamp": 1761608958635,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "108.33",
        "tif": "Alo",
        "cloid": "0x00000000000003726922614053233494"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ZRO",
        "side": "A",
        "limitPx": "1.6951",
        "sz": "1192.3",
        "oid": 214381795550,
        "timestamp": 1761608961631,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1192.3",
        "tif": "Alo",
        "cloid": "0x00000000000010698213411120389835"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "OP",
        "side": "B",
        "limitPx": "0.4509",
        "sz": "6589.4",
        "oid": 214381794523,
        "timestamp": 1761608961429,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "6589.4",
        "tif": "Alo",
        "cloid": "0x00000000000010705250285537171115"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.47",
        "sz": "47.48",
        "oid": 214381797666,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "47.48",
        "tif": "Alo",
        "cloid": "0x10001d961e3f73df10eeb5fc1387d10f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ONDO",
        "side": "B",
        "limitPx": "0.74536",
        "sz": "6760.0",
        "oid": 214381764645,
        "timestamp": 1761608959011,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "6760.0",
        "tif": "Alo",
        "cloid": "0x00000000000003730300313759376712"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x01783e91d31c69633153c2aadbf9b7901049d03a",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "WIF",
        "side": "A",
        "limitPx": "0.55275",
        "sz": "704.0",
        "oid": 214381776185,
        "timestamp": 1761608959767,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "704.0",
        "tif": "Gtc",
        "cloid": "0x420a6a84a6fe5a3709936a4c190818bd"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4123.8",
        "sz": "3.6374",
        "oid": 214381793011,
        "timestamp": 1761608961342,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "3.6374",
        "tif": "Alo",
        "cloid": "0x1011c00655d44b97ecf6f89f6d1bf898"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "46.467",
        "sz": "32.28",
        "oid": 214381798444,
        "timestamp": 1761608961980,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "32.28",
        "tif": "Alo",
        "cloid": "0x65909cf9556485adb76583982762f0af"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.502",
        "sz": "32.26",
        "oid": 214381797850,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "32.26",
        "tif": "Alo",
        "cloid": "0x65901d1f5d851a54b8c69b4f5dd60195"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "OM",
        "side": "A",
        "limitPx": "0.1135",
        "sz": "17310.8",
        "oid": 214381710476,
        "timestamp": 1761608954658,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "17310.8",
        "tif": "Alo",
        "cloid": "0x00000000000000007d066f289a798127"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "OM",
        "side": "A",
        "limitPx": "0.11352",
        "sz": "17310.8",
        "oid": 214381669224,
        "timestamp": 1761608951746,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "17310.8",
        "tif": "Alo",
        "cloid": "0x00000000000000007d066f289a7975b7"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x27c9fa86c91b84ddfa15de58c482ff662498d65d",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "@142",
        "side": "A",
        "limitPx": "114263.0",
        "sz": "0.02552",
        "oid": 214381663042,
        "timestamp": 1761608951565,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.02552",
        "tif": "Alo",
        "cloid": "0xb2461f2d84dfc391e942107637bf232b"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x49cba59fce90bfdb2ac76f84b6d5144a84bcb3a8",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ETH",
        "side": "A",
        "limitPx": "4125.6",
        "sz": "9.7494",
        "oid": 214381788456,
        "timestamp": 1761608960760,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "9.7494",
        "tif": "Alo",
        "cloid": "0x20251028000000000000000000027024"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "REZ",
        "side": "A",
        "limitPx": "0.010347",
        "sz": "164015.0",
        "oid": 214381566728,
        "timestamp": 1761608946472,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "164015.0",
        "tif": "Alo",
        "cloid": "0x14441490859189010000000003438082"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "NIL",
        "side": "B",
        "limitPx": "0.29308",
        "sz": "10178.0",
        "oid": 214381795943,
        "timestamp": 1761608961631,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "10178.0",
        "tif": "Alo",
        "cloid": "0x0000019a28138df706de0fed67d8e76d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "NIL",
        "side": "A",
        "limitPx": "0.29449",
        "sz": "10344.0",
        "oid": 214381718429,
        "timestamp": 1761608955094,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "10344.0",
        "tif": "Alo",
        "cloid": "0x00000000000010705250285537171057"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x543356c584e7931233afe07ecd982cda5f0f3f3b",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ETH",
        "side": "B",
        "limitPx": "4124.3",
        "sz": "13.2054",
        "oid": 214381786230,
        "timestamp": 1761608960628,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "13.2054",
        "tif": "Alo",
        "cloid": "0x0000000000000000000000000132fe24"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "NIL",
        "side": "A",
        "limitPx": "0.29409",
        "sz": "2333.0",
        "oid": 214381787271,
        "timestamp": 1761608960628,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2333.0",
        "tif": "Alo",
        "cloid": "0x3a46f6a9f96fe583bcbb81fc5bf25588"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.539",
        "sz": "120.0",
        "oid": 214381776556,
        "timestamp": 1761608959842,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "120.0",
        "tif": "Alo",
        "cloid": "0x020d0018fe4a30400000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.544",
        "sz": "120.0",
        "oid": 214381741728,
        "timestamp": 1761608957218,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "120.0",
        "tif": "Alo",
        "cloid": "0x020d0018f14a30400000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc926ddba8b7617dbc65712f20cf8e1b58b8598d3",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.549",
        "sz": "120.0",
        "oid": 214381686698,
        "timestamp": 1761608952550,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "120.0",
        "tif": "Alo",
        "cloid": "0x020d0018d9ca30400000000000000000"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "PURR/USDC",
        "side": "B",
        "limitPx": "0.13574",
        "sz": "2051.0",
        "oid": 214381796920,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2051.0",
        "tif": "Alo",
        "cloid": "0x10108c15eb380e781fc148fa201296cd"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "PURR/USDC",
        "side": "A",
        "limitPx": "0.13619",
        "sz": "3530.0",
        "oid": 214381797447,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "3530.0",
        "tif": "Alo",
        "cloid": "0x1011033d92830790f3e7fd95b458c66f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "PURR/USDC",
        "side": "B",
        "limitPx": "0.13553",
        "sz": "4622.0",
        "oid": 214381797448,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "4622.0",
        "tif": "Alo",
        "cloid": "0x10115083afd4d4b3faab2ee2b9a1dd77"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "LINK",
        "side": "A",
        "limitPx": "18.24",
        "sz": "1172.7",
        "oid": 214381750355,
        "timestamp": 1761608958191,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1172.7",
        "tif": "Alo",
        "cloid": "0x0000019a28137f9001de0fed67d8e5e4"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "REZ",
        "side": "A",
        "limitPx": "0.010219",
        "sz": "1807101.0",
        "oid": 214381569395,
        "timestamp": 1761608946541,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1807101.0",
        "tif": "Alo",
        "cloid": "0x0000019a2813519e01de0fed67d8de81"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "REZ",
        "side": "B",
        "limitPx": "0.010135",
        "sz": "2039664.0",
        "oid": 214381649416,
        "timestamp": 1761608951004,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2039664.0",
        "tif": "Alo",
        "cloid": "0x0000019a281363f003de0fed67d8e16f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf5c48df6137af44ab93ee34d1d94db206547bc78",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.597",
        "sz": "19.57",
        "oid": 214381794019,
        "timestamp": 1761608961429,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "19.57",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf5c48df6137af44ab93ee34d1d94db206547bc78",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.615",
        "sz": "19.57",
        "oid": 214381794020,
        "timestamp": 1761608961429,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "19.57",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf5c48df6137af44ab93ee34d1d94db206547bc78",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "@107",
        "side": "A",
        "limitPx": "46.634",
        "sz": "19.57",
        "oid": 214381794021,
        "timestamp": 1761608961429,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "19.57",
        "tif": "Alo",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "LINK",
        "side": "B",
        "limitPx": "18.205",
        "sz": "1155.3",
        "oid": 214381754577,
        "timestamp": 1761608958473,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1155.3",
        "tif": "Alo",
        "cloid": "0x0000019a28137f5a03de0fed67d8e5d4"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "LINK",
        "side": "A",
        "limitPx": "18.246",
        "sz": "1144.1",
        "oid": 214381750864,
        "timestamp": 1761608958191,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1144.1",
        "tif": "Alo",
        "cloid": "0x0000019a2813805500de0fed67d8e600"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ZEC",
        "side": "A",
        "limitPx": "341.41",
        "sz": "2.96",
        "oid": 214381785226,
        "timestamp": 1761608960425,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2.96",
        "tif": "Alo",
        "cloid": "0x00000000000010698213411120389834"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "REZ",
        "side": "B",
        "limitPx": "0.010149",
        "sz": "685704.0",
        "oid": 214381614995,
        "timestamp": 1761608949226,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "685704.0",
        "tif": "Alo",
        "cloid": "0x0000019a28135d0d05de0fed67d8e036"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "LINK",
        "side": "B",
        "limitPx": "18.221",
        "sz": "55.1",
        "oid": 214381797449,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "55.1",
        "tif": "Alo",
        "cloid": "0x00000000000003733959488471475747"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "DOGE",
        "side": "B",
        "limitPx": "0.20039",
        "sz": "124912.0",
        "oid": 214381778949,
        "timestamp": 1761608960021,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "124912.0",
        "tif": "Alo",
        "cloid": "0x00000000000003728329988933759295"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6da22bdd4ed69704c4912ced8ba988e93e6b0a83",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "MAVIA",
        "side": "A",
        "limitPx": "0.24762",
        "sz": "331.4",
        "oid": 214381789713,
        "timestamp": 1761608961027,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "331.4",
        "tif": "Alo",
        "cloid": "0x4c69873c59274267aaa73e87967c5cad"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ZEC",
        "side": "B",
        "limitPx": "341.31",
        "sz": "1.68",
        "oid": 214381791997,
        "timestamp": 1761608961246,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2.92",
        "tif": "Alo",
        "cloid": "0x8050b1316577ac5143e0077d97527b2f"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "ASTER",
        "side": "B",
        "limitPx": "1.0808",
        "sz": "3980.0",
        "oid": 214381770371,
        "timestamp": 1761608959417,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "3980.0",
        "tif": "Alo",
        "cloid": "0x00000000000012180179158673789177"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x365e0c115f1ca1adcb42fd21142873493df7f880",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "DOT",
        "side": "A",
        "limitPx": "3.1412",
        "sz": "369.5",
        "oid": 214381648057,
        "timestamp": 1761608950929,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "369.5",
        "tif": "Alo",
        "cloid": "0x0000000000000000000eadf34359b4c0"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "46.472",
        "sz": "21.52",
        "oid": 214381797776,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "21.52",
        "tif": "Alo",
        "cloid": "0xa68f34bf783698c0dd5de93c93d19de9"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x6a0b567be46b51ad02db4ebdf90d057ecb02f864",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "MAVIA",
        "side": "B",
        "limitPx": "0.24744",
        "sz": "222.2",
        "oid": 214381797443,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "222.2",
        "tif": "Alo",
        "cloid": "0x20251028000000000000000000402567"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x57dd78cd36e76e2011e8f6dc25cabbaba994494b",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "PUMP",
        "side": "A",
        "limitPx": "0.00481",
        "sz": "11317540.0",
        "oid": 214381356432,
        "timestamp": 1761608928560,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "11317540.0",
        "tif": "Alo",
        "cloid": "0x00000000000000104562302094015990"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "PEOPLE",
        "side": "A",
        "limitPx": "0.012544",
        "sz": "561952.0",
        "oid": 214381709218,
        "timestamp": 1761608954588,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "561952.0",
        "tif": "Alo",
        "cloid": "0x0000019a281371e100de0fed67d8e478"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xd77a0944d0039e92f3309f9cf52b6926b7c14b6d",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "PUMP",
        "side": "B",
        "limitPx": "0.00477",
        "sz": "437583.0",
        "oid": 214381358805,
        "timestamp": 1761608928760,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "437583.0",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000004572de3"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xa289ee1e56c0d5d041db762e6123e78af0f7d9ad",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "MAVIA",
        "side": "B",
        "limitPx": "0.24727",
        "sz": "2529.0",
        "oid": 214381500777,
        "timestamp": 1761608940937,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "2529.0",
        "tif": "Alo",
        "cloid": "0x00000000000000009c5ad588c505f4b8"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "MELANIA",
        "side": "A",
        "limitPx": "0.13308",
        "sz": "62833.8",
        "oid": 214381315882,
        "timestamp": 1761608924537,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "62833.8",
        "tif": "Alo",
        "cloid": "0x0000000000000000e0f5526e6237493d"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "MELANIA",
        "side": "A",
        "limitPx": "0.13151",
        "sz": "3802.0",
        "oid": 214381769788,
        "timestamp": 1761608959417,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "3802.0",
        "tif": "Alo",
        "cloid": "0x71711c27f8cbd8fccb20b13f3810e351"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "WIF",
        "side": "A",
        "limitPx": "0.55228",
        "sz": "1493.0",
        "oid": 214381796783,
        "timestamp": 1761608961833,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "1493.0",
        "tif": "Alo",
        "cloid": "0xf3ce430381f99e53b61253b9b2518794"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7b4da0eb8dcca7a9125780620172f60cc40fb4e1",
      "hash": "0x93c28b5055858260953c042e55063c01c400a335f088a132378b36a314895c4b",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "47.52",
        "sz": "116.72",
        "oid": 214381798971,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "116.72",
        "tif": "Gtc",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x4fb31d5ffaa08be2128cb14485584d76eebef173",
      "hash": null,
      "builder": null,
      "status": "filled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.473",
        "sz": "0.0",
        "oid": 214381798720,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "85.67",
        "tif": "Alo",
        "cloid": "0x00000000000000000000000115002ba0"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x20faffd0045455e86b36b827d68f8f99f9c929b4",
      "hash": null,
      "builder": null,
      "status": "filled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.473",
        "sz": "0.0",
        "oid": 214381798785,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "4.44",
        "tif": "Alo",
        "cloid": "0x874fb12c53934f5c8ad09a59a00fe1ab"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x162cc7c861ebd0c06b3d72319201150482518185",
      "hash": null,
      "builder": null,
      "status": "filled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.483",
        "sz": "0.0",
        "oid": 214381777986,
        "timestamp": 1761608960021,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "0.43",
        "tif": "Alo",
        "cloid": "0x00000000000000000b56a216233de146"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xae7a0f9f663bb54bfe95378c7e3de7dd6e28d6bc",
      "hash": null,
      "builder": null,
      "status": "filled",
      "order": {
        "coin": "HYPE",
        "side": "A",
        "limitPx": "46.483",
        "sz": "0.0",
        "oid": 214381795808,
        "timestamp": 1761608961631,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "10.75",
        "tif": "Alo",
        "cloid": "0x0014870063923c0000019a28138dd226"
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0x7b4da0eb8dcca7a9125780620172f60cc40fb4e1",
      "hash": null,
      "builder": null,
      "status": "filled",
      "order": {
        "coin": "HYPE",
        "side": "B",
        "limitPx": "47.52",
        "sz": "0.0",
        "oid": 214381798971,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "116.72",
        "tif": "Gtc",
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc7468a9aee4b4623d298a9cf6d43241e1a52d0be",
      "hash": "0x469601eb6f16c4dc480f042e55063c01c50019d10a19e3aeea5ead3e2e1a9ec6",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "MELANIA",
        "side": "A",
        "limitPx": "0.1205",
        "sz": "609.4",
        "oid": 214381798972,
        "timestamp": 1761608962031,
        "triggerCondition": "Price below 0.13098",
        "isTrigger": true,
        "triggerPx": "0.13098",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": true,
        "orderType": "Stop Market",
        "origSz": "609.4",
        "tif": null,
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xc7468a9aee4b4623d298a9cf6d43241e1a52d0be",
      "hash": null,
      "builder": null,
      "status": "canceled",
      "order": {
        "coin": "MELANIA",
        "side": "A",
        "limitPx": "0.11964",
        "sz": "609.4",
        "oid": 214381636190,
        "timestamp": 1761608950242,
        "triggerCondition": "Price below 0.13005",
        "isTrigger": true,
        "triggerPx": "0.13005",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": true,
        "orderType": "Stop Market",
        "origSz": "609.4",
        "tif": null,
        "cloid": null
      }
    },
    {
      "time": "2025-10-27T23:49:22.031892297",
      "user": "0xba6de73b64eca3ff95c4ab7e5328912c360bd003",
      "hash": "0xf969788688a80757fae3042e55063c01c600906c23ab262a9d3223d947abe142",
      "builder": null,
      "status": "open",
      "order": {
        "coin": "CRV",
        "side": "A",
        "limitPx": "0.55809",
        "sz": "696.9",
        "oid": 214381798973,
        "timestamp": 1761608962031,
        "triggerCondition": "N/A",
        "isTrigger": false,
        "triggerPx": "0.0",
        "children": [],
        "isPositionTpsl": false,
        "reduceOnly": false,
        "orderType": "Limit",
        "origSz": "696.9",
        "tif": "Gtc",
        "cloid": "0x1bd17e93d2f7a367355b991ea50f9723"
      }
    }
  ]
}

`

var blockOrderBookDiffSample = `
{
    "local_time": "2025-12-01T12:59:40.285491771",
    "block_time": "2025-12-01T12:59:40.101046556",
    "block_number": 814448093,
    "events": [
        {
            "user": "0xbda495ebf6fc94d61f9a7f22b9877273378ed82b",
            "oid": 254433008712,
            "coin": "APT",
            "side": "B",
            "px": "1.8549",
            "raw_book_diff": {
                "new": {
                    "sz": "1407.89"
                }
            }
        },
        {
            "user": "0x34827044cbd4b808fc1b189fce9f50e6dafae7c9",
            "oid": 254433008713,
            "coin": "BIGTIME",
            "side": "B",
            "px": "0.021663",
            "raw_book_diff": {
                "new": {
                    "sz": "22510.0"
                }
            }
        },
        {
            "user": "0x6628f829994f1756257e6c2468dabd675a530a1d",
            "oid": 254433008714,
            "coin": "AVAX",
            "side": "A",
            "px": "12.921",
            "raw_book_diff": {
                "new": {
                    "sz": "29.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008715,
            "coin": "ZK",
            "side": "B",
            "px": "0.033829",
            "raw_book_diff": {
                "new": {
                    "sz": "16643.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008716,
            "coin": "ZK",
            "side": "B",
            "px": "0.033817",
            "raw_book_diff": {
                "new": {
                    "sz": "33140.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008717,
            "coin": "ZK",
            "side": "B",
            "px": "0.033801",
            "raw_book_diff": {
                "new": {
                    "sz": "35239.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008718,
            "coin": "ZK",
            "side": "A",
            "px": "0.033856",
            "raw_book_diff": {
                "new": {
                    "sz": "50465.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008719,
            "coin": "ZK",
            "side": "A",
            "px": "0.033871",
            "raw_book_diff": {
                "new": {
                    "sz": "49124.0"
                }
            }
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433008721,
            "coin": "APT",
            "side": "A",
            "px": "1.8676",
            "raw_book_diff": {
                "new": {
                    "sz": "674.0"
                }
            }
        },
        {
            "user": "0xfcf104006bfff47695c1dc21dad3e9de1e72098e",
            "oid": 254433008722,
            "coin": "BTC",
            "side": "A",
            "px": "85423.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.13919"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008728,
            "coin": "AVNT",
            "side": "A",
            "px": "0.33532",
            "raw_book_diff": {
                "new": {
                    "sz": "2008.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008729,
            "coin": "PNUT",
            "side": "A",
            "px": "0.0808",
            "raw_book_diff": {
                "new": {
                    "sz": "9429.9"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008730,
            "coin": "MAVIA",
            "side": "B",
            "px": "0.04734",
            "raw_book_diff": {
                "new": {
                    "sz": "4916.7"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008731,
            "coin": "MAVIA",
            "side": "A",
            "px": "0.0474",
            "raw_book_diff": {
                "new": {
                    "sz": "6818.6"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008732,
            "coin": "MAVIA",
            "side": "A",
            "px": "0.04743",
            "raw_book_diff": {
                "new": {
                    "sz": "13813.8"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008733,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33973",
            "raw_book_diff": {
                "new": {
                    "sz": "1640.5"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008734,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33961",
            "raw_book_diff": {
                "new": {
                    "sz": "3217.5"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008735,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33944",
            "raw_book_diff": {
                "new": {
                    "sz": "2896.7"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008736,
            "coin": "SUSHI",
            "side": "A",
            "px": "0.3399",
            "raw_book_diff": {
                "new": {
                    "sz": "1796.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008737,
            "coin": "SUSHI",
            "side": "A",
            "px": "0.34001",
            "raw_book_diff": {
                "new": {
                    "sz": "3433.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008738,
            "coin": "USUAL",
            "side": "B",
            "px": "0.02466",
            "raw_book_diff": {
                "new": {
                    "sz": "25341.2"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008739,
            "coin": "USUAL",
            "side": "A",
            "px": "0.02478",
            "raw_book_diff": {
                "new": {
                    "sz": "11687.4"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008740,
            "coin": "EIGEN",
            "side": "B",
            "px": "0.5045",
            "raw_book_diff": {
                "new": {
                    "sz": "2579.7"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008741,
            "coin": "EIGEN",
            "side": "A",
            "px": "0.505",
            "raw_book_diff": {
                "new": {
                    "sz": "1586.54"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008742,
            "coin": "EIGEN",
            "side": "A",
            "px": "0.5051",
            "raw_book_diff": {
                "new": {
                    "sz": "2971.84"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008743,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.076455",
            "raw_book_diff": {
                "new": {
                    "sz": "3131.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008744,
            "coin": "RESOLV",
            "side": "A",
            "px": "0.077867",
            "raw_book_diff": {
                "new": {
                    "sz": "6535.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008745,
            "coin": "ZEREBRO",
            "side": "B",
            "px": "0.027716",
            "raw_book_diff": {
                "new": {
                    "sz": "5682.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433008746,
            "coin": "ZEREBRO",
            "side": "A",
            "px": "0.028282",
            "raw_book_diff": {
                "new": {
                    "sz": "2794.0"
                }
            }
        },
        {
            "user": "0x14231bf8d1a5c2b9efa36d1c1f780336a3edc719",
            "oid": 254433008748,
            "coin": "ETH",
            "side": "A",
            "px": "2822.1",
            "raw_book_diff": {
                "new": {
                    "sz": "2.7031"
                }
            }
        },
        {
            "user": "0xfc8c156428a8e48cb8d0356db16e59bec4c0ecea",
            "oid": 254433008749,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.042001",
            "raw_book_diff": {
                "new": {
                    "sz": "221274.0"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008754,
            "coin": "@151",
            "side": "B",
            "px": "2822.0",
            "raw_book_diff": {
                "new": {
                    "sz": "17.7179"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008755,
            "coin": "@151",
            "side": "B",
            "px": "2821.7",
            "raw_book_diff": {
                "new": {
                    "sz": "17.7198"
                }
            }
        },
        {
            "user": "0xbf1935fe7ab6d0aa3ee8d3da47c2f80e215b2a1c",
            "oid": 254433008756,
            "coin": "ETH",
            "side": "A",
            "px": "2824.3",
            "raw_book_diff": {
                "new": {
                    "sz": "0.0071"
                }
            }
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433008757,
            "coin": "APT",
            "side": "B",
            "px": "1.8564",
            "raw_book_diff": {
                "new": {
                    "sz": "1555.69"
                }
            }
        },
        {
            "user": "0xa75f44732595c1e86c0fea738fc5489cc71726ca",
            "oid": 254433005731,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.076156",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa75f44732595c1e86c0fea738fc5489cc71726ca",
            "oid": 254433008761,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.076326",
            "raw_book_diff": {
                "new": {
                    "sz": "110162.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433008763,
            "coin": "ZEREBRO",
            "side": "A",
            "px": "0.028238",
            "raw_book_diff": {
                "new": {
                    "sz": "25000.0"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008783,
            "coin": "PYTH",
            "side": "A",
            "px": "0.067765",
            "raw_book_diff": {
                "new": {
                    "sz": "3690.0"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008784,
            "coin": "PYTH",
            "side": "A",
            "px": "0.067868",
            "raw_book_diff": {
                "new": {
                    "sz": "14738.0"
                }
            }
        },
        {
            "user": "0x365e0c115f1ca1adcb42fd21142873493df7f880",
            "oid": 254433008785,
            "coin": "DOT",
            "side": "B",
            "px": "2.0259",
            "raw_book_diff": {
                "new": {
                    "sz": "684.9"
                }
            }
        },
        {
            "user": "0x365e0c115f1ca1adcb42fd21142873493df7f880",
            "oid": 254433008786,
            "coin": "DOT",
            "side": "A",
            "px": "2.0309",
            "raw_book_diff": {
                "new": {
                    "sz": "517.9"
                }
            }
        },
        {
            "user": "0x365e0c115f1ca1adcb42fd21142873493df7f880",
            "oid": 254433008788,
            "coin": "BTC",
            "side": "A",
            "px": "85490.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.03351"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008789,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13609",
            "raw_book_diff": {
                "new": {
                    "sz": "14783.0"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008790,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13603",
            "raw_book_diff": {
                "new": {
                    "sz": "16923.0"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008791,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13599",
            "raw_book_diff": {
                "new": {
                    "sz": "37149.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433008792,
            "coin": "GRIFFAIN",
            "side": "A",
            "px": "0.019777",
            "raw_book_diff": {
                "new": {
                    "sz": "34577.0"
                }
            }
        },
        {
            "user": "0x704a8267e9bcd42449909298b60ec05ff1f50865",
            "oid": 254433005667,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.075881",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x704a8267e9bcd42449909298b60ec05ff1f50865",
            "oid": 254433008793,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.076072",
            "raw_book_diff": {
                "new": {
                    "sz": "126646.0"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433008795,
            "coin": "JTO",
            "side": "B",
            "px": "0.41721",
            "raw_book_diff": {
                "new": {
                    "sz": "28976.0"
                }
            }
        },
        {
            "user": "0xa775d1bd91ee2cc4cae6113e5fd9c0c17b355277",
            "oid": 254433008797,
            "coin": "XRP",
            "side": "B",
            "px": "2.0265",
            "raw_book_diff": {
                "new": {
                    "sz": "1988.0"
                }
            }
        },
        {
            "user": "0xe29a71644928e2a63312fba2556536f3b6aa220b",
            "oid": 254433005631,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.075879",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe29a71644928e2a63312fba2556536f3b6aa220b",
            "oid": 254433008798,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.07607",
            "raw_book_diff": {
                "new": {
                    "sz": "128394.0"
                }
            }
        },
        {
            "user": "0xdab4bc9946e470fd56f7c1a5d79854cba2c447c6",
            "oid": 254433008799,
            "coin": "HYPE",
            "side": "B",
            "px": "30.214",
            "raw_book_diff": {
                "new": {
                    "sz": "33.94"
                }
            }
        },
        {
            "user": "0xc2e137ea6b79a7992dd54d44ee8223f2c7e13e22",
            "oid": 254433008805,
            "coin": "ZEC",
            "side": "B",
            "px": "348.15",
            "raw_book_diff": {
                "new": {
                    "sz": "0.51"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433008806,
            "coin": "IP",
            "side": "B",
            "px": "2.2605",
            "raw_book_diff": {
                "new": {
                    "sz": "1706.4"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008807,
            "coin": "ZEC",
            "side": "A",
            "px": "348.3",
            "raw_book_diff": {
                "new": {
                    "sz": "4.3"
                }
            }
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008808,
            "coin": "ZEC",
            "side": "A",
            "px": "348.36",
            "raw_book_diff": {
                "new": {
                    "sz": "4.3"
                }
            }
        },
        {
            "user": "0x9e7cdb495b4864448a89bdd959e0fc070e47c621",
            "oid": 254433008810,
            "coin": "MON",
            "side": "A",
            "px": "0.023601",
            "raw_book_diff": {
                "new": {
                    "sz": "8480.0"
                }
            }
        },
        {
            "user": "0x352deb23bebae8b4c57d0ae341d9c1951fd8425a",
            "oid": 254433008812,
            "coin": "xyz:AMZN",
            "side": "A",
            "px": "233.65",
            "raw_book_diff": {
                "new": {
                    "sz": "4.323"
                }
            }
        },
        {
            "user": "0x9e7cdb495b4864448a89bdd959e0fc070e47c621",
            "oid": 254433008813,
            "coin": "MON",
            "side": "B",
            "px": "0.023552",
            "raw_book_diff": {
                "new": {
                    "sz": "8480.0"
                }
            }
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254433008814,
            "coin": "ETH",
            "side": "A",
            "px": "2822.9",
            "raw_book_diff": {
                "new": {
                    "sz": "7.4273"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433008821,
            "coin": "BOME",
            "side": "B",
            "px": "0.000645",
            "raw_book_diff": {
                "new": {
                    "sz": "19141784.0"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433005923,
            "coin": "ENA",
            "side": "B",
            "px": "0.24104",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433008822,
            "coin": "ENA",
            "side": "B",
            "px": "0.24176",
            "raw_book_diff": {
                "new": {
                    "sz": "3698.0"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433005924,
            "coin": "ETH",
            "side": "B",
            "px": "2815.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433008823,
            "coin": "ETH",
            "side": "B",
            "px": "2820.1",
            "raw_book_diff": {
                "new": {
                    "sz": "1.0402"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433005925,
            "coin": "LTC",
            "side": "B",
            "px": "76.676",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433008824,
            "coin": "LTC",
            "side": "B",
            "px": "76.843",
            "raw_book_diff": {
                "new": {
                    "sz": "19.44"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254432958159,
            "coin": "CRV",
            "side": "A",
            "px": "0.38321",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433008825,
            "coin": "CRV",
            "side": "A",
            "px": "0.38299",
            "raw_book_diff": {
                "new": {
                    "sz": "3044.0"
                }
            }
        },
        {
            "user": "0xb30bda7a381f676e7ccbc22c22c9e594a10d2ead",
            "oid": 254433006007,
            "coin": "IP",
            "side": "A",
            "px": "2.2805",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb30bda7a381f676e7ccbc22c22c9e594a10d2ead",
            "oid": 254433008826,
            "coin": "IP",
            "side": "A",
            "px": "2.2865",
            "raw_book_diff": {
                "new": {
                    "sz": "4410.1"
                }
            }
        },
        {
            "user": "0xfc8c156428a8e48cb8d0356db16e59bec4c0ecea",
            "oid": 254433008827,
            "coin": "ME",
            "side": "A",
            "px": "0.30305",
            "raw_book_diff": {
                "new": {
                    "sz": "17404.2"
                }
            }
        },
        {
            "user": "0xb1d488187aea94e8bdb3f210a5c32626afe14c5a",
            "oid": 254433008828,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.56",
            "raw_book_diff": {
                "new": {
                    "sz": "1.803"
                }
            }
        },
        {
            "user": "0xfc667adba8d4837586078f4fdcdc29804337ca06",
            "oid": 254433008830,
            "coin": "BTC",
            "side": "A",
            "px": "85431.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.1"
                }
            }
        },
        {
            "user": "0x7a9d45682babd8a2d79181623c75a7f615990b27",
            "oid": 254433008831,
            "coin": "SNX",
            "side": "A",
            "px": "0.51074",
            "raw_book_diff": {
                "new": {
                    "sz": "400.0"
                }
            }
        },
        {
            "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
            "oid": 254433008832,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13608",
            "raw_book_diff": {
                "new": {
                    "sz": "3566.0"
                }
            }
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254433008833,
            "coin": "ETH",
            "side": "B",
            "px": "2820.9",
            "raw_book_diff": {
                "new": {
                    "sz": "3.6"
                }
            }
        },
        {
            "user": "0x31b1ac74a39fa92604c3bccac20453f2e3eebfaa",
            "oid": 254433008856,
            "coin": "KAS",
            "side": "B",
            "px": "0.051291",
            "raw_book_diff": {
                "new": {
                    "sz": "30840.0"
                }
            }
        },
        {
            "user": "0xef5fdcefdb38803a654ee9f3539bed8058443d99",
            "oid": 254433008857,
            "coin": "XRP",
            "side": "A",
            "px": "2.028",
            "raw_book_diff": {
                "new": {
                    "sz": "192.0"
                }
            }
        },
        {
            "user": "0x9266865bb6afb4c4f618544dd3b8c970f17aa664",
            "oid": 254433008872,
            "coin": "COMP",
            "side": "A",
            "px": "34.84",
            "raw_book_diff": {
                "new": {
                    "sz": "20.8"
                }
            }
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254433008873,
            "coin": "ETH",
            "side": "B",
            "px": "2820.6",
            "raw_book_diff": {
                "new": {
                    "sz": "9.0"
                }
            }
        },
        {
            "user": "0xc42a4e4bcfb9e94139cb92d7c676c27244a1238c",
            "oid": 254433008884,
            "coin": "xyz:PLTR",
            "side": "B",
            "px": "165.24",
            "raw_book_diff": {
                "new": {
                    "sz": "11.477"
                }
            }
        },
        {
            "user": "0xc1fce740d83a60de67d039aa927a678ff78c202f",
            "oid": 254433008886,
            "coin": "BCH",
            "side": "B",
            "px": "521.0",
            "raw_book_diff": {
                "new": {
                    "sz": "2.0"
                }
            }
        },
        {
            "user": "0xcacf02ccefefe60f10762556ab996587d4f08f67",
            "oid": 254433008897,
            "coin": "RESOLV",
            "side": "A",
            "px": "0.076749",
            "raw_book_diff": {
                "new": {
                    "sz": "894.0"
                }
            }
        },
        {
            "user": "0x9f052eabbe59d1d6e324232a73ef7e1faa66bd54",
            "oid": 254433008898,
            "coin": "CELO",
            "side": "A",
            "px": "0.15374",
            "raw_book_diff": {
                "new": {
                    "sz": "738.0"
                }
            }
        },
        {
            "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
            "oid": 254433008899,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13608",
            "raw_book_diff": {
                "new": {
                    "sz": "6896.0"
                }
            }
        },
        {
            "user": "0x4334ff1dd9fa46667b718a6f8a489c3d0079b44b",
            "oid": 254433008900,
            "coin": "GALA",
            "side": "B",
            "px": "0.006798",
            "raw_book_diff": {
                "new": {
                    "sz": "7532.0"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008906,
            "coin": "xyz:XYZ100",
            "side": "B",
            "px": "25101.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.992"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008907,
            "coin": "xyz:XYZ100",
            "side": "A",
            "px": "25252.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.99"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008908,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "277.1",
            "raw_book_diff": {
                "new": {
                    "sz": "24.059"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008909,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "277.1",
            "raw_book_diff": {
                "new": {
                    "sz": "24.059"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008910,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "276.8",
            "raw_book_diff": {
                "new": {
                    "sz": "24.085"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008911,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.7",
            "raw_book_diff": {
                "new": {
                    "sz": "24.007"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008912,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.7",
            "raw_book_diff": {
                "new": {
                    "sz": "24.007"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008913,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "278.1",
            "raw_book_diff": {
                "new": {
                    "sz": "23.972"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008914,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "277.0",
            "raw_book_diff": {
                "new": {
                    "sz": "120.337"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008915,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.7",
            "raw_book_diff": {
                "new": {
                    "sz": "120.034"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008916,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.8",
            "raw_book_diff": {
                "new": {
                    "sz": "119.99"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008917,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "273.0",
            "raw_book_diff": {
                "new": {
                    "sz": "183.15"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008918,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "280.4",
            "raw_book_diff": {
                "new": {
                    "sz": "178.317"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433008919,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "285.09",
            "raw_book_diff": {
                "new": {
                    "sz": "350.766"
                }
            }
        },
        {
            "user": "0x9e5c7a2fed36727acf6306e3e5ae12bce86856c7",
            "oid": 254433008920,
            "coin": "ALT",
            "side": "A",
            "px": "0.012523",
            "raw_book_diff": {
                "new": {
                    "sz": "99816.0"
                }
            }
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433008927,
            "coin": "BERA",
            "side": "B",
            "px": "0.94342",
            "raw_book_diff": {
                "new": {
                    "sz": "2160.5"
                }
            }
        },
        {
            "user": "0x49208e12b02effd27afe01a78ea871fd73be2b12",
            "oid": 254433008932,
            "coin": "xyz:XYZ100",
            "side": "B",
            "px": "25175.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.0004"
                }
            }
        },
        {
            "user": "0x1d2a3b568e82678f22df040485b8cc006f3d0ea6",
            "oid": 254433008946,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13617",
            "raw_book_diff": {
                "new": {
                    "sz": "7344.0"
                }
            }
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433008947,
            "coin": "ASTER",
            "side": "B",
            "px": "0.93948",
            "raw_book_diff": {
                "new": {
                    "sz": "3122.0"
                }
            }
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433008948,
            "coin": "STRK",
            "side": "A",
            "px": "0.11795",
            "raw_book_diff": {
                "new": {
                    "sz": "1040.0"
                }
            }
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433008955,
            "coin": "XLM",
            "side": "A",
            "px": "0.23116",
            "raw_book_diff": {
                "new": {
                    "sz": "2773.0"
                }
            }
        },
        {
            "user": "0xfc8c156428a8e48cb8d0356db16e59bec4c0ecea",
            "oid": 254433008956,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6236",
            "raw_book_diff": {
                "new": {
                    "sz": "6201.1"
                }
            }
        },
        {
            "user": "0x335f45392f8d87745aaae68f5c192849afd9b60e",
            "oid": 254433008957,
            "coin": "BTC",
            "side": "A",
            "px": "85423.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.70246"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433008958,
            "coin": "ASTER",
            "side": "A",
            "px": "0.94107",
            "raw_book_diff": {
                "new": {
                    "sz": "500.0"
                }
            }
        },
        {
            "user": "0xeaa123246b396028fc8fea0175f013c13e097157",
            "oid": 254433008959,
            "coin": "KAS",
            "side": "A",
            "px": "0.05153",
            "raw_book_diff": {
                "new": {
                    "sz": "16704.0"
                }
            }
        },
        {
            "user": "0xb730e6f15b5d89254db2e84cf0a7fc3bcc2129cc",
            "oid": 254433000180,
            "coin": "MEGA",
            "side": "A",
            "px": "0.33321",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb730e6f15b5d89254db2e84cf0a7fc3bcc2129cc",
            "oid": 254433008960,
            "coin": "MEGA",
            "side": "A",
            "px": "0.33318",
            "raw_book_diff": {
                "new": {
                    "sz": "221.0"
                }
            }
        },
        {
            "user": "0xa242b590f63941749be337eed817af40b61fdd7c",
            "oid": 254432998358,
            "coin": "xyz:NVDA",
            "side": "A",
            "px": "173.51",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa242b590f63941749be337eed817af40b61fdd7c",
            "oid": 254433008963,
            "coin": "xyz:NVDA",
            "side": "A",
            "px": "173.53",
            "raw_book_diff": {
                "new": {
                    "sz": "14.414"
                }
            }
        },
        {
            "user": "0x49208e12b02effd27afe01a78ea871fd73be2b12",
            "oid": 254433008970,
            "coin": "xyz:XYZ100",
            "side": "B",
            "px": "25176.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.0285"
                }
            }
        },
        {
            "user": "0x8ae62e9a4775350022c1577b2eff694e2b72be8d",
            "oid": 254433008975,
            "coin": "ETH",
            "side": "B",
            "px": "2821.1",
            "raw_book_diff": {
                "new": {
                    "sz": "0.0354"
                }
            }
        },
        {
            "user": "0xeaa123246b396028fc8fea0175f013c13e097157",
            "oid": 254433008977,
            "coin": "ZK",
            "side": "A",
            "px": "0.033862",
            "raw_book_diff": {
                "new": {
                    "sz": "21969.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008978,
            "coin": "XRP",
            "side": "A",
            "px": "2.0276",
            "raw_book_diff": {
                "new": {
                    "sz": "247.0"
                }
            }
        },
        {
            "user": "0x727956612a8700627451204a3ae26268bd1a1525",
            "oid": 254433008979,
            "coin": "FIL",
            "side": "A",
            "px": "1.4526",
            "raw_book_diff": {
                "new": {
                    "sz": "341.8"
                }
            }
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254433008991,
            "coin": "SUI",
            "side": "A",
            "px": "1.3513",
            "raw_book_diff": {
                "new": {
                    "sz": "555.2"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008993,
            "coin": "SOL",
            "side": "B",
            "px": "126.29",
            "raw_book_diff": {
                "new": {
                    "sz": "23.71"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008994,
            "coin": "SUI",
            "side": "B",
            "px": "1.3502",
            "raw_book_diff": {
                "new": {
                    "sz": "1480.9"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008995,
            "coin": "SUI",
            "side": "B",
            "px": "1.3503",
            "raw_book_diff": {
                "new": {
                    "sz": "1480.9"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008996,
            "coin": "SUI",
            "side": "A",
            "px": "1.3509",
            "raw_book_diff": {
                "new": {
                    "sz": "1480.9"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008997,
            "coin": "SUI",
            "side": "B",
            "px": "1.3501",
            "raw_book_diff": {
                "new": {
                    "sz": "1480.9"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009008,
            "coin": "SOL",
            "side": "A",
            "px": "126.8",
            "raw_book_diff": {
                "new": {
                    "sz": "23.71"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009009,
            "coin": "SOL",
            "side": "B",
            "px": "126.41",
            "raw_book_diff": {
                "new": {
                    "sz": "15.8"
                }
            }
        },
        {
            "user": "0x543356c584e7931233afe07ecd982cda5f0f3f3b",
            "oid": 254433009010,
            "coin": "SOL",
            "side": "A",
            "px": "126.56",
            "raw_book_diff": {
                "new": {
                    "sz": "157.97"
                }
            }
        },
        {
            "user": "0x727956612a8700627451204a3ae26268bd1a1525",
            "oid": 254433009011,
            "coin": "SUI",
            "side": "A",
            "px": "1.3512",
            "raw_book_diff": {
                "new": {
                    "sz": "996.1"
                }
            }
        },
        {
            "user": "0xb2b1d1a7a034d9209a627e754ebb8d24988ca23a",
            "oid": 254433009012,
            "coin": "FET",
            "side": "B",
            "px": "0.22756",
            "raw_book_diff": {
                "new": {
                    "sz": "29618.0"
                }
            }
        },
        {
            "user": "0xb2b1d1a7a034d9209a627e754ebb8d24988ca23a",
            "oid": 254433009013,
            "coin": "FET",
            "side": "B",
            "px": "0.22875",
            "raw_book_diff": {
                "new": {
                    "sz": "1975.0"
                }
            }
        },
        {
            "user": "0xf21cd17b9db49fdb21bfdf714558cf8bde3221bd",
            "oid": 254433009014,
            "coin": "LTC",
            "side": "B",
            "px": "76.834",
            "raw_book_diff": {
                "new": {
                    "sz": "36.7"
                }
            }
        },
        {
            "user": "0x7a5a42bd6b1fd13368732be8a7b63e06064e0394",
            "oid": 254433009015,
            "coin": "SOL",
            "side": "B",
            "px": "126.53",
            "raw_book_diff": {
                "new": {
                    "sz": "118.54"
                }
            }
        },
        {
            "user": "0xb4321b142b2a03ce20fcab2007ff6990b9acba93",
            "oid": 254433009016,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008751",
            "raw_book_diff": {
                "new": {
                    "sz": "281302.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009018,
            "coin": "SOL",
            "side": "A",
            "px": "126.67",
            "raw_book_diff": {
                "new": {
                    "sz": "15.8"
                }
            }
        },
        {
            "user": "0x26f68941b957f1edf0b2702296dc8d7a5af70686",
            "oid": 254433009019,
            "coin": "LTC",
            "side": "A",
            "px": "76.9",
            "raw_book_diff": {
                "new": {
                    "sz": "1.9"
                }
            }
        },
        {
            "user": "0x5c9c9ab381c841530464ef9ee402568f84c3b676",
            "oid": 254433009020,
            "coin": "kBONK",
            "side": "A",
            "px": "0.00875",
            "raw_book_diff": {
                "new": {
                    "sz": "562618.0"
                }
            }
        },
        {
            "user": "0x5c9c9ab381c841530464ef9ee402568f84c3b676",
            "oid": 254433009021,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13609",
            "raw_book_diff": {
                "new": {
                    "sz": "19119.0"
                }
            }
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433009022,
            "coin": "ASTER",
            "side": "B",
            "px": "0.93884",
            "raw_book_diff": {
                "new": {
                    "sz": "5342.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009024,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13565",
            "raw_book_diff": {
                "new": {
                    "sz": "220484.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009025,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13675",
            "raw_book_diff": {
                "new": {
                    "sz": "293979.0"
                }
            }
        },
        {
            "user": "0x95995f302ad58138d791ce49f9f3b1274e80c60a",
            "oid": 254433009026,
            "coin": "SOL",
            "side": "B",
            "px": "126.5",
            "raw_book_diff": {
                "new": {
                    "sz": "765.31"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009027,
            "coin": "SOL",
            "side": "A",
            "px": "126.62",
            "raw_book_diff": {
                "new": {
                    "sz": "637.3"
                }
            }
        },
        {
            "user": "0x48ea62a2cc8391fbbe210e8ee89db573a8ec145f",
            "oid": 254433009028,
            "coin": "TIA",
            "side": "A",
            "px": "0.56891",
            "raw_book_diff": {
                "new": {
                    "sz": "15880.8"
                }
            }
        },
        {
            "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
            "oid": 254433009041,
            "coin": "PENGU",
            "side": "B",
            "px": "0.009789",
            "raw_book_diff": {
                "new": {
                    "sz": "381194.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009042,
            "coin": "ORDI",
            "side": "A",
            "px": "3.5445",
            "raw_book_diff": {
                "new": {
                    "sz": "215.79"
                }
            }
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254433009043,
            "coin": "ETH",
            "side": "B",
            "px": "2819.4",
            "raw_book_diff": {
                "new": {
                    "sz": "14.8936"
                }
            }
        },
        {
            "user": "0x72cb918356c4f6d3f1b2e532928110ba6995f139",
            "oid": 254433009047,
            "coin": "SOL",
            "side": "B",
            "px": "126.52",
            "raw_book_diff": {
                "new": {
                    "sz": "948.48"
                }
            }
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254433009061,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.0946",
            "raw_book_diff": {
                "new": {
                    "sz": "7053.0"
                }
            }
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254433009062,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13603",
            "raw_book_diff": {
                "new": {
                    "sz": "5810.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009064,
            "coin": "NEAR",
            "side": "A",
            "px": "1.6306",
            "raw_book_diff": {
                "new": {
                    "sz": "595.4"
                }
            }
        },
        {
            "user": "0xa880d6cc607a05ea617307ab3b0d335e8d8424ee",
            "oid": 254433009070,
            "coin": "RESOLV",
            "side": "A",
            "px": "0.076796",
            "raw_book_diff": {
                "new": {
                    "sz": "3170.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009071,
            "coin": "ZK",
            "side": "A",
            "px": "0.033845",
            "raw_book_diff": {
                "new": {
                    "sz": "25691.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009072,
            "coin": "TIA",
            "side": "B",
            "px": "0.56867",
            "raw_book_diff": {
                "new": {
                    "sz": "1470.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009073,
            "coin": "TIA",
            "side": "B",
            "px": "0.56853",
            "raw_book_diff": {
                "new": {
                    "sz": "2617.6"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009074,
            "coin": "TIA",
            "side": "B",
            "px": "0.56834",
            "raw_book_diff": {
                "new": {
                    "sz": "2782.7"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009075,
            "coin": "TIA",
            "side": "A",
            "px": "0.56885",
            "raw_book_diff": {
                "new": {
                    "sz": "1465.6"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009076,
            "coin": "TIA",
            "side": "A",
            "px": "0.56899",
            "raw_book_diff": {
                "new": {
                    "sz": "2615.2"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009077,
            "coin": "TIA",
            "side": "A",
            "px": "0.56917",
            "raw_book_diff": {
                "new": {
                    "sz": "2724.2"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009078,
            "coin": "TIA",
            "side": "A",
            "px": "0.57018",
            "raw_book_diff": {
                "new": {
                    "sz": "2654.5"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009079,
            "coin": "WLD",
            "side": "B",
            "px": "0.57116",
            "raw_book_diff": {
                "new": {
                    "sz": "1437.3"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009080,
            "coin": "WLD",
            "side": "A",
            "px": "0.57347",
            "raw_book_diff": {
                "new": {
                    "sz": "1349.5"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009081,
            "coin": "SUI",
            "side": "A",
            "px": "1.3509",
            "raw_book_diff": {
                "new": {
                    "sz": "436.3"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009082,
            "coin": "SUI",
            "side": "A",
            "px": "1.3561",
            "raw_book_diff": {
                "new": {
                    "sz": "814.7"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009083,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13609",
            "raw_book_diff": {
                "new": {
                    "sz": "5094.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009084,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.093946",
            "raw_book_diff": {
                "new": {
                    "sz": "6629.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009085,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094809",
            "raw_book_diff": {
                "new": {
                    "sz": "12054.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009086,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.095129",
            "raw_book_diff": {
                "new": {
                    "sz": "11261.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009087,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.095472",
            "raw_book_diff": {
                "new": {
                    "sz": "12091.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009088,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.095859",
            "raw_book_diff": {
                "new": {
                    "sz": "13916.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009089,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.096358",
            "raw_book_diff": {
                "new": {
                    "sz": "14044.0"
                }
            }
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433009091,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.8262",
            "raw_book_diff": {
                "new": {
                    "sz": "3675.2"
                }
            }
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254433009093,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6289",
            "raw_book_diff": {
                "new": {
                    "sz": "3571.7"
                }
            }
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254433009094,
            "coin": "NEAR",
            "side": "A",
            "px": "1.6314",
            "raw_book_diff": {
                "new": {
                    "sz": "3566.2"
                }
            }
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433009106,
            "coin": "APT",
            "side": "A",
            "px": "1.8588",
            "raw_book_diff": {
                "new": {
                    "sz": "156.0"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009107,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094168",
            "raw_book_diff": {
                "new": {
                    "sz": "129301.0"
                }
            }
        },
        {
            "user": "0x727956612a8700627451204a3ae26268bd1a1525",
            "oid": 254433009109,
            "coin": "SUI",
            "side": "A",
            "px": "1.3511",
            "raw_book_diff": {
                "new": {
                    "sz": "996.1"
                }
            }
        },
        {
            "user": "0x588432f54227364caf999a6df88dbf86a90f4772",
            "oid": 254433009110,
            "coin": "BTC",
            "side": "A",
            "px": "85423.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.13951"
                }
            }
        },
        {
            "user": "0x92c7d53dc3bac14722b281e74a3feba8f7ba040e",
            "oid": 254433009111,
            "coin": "TON",
            "side": "A",
            "px": "1.4906",
            "raw_book_diff": {
                "new": {
                    "sz": "6044.7"
                }
            }
        },
        {
            "user": "0x92c7d53dc3bac14722b281e74a3feba8f7ba040e",
            "oid": 254433009113,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094659",
            "raw_book_diff": {
                "new": {
                    "sz": "8159.0"
                }
            }
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254433009114,
            "coin": "ETH",
            "side": "A",
            "px": "2825.8",
            "raw_book_diff": {
                "new": {
                    "sz": "33.1112"
                }
            }
        },
        {
            "user": "0xb2b1d1a7a034d9209a627e754ebb8d24988ca23a",
            "oid": 254433009115,
            "coin": "VIRTUAL",
            "side": "A",
            "px": "0.83039",
            "raw_book_diff": {
                "new": {
                    "sz": "518.3"
                }
            }
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254433009140,
            "coin": "BANANA",
            "side": "B",
            "px": "8.16",
            "raw_book_diff": {
                "new": {
                    "sz": "121.0"
                }
            }
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254433009141,
            "coin": "BANANA",
            "side": "B",
            "px": "8.1563",
            "raw_book_diff": {
                "new": {
                    "sz": "121.0"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009142,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094725",
            "raw_book_diff": {
                "new": {
                    "sz": "130677.0"
                }
            }
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254433009143,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094574",
            "raw_book_diff": {
                "new": {
                    "sz": "1090.0"
                }
            }
        },
        {
            "user": "0x43b29051e7b4ce9fcbc893aa6bd07c32df7399f9",
            "oid": 254433009144,
            "coin": "LINK",
            "side": "B",
            "px": "12.093",
            "raw_book_diff": {
                "new": {
                    "sz": "103.4"
                }
            }
        },
        {
            "user": "0xde236b0b3b367d5ad2a2c42b914bded79d03fe28",
            "oid": 254433009146,
            "coin": "WLD",
            "side": "A",
            "px": "0.57516",
            "raw_book_diff": {
                "new": {
                    "sz": "4023.5"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254432965298,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.29775",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433009147,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.29742",
            "raw_book_diff": {
                "new": {
                    "sz": "2920.8"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254432946124,
            "coin": "SOL",
            "side": "B",
            "px": "126.55",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433009148,
            "coin": "SOL",
            "side": "B",
            "px": "126.49",
            "raw_book_diff": {
                "new": {
                    "sz": "18.89"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254432822531,
            "coin": "kPEPE",
            "side": "A",
            "px": "0.004127",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433009149,
            "coin": "kPEPE",
            "side": "A",
            "px": "0.004125",
            "raw_book_diff": {
                "new": {
                    "sz": "569958.0"
                }
            }
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254432882482,
            "coin": "XRP",
            "side": "A",
            "px": "2.0291",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6ba889db7f923622d3548f621ecc2054b80c1817",
            "oid": 254433009150,
            "coin": "XRP",
            "side": "A",
            "px": "2.0282",
            "raw_book_diff": {
                "new": {
                    "sz": "1983.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009151,
            "coin": "PENGU",
            "side": "B",
            "px": "0.009816",
            "raw_book_diff": {
                "new": {
                    "sz": "247550.0"
                }
            }
        },
        {
            "user": "0x80196a7475464e06de641df6c66e6e278cd364d5",
            "oid": 254433009154,
            "coin": "@107",
            "side": "A",
            "px": "30.226",
            "raw_book_diff": {
                "new": {
                    "sz": "396.9"
                }
            }
        },
        {
            "user": "0x80196a7475464e06de641df6c66e6e278cd364d5",
            "oid": 254433009155,
            "coin": "@107",
            "side": "A",
            "px": "30.219",
            "raw_book_diff": {
                "new": {
                    "sz": "264.6"
                }
            }
        },
        {
            "user": "0xc2e137ea6b79a7992dd54d44ee8223f2c7e13e22",
            "oid": 254433009156,
            "coin": "PYTH",
            "side": "B",
            "px": "0.067659",
            "raw_book_diff": {
                "new": {
                    "sz": "739.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009157,
            "coin": "ENS",
            "side": "B",
            "px": "10.536",
            "raw_book_diff": {
                "new": {
                    "sz": "98.8"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009158,
            "coin": "LTC",
            "side": "B",
            "px": "76.653",
            "raw_book_diff": {
                "new": {
                    "sz": "9.77"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009159,
            "coin": "LTC",
            "side": "A",
            "px": "76.894",
            "raw_book_diff": {
                "new": {
                    "sz": "10.96"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009160,
            "coin": "APT",
            "side": "B",
            "px": "1.8567",
            "raw_book_diff": {
                "new": {
                    "sz": "339.79"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009161,
            "coin": "APT",
            "side": "B",
            "px": "1.8563",
            "raw_book_diff": {
                "new": {
                    "sz": "653.67"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009162,
            "coin": "APT",
            "side": "B",
            "px": "1.8556",
            "raw_book_diff": {
                "new": {
                    "sz": "625.33"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009163,
            "coin": "APT",
            "side": "B",
            "px": "1.8514",
            "raw_book_diff": {
                "new": {
                    "sz": "948.96"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009164,
            "coin": "APT",
            "side": "B",
            "px": "1.8465",
            "raw_book_diff": {
                "new": {
                    "sz": "944.26"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009165,
            "coin": "APT",
            "side": "A",
            "px": "1.8574",
            "raw_book_diff": {
                "new": {
                    "sz": "435.96"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009166,
            "coin": "APT",
            "side": "A",
            "px": "1.8578",
            "raw_book_diff": {
                "new": {
                    "sz": "787.6"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009167,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5397",
            "raw_book_diff": {
                "new": {
                    "sz": "200.49"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009168,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5385",
            "raw_book_diff": {
                "new": {
                    "sz": "341.59"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009169,
            "coin": "ORDI",
            "side": "B",
            "px": "3.537",
            "raw_book_diff": {
                "new": {
                    "sz": "373.5"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009170,
            "coin": "ORDI",
            "side": "A",
            "px": "3.5413",
            "raw_book_diff": {
                "new": {
                    "sz": "229.4"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009171,
            "coin": "ORDI",
            "side": "A",
            "px": "3.5425",
            "raw_book_diff": {
                "new": {
                    "sz": "488.94"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009172,
            "coin": "ORDI",
            "side": "A",
            "px": "3.5439",
            "raw_book_diff": {
                "new": {
                    "sz": "469.84"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009173,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6224",
            "raw_book_diff": {
                "new": {
                    "sz": "506.7"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009174,
            "coin": "NEAR",
            "side": "A",
            "px": "1.6345",
            "raw_book_diff": {
                "new": {
                    "sz": "458.2"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009175,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008741",
            "raw_book_diff": {
                "new": {
                    "sz": "95673.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009176,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008772",
            "raw_book_diff": {
                "new": {
                    "sz": "166095.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009177,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008812",
            "raw_book_diff": {
                "new": {
                    "sz": "181636.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009178,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008904",
            "raw_book_diff": {
                "new": {
                    "sz": "963639.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009179,
            "coin": "PENGU",
            "side": "A",
            "px": "0.010075",
            "raw_book_diff": {
                "new": {
                    "sz": "751893.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009180,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82891",
            "raw_book_diff": {
                "new": {
                    "sz": "829.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009181,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82798",
            "raw_book_diff": {
                "new": {
                    "sz": "1827.8"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009182,
            "coin": "VIRTUAL",
            "side": "A",
            "px": "0.83038",
            "raw_book_diff": {
                "new": {
                    "sz": "1879.1"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009183,
            "coin": "CHILLGUY",
            "side": "A",
            "px": "0.018696",
            "raw_book_diff": {
                "new": {
                    "sz": "650824.0"
                }
            }
        },
        {
            "user": "0x4a1f71d91ee321209939a227337f72761b53f701",
            "oid": 254433009184,
            "coin": "xyz:XYZ100",
            "side": "A",
            "px": "25179.0",
            "raw_book_diff": {
                "new": {
                    "sz": "3.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009195,
            "coin": "ME",
            "side": "B",
            "px": "0.30171",
            "raw_book_diff": {
                "new": {
                    "sz": "1956.5"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009196,
            "coin": "ME",
            "side": "B",
            "px": "0.30161",
            "raw_book_diff": {
                "new": {
                    "sz": "3788.5"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009197,
            "coin": "ME",
            "side": "B",
            "px": "0.30148",
            "raw_book_diff": {
                "new": {
                    "sz": "4073.8"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009198,
            "coin": "ME",
            "side": "A",
            "px": "0.30186",
            "raw_book_diff": {
                "new": {
                    "sz": "1953.6"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009199,
            "coin": "ME",
            "side": "A",
            "px": "0.30196",
            "raw_book_diff": {
                "new": {
                    "sz": "3847.8"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009200,
            "coin": "ME",
            "side": "A",
            "px": "0.3021",
            "raw_book_diff": {
                "new": {
                    "sz": "4448.6"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009201,
            "coin": "IP",
            "side": "A",
            "px": "2.267",
            "raw_book_diff": {
                "new": {
                    "sz": "303.4"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009202,
            "coin": "2Z",
            "side": "B",
            "px": "0.10234",
            "raw_book_diff": {
                "new": {
                    "sz": "12144.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009203,
            "coin": "ZEC",
            "side": "B",
            "px": "348.07",
            "raw_book_diff": {
                "new": {
                    "sz": "0.95"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009204,
            "coin": "BIO",
            "side": "B",
            "px": "0.050922",
            "raw_book_diff": {
                "new": {
                    "sz": "14646.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009205,
            "coin": "BIO",
            "side": "B",
            "px": "0.050908",
            "raw_book_diff": {
                "new": {
                    "sz": "30001.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009206,
            "coin": "BIO",
            "side": "B",
            "px": "0.050892",
            "raw_book_diff": {
                "new": {
                    "sz": "26673.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009207,
            "coin": "BIO",
            "side": "B",
            "px": "0.050852",
            "raw_book_diff": {
                "new": {
                    "sz": "31594.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009208,
            "coin": "BIO",
            "side": "A",
            "px": "0.050953",
            "raw_book_diff": {
                "new": {
                    "sz": "31465.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009209,
            "coin": "BIO",
            "side": "A",
            "px": "0.050971",
            "raw_book_diff": {
                "new": {
                    "sz": "32463.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009210,
            "coin": "BIO",
            "side": "A",
            "px": "0.051013",
            "raw_book_diff": {
                "new": {
                    "sz": "30752.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009211,
            "coin": "KAS",
            "side": "A",
            "px": "0.051634",
            "raw_book_diff": {
                "new": {
                    "sz": "25752.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009212,
            "coin": "KAS",
            "side": "A",
            "px": "0.052015",
            "raw_book_diff": {
                "new": {
                    "sz": "24892.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009213,
            "coin": "FXS",
            "side": "A",
            "px": "0.77832",
            "raw_book_diff": {
                "new": {
                    "sz": "760.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009214,
            "coin": "NOT",
            "side": "A",
            "px": "0.000518",
            "raw_book_diff": {
                "new": {
                    "sz": "3103983.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009215,
            "coin": "NOT",
            "side": "A",
            "px": "0.000519",
            "raw_book_diff": {
                "new": {
                    "sz": "3050533.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009216,
            "coin": "FTT",
            "side": "B",
            "px": "0.55273",
            "raw_book_diff": {
                "new": {
                    "sz": "291.3"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009217,
            "coin": "FTT",
            "side": "B",
            "px": "0.55096",
            "raw_book_diff": {
                "new": {
                    "sz": "1176.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009218,
            "coin": "FTT",
            "side": "A",
            "px": "0.55418",
            "raw_book_diff": {
                "new": {
                    "sz": "330.8"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009219,
            "coin": "BSV",
            "side": "B",
            "px": "19.693",
            "raw_book_diff": {
                "new": {
                    "sz": "16.63"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009220,
            "coin": "BSV",
            "side": "B",
            "px": "19.671",
            "raw_book_diff": {
                "new": {
                    "sz": "35.24"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009221,
            "coin": "ZORA",
            "side": "A",
            "px": "0.043826",
            "raw_book_diff": {
                "new": {
                    "sz": "4111.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009222,
            "coin": "ZORA",
            "side": "A",
            "px": "0.043925",
            "raw_book_diff": {
                "new": {
                    "sz": "16873.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009223,
            "coin": "SAND",
            "side": "B",
            "px": "0.14118",
            "raw_book_diff": {
                "new": {
                    "sz": "4806.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009224,
            "coin": "SAND",
            "side": "A",
            "px": "0.14196",
            "raw_book_diff": {
                "new": {
                    "sz": "11876.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009225,
            "coin": "MEME",
            "side": "B",
            "px": "0.001178",
            "raw_book_diff": {
                "new": {
                    "sz": "562324.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009226,
            "coin": "MEME",
            "side": "B",
            "px": "0.001178",
            "raw_book_diff": {
                "new": {
                    "sz": "1069030.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009227,
            "coin": "MEME",
            "side": "B",
            "px": "0.001177",
            "raw_book_diff": {
                "new": {
                    "sz": "1087412.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009228,
            "coin": "MEME",
            "side": "A",
            "px": "0.00118",
            "raw_book_diff": {
                "new": {
                    "sz": "529601.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009229,
            "coin": "MEME",
            "side": "A",
            "px": "0.001181",
            "raw_book_diff": {
                "new": {
                    "sz": "1117348.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009230,
            "coin": "MEME",
            "side": "A",
            "px": "0.001182",
            "raw_book_diff": {
                "new": {
                    "sz": "1171363.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009231,
            "coin": "MEME",
            "side": "A",
            "px": "0.001183",
            "raw_book_diff": {
                "new": {
                    "sz": "1087489.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009232,
            "coin": "MEME",
            "side": "A",
            "px": "0.001185",
            "raw_book_diff": {
                "new": {
                    "sz": "1035445.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009233,
            "coin": "COMP",
            "side": "B",
            "px": "34.752",
            "raw_book_diff": {
                "new": {
                    "sz": "36.46"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009234,
            "coin": "COMP",
            "side": "A",
            "px": "34.9",
            "raw_book_diff": {
                "new": {
                    "sz": "24.49"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009235,
            "coin": "SUSHI",
            "side": "A",
            "px": "0.34014",
            "raw_book_diff": {
                "new": {
                    "sz": "3455.1"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009236,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001874",
            "raw_book_diff": {
                "new": {
                    "sz": "496397.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009237,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001872",
            "raw_book_diff": {
                "new": {
                    "sz": "533030.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009238,
            "coin": "TURBO",
            "side": "A",
            "px": "0.001876",
            "raw_book_diff": {
                "new": {
                    "sz": "647685.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009239,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62148",
            "raw_book_diff": {
                "new": {
                    "sz": "919.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009240,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62126",
            "raw_book_diff": {
                "new": {
                    "sz": "1747.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009241,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62099",
            "raw_book_diff": {
                "new": {
                    "sz": "1827.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009242,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62042",
            "raw_book_diff": {
                "new": {
                    "sz": "1781.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009243,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041481",
            "raw_book_diff": {
                "new": {
                    "sz": "13486.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009244,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041468",
            "raw_book_diff": {
                "new": {
                    "sz": "25668.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009245,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041405",
            "raw_book_diff": {
                "new": {
                    "sz": "27562.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009246,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041502",
            "raw_book_diff": {
                "new": {
                    "sz": "15751.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009247,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041514",
            "raw_book_diff": {
                "new": {
                    "sz": "31034.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009248,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041539",
            "raw_book_diff": {
                "new": {
                    "sz": "29397.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009249,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041574",
            "raw_book_diff": {
                "new": {
                    "sz": "29585.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009250,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014519",
            "raw_book_diff": {
                "new": {
                    "sz": "22475.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009251,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01451",
            "raw_book_diff": {
                "new": {
                    "sz": "50990.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009252,
            "coin": "BRETT",
            "side": "B",
            "px": "0.0145",
            "raw_book_diff": {
                "new": {
                    "sz": "51291.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009253,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01448",
            "raw_book_diff": {
                "new": {
                    "sz": "51737.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009254,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014539",
            "raw_book_diff": {
                "new": {
                    "sz": "93734.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009255,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014549",
            "raw_book_diff": {
                "new": {
                    "sz": "83297.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009256,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014564",
            "raw_book_diff": {
                "new": {
                    "sz": "85603.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009257,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014701",
            "raw_book_diff": {
                "new": {
                    "sz": "95115.0"
                }
            }
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433009258,
            "coin": "POLYX",
            "side": "A",
            "px": "0.061343",
            "raw_book_diff": {
                "new": {
                    "sz": "10970.0"
                }
            }
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433009269,
            "coin": "IP",
            "side": "B",
            "px": "2.2639",
            "raw_book_diff": {
                "new": {
                    "sz": "833.9"
                }
            }
        },
        {
            "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
            "oid": 254433009292,
            "coin": "ZORA",
            "side": "A",
            "px": "0.043892",
            "raw_book_diff": {
                "new": {
                    "sz": "41857.0"
                }
            }
        },
        {
            "user": "0xd911e53d53b663972254e086450fd6198a25961e",
            "oid": 254433009293,
            "coin": "@151",
            "side": "B",
            "px": "2821.8",
            "raw_book_diff": {
                "new": {
                    "sz": "16.9207"
                }
            }
        },
        {
            "user": "0x0fd468a73084daa6ea77a9261e40fdec3e67e0c7",
            "oid": 254433009305,
            "coin": "BTC",
            "side": "B",
            "px": "85402.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.46377"
                }
            }
        },
        {
            "user": "0xd6e1a6f7bb47aca69aaa511210227d40926ff256",
            "oid": 254433009330,
            "coin": "DOGE",
            "side": "B",
            "px": "0.1359",
            "raw_book_diff": {
                "new": {
                    "sz": "266558.0"
                }
            }
        },
        {
            "user": "0xd6e1a6f7bb47aca69aaa511210227d40926ff256",
            "oid": 254433009331,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13574",
            "raw_book_diff": {
                "new": {
                    "sz": "2965247.0"
                }
            }
        },
        {
            "user": "0xd6e1a6f7bb47aca69aaa511210227d40926ff256",
            "oid": 254433009332,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13633",
            "raw_book_diff": {
                "new": {
                    "sz": "2915310.0"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009381,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07075",
            "raw_book_diff": {
                "new": {
                    "sz": "26921.7"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009387,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13618",
            "raw_book_diff": {
                "new": {
                    "sz": "122191.0"
                }
            }
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433009388,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07077",
            "raw_book_diff": {
                "new": {
                    "sz": "28653.3"
                }
            }
        },
        {
            "user": "0x92c7d53dc3bac14722b281e74a3feba8f7ba040e",
            "oid": 254433009389,
            "coin": "kPEPE",
            "side": "B",
            "px": "0.00411",
            "raw_book_diff": {
                "new": {
                    "sz": "4372647.0"
                }
            }
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254433009390,
            "coin": "kPEPE",
            "side": "A",
            "px": "0.004124",
            "raw_book_diff": {
                "new": {
                    "sz": "960730.0"
                }
            }
        },
        {
            "user": "0x4537e1a5b7bcadd39f6d6211cb757ce431798732",
            "oid": 254433009394,
            "coin": "XRP",
            "side": "A",
            "px": "2.0278",
            "raw_book_diff": {
                "new": {
                    "sz": "24659.0"
                }
            }
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254433009395,
            "coin": "ARB",
            "side": "A",
            "px": "0.19325",
            "raw_book_diff": {
                "new": {
                    "sz": "1725.6"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009396,
            "coin": "kPEPE",
            "side": "A",
            "px": "0.004125",
            "raw_book_diff": {
                "new": {
                    "sz": "1920538.0"
                }
            }
        },
        {
            "user": "0xb4321b142b2a03ce20fcab2007ff6990b9acba93",
            "oid": 254433009397,
            "coin": "kSHIB",
            "side": "A",
            "px": "0.008002",
            "raw_book_diff": {
                "new": {
                    "sz": "37332.0"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009398,
            "coin": "ARB",
            "side": "A",
            "px": "0.19331",
            "raw_book_diff": {
                "new": {
                    "sz": "5247.8"
                }
            }
        },
        {
            "user": "0x162cc7c861ebd0c06b3d72319201150482518185",
            "oid": 254433009399,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.00257"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009400,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "277.1",
            "raw_book_diff": {
                "new": {
                    "sz": "24.059"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009401,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "277.1",
            "raw_book_diff": {
                "new": {
                    "sz": "24.059"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009402,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "276.8",
            "raw_book_diff": {
                "new": {
                    "sz": "24.085"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009403,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.7",
            "raw_book_diff": {
                "new": {
                    "sz": "24.007"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009404,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.7",
            "raw_book_diff": {
                "new": {
                    "sz": "24.007"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009405,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "278.1",
            "raw_book_diff": {
                "new": {
                    "sz": "23.972"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009406,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "277.0",
            "raw_book_diff": {
                "new": {
                    "sz": "120.337"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009407,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.7",
            "raw_book_diff": {
                "new": {
                    "sz": "120.034"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009408,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "277.8",
            "raw_book_diff": {
                "new": {
                    "sz": "119.99"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009409,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "273.0",
            "raw_book_diff": {
                "new": {
                    "sz": "183.15"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009410,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "280.4",
            "raw_book_diff": {
                "new": {
                    "sz": "178.317"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009411,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "285.09",
            "raw_book_diff": {
                "new": {
                    "sz": "350.766"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009412,
            "coin": "xyz:XYZ100",
            "side": "B",
            "px": "25101.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.992"
                }
            }
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433009413,
            "coin": "xyz:XYZ100",
            "side": "A",
            "px": "25252.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.99"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009419,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.00257"
                }
            }
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254433009420,
            "coin": "BTC",
            "side": "B",
            "px": "85412.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.17563"
                }
            }
        },
        {
            "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
            "oid": 254433009421,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.00257"
                }
            }
        },
        {
            "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
            "oid": 254433009422,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.00257"
                }
            }
        },
        {
            "user": "0xe3b6e3443c8f2080704e7421bad9340f13950acb",
            "oid": 254433009423,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.00257"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009424,
            "coin": "ETHFI",
            "side": "B",
            "px": "0.73604",
            "raw_book_diff": {
                "new": {
                    "sz": "5406.7"
                }
            }
        },
        {
            "user": "0xc6ac58a7a63339898aeda32499a8238a46d88e84",
            "oid": 254433009425,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.00257"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009426,
            "coin": "WLD",
            "side": "B",
            "px": "0.57235",
            "raw_book_diff": {
                "new": {
                    "sz": "21361.2"
                }
            }
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254433009427,
            "coin": "ARB",
            "side": "A",
            "px": "0.19327",
            "raw_book_diff": {
                "new": {
                    "sz": "80893.1"
                }
            }
        },
        {
            "user": "0xa289ee1e56c0d5d041db762e6123e78af0f7d9ad",
            "oid": 254433009428,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.00257"
                }
            }
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254433009429,
            "coin": "UNI",
            "side": "B",
            "px": "5.465",
            "raw_book_diff": {
                "new": {
                    "sz": "90.6"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009430,
            "coin": "UNI",
            "side": "B",
            "px": "5.5097",
            "raw_book_diff": {
                "new": {
                    "sz": "2189.1"
                }
            }
        },
        {
            "user": "0xe056eebc2c7acfc782d47ebe18ad71735480f72a",
            "oid": 254433009431,
            "coin": "ETH",
            "side": "A",
            "px": "2822.3",
            "raw_book_diff": {
                "new": {
                    "sz": "0.0443"
                }
            }
        },
        {
            "user": "0xaf9f722a676230cc44045efe26fe9a85801ca4fa",
            "oid": 254433009432,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14612",
            "raw_book_diff": {
                "new": {
                    "sz": "66097.0"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009434,
            "coin": "AR",
            "side": "B",
            "px": "3.9564",
            "raw_book_diff": {
                "new": {
                    "sz": "3035.53"
                }
            }
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254433009435,
            "coin": "BTC",
            "side": "A",
            "px": "85423.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.17561"
                }
            }
        },
        {
            "user": "0xef5fdcefdb38803a654ee9f3539bed8058443d99",
            "oid": 254433009436,
            "coin": "SOL",
            "side": "A",
            "px": "126.6",
            "raw_book_diff": {
                "new": {
                    "sz": "1.95"
                }
            }
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254433009437,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13602",
            "raw_book_diff": {
                "new": {
                    "sz": "5810.0"
                }
            }
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254433009438,
            "coin": "SUI",
            "side": "A",
            "px": "1.3509",
            "raw_book_diff": {
                "new": {
                    "sz": "363.4"
                }
            }
        },
        {
            "user": "0xb4321b142b2a03ce20fcab2007ff6990b9acba93",
            "oid": 254433009439,
            "coin": "STX",
            "side": "A",
            "px": "0.28286",
            "raw_book_diff": {
                "new": {
                    "sz": "4620.4"
                }
            }
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254433009440,
            "coin": "TON",
            "side": "A",
            "px": "1.4903",
            "raw_book_diff": {
                "new": {
                    "sz": "1336.4"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009441,
            "coin": "ARB",
            "side": "A",
            "px": "0.19331",
            "raw_book_diff": {
                "new": {
                    "sz": "6736.9"
                }
            }
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254433009442,
            "coin": "kNEIRO",
            "side": "A",
            "px": "0.12033",
            "raw_book_diff": {
                "new": {
                    "sz": "4951.3"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009443,
            "coin": "SAND",
            "side": "A",
            "px": "0.14134",
            "raw_book_diff": {
                "new": {
                    "sz": "5282.0"
                }
            }
        },
        {
            "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
            "oid": 254433009444,
            "coin": "FARTCOIN",
            "side": "A",
            "px": "0.29822",
            "raw_book_diff": {
                "new": {
                    "sz": "15060.2"
                }
            }
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254433009445,
            "coin": "SUI",
            "side": "A",
            "px": "1.3508",
            "raw_book_diff": {
                "new": {
                    "sz": "500.8"
                }
            }
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254433009446,
            "coin": "SOL",
            "side": "B",
            "px": "126.51",
            "raw_book_diff": {
                "new": {
                    "sz": "474.29"
                }
            }
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254433009448,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094666",
            "raw_book_diff": {
                "new": {
                    "sz": "118279.0"
                }
            }
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254433009449,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.09469",
            "raw_book_diff": {
                "new": {
                    "sz": "182912.0"
                }
            }
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254433009450,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094713",
            "raw_book_diff": {
                "new": {
                    "sz": "182868.0"
                }
            }
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254433009451,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094737",
            "raw_book_diff": {
                "new": {
                    "sz": "182821.0"
                }
            }
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433009453,
            "coin": "@162",
            "side": "B",
            "px": "0.29698",
            "raw_book_diff": {
                "new": {
                    "sz": "16514.5"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009475,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14627",
            "raw_book_diff": {
                "new": {
                    "sz": "1367.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009476,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14626",
            "raw_book_diff": {
                "new": {
                    "sz": "1367.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009477,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14625",
            "raw_book_diff": {
                "new": {
                    "sz": "1367.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009478,
            "coin": "WLFI",
            "side": "A",
            "px": "0.14636",
            "raw_book_diff": {
                "new": {
                    "sz": "1367.0"
                }
            }
        },
        {
            "user": "0x3ff54fd26855db3758d0b5ae7aed47440c47f705",
            "oid": 254433009479,
            "coin": "BTC",
            "side": "A",
            "px": "85446.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.40439"
                }
            }
        },
        {
            "user": "0xbb475febf78b528f4f0d4ae1a12541406565199c",
            "oid": 254433009481,
            "coin": "BTC",
            "side": "B",
            "px": "85393.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.7027"
                }
            }
        },
        {
            "user": "0xbb475febf78b528f4f0d4ae1a12541406565199c",
            "oid": 254433009482,
            "coin": "BTC",
            "side": "A",
            "px": "85443.0",
            "raw_book_diff": {
                "new": {
                    "sz": "2.92623"
                }
            }
        },
        {
            "user": "0xbb475febf78b528f4f0d4ae1a12541406565199c",
            "oid": 254433009483,
            "coin": "BTC",
            "side": "B",
            "px": "85404.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.7026"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009484,
            "coin": "BABY",
            "side": "A",
            "px": "0.018348",
            "raw_book_diff": {
                "new": {
                    "sz": "38399.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009485,
            "coin": "ETH",
            "side": "A",
            "px": "2821.6",
            "raw_book_diff": {
                "new": {
                    "sz": "1.7721"
                }
            }
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254433009487,
            "coin": "BTC",
            "side": "B",
            "px": "85401.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.75658"
                }
            }
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254433009488,
            "coin": "BTC",
            "side": "A",
            "px": "85423.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.75614"
                }
            }
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433009489,
            "coin": "ARB",
            "side": "A",
            "px": "0.19328",
            "raw_book_diff": {
                "new": {
                    "sz": "800.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009490,
            "coin": "WLFI",
            "side": "A",
            "px": "0.1464",
            "raw_book_diff": {
                "new": {
                    "sz": "5628.0"
                }
            }
        },
        {
            "user": "0x38f96bba487a628d2a84afc3fe5c7d77eb07cb04",
            "oid": 254433009491,
            "coin": "ENA",
            "side": "B",
            "px": "0.24184",
            "raw_book_diff": {
                "new": {
                    "sz": "248096.0"
                }
            }
        },
        {
            "user": "0x38f96bba487a628d2a84afc3fe5c7d77eb07cb04",
            "oid": 254433009492,
            "coin": "ENA",
            "side": "A",
            "px": "0.24219",
            "raw_book_diff": {
                "new": {
                    "sz": "41294.0"
                }
            }
        },
        {
            "user": "0x38f96bba487a628d2a84afc3fe5c7d77eb07cb04",
            "oid": 254433009493,
            "coin": "ENA",
            "side": "A",
            "px": "0.24229",
            "raw_book_diff": {
                "new": {
                    "sz": "247665.0"
                }
            }
        },
        {
            "user": "0x6a072733f7249e6b794cae88c6d3cc7eb756a93c",
            "oid": 254433009494,
            "coin": "@142",
            "side": "A",
            "px": "85467.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.01"
                }
            }
        },
        {
            "user": "0x34827044cbd4b808fc1b189fce9f50e6dafae7c9",
            "oid": 254432824918,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014518",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x34827044cbd4b808fc1b189fce9f50e6dafae7c9",
            "oid": 254433009495,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014508",
            "raw_book_diff": {
                "new": {
                    "sz": "72580.0"
                }
            }
        },
        {
            "user": "0xaf9f722a676230cc44045efe26fe9a85801ca4fa",
            "oid": 254433009508,
            "coin": "ETH",
            "side": "A",
            "px": "2824.5",
            "raw_book_diff": {
                "new": {
                    "sz": "69.5149"
                }
            }
        },
        {
            "user": "0xa33a4a057334c7811ad5f45f3c4f0dfa3d081ff8",
            "oid": 254433009511,
            "coin": "LINK",
            "side": "B",
            "px": "12.111",
            "raw_book_diff": {
                "new": {
                    "sz": "825.7"
                }
            }
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254433009514,
            "coin": "WLFI",
            "side": "B",
            "px": "0.1462",
            "raw_book_diff": {
                "new": {
                    "sz": "6840.0"
                }
            }
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254433009515,
            "coin": "TON",
            "side": "A",
            "px": "1.4903",
            "raw_book_diff": {
                "new": {
                    "sz": "1340.3"
                }
            }
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254433009516,
            "coin": "TON",
            "side": "A",
            "px": "1.4903",
            "raw_book_diff": {
                "new": {
                    "sz": "1347.4"
                }
            }
        },
        {
            "user": "0x4537e1a5b7bcadd39f6d6211cb757ce431798732",
            "oid": 254433009517,
            "coin": "XRP",
            "side": "B",
            "px": "2.0267",
            "raw_book_diff": {
                "new": {
                    "sz": "24669.0"
                }
            }
        },
        {
            "user": "0xfad17af4900f3d3afa663103f58b41607ee45d70",
            "oid": 254433009518,
            "coin": "ENA",
            "side": "A",
            "px": "0.24216",
            "raw_book_diff": {
                "new": {
                    "sz": "4000.0"
                }
            }
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254433009525,
            "coin": "SOL",
            "side": "A",
            "px": "126.57",
            "raw_book_diff": {
                "new": {
                    "sz": "474.09"
                }
            }
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254433009526,
            "coin": "SOL",
            "side": "A",
            "px": "126.56",
            "raw_book_diff": {
                "new": {
                    "sz": "474.14"
                }
            }
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433009528,
            "coin": "FIL",
            "side": "A",
            "px": "1.4518",
            "raw_book_diff": {
                "new": {
                    "sz": "240.0"
                }
            }
        },
        {
            "user": "0x92c7d53dc3bac14722b281e74a3feba8f7ba040e",
            "oid": 254433009530,
            "coin": "UNI",
            "side": "A",
            "px": "5.5227",
            "raw_book_diff": {
                "new": {
                    "sz": "70.7"
                }
            }
        },
        {
            "user": "0x4129c62faf652fea61375dcd9ca8ce24b2bb8b95",
            "oid": 254433009531,
            "coin": "ETH",
            "side": "B",
            "px": "2821.4",
            "raw_book_diff": {
                "new": {
                    "sz": "0.7975"
                }
            }
        },
        {
            "user": "0x621c5551678189b9a6c94d929924c225ff1d63ab",
            "oid": 254433009532,
            "coin": "ETH",
            "side": "A",
            "px": "2824.0",
            "raw_book_diff": {
                "new": {
                    "sz": "69.5149"
                }
            }
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433009539,
            "coin": "ADA",
            "side": "A",
            "px": "0.37923",
            "raw_book_diff": {
                "new": {
                    "sz": "1600.0"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009540,
            "coin": "BABY",
            "side": "B",
            "px": "0.018303",
            "raw_book_diff": {
                "new": {
                    "sz": "220890.0"
                }
            }
        },
        {
            "user": "0x48ea62a2cc8391fbbe210e8ee89db573a8ec145f",
            "oid": 254433009542,
            "coin": "SUI",
            "side": "A",
            "px": "1.3508",
            "raw_book_diff": {
                "new": {
                    "sz": "7403.3"
                }
            }
        },
        {
            "user": "0x0fd468a73084daa6ea77a9261e40fdec3e67e0c7",
            "oid": 254433009543,
            "coin": "BTC",
            "side": "B",
            "px": "85409.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.17094"
                }
            }
        },
        {
            "user": "0x0fd468a73084daa6ea77a9261e40fdec3e67e0c7",
            "oid": 254433009544,
            "coin": "BTC",
            "side": "A",
            "px": "85423.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.46346"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009547,
            "coin": "FARTCOIN",
            "side": "A",
            "px": "0.29794",
            "raw_book_diff": {
                "new": {
                    "sz": "15101.7"
                }
            }
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254433009548,
            "coin": "BTC",
            "side": "B",
            "px": "85386.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.75688"
                }
            }
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254433009549,
            "coin": "BTC",
            "side": "B",
            "px": "85395.0",
            "raw_book_diff": {
                "new": {
                    "sz": "1.7567"
                }
            }
        },
        {
            "user": "0x39475d17bcd20adc540e647dae6781b153fbf3b1",
            "oid": 254433009562,
            "coin": "ETH",
            "side": "A",
            "px": "2823.6",
            "raw_book_diff": {
                "new": {
                    "sz": "69.5149"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009576,
            "coin": "WIF",
            "side": "B",
            "px": "0.33697",
            "raw_book_diff": {
                "new": {
                    "sz": "12042.0"
                }
            }
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433009603,
            "coin": "SUI",
            "side": "A",
            "px": "1.3515",
            "raw_book_diff": {
                "new": {
                    "sz": "517.9"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009604,
            "coin": "COMP",
            "side": "A",
            "px": "34.893",
            "raw_book_diff": {
                "new": {
                    "sz": "112.59"
                }
            }
        },
        {
            "user": "0x2db6f50126cf628c138de573a13ca3fcd2e22ed4",
            "oid": 254433009605,
            "coin": "MANTA",
            "side": "B",
            "px": "0.09715",
            "raw_book_diff": {
                "new": {
                    "sz": "118.4"
                }
            }
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254433009606,
            "coin": "XRP",
            "side": "B",
            "px": "2.0245",
            "raw_book_diff": {
                "new": {
                    "sz": "23331.0"
                }
            }
        },
        {
            "user": "0x9f052eabbe59d1d6e324232a73ef7e1faa66bd54",
            "oid": 254433009607,
            "coin": "RESOLV",
            "side": "A",
            "px": "0.076796",
            "raw_book_diff": {
                "new": {
                    "sz": "526.0"
                }
            }
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433009634,
            "coin": "WIF",
            "side": "A",
            "px": "0.33898",
            "raw_book_diff": {
                "new": {
                    "sz": "2346.0"
                }
            }
        },
        {
            "user": "0x5c9c9ab381c841530464ef9ee402568f84c3b676",
            "oid": 254433009635,
            "coin": "BNB",
            "side": "A",
            "px": "821.14",
            "raw_book_diff": {
                "new": {
                    "sz": "3.378"
                }
            }
        },
        {
            "user": "0x75f3665323cce024368870bfd89db0c21e1ce76b",
            "oid": 254433009639,
            "coin": "AVAX",
            "side": "B",
            "px": "12.902",
            "raw_book_diff": {
                "new": {
                    "sz": "240.62"
                }
            }
        },
        {
            "user": "0x399965e15d4e61ec3529cc98b7f7ebb93b733336",
            "oid": 254433009640,
            "coin": "@162",
            "side": "A",
            "px": "0.29787",
            "raw_book_diff": {
                "new": {
                    "sz": "2303.2"
                }
            }
        },
        {
            "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
            "oid": 254433009641,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.053369",
            "raw_book_diff": {
                "new": {
                    "sz": "15517.0"
                }
            }
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433009645,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82478",
            "raw_book_diff": {
                "new": {
                    "sz": "14630.5"
                }
            }
        },
        {
            "user": "0xaf9f722a676230cc44045efe26fe9a85801ca4fa",
            "oid": 254433009646,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.29708",
            "raw_book_diff": {
                "new": {
                    "sz": "34036.2"
                }
            }
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254433009651,
            "coin": "BTC",
            "side": "B",
            "px": "85410.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.17564"
                }
            }
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254433009652,
            "coin": "BTC",
            "side": "B",
            "px": "85398.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.17566"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009654,
            "coin": "WIF",
            "side": "B",
            "px": "0.33865",
            "raw_book_diff": {
                "new": {
                    "sz": "2952.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009655,
            "coin": "WIF",
            "side": "B",
            "px": "0.33863",
            "raw_book_diff": {
                "new": {
                    "sz": "2952.0"
                }
            }
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009656,
            "coin": "WIF",
            "side": "B",
            "px": "0.33861",
            "raw_book_diff": {
                "new": {
                    "sz": "2952.0"
                }
            }
        },
        {
            "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
            "oid": 254433009657,
            "coin": "VIRTUAL",
            "side": "A",
            "px": "0.82935",
            "raw_book_diff": {
                "new": {
                    "sz": "745.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009658,
            "coin": "IP",
            "side": "A",
            "px": "2.267",
            "raw_book_diff": {
                "new": {
                    "sz": "304.9"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009659,
            "coin": "ZEC",
            "side": "B",
            "px": "348.06",
            "raw_book_diff": {
                "new": {
                    "sz": "0.93"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009660,
            "coin": "ZEC",
            "side": "A",
            "px": "349.02",
            "raw_book_diff": {
                "new": {
                    "sz": "1.74"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009661,
            "coin": "KAS",
            "side": "B",
            "px": "0.051481",
            "raw_book_diff": {
                "new": {
                    "sz": "10846.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009662,
            "coin": "KAS",
            "side": "A",
            "px": "0.051577",
            "raw_book_diff": {
                "new": {
                    "sz": "23812.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009663,
            "coin": "ZORA",
            "side": "A",
            "px": "0.043826",
            "raw_book_diff": {
                "new": {
                    "sz": "4441.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009664,
            "coin": "ZORA",
            "side": "A",
            "px": "0.043924",
            "raw_book_diff": {
                "new": {
                    "sz": "16621.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009665,
            "coin": "AVNT",
            "side": "B",
            "px": "0.33398",
            "raw_book_diff": {
                "new": {
                    "sz": "877.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009666,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001874",
            "raw_book_diff": {
                "new": {
                    "sz": "526235.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009667,
            "coin": "TIA",
            "side": "B",
            "px": "0.56867",
            "raw_book_diff": {
                "new": {
                    "sz": "1399.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009668,
            "coin": "TIA",
            "side": "B",
            "px": "0.56853",
            "raw_book_diff": {
                "new": {
                    "sz": "2572.9"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009669,
            "coin": "TIA",
            "side": "B",
            "px": "0.56834",
            "raw_book_diff": {
                "new": {
                    "sz": "2672.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009670,
            "coin": "TIA",
            "side": "B",
            "px": "0.56726",
            "raw_book_diff": {
                "new": {
                    "sz": "3079.8"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009671,
            "coin": "TIA",
            "side": "A",
            "px": "0.56885",
            "raw_book_diff": {
                "new": {
                    "sz": "1328.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009672,
            "coin": "TIA",
            "side": "A",
            "px": "0.56899",
            "raw_book_diff": {
                "new": {
                    "sz": "2644.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009673,
            "coin": "TIA",
            "side": "A",
            "px": "0.56917",
            "raw_book_diff": {
                "new": {
                    "sz": "3067.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009674,
            "coin": "TIA",
            "side": "A",
            "px": "0.57018",
            "raw_book_diff": {
                "new": {
                    "sz": "2875.6"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009675,
            "coin": "ARB",
            "side": "B",
            "px": "0.19305",
            "raw_book_diff": {
                "new": {
                    "sz": "4083.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009676,
            "coin": "ARB",
            "side": "B",
            "px": "0.19292",
            "raw_book_diff": {
                "new": {
                    "sz": "8586.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009677,
            "coin": "ARB",
            "side": "B",
            "px": "0.19274",
            "raw_book_diff": {
                "new": {
                    "sz": "7659.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009678,
            "coin": "ARB",
            "side": "B",
            "px": "0.19252",
            "raw_book_diff": {
                "new": {
                    "sz": "8450.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009679,
            "coin": "ARB",
            "side": "B",
            "px": "0.19221",
            "raw_book_diff": {
                "new": {
                    "sz": "8851.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009680,
            "coin": "ARB",
            "side": "B",
            "px": "0.19169",
            "raw_book_diff": {
                "new": {
                    "sz": "8083.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009681,
            "coin": "ARB",
            "side": "B",
            "px": "0.19075",
            "raw_book_diff": {
                "new": {
                    "sz": "7735.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009682,
            "coin": "ARB",
            "side": "A",
            "px": "0.19359",
            "raw_book_diff": {
                "new": {
                    "sz": "8727.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009683,
            "coin": "ARB",
            "side": "A",
            "px": "0.1938",
            "raw_book_diff": {
                "new": {
                    "sz": "8724.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009684,
            "coin": "WLD",
            "side": "B",
            "px": "0.57117",
            "raw_book_diff": {
                "new": {
                    "sz": "1310.9"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009685,
            "coin": "WLD",
            "side": "A",
            "px": "0.57347",
            "raw_book_diff": {
                "new": {
                    "sz": "1365.9"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009686,
            "coin": "SUI",
            "side": "A",
            "px": "1.3509",
            "raw_book_diff": {
                "new": {
                    "sz": "385.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009687,
            "coin": "SUI",
            "side": "A",
            "px": "1.3561",
            "raw_book_diff": {
                "new": {
                    "sz": "758.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009688,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13626",
            "raw_book_diff": {
                "new": {
                    "sz": "10100.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009689,
            "coin": "kPEPE",
            "side": "A",
            "px": "0.004124",
            "raw_book_diff": {
                "new": {
                    "sz": "156664.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009690,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.09394",
            "raw_book_diff": {
                "new": {
                    "sz": "6733.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009691,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094813",
            "raw_book_diff": {
                "new": {
                    "sz": "13076.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009692,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.095136",
            "raw_book_diff": {
                "new": {
                    "sz": "11240.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009693,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.095482",
            "raw_book_diff": {
                "new": {
                    "sz": "11554.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009694,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.095872",
            "raw_book_diff": {
                "new": {
                    "sz": "14496.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009695,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008762",
            "raw_book_diff": {
                "new": {
                    "sz": "148117.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009696,
            "coin": "PENGU",
            "side": "A",
            "px": "0.010275",
            "raw_book_diff": {
                "new": {
                    "sz": "711804.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009697,
            "coin": "kFLOKI",
            "side": "B",
            "px": "0.042736",
            "raw_book_diff": {
                "new": {
                    "sz": "20184.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009698,
            "coin": "kFLOKI",
            "side": "A",
            "px": "0.042913",
            "raw_book_diff": {
                "new": {
                    "sz": "17239.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009699,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12011",
            "raw_book_diff": {
                "new": {
                    "sz": "4718.8"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009700,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.11999",
            "raw_book_diff": {
                "new": {
                    "sz": "9973.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009701,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.11985",
            "raw_book_diff": {
                "new": {
                    "sz": "9430.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009702,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.11965",
            "raw_book_diff": {
                "new": {
                    "sz": "9570.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009703,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.11939",
            "raw_book_diff": {
                "new": {
                    "sz": "13332.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009704,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.11896",
            "raw_book_diff": {
                "new": {
                    "sz": "12764.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009705,
            "coin": "kNEIRO",
            "side": "A",
            "px": "0.12046",
            "raw_book_diff": {
                "new": {
                    "sz": "13080.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009706,
            "coin": "kNEIRO",
            "side": "A",
            "px": "0.12059",
            "raw_book_diff": {
                "new": {
                    "sz": "12059.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009707,
            "coin": "kNEIRO",
            "side": "A",
            "px": "0.12078",
            "raw_book_diff": {
                "new": {
                    "sz": "12801.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009708,
            "coin": "kNEIRO",
            "side": "A",
            "px": "0.12102",
            "raw_book_diff": {
                "new": {
                    "sz": "12330.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009709,
            "coin": "kNEIRO",
            "side": "A",
            "px": "0.12141",
            "raw_book_diff": {
                "new": {
                    "sz": "13619.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009710,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82891",
            "raw_book_diff": {
                "new": {
                    "sz": "913.8"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009711,
            "coin": "VIRTUAL",
            "side": "A",
            "px": "0.82946",
            "raw_book_diff": {
                "new": {
                    "sz": "906.8"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009712,
            "coin": "IO",
            "side": "B",
            "px": "0.20048",
            "raw_book_diff": {
                "new": {
                    "sz": "3295.8"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009713,
            "coin": "IO",
            "side": "B",
            "px": "0.20039",
            "raw_book_diff": {
                "new": {
                    "sz": "7146.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009714,
            "coin": "IO",
            "side": "B",
            "px": "0.20028",
            "raw_book_diff": {
                "new": {
                    "sz": "6284.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009715,
            "coin": "IO",
            "side": "A",
            "px": "0.20061",
            "raw_book_diff": {
                "new": {
                    "sz": "4211.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009716,
            "coin": "IO",
            "side": "A",
            "px": "0.2007",
            "raw_book_diff": {
                "new": {
                    "sz": "7478.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009717,
            "coin": "IO",
            "side": "A",
            "px": "0.20081",
            "raw_book_diff": {
                "new": {
                    "sz": "7411.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009718,
            "coin": "IO",
            "side": "A",
            "px": "0.201",
            "raw_book_diff": {
                "new": {
                    "sz": "8058.8"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009719,
            "coin": "IO",
            "side": "A",
            "px": "0.20182",
            "raw_book_diff": {
                "new": {
                    "sz": "7268.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009720,
            "coin": "ETC",
            "side": "B",
            "px": "12.892",
            "raw_book_diff": {
                "new": {
                    "sz": "52.19"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009721,
            "coin": "ETC",
            "side": "B",
            "px": "12.889",
            "raw_book_diff": {
                "new": {
                    "sz": "105.59"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009722,
            "coin": "ETC",
            "side": "A",
            "px": "12.896",
            "raw_book_diff": {
                "new": {
                    "sz": "56.56"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009723,
            "coin": "ETC",
            "side": "A",
            "px": "12.899",
            "raw_book_diff": {
                "new": {
                    "sz": "135.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009724,
            "coin": "ETC",
            "side": "A",
            "px": "12.903",
            "raw_book_diff": {
                "new": {
                    "sz": "124.99"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009725,
            "coin": "LTC",
            "side": "B",
            "px": "76.653",
            "raw_book_diff": {
                "new": {
                    "sz": "11.28"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009726,
            "coin": "LTC",
            "side": "A",
            "px": "76.894",
            "raw_book_diff": {
                "new": {
                    "sz": "9.49"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009727,
            "coin": "XAI",
            "side": "B",
            "px": "0.01725",
            "raw_book_diff": {
                "new": {
                    "sz": "5995.8"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009728,
            "coin": "XAI",
            "side": "B",
            "px": "0.01723",
            "raw_book_diff": {
                "new": {
                    "sz": "37229.6"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009729,
            "coin": "XAI",
            "side": "A",
            "px": "0.01729",
            "raw_book_diff": {
                "new": {
                    "sz": "44056.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009730,
            "coin": "TAO",
            "side": "A",
            "px": "264.89",
            "raw_book_diff": {
                "new": {
                    "sz": "3.272"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009731,
            "coin": "APT",
            "side": "B",
            "px": "1.8568",
            "raw_book_diff": {
                "new": {
                    "sz": "329.9"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009732,
            "coin": "APT",
            "side": "B",
            "px": "1.8563",
            "raw_book_diff": {
                "new": {
                    "sz": "757.38"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009733,
            "coin": "APT",
            "side": "B",
            "px": "1.8517",
            "raw_book_diff": {
                "new": {
                    "sz": "928.89"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009734,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5397",
            "raw_book_diff": {
                "new": {
                    "sz": "189.91"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009735,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5385",
            "raw_book_diff": {
                "new": {
                    "sz": "400.46"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009736,
            "coin": "ORDI",
            "side": "B",
            "px": "3.537",
            "raw_book_diff": {
                "new": {
                    "sz": "369.78"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009737,
            "coin": "ORDI",
            "side": "A",
            "px": "3.5413",
            "raw_book_diff": {
                "new": {
                    "sz": "228.34"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009738,
            "coin": "ORDI",
            "side": "A",
            "px": "3.5425",
            "raw_book_diff": {
                "new": {
                    "sz": "454.39"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009739,
            "coin": "ORDI",
            "side": "A",
            "px": "3.5439",
            "raw_book_diff": {
                "new": {
                    "sz": "485.59"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009740,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6224",
            "raw_book_diff": {
                "new": {
                    "sz": "536.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009741,
            "coin": "NEAR",
            "side": "A",
            "px": "1.6345",
            "raw_book_diff": {
                "new": {
                    "sz": "445.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009742,
            "coin": "MANTA",
            "side": "A",
            "px": "0.0972",
            "raw_book_diff": {
                "new": {
                    "sz": "7658.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009743,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007996",
            "raw_book_diff": {
                "new": {
                    "sz": "93083.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009744,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007994",
            "raw_book_diff": {
                "new": {
                    "sz": "178667.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009745,
            "coin": "kSHIB",
            "side": "A",
            "px": "0.007999",
            "raw_book_diff": {
                "new": {
                    "sz": "92185.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009746,
            "coin": "kSHIB",
            "side": "A",
            "px": "0.008001",
            "raw_book_diff": {
                "new": {
                    "sz": "219128.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009747,
            "coin": "kSHIB",
            "side": "A",
            "px": "0.00802",
            "raw_book_diff": {
                "new": {
                    "sz": "198732.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009748,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008741",
            "raw_book_diff": {
                "new": {
                    "sz": "100435.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009749,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008735",
            "raw_book_diff": {
                "new": {
                    "sz": "197325.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009750,
            "coin": "kBONK",
            "side": "A",
            "px": "0.00875",
            "raw_book_diff": {
                "new": {
                    "sz": "77320.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009751,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008756",
            "raw_book_diff": {
                "new": {
                    "sz": "148465.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009752,
            "coin": "W",
            "side": "B",
            "px": "0.04029",
            "raw_book_diff": {
                "new": {
                    "sz": "19755.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009753,
            "coin": "W",
            "side": "A",
            "px": "0.04032",
            "raw_book_diff": {
                "new": {
                    "sz": "19429.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009754,
            "coin": "W",
            "side": "A",
            "px": "0.04033",
            "raw_book_diff": {
                "new": {
                    "sz": "36385.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009755,
            "coin": "W",
            "side": "A",
            "px": "0.04035",
            "raw_book_diff": {
                "new": {
                    "sz": "37066.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009756,
            "coin": "ME",
            "side": "B",
            "px": "0.30171",
            "raw_book_diff": {
                "new": {
                    "sz": "1975.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009757,
            "coin": "ME",
            "side": "B",
            "px": "0.30161",
            "raw_book_diff": {
                "new": {
                    "sz": "3771.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009758,
            "coin": "ME",
            "side": "B",
            "px": "0.30148",
            "raw_book_diff": {
                "new": {
                    "sz": "3731.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009759,
            "coin": "ME",
            "side": "A",
            "px": "0.30186",
            "raw_book_diff": {
                "new": {
                    "sz": "2162.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009760,
            "coin": "ME",
            "side": "A",
            "px": "0.30196",
            "raw_book_diff": {
                "new": {
                    "sz": "4295.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009761,
            "coin": "ME",
            "side": "A",
            "px": "0.3021",
            "raw_book_diff": {
                "new": {
                    "sz": "3896.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009762,
            "coin": "OM",
            "side": "A",
            "px": "0.07256",
            "raw_book_diff": {
                "new": {
                    "sz": "8105.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009763,
            "coin": "AR",
            "side": "B",
            "px": "3.9629",
            "raw_book_diff": {
                "new": {
                    "sz": "267.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009764,
            "coin": "AR",
            "side": "B",
            "px": "3.9608",
            "raw_book_diff": {
                "new": {
                    "sz": "302.42"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009765,
            "coin": "AR",
            "side": "B",
            "px": "3.9574",
            "raw_book_diff": {
                "new": {
                    "sz": "255.38"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009766,
            "coin": "AR",
            "side": "A",
            "px": "3.9673",
            "raw_book_diff": {
                "new": {
                    "sz": "321.76"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009767,
            "coin": "AR",
            "side": "A",
            "px": "3.9694",
            "raw_book_diff": {
                "new": {
                    "sz": "330.31"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009768,
            "coin": "2Z",
            "side": "B",
            "px": "0.10238",
            "raw_book_diff": {
                "new": {
                    "sz": "12353.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009769,
            "coin": "2Z",
            "side": "A",
            "px": "0.10346",
            "raw_book_diff": {
                "new": {
                    "sz": "12527.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009770,
            "coin": "UMA",
            "side": "A",
            "px": "0.77645",
            "raw_book_diff": {
                "new": {
                    "sz": "968.3"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009771,
            "coin": "UMA",
            "side": "A",
            "px": "0.77913",
            "raw_book_diff": {
                "new": {
                    "sz": "1889.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009772,
            "coin": "UMA",
            "side": "A",
            "px": "0.78085",
            "raw_book_diff": {
                "new": {
                    "sz": "1686.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009773,
            "coin": "UMA",
            "side": "A",
            "px": "0.78599",
            "raw_book_diff": {
                "new": {
                    "sz": "2200.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009774,
            "coin": "ACE",
            "side": "B",
            "px": "0.2156",
            "raw_book_diff": {
                "new": {
                    "sz": "2735.78"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009775,
            "coin": "ACE",
            "side": "B",
            "px": "0.2155",
            "raw_book_diff": {
                "new": {
                    "sz": "6128.87"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009776,
            "coin": "ACE",
            "side": "B",
            "px": "0.2154",
            "raw_book_diff": {
                "new": {
                    "sz": "5240.61"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009777,
            "coin": "ACE",
            "side": "A",
            "px": "0.2158",
            "raw_book_diff": {
                "new": {
                    "sz": "3234.25"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009778,
            "coin": "ACE",
            "side": "A",
            "px": "0.2159",
            "raw_book_diff": {
                "new": {
                    "sz": "5471.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009779,
            "coin": "ZEN",
            "side": "B",
            "px": "9.5246",
            "raw_book_diff": {
                "new": {
                    "sz": "59.95"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009780,
            "coin": "ZEN",
            "side": "B",
            "px": "9.5227",
            "raw_book_diff": {
                "new": {
                    "sz": "145.02"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009781,
            "coin": "ZEN",
            "side": "A",
            "px": "9.5271",
            "raw_book_diff": {
                "new": {
                    "sz": "65.97"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009782,
            "coin": "ZEN",
            "side": "A",
            "px": "9.529",
            "raw_book_diff": {
                "new": {
                    "sz": "135.47"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009783,
            "coin": "BIO",
            "side": "B",
            "px": "0.050922",
            "raw_book_diff": {
                "new": {
                    "sz": "14319.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009784,
            "coin": "BIO",
            "side": "B",
            "px": "0.050908",
            "raw_book_diff": {
                "new": {
                    "sz": "31321.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009785,
            "coin": "BIO",
            "side": "B",
            "px": "0.050892",
            "raw_book_diff": {
                "new": {
                    "sz": "29038.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009786,
            "coin": "BIO",
            "side": "B",
            "px": "0.050852",
            "raw_book_diff": {
                "new": {
                    "sz": "31555.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009787,
            "coin": "BIO",
            "side": "A",
            "px": "0.050953",
            "raw_book_diff": {
                "new": {
                    "sz": "28336.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009788,
            "coin": "BIO",
            "side": "A",
            "px": "0.050971",
            "raw_book_diff": {
                "new": {
                    "sz": "29714.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009789,
            "coin": "BIO",
            "side": "A",
            "px": "0.051013",
            "raw_book_diff": {
                "new": {
                    "sz": "28584.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009790,
            "coin": "SCR",
            "side": "B",
            "px": "0.08823",
            "raw_book_diff": {
                "new": {
                    "sz": "6846.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009791,
            "coin": "SCR",
            "side": "A",
            "px": "0.08835",
            "raw_book_diff": {
                "new": {
                    "sz": "8152.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009792,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041502",
            "raw_book_diff": {
                "new": {
                    "sz": "15517.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009793,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041519",
            "raw_book_diff": {
                "new": {
                    "sz": "32017.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009794,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041549",
            "raw_book_diff": {
                "new": {
                    "sz": "33607.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009795,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014519",
            "raw_book_diff": {
                "new": {
                    "sz": "23756.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009796,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01451",
            "raw_book_diff": {
                "new": {
                    "sz": "45779.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009797,
            "coin": "BRETT",
            "side": "B",
            "px": "0.0145",
            "raw_book_diff": {
                "new": {
                    "sz": "43724.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009798,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01448",
            "raw_book_diff": {
                "new": {
                    "sz": "46148.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009799,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014539",
            "raw_book_diff": {
                "new": {
                    "sz": "85598.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009800,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014549",
            "raw_book_diff": {
                "new": {
                    "sz": "85476.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009801,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014564",
            "raw_book_diff": {
                "new": {
                    "sz": "91346.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009802,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014701",
            "raw_book_diff": {
                "new": {
                    "sz": "92411.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009803,
            "coin": "POLYX",
            "side": "A",
            "px": "0.061352",
            "raw_book_diff": {
                "new": {
                    "sz": "10985.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009804,
            "coin": "PENDLE",
            "side": "B",
            "px": "2.4244",
            "raw_book_diff": {
                "new": {
                    "sz": "213.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009805,
            "coin": "PENDLE",
            "side": "B",
            "px": "2.4239",
            "raw_book_diff": {
                "new": {
                    "sz": "457.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009806,
            "coin": "PENDLE",
            "side": "A",
            "px": "2.4253",
            "raw_book_diff": {
                "new": {
                    "sz": "274.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009807,
            "coin": "PENDLE",
            "side": "A",
            "px": "2.4258",
            "raw_book_diff": {
                "new": {
                    "sz": "541.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009808,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053261",
            "raw_book_diff": {
                "new": {
                    "sz": "11422.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009809,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053213",
            "raw_book_diff": {
                "new": {
                    "sz": "25504.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009810,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053085",
            "raw_book_diff": {
                "new": {
                    "sz": "22818.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009811,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.052676",
            "raw_book_diff": {
                "new": {
                    "sz": "26304.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009812,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.052408",
            "raw_book_diff": {
                "new": {
                    "sz": "129410.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009813,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.053317",
            "raw_book_diff": {
                "new": {
                    "sz": "19621.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009814,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.053343",
            "raw_book_diff": {
                "new": {
                    "sz": "20804.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009815,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.053397",
            "raw_book_diff": {
                "new": {
                    "sz": "21059.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009816,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.053477",
            "raw_book_diff": {
                "new": {
                    "sz": "22780.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009817,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.053624",
            "raw_book_diff": {
                "new": {
                    "sz": "22128.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009818,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.053907",
            "raw_book_diff": {
                "new": {
                    "sz": "21956.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009819,
            "coin": "PROMPT",
            "side": "A",
            "px": "0.055226",
            "raw_book_diff": {
                "new": {
                    "sz": "109575.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009820,
            "coin": "RESOLV",
            "side": "A",
            "px": "0.077863",
            "raw_book_diff": {
                "new": {
                    "sz": "5749.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009821,
            "coin": "MOODENG",
            "side": "B",
            "px": "0.069535",
            "raw_book_diff": {
                "new": {
                    "sz": "10703.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009822,
            "coin": "MOODENG",
            "side": "B",
            "px": "0.069517",
            "raw_book_diff": {
                "new": {
                    "sz": "20319.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009823,
            "coin": "MOODENG",
            "side": "A",
            "px": "0.069557",
            "raw_book_diff": {
                "new": {
                    "sz": "11170.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009824,
            "coin": "MOODENG",
            "side": "A",
            "px": "0.069574",
            "raw_book_diff": {
                "new": {
                    "sz": "25111.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009825,
            "coin": "MOODENG",
            "side": "A",
            "px": "0.069601",
            "raw_book_diff": {
                "new": {
                    "sz": "21654.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009826,
            "coin": "KAS",
            "side": "A",
            "px": "0.051634",
            "raw_book_diff": {
                "new": {
                    "sz": "25868.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009827,
            "coin": "KAS",
            "side": "A",
            "px": "0.052015",
            "raw_book_diff": {
                "new": {
                    "sz": "22527.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009828,
            "coin": "ENS",
            "side": "A",
            "px": "10.547",
            "raw_book_diff": {
                "new": {
                    "sz": "59.02"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009829,
            "coin": "FXS",
            "side": "A",
            "px": "0.77832",
            "raw_book_diff": {
                "new": {
                    "sz": "855.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009830,
            "coin": "NOT",
            "side": "A",
            "px": "0.000518",
            "raw_book_diff": {
                "new": {
                    "sz": "3312366.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009831,
            "coin": "NOT",
            "side": "A",
            "px": "0.000519",
            "raw_book_diff": {
                "new": {
                    "sz": "3281189.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009832,
            "coin": "BSV",
            "side": "B",
            "px": "19.694",
            "raw_book_diff": {
                "new": {
                    "sz": "17.44"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009833,
            "coin": "BSV",
            "side": "B",
            "px": "19.672",
            "raw_book_diff": {
                "new": {
                    "sz": "36.84"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009834,
            "coin": "VVV",
            "side": "B",
            "px": "0.9478",
            "raw_book_diff": {
                "new": {
                    "sz": "436.42"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009835,
            "coin": "VVV",
            "side": "A",
            "px": "0.9491",
            "raw_book_diff": {
                "new": {
                    "sz": "708.58"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009836,
            "coin": "STX",
            "side": "B",
            "px": "0.28266",
            "raw_book_diff": {
                "new": {
                    "sz": "2190.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009837,
            "coin": "STX",
            "side": "B",
            "px": "0.28256",
            "raw_book_diff": {
                "new": {
                    "sz": "3935.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009838,
            "coin": "STX",
            "side": "B",
            "px": "0.28244",
            "raw_book_diff": {
                "new": {
                    "sz": "4343.6"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009839,
            "coin": "STX",
            "side": "A",
            "px": "0.28281",
            "raw_book_diff": {
                "new": {
                    "sz": "2055.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009840,
            "coin": "STX",
            "side": "A",
            "px": "0.28291",
            "raw_book_diff": {
                "new": {
                    "sz": "4110.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009841,
            "coin": "STX",
            "side": "A",
            "px": "0.28304",
            "raw_book_diff": {
                "new": {
                    "sz": "4115.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009842,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07056",
            "raw_book_diff": {
                "new": {
                    "sz": "7800.2"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009843,
            "coin": "SAGA",
            "side": "B",
            "px": "0.0705",
            "raw_book_diff": {
                "new": {
                    "sz": "16828.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009844,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07043",
            "raw_book_diff": {
                "new": {
                    "sz": "14617.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009845,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07032",
            "raw_book_diff": {
                "new": {
                    "sz": "14680.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009846,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07019",
            "raw_book_diff": {
                "new": {
                    "sz": "20036.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009847,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07068",
            "raw_book_diff": {
                "new": {
                    "sz": "9557.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009848,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07074",
            "raw_book_diff": {
                "new": {
                    "sz": "16417.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009849,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07081",
            "raw_book_diff": {
                "new": {
                    "sz": "17146.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009850,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07091",
            "raw_book_diff": {
                "new": {
                    "sz": "19754.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009851,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07105",
            "raw_book_diff": {
                "new": {
                    "sz": "17341.9"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009852,
            "coin": "ZETA",
            "side": "B",
            "px": "0.08187",
            "raw_book_diff": {
                "new": {
                    "sz": "8468.9"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009853,
            "coin": "SAND",
            "side": "B",
            "px": "0.14117",
            "raw_book_diff": {
                "new": {
                    "sz": "4773.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009854,
            "coin": "SAND",
            "side": "B",
            "px": "0.14107",
            "raw_book_diff": {
                "new": {
                    "sz": "10784.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009855,
            "coin": "SAND",
            "side": "A",
            "px": "0.14123",
            "raw_book_diff": {
                "new": {
                    "sz": "5669.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009856,
            "coin": "DOOD",
            "side": "A",
            "px": "0.003645",
            "raw_book_diff": {
                "new": {
                    "sz": "145076.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009857,
            "coin": "DOOD",
            "side": "A",
            "px": "0.003675",
            "raw_book_diff": {
                "new": {
                    "sz": "380786.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009858,
            "coin": "MEME",
            "side": "B",
            "px": "0.001178",
            "raw_book_diff": {
                "new": {
                    "sz": "546841.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009859,
            "coin": "MEME",
            "side": "B",
            "px": "0.001178",
            "raw_book_diff": {
                "new": {
                    "sz": "981923.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009860,
            "coin": "MEME",
            "side": "B",
            "px": "0.001177",
            "raw_book_diff": {
                "new": {
                    "sz": "995010.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009861,
            "coin": "MEME",
            "side": "A",
            "px": "0.00118",
            "raw_book_diff": {
                "new": {
                    "sz": "521606.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009862,
            "coin": "MEME",
            "side": "A",
            "px": "0.001181",
            "raw_book_diff": {
                "new": {
                    "sz": "1030897.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009863,
            "coin": "MEME",
            "side": "A",
            "px": "0.001182",
            "raw_book_diff": {
                "new": {
                    "sz": "1085823.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009864,
            "coin": "MEME",
            "side": "A",
            "px": "0.001188",
            "raw_book_diff": {
                "new": {
                    "sz": "984823.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009865,
            "coin": "PYTH",
            "side": "B",
            "px": "0.067504",
            "raw_book_diff": {
                "new": {
                    "sz": "10141.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009866,
            "coin": "PYTH",
            "side": "A",
            "px": "0.067718",
            "raw_book_diff": {
                "new": {
                    "sz": "5586.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009867,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14616",
            "raw_book_diff": {
                "new": {
                    "sz": "4416.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009868,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14609",
            "raw_book_diff": {
                "new": {
                    "sz": "7936.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009869,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14596",
            "raw_book_diff": {
                "new": {
                    "sz": "9436.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009870,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14582",
            "raw_book_diff": {
                "new": {
                    "sz": "8313.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009871,
            "coin": "WLFI",
            "side": "A",
            "px": "0.14638",
            "raw_book_diff": {
                "new": {
                    "sz": "8680.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009872,
            "coin": "WLFI",
            "side": "A",
            "px": "0.14651",
            "raw_book_diff": {
                "new": {
                    "sz": "8455.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009873,
            "coin": "WLFI",
            "side": "A",
            "px": "0.14665",
            "raw_book_diff": {
                "new": {
                    "sz": "8558.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009874,
            "coin": "COMP",
            "side": "B",
            "px": "34.812",
            "raw_book_diff": {
                "new": {
                    "sz": "19.44"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009875,
            "coin": "COMP",
            "side": "B",
            "px": "34.801",
            "raw_book_diff": {
                "new": {
                    "sz": "39.27"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009876,
            "coin": "COMP",
            "side": "B",
            "px": "34.782",
            "raw_book_diff": {
                "new": {
                    "sz": "35.31"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009877,
            "coin": "COMP",
            "side": "B",
            "px": "34.752",
            "raw_book_diff": {
                "new": {
                    "sz": "33.64"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009878,
            "coin": "COMP",
            "side": "A",
            "px": "34.843",
            "raw_book_diff": {
                "new": {
                    "sz": "22.28"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009879,
            "coin": "COMP",
            "side": "A",
            "px": "34.864",
            "raw_book_diff": {
                "new": {
                    "sz": "23.94"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009880,
            "coin": "COMP",
            "side": "A",
            "px": "34.9",
            "raw_book_diff": {
                "new": {
                    "sz": "24.64"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009881,
            "coin": "PNUT",
            "side": "A",
            "px": "0.08078",
            "raw_book_diff": {
                "new": {
                    "sz": "9111.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009882,
            "coin": "MAVIA",
            "side": "B",
            "px": "0.04734",
            "raw_book_diff": {
                "new": {
                    "sz": "5096.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009883,
            "coin": "MAVIA",
            "side": "A",
            "px": "0.0474",
            "raw_book_diff": {
                "new": {
                    "sz": "6272.4"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009884,
            "coin": "MAVIA",
            "side": "A",
            "px": "0.04743",
            "raw_book_diff": {
                "new": {
                    "sz": "14926.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009885,
            "coin": "ETHFI",
            "side": "A",
            "px": "0.74049",
            "raw_book_diff": {
                "new": {
                    "sz": "711.5"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009886,
            "coin": "USUAL",
            "side": "B",
            "px": "0.02466",
            "raw_book_diff": {
                "new": {
                    "sz": "27467.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009887,
            "coin": "USUAL",
            "side": "A",
            "px": "0.02478",
            "raw_book_diff": {
                "new": {
                    "sz": "14182.6"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009888,
            "coin": "EIGEN",
            "side": "A",
            "px": "0.505",
            "raw_book_diff": {
                "new": {
                    "sz": "1519.58"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009889,
            "coin": "EIGEN",
            "side": "A",
            "px": "0.5051",
            "raw_book_diff": {
                "new": {
                    "sz": "2876.06"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009890,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001874",
            "raw_book_diff": {
                "new": {
                    "sz": "497180.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009891,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001872",
            "raw_book_diff": {
                "new": {
                    "sz": "525019.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009892,
            "coin": "TURBO",
            "side": "A",
            "px": "0.001876",
            "raw_book_diff": {
                "new": {
                    "sz": "614773.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009893,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62148",
            "raw_book_diff": {
                "new": {
                    "sz": "890.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009894,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62126",
            "raw_book_diff": {
                "new": {
                    "sz": "1750.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009895,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62099",
            "raw_book_diff": {
                "new": {
                    "sz": "1877.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009896,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62042",
            "raw_book_diff": {
                "new": {
                    "sz": "1857.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009897,
            "coin": "GRASS",
            "side": "B",
            "px": "0.30312",
            "raw_book_diff": {
                "new": {
                    "sz": "1722.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009898,
            "coin": "GRASS",
            "side": "B",
            "px": "0.30301",
            "raw_book_diff": {
                "new": {
                    "sz": "3430.7"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009899,
            "coin": "GRASS",
            "side": "B",
            "px": "0.30286",
            "raw_book_diff": {
                "new": {
                    "sz": "3313.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009900,
            "coin": "GRASS",
            "side": "A",
            "px": "0.30338",
            "raw_book_diff": {
                "new": {
                    "sz": "3885.1"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009901,
            "coin": "GRASS",
            "side": "A",
            "px": "0.30353",
            "raw_book_diff": {
                "new": {
                    "sz": "4455.6"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009902,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041471",
            "raw_book_diff": {
                "new": {
                    "sz": "13358.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009903,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041453",
            "raw_book_diff": {
                "new": {
                    "sz": "27452.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009904,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041422",
            "raw_book_diff": {
                "new": {
                    "sz": "23481.0"
                }
            }
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433009905,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041381",
            "raw_book_diff": {
                "new": {
                    "sz": "25077.0"
                }
            }
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005961,
            "coin": "ZEC",
            "side": "B",
            "px": "347.64",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254433004805,
            "coin": "BANANA",
            "side": "B",
            "px": "8.0256",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254432959249,
            "coin": "ETH",
            "side": "A",
            "px": "2871.7",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433003544,
            "coin": "XRP",
            "side": "B",
            "px": "2.0266",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433002113,
            "coin": "XRP",
            "side": "A",
            "px": "2.0282",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433002208,
            "coin": "OM",
            "side": "A",
            "px": "0.07253",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433002220,
            "coin": "FXS",
            "side": "B",
            "px": "0.77788",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001070,
            "coin": "HBAR",
            "side": "B",
            "px": "0.1324",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003701,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33984",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003702,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33972",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005573,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6273",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005572,
            "coin": "NEAR",
            "side": "A",
            "px": "1.6312",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007429,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6261",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254432958156,
            "coin": "ETH",
            "side": "B",
            "px": "2795.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254432955168,
            "coin": "ETH",
            "side": "A",
            "px": "2834.7",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432928597,
            "coin": "@188",
            "side": "B",
            "px": "0.002647",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xae7a0f9f663bb54bfe95378c7e3de7dd6e28d6bc",
            "oid": 254432973656,
            "coin": "LINEA",
            "side": "B",
            "px": "0.009068",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432900580,
            "coin": "SNX",
            "side": "B",
            "px": "0.5096",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432954670,
            "coin": "HYPE",
            "side": "B",
            "px": "30.165",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432995963,
            "coin": "HYPE",
            "side": "A",
            "px": "30.289",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432960179,
            "coin": "HYPE",
            "side": "B",
            "px": "30.199",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc1914d36f97dc5557e4df26cbdab98e9c988ef37",
            "oid": 254432455776,
            "coin": "PENGU",
            "side": "A",
            "px": "0.009952",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf36c562be0598b8ce109998b6aa2d53ffcd8316b",
            "oid": 254432863782,
            "coin": "PUMP",
            "side": "B",
            "px": "0.002638",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x59d00a5a603ebe61cdce3f5ff54e9cfabe64ad7c",
            "oid": 254432975524,
            "coin": "DYM",
            "side": "A",
            "px": "0.09732",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433002185,
            "coin": "HYPE",
            "side": "A",
            "px": "30.268",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7fdafde5cfb5465924316eced2d3715494c517d1",
            "oid": 254433003818,
            "coin": "FARTCOIN",
            "side": "A",
            "px": "0.29812",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xae7a0f9f663bb54bfe95378c7e3de7dd6e28d6bc",
            "oid": 254433003543,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13611",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433007675,
            "coin": "@243",
            "side": "A",
            "px": "0.023611",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433007676,
            "coin": "@243",
            "side": "B",
            "px": "0.023573",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007651,
            "coin": "DOGE",
            "side": "A",
            "px": "0.1362",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd393fc74c7d944b8bb3c8e922cc9fa7fbcfa8c5d",
            "oid": 254432861092,
            "coin": "UNI",
            "side": "B",
            "px": "5.5145",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa716706dcad2e4460a53b304d703b9a2c4450b22",
            "oid": 254432851649,
            "coin": "UNI",
            "side": "B",
            "px": "5.5155",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdc3cc11485ecd459477afe0781ccc292de1f6a6c",
            "oid": 254432647193,
            "coin": "xyz:XYZ100",
            "side": "B",
            "px": "25172.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254432838414,
            "coin": "PENGU",
            "side": "A",
            "px": "0.009851",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254432975534,
            "coin": "ETH",
            "side": "A",
            "px": "2822.9",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31b1ac74a39fa92604c3bccac20453f2e3eebfaa",
            "oid": 254433003750,
            "coin": "FARTCOIN",
            "side": "A",
            "px": "0.29896",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe2d2dbae54981833479038c70425e1e3ce6680ec",
            "oid": 254432928353,
            "coin": "STRK",
            "side": "B",
            "px": "0.11782",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb2b1d1a7a034d9209a627e754ebb8d24988ca23a",
            "oid": 254432945017,
            "coin": "FET",
            "side": "A",
            "px": "0.23026",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xfc8c156428a8e48cb8d0356db16e59bec4c0ecea",
            "oid": 254432925488,
            "coin": "OP",
            "side": "A",
            "px": "0.2919",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433002186,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13615",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432996082,
            "coin": "MOVE",
            "side": "B",
            "px": "0.045801",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432998498,
            "coin": "TON",
            "side": "B",
            "px": "1.4877",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433003822,
            "coin": "TON",
            "side": "B",
            "px": "1.4869",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432862784,
            "coin": "LDO",
            "side": "A",
            "px": "0.57787",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005426,
            "coin": "kFLOKI",
            "side": "B",
            "px": "0.042871",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0631897de2f668ff8dfe5bfe23d63d5634035ce5",
            "oid": 254433008128,
            "coin": "BTC",
            "side": "B",
            "px": "85410.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432903941,
            "coin": "AVAX",
            "side": "A",
            "px": "12.98",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432992743,
            "coin": "ZETA",
            "side": "B",
            "px": "0.08189",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432980987,
            "coin": "ZRO",
            "side": "B",
            "px": "1.2204",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433001728,
            "coin": "MOVE",
            "side": "B",
            "px": "0.04574",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432981041,
            "coin": "ZRO",
            "side": "A",
            "px": "1.2219",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432825207,
            "coin": "ALT",
            "side": "A",
            "px": "0.012513",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432999697,
            "coin": "TON",
            "side": "A",
            "px": "1.4924",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432944368,
            "coin": "JTO",
            "side": "A",
            "px": "0.41857",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432940443,
            "coin": "ALT",
            "side": "B",
            "px": "0.012488",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x34827044cbd4b808fc1b189fce9f50e6dafae7c9",
            "oid": 254433007406,
            "coin": "COMP",
            "side": "B",
            "px": "34.812",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9f052eabbe59d1d6e324232a73ef7e1faa66bd54",
            "oid": 254432867339,
            "coin": "CELO",
            "side": "A",
            "px": "0.15384",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432965962,
            "coin": "BNB",
            "side": "B",
            "px": "820.34",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e5c7a2fed36727acf6306e3e5ae12bce86856c7",
            "oid": 254432913469,
            "coin": "ALT",
            "side": "A",
            "px": "0.012533",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007677,
            "coin": "XPL",
            "side": "B",
            "px": "0.17749",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xde236b0b3b367d5ad2a2c42b914bded79d03fe28",
            "oid": 254432375023,
            "coin": "KAS",
            "side": "A",
            "px": "0.051755",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432988047,
            "coin": "HYPE",
            "side": "B",
            "px": "30.192",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433003492,
            "coin": "LINK",
            "side": "A",
            "px": "12.126",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x6c2e4912a03c663eaa74a0bd0bbc3bc4c7c3d43d",
            "oid": 254432881697,
            "coin": "KAS",
            "side": "B",
            "px": "0.05138",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007712,
            "coin": "@162",
            "side": "A",
            "px": "0.29827",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433005849,
            "coin": "xyz:XYZ100",
            "side": "A",
            "px": "25241.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433005851,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "277.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433005852,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "278.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433005855,
            "coin": "xyz:AAPL",
            "side": "B",
            "px": "269.11",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x2025137a136bea7446deba681cbfc7cf1970840e",
            "oid": 254433003515,
            "coin": "xyz:AAPL",
            "side": "A",
            "px": "278.19",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd9a5ae508abffe8231b8d6ab90c2738052f3424a",
            "oid": 254432997670,
            "coin": "SPX",
            "side": "A",
            "px": "0.6104",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9266865bb6afb4c4f618544dd3b8c970f17aa664",
            "oid": 254433006045,
            "coin": "COMP",
            "side": "A",
            "px": "34.849",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433005757,
            "coin": "PYTH",
            "side": "B",
            "px": "0.06766",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xeaa123246b396028fc8fea0175f013c13e097157",
            "oid": 254432672632,
            "coin": "ONDO",
            "side": "A",
            "px": "0.46081",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7f3d2fcb1d78872135ea9585b354c49b63092110",
            "oid": 254432680896,
            "coin": "vntl:SPACEX",
            "side": "A",
            "px": "515.16",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432974670,
            "coin": "BNB",
            "side": "A",
            "px": "821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432866864,
            "coin": "AVAX",
            "side": "A",
            "px": "12.945",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
            "oid": 254433005323,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968952,
            "coin": "KAS",
            "side": "B",
            "px": "0.051471",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432923036,
            "coin": "ZRO",
            "side": "B",
            "px": "1.2183",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x8ae62e9a4775350022c1577b2eff694e2b72be8d",
            "oid": 254433005038,
            "coin": "ETH",
            "side": "B",
            "px": "2821.1",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433003589,
            "coin": "TAO",
            "side": "A",
            "px": "264.69",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xfe72c242af1a4ca8114c6a170ba4a823785241ce",
            "oid": 254433004691,
            "coin": "TAO",
            "side": "A",
            "px": "264.66",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xef5fdcefdb38803a654ee9f3539bed8058443d99",
            "oid": 254432959668,
            "coin": "TRUMP",
            "side": "A",
            "px": "5.7448",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001000,
            "coin": "COMP",
            "side": "B",
            "px": "34.802",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7f3d2fcb1d78872135ea9585b354c49b63092110",
            "oid": 254432681534,
            "coin": "vntl:SPACEX",
            "side": "A",
            "px": "515.16",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x399965e15d4e61ec3529cc98b7f7ebb93b733336",
            "oid": 254432963443,
            "coin": "@142",
            "side": "B",
            "px": "85455.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd63c228ea09e4d999f5fb3ed7d3529756e3a724f",
            "oid": 254433006397,
            "coin": "PENGU",
            "side": "A",
            "px": "0.009857",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd63c228ea09e4d999f5fb3ed7d3529756e3a724f",
            "oid": 254433006398,
            "coin": "PENGU",
            "side": "A",
            "px": "0.009852",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd63c228ea09e4d999f5fb3ed7d3529756e3a724f",
            "oid": 254433008036,
            "coin": "COMP",
            "side": "A",
            "px": "34.94",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd63c228ea09e4d999f5fb3ed7d3529756e3a724f",
            "oid": 254433008037,
            "coin": "COMP",
            "side": "A",
            "px": "34.923",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432922968,
            "coin": "ENS",
            "side": "B",
            "px": "10.528",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432842784,
            "coin": "ENS",
            "side": "A",
            "px": "10.553",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432872003,
            "coin": "ENS",
            "side": "A",
            "px": "10.574",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432985904,
            "coin": "SUI",
            "side": "A",
            "px": "1.3512",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe2d2dbae54981833479038c70425e1e3ce6680ec",
            "oid": 254433005074,
            "coin": "PENGU",
            "side": "B",
            "px": "0.009824",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432977709,
            "coin": "SUI",
            "side": "B",
            "px": "1.3505",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432989610,
            "coin": "SUI",
            "side": "B",
            "px": "1.3506",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432989866,
            "coin": "ENS",
            "side": "A",
            "px": "10.562",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433004152,
            "coin": "@184",
            "side": "B",
            "px": "0.009816",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb4321b142b2a03ce20fcab2007ff6990b9acba93",
            "oid": 254432965485,
            "coin": "PEOPLE",
            "side": "B",
            "px": "0.00914",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432965347,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5405",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432968161,
            "coin": "WLFI",
            "side": "A",
            "px": "0.14644",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254432960036,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008749",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433003198,
            "coin": "kBONK",
            "side": "B",
            "px": "0.00873",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432897082,
            "coin": "BTC",
            "side": "A",
            "px": "85504.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432906805,
            "coin": "BTC",
            "side": "B",
            "px": "85378.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432965898,
            "coin": "BTC",
            "side": "A",
            "px": "85449.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432948467,
            "coin": "SOL",
            "side": "B",
            "px": "126.55",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432957081,
            "coin": "SOL",
            "side": "B",
            "px": "126.43",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432956166,
            "coin": "SOL",
            "side": "A",
            "px": "126.82",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432997303,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041401",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433003107,
            "coin": "LTC",
            "side": "B",
            "px": "76.813",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x97f20f91c18093c7f738a00ffabc03eea04ec668",
            "oid": 254432998269,
            "coin": "PENGU",
            "side": "A",
            "px": "0.009848",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433003539,
            "coin": "LTC",
            "side": "A",
            "px": "76.929",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432996132,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041421",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
            "oid": 254432987582,
            "coin": "BIO",
            "side": "B",
            "px": "0.050903",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7a5a42bd6b1fd13368732be8a7b63e06064e0394",
            "oid": 254432994760,
            "coin": "SOL",
            "side": "B",
            "px": "126.54",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xca163b98cbc2ad5b08b01ae7bcb4bef68b9161ad",
            "oid": 254433007625,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008746",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432976546,
            "coin": "SOL",
            "side": "A",
            "px": "126.69",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432946408,
            "coin": "JUP",
            "side": "A",
            "px": "0.22295",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x3f3e7878cbde6da00629f6f6115f367f8ef316d2",
            "oid": 254432974540,
            "coin": "MANTA",
            "side": "B",
            "px": "0.09715",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432833563,
            "coin": "DOGE",
            "side": "A",
            "px": "0.1368",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432833564,
            "coin": "DOGE",
            "side": "B",
            "px": "0.1357",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433002182,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13605",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432897699,
            "coin": "BCH",
            "side": "B",
            "px": "520.09",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433004479,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82923",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432944328,
            "coin": "SUI",
            "side": "B",
            "px": "1.3504",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433001508,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6287",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254432989415,
            "coin": "SUI",
            "side": "A",
            "px": "1.351",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254431271600,
            "coin": "SUI",
            "side": "B",
            "px": "1.3502",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008431,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.829",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432880339,
            "coin": "SEI",
            "side": "A",
            "px": "0.12455",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432972684,
            "coin": "@206",
            "side": "B",
            "px": "0.24157",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x97f20f91c18093c7f738a00ffabc03eea04ec668",
            "oid": 254432877535,
            "coin": "SOL",
            "side": "A",
            "px": "126.72",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432917975,
            "coin": "BCH",
            "side": "B",
            "px": "520.65",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433005268,
            "coin": "@156",
            "side": "B",
            "px": "126.6",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254432965610,
            "coin": "TIA",
            "side": "B",
            "px": "0.56809",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432974821,
            "coin": "SUI",
            "side": "B",
            "px": "1.3503",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x20faffd0045455e86b36b827d68f8f99f9c929b4",
            "oid": 254432871557,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
            "oid": 254432966298,
            "coin": "SUI",
            "side": "B",
            "px": "1.35",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433000315,
            "coin": "PNUT",
            "side": "B",
            "px": "0.0806",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254432965629,
            "coin": "TIA",
            "side": "A",
            "px": "0.57035",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432909565,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094551",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432742850,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094465",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432863969,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014523",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x72cb918356c4f6d3f1b2e532928110ba6995f139",
            "oid": 254433008196,
            "coin": "@156",
            "side": "B",
            "px": "126.46",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254432796153,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094311",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432973484,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014535",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254432801854,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094616",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254432186188,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.09449",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007366,
            "coin": "ASTER",
            "side": "B",
            "px": "0.93755",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432975931,
            "coin": "UNI",
            "side": "B",
            "px": "5.5172",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432998429,
            "coin": "MOODENG",
            "side": "B",
            "px": "0.069451",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433005336,
            "coin": "TIA",
            "side": "A",
            "px": "0.56911",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254433001745,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6277",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007352,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5386",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf4e62d9d431a07b9b1efb5dd548e3ad9aadd07ba",
            "oid": 254432966011,
            "coin": "NEAR",
            "side": "B",
            "px": "1.6301",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254432923621,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13614",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433007160,
            "coin": "APT",
            "side": "A",
            "px": "1.8584",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433007161,
            "coin": "APT",
            "side": "B",
            "px": "1.8562",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254431936515,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.09443",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x72cb918356c4f6d3f1b2e532928110ba6995f139",
            "oid": 254432972383,
            "coin": "SOL",
            "side": "B",
            "px": "126.53",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254433008361,
            "coin": "PENGU",
            "side": "B",
            "px": "0.009821",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433006347,
            "coin": "ENA",
            "side": "A",
            "px": "0.24217",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433003807,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008755",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433004615,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13605",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254432852262,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094498",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xae7a0f9f663bb54bfe95378c7e3de7dd6e28d6bc",
            "oid": 254432958504,
            "coin": "KAITO",
            "side": "B",
            "px": "0.6216",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433006490,
            "coin": "TIA",
            "side": "B",
            "px": "0.5688",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432979694,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014495",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254432836040,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432997450,
            "coin": "SUI",
            "side": "B",
            "px": "1.3507",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005545,
            "coin": "MOODENG",
            "side": "A",
            "px": "0.069665",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432878901,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014507",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433009022,
            "coin": "ASTER",
            "side": "B",
            "px": "0.93884",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432962087,
            "coin": "CFX",
            "side": "B",
            "px": "0.069643",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254432796392,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094808",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008995,
            "coin": "SUI",
            "side": "B",
            "px": "1.3503",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433008994,
            "coin": "SUI",
            "side": "B",
            "px": "1.3502",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254432802607,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094481",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254433003265,
            "coin": "APT",
            "side": "A",
            "px": "1.8595",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xae7a0f9f663bb54bfe95378c7e3de7dd6e28d6bc",
            "oid": 254432999207,
            "coin": "HYPE",
            "side": "A",
            "px": "30.232",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433002368,
            "coin": "SUI",
            "side": "B",
            "px": "1.3506",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432985611,
            "coin": "SUI",
            "side": "A",
            "px": "1.3511",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc1fce740d83a60de67d039aa927a678ff78c202f",
            "oid": 254432996603,
            "coin": "UNI",
            "side": "B",
            "px": "5.5173",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31b1ac74a39fa92604c3bccac20453f2e3eebfaa",
            "oid": 254432454225,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01449",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xac7476e14f768e3e67c195e79c2490dd20c70127",
            "oid": 254432996104,
            "coin": "LTC",
            "side": "A",
            "px": "76.909",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf3ea98c72ff0b195db5512abe4c6e4e6fe8534b4",
            "oid": 254433004688,
            "coin": "BSV",
            "side": "B",
            "px": "19.718",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc2e137ea6b79a7992dd54d44ee8223f2c7e13e22",
            "oid": 254433002700,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014535",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432965978,
            "coin": "ENS",
            "side": "B",
            "px": "10.54",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432907158,
            "coin": "AAVE",
            "side": "B",
            "px": "168.06",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432989450,
            "coin": "TIA",
            "side": "A",
            "px": "0.56988",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432917496,
            "coin": "AAVE",
            "side": "A",
            "px": "168.66",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xeaa123246b396028fc8fea0175f013c13e097157",
            "oid": 254432481633,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014491",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432944635,
            "coin": "AAVE",
            "side": "A",
            "px": "169.02",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xfe96dd504543165efeb7ab545d161d0ba112f84e",
            "oid": 254433003853,
            "coin": "EIGEN",
            "side": "B",
            "px": "0.5045",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x27c9fa86c91b84ddfa15de58c482ff662498d65d",
            "oid": 254432887650,
            "coin": "@151",
            "side": "A",
            "px": "2825.4",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x34fb5ec7d4e939161946340ea2a1f29254b893de",
            "oid": 254431210939,
            "coin": "APT",
            "side": "A",
            "px": "1.8644",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432933035,
            "coin": "AAVE",
            "side": "B",
            "px": "168.38",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x92c7d53dc3bac14722b281e74a3feba8f7ba040e",
            "oid": 254432517791,
            "coin": "CRV",
            "side": "B",
            "px": "0.38208",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x34827044cbd4b808fc1b189fce9f50e6dafae7c9",
            "oid": 254432830268,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014527",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433008947,
            "coin": "ASTER",
            "side": "B",
            "px": "0.93948",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432968210,
            "coin": "XPL",
            "side": "B",
            "px": "0.17745",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433000039,
            "coin": "HYPE",
            "side": "B",
            "px": "30.226",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc1fce740d83a60de67d039aa927a678ff78c202f",
            "oid": 254432985724,
            "coin": "UNI",
            "side": "B",
            "px": "5.5174",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003590,
            "coin": "TIA",
            "side": "B",
            "px": "0.56877",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc1fce740d83a60de67d039aa927a678ff78c202f",
            "oid": 254432960142,
            "coin": "UNI",
            "side": "B",
            "px": "5.5176",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432948208,
            "coin": "CHILLGUY",
            "side": "B",
            "px": "0.018615",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x80196a7475464e06de641df6c66e6e278cd364d5",
            "oid": 254433005884,
            "coin": "@107",
            "side": "A",
            "px": "30.219",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x80196a7475464e06de641df6c66e6e278cd364d5",
            "oid": 254433005883,
            "coin": "@107",
            "side": "A",
            "px": "30.226",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9046c5826a7e3a3c2366be1307ed604d64deffda",
            "oid": 254432994614,
            "coin": "PURR/USDC",
            "side": "B",
            "px": "0.074759",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0ab2ff3a85eb341ff9f9307858b94904a42c5b45",
            "oid": 254432752407,
            "coin": "BTC",
            "side": "A",
            "px": "85544.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0ab2ff3a85eb341ff9f9307858b94904a42c5b45",
            "oid": 254432799675,
            "coin": "BTC",
            "side": "A",
            "px": "85544.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003771,
            "coin": "APT",
            "side": "B",
            "px": "1.857",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968518,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5407",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968519,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5395",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432837663,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008743",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432999273,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82893",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432949999,
            "coin": "SUI",
            "side": "A",
            "px": "1.3546",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432849344,
            "coin": "SUI",
            "side": "A",
            "px": "1.3531",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432951607,
            "coin": "CHILLGUY",
            "side": "B",
            "px": "0.018586",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x5fffee2555a15899ad656c1a80f1b35cd0b2c0c1",
            "oid": 254432945024,
            "coin": "@156",
            "side": "B",
            "px": "126.61",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa775d1bd91ee2cc4cae6113e5fd9c0c17b355277",
            "oid": 254432882206,
            "coin": "SOL",
            "side": "A",
            "px": "126.61",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432909378,
            "coin": "SUI",
            "side": "B",
            "px": "1.3479",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432952933,
            "coin": "ME",
            "side": "B",
            "px": "0.30181",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432952934,
            "coin": "ME",
            "side": "B",
            "px": "0.30171",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001058,
            "coin": "BIO",
            "side": "B",
            "px": "0.050942",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001059,
            "coin": "BIO",
            "side": "B",
            "px": "0.050928",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001060,
            "coin": "BIO",
            "side": "B",
            "px": "0.050912",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432907290,
            "coin": "SAND",
            "side": "B",
            "px": "0.1412",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432768942,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432480826,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432705718,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432832699,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001875",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432445458,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001874",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432958271,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62158",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001024,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.04149",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825816,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014528",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825817,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01452",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x621c5551678189b9a6c94d929924c225ff1d63ab",
            "oid": 254432966625,
            "coin": "HYPE",
            "side": "B",
            "px": "30.199",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x20faffd0045455e86b36b827d68f8f99f9c929b4",
            "oid": 254432962379,
            "coin": "ENS",
            "side": "B",
            "px": "10.544",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432897713,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014574",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd9832441317a2dfaee10effcecfacc7e839ad37e",
            "oid": 254432956340,
            "coin": "SOL",
            "side": "B",
            "px": "126.52",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xeaa123246b396028fc8fea0175f013c13e097157",
            "oid": 254432824821,
            "coin": "kBONK",
            "side": "A",
            "px": "0.008784",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xde236b0b3b367d5ad2a2c42b914bded79d03fe28",
            "oid": 254432481654,
            "coin": "SUI",
            "side": "B",
            "px": "1.348",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd911e53d53b663972254e086450fd6198a25961e",
            "oid": 254433008121,
            "coin": "@151",
            "side": "B",
            "px": "2821.1",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xfc667adba8d4837586078f4fdcdc29804337ca06",
            "oid": 254433005542,
            "coin": "BTC",
            "side": "B",
            "px": "85384.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0fd468a73084daa6ea77a9261e40fdec3e67e0c7",
            "oid": 254433000093,
            "coin": "BTC",
            "side": "B",
            "px": "85406.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001068,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07066",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001069,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07078",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432987769,
            "coin": "ZORA",
            "side": "A",
            "px": "0.043836",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432946847,
            "coin": "ZORA",
            "side": "A",
            "px": "0.043953",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432488365,
            "coin": "SAND",
            "side": "A",
            "px": "0.14211",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432793816,
            "coin": "MEME",
            "side": "A",
            "px": "0.001181",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432768944,
            "coin": "MEME",
            "side": "A",
            "px": "0.001182",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432480832,
            "coin": "MEME",
            "side": "A",
            "px": "0.001183",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432412062,
            "coin": "MEME",
            "side": "A",
            "px": "0.001185",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432412063,
            "coin": "MEME",
            "side": "A",
            "px": "0.001187",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432977310,
            "coin": "MOVE",
            "side": "B",
            "px": "0.045837",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432977311,
            "coin": "MOVE",
            "side": "B",
            "px": "0.045812",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968442,
            "coin": "MOVE",
            "side": "A",
            "px": "0.045895",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968443,
            "coin": "MOVE",
            "side": "A",
            "px": "0.045912",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968444,
            "coin": "MOVE",
            "side": "A",
            "px": "0.045931",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974356,
            "coin": "MOVE",
            "side": "A",
            "px": "0.046049",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432906457,
            "coin": "PAXG",
            "side": "B",
            "px": "4262.1",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974358,
            "coin": "PAXG",
            "side": "B",
            "px": "4259.3",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432906459,
            "coin": "PAXG",
            "side": "B",
            "px": "4255.4",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432949657,
            "coin": "PAXG",
            "side": "A",
            "px": "4263.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974359,
            "coin": "PAXG",
            "side": "A",
            "px": "4264.2",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974360,
            "coin": "PAXG",
            "side": "A",
            "px": "4266.4",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432906463,
            "coin": "PAXG",
            "side": "A",
            "px": "4271.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432836857,
            "coin": "PYTH",
            "side": "B",
            "px": "0.067501",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432897797,
            "coin": "PYTH",
            "side": "A",
            "px": "0.067709",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945226,
            "coin": "STBL",
            "side": "A",
            "px": "0.061545",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945227,
            "coin": "STBL",
            "side": "A",
            "px": "0.061595",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001001,
            "coin": "COMP",
            "side": "A",
            "px": "34.85",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001002,
            "coin": "COMP",
            "side": "A",
            "px": "34.863",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001003,
            "coin": "COMP",
            "side": "A",
            "px": "34.884",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974364,
            "coin": "COMP",
            "side": "B",
            "px": "34.771",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974369,
            "coin": "COMP",
            "side": "A",
            "px": "34.921",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001009,
            "coin": "AVNT",
            "side": "B",
            "px": "0.33381",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001010,
            "coin": "AVNT",
            "side": "B",
            "px": "0.33355",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974373,
            "coin": "AVNT",
            "side": "A",
            "px": "0.33426",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968447,
            "coin": "AVNT",
            "side": "A",
            "px": "0.33437",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974374,
            "coin": "AVNT",
            "side": "A",
            "px": "0.33458",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001012,
            "coin": "AVNT",
            "side": "A",
            "px": "0.33552",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001013,
            "coin": "PNUT",
            "side": "B",
            "px": "0.08045",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001072,
            "coin": "PNUT",
            "side": "A",
            "px": "0.0808",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432993358,
            "coin": "APEX",
            "side": "B",
            "px": "0.53073",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432993359,
            "coin": "APEX",
            "side": "B",
            "px": "0.53043",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432993360,
            "coin": "APEX",
            "side": "B",
            "px": "0.53005",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432993361,
            "coin": "APEX",
            "side": "B",
            "px": "0.52946",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432993362,
            "coin": "APEX",
            "side": "B",
            "px": "0.5286",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974378,
            "coin": "APEX",
            "side": "B",
            "px": "0.52697",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974383,
            "coin": "APEX",
            "side": "A",
            "px": "0.53239",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974384,
            "coin": "APEX",
            "side": "A",
            "px": "0.53323",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432965467,
            "coin": "MAVIA",
            "side": "A",
            "px": "0.04743",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001073,
            "coin": "MAVIA",
            "side": "B",
            "px": "0.04733",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001074,
            "coin": "MAVIA",
            "side": "A",
            "px": "0.0474",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433002230,
            "coin": "MAVIA",
            "side": "A",
            "px": "0.04743",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825847,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33974",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825848,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33962",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825849,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33948",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825852,
            "coin": "SUSHI",
            "side": "A",
            "px": "0.34014",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003703,
            "coin": "SUSHI",
            "side": "B",
            "px": "0.33958",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003704,
            "coin": "SUSHI",
            "side": "A",
            "px": "0.33999",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003705,
            "coin": "SUSHI",
            "side": "A",
            "px": "0.3401",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003706,
            "coin": "SUSHI",
            "side": "A",
            "px": "0.34024",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001075,
            "coin": "USUAL",
            "side": "B",
            "px": "0.02466",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001076,
            "coin": "USUAL",
            "side": "A",
            "px": "0.02478",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001020,
            "coin": "EIGEN",
            "side": "B",
            "px": "0.5045",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001077,
            "coin": "EIGEN",
            "side": "A",
            "px": "0.5051",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001078,
            "coin": "EIGEN",
            "side": "A",
            "px": "0.5052",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432829986,
            "coin": "TURBO",
            "side": "A",
            "px": "0.001877",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432958272,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62141",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432980813,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62119",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432958274,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62067",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968965,
            "coin": "HYPER",
            "side": "B",
            "px": "0.12029",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968966,
            "coin": "HYPER",
            "side": "B",
            "px": "0.12016",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968968,
            "coin": "HYPER",
            "side": "A",
            "px": "0.12058",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968969,
            "coin": "HYPER",
            "side": "A",
            "px": "0.12065",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968970,
            "coin": "HYPER",
            "side": "A",
            "px": "0.12075",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968971,
            "coin": "HYPER",
            "side": "A",
            "px": "0.12088",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968972,
            "coin": "HYPER",
            "side": "A",
            "px": "0.12108",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001025,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041478",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974405,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.04143",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001027,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041512",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001028,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041524",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001029,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.04155",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968405,
            "coin": "AIXBT",
            "side": "A",
            "px": "0.041593",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825818,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014509",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825819,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01449",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432848798,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014542",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825821,
            "coin": "BRETT",
            "side": "A",
            "px": "0.01455",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825822,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014559",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825823,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014574",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432251736,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014744",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432866500,
            "coin": "POLYX",
            "side": "A",
            "px": "0.061348",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432995961,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.076533",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001031,
            "coin": "RESOLV",
            "side": "A",
            "px": "0.077926",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003707,
            "coin": "RESOLV",
            "side": "B",
            "px": "0.076445",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432944742,
            "coin": "BIGTIME",
            "side": "B",
            "px": "0.022328",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432944743,
            "coin": "BIGTIME",
            "side": "B",
            "px": "0.022307",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432944744,
            "coin": "BIGTIME",
            "side": "A",
            "px": "0.022362",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432944745,
            "coin": "BIGTIME",
            "side": "A",
            "px": "0.022371",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432944746,
            "coin": "BIGTIME",
            "side": "A",
            "px": "0.022381",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432960072,
            "coin": "BIGTIME",
            "side": "A",
            "px": "0.022405",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432996034,
            "coin": "ZEREBRO",
            "side": "B",
            "px": "0.027715",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001047,
            "coin": "W",
            "side": "B",
            "px": "0.04029",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001048,
            "coin": "W",
            "side": "A",
            "px": "0.04032",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001049,
            "coin": "W",
            "side": "A",
            "px": "0.04033",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001050,
            "coin": "W",
            "side": "A",
            "px": "0.04035",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432952935,
            "coin": "ME",
            "side": "B",
            "px": "0.30158",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945254,
            "coin": "ME",
            "side": "A",
            "px": "0.30196",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432952938,
            "coin": "ME",
            "side": "A",
            "px": "0.30206",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432952939,
            "coin": "ME",
            "side": "A",
            "px": "0.3022",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432909268,
            "coin": "OM",
            "side": "A",
            "px": "0.07264",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432909269,
            "coin": "OM",
            "side": "A",
            "px": "0.07273",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433000974,
            "coin": "OM",
            "side": "B",
            "px": "0.07247",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433000975,
            "coin": "OM",
            "side": "B",
            "px": "0.07237",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001051,
            "coin": "OM",
            "side": "B",
            "px": "0.0725",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001052,
            "coin": "OM",
            "side": "B",
            "px": "0.07245",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003678,
            "coin": "OM",
            "side": "A",
            "px": "0.07263",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003679,
            "coin": "OM",
            "side": "A",
            "px": "0.0727",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433003680,
            "coin": "OM",
            "side": "A",
            "px": "0.0728",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432980792,
            "coin": "IP",
            "side": "B",
            "px": "2.2638",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432985539,
            "coin": "IP",
            "side": "A",
            "px": "2.268",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432980795,
            "coin": "IP",
            "side": "A",
            "px": "2.2689",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432980796,
            "coin": "IP",
            "side": "A",
            "px": "2.2705",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432980797,
            "coin": "IP",
            "side": "A",
            "px": "2.2726",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968489,
            "coin": "IP",
            "side": "A",
            "px": "2.2768",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432990916,
            "coin": "AR",
            "side": "B",
            "px": "3.9634",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432990917,
            "coin": "AR",
            "side": "A",
            "px": "3.9681",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432990918,
            "coin": "AR",
            "side": "A",
            "px": "3.9698",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432996030,
            "coin": "AR",
            "side": "B",
            "px": "3.9618",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432990919,
            "coin": "AR",
            "side": "A",
            "px": "3.9724",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968947,
            "coin": "2Z",
            "side": "B",
            "px": "0.10231",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968499,
            "coin": "TRB",
            "side": "B",
            "px": "19.929",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968503,
            "coin": "TRB",
            "side": "A",
            "px": "19.948",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968504,
            "coin": "TRB",
            "side": "A",
            "px": "19.954",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432987768,
            "coin": "ZEC",
            "side": "A",
            "px": "349.05",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432987767,
            "coin": "ZEC",
            "side": "B",
            "px": "348.1",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001053,
            "coin": "ACE",
            "side": "B",
            "px": "0.2156",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001054,
            "coin": "ACE",
            "side": "B",
            "px": "0.2155",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001055,
            "coin": "ACE",
            "side": "B",
            "px": "0.2154",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001056,
            "coin": "ACE",
            "side": "A",
            "px": "0.2158",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001057,
            "coin": "ACE",
            "side": "A",
            "px": "0.2159",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433000978,
            "coin": "XLM",
            "side": "B",
            "px": "0.23005",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433000980,
            "coin": "XLM",
            "side": "A",
            "px": "0.2311",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432974355,
            "coin": "MON",
            "side": "B",
            "px": "0.023367",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432985540,
            "coin": "MON",
            "side": "A",
            "px": "0.023599",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001061,
            "coin": "BIO",
            "side": "B",
            "px": "0.050872",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432988497,
            "coin": "BIO",
            "side": "A",
            "px": "0.050959",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001062,
            "coin": "BIO",
            "side": "A",
            "px": "0.050973",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001063,
            "coin": "BIO",
            "side": "A",
            "px": "0.050991",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432906455,
            "coin": "BIO",
            "side": "A",
            "px": "0.051038",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842006,
            "coin": "SCR",
            "side": "B",
            "px": "0.08811",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842007,
            "coin": "SCR",
            "side": "B",
            "px": "0.08799",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842008,
            "coin": "SCR",
            "side": "B",
            "px": "0.08782",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842009,
            "coin": "SCR",
            "side": "B",
            "px": "0.0876",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842010,
            "coin": "SCR",
            "side": "B",
            "px": "0.08727",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842011,
            "coin": "SCR",
            "side": "B",
            "px": "0.08614",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842012,
            "coin": "SCR",
            "side": "A",
            "px": "0.08857",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842013,
            "coin": "SCR",
            "side": "A",
            "px": "0.08869",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842014,
            "coin": "SCR",
            "side": "A",
            "px": "0.08885",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842015,
            "coin": "SCR",
            "side": "A",
            "px": "0.08906",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842016,
            "coin": "SCR",
            "side": "A",
            "px": "0.08937",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001064,
            "coin": "SCR",
            "side": "B",
            "px": "0.08823",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001065,
            "coin": "SCR",
            "side": "A",
            "px": "0.08845",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842018,
            "coin": "GAS",
            "side": "B",
            "px": "2.0801",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842019,
            "coin": "GAS",
            "side": "B",
            "px": "2.0781",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432830343,
            "coin": "GAS",
            "side": "A",
            "px": "2.0834",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432830344,
            "coin": "GAS",
            "side": "A",
            "px": "2.0843",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842021,
            "coin": "GAS",
            "side": "A",
            "px": "2.0853",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842022,
            "coin": "GAS",
            "side": "A",
            "px": "2.0873",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432480790,
            "coin": "GAS",
            "side": "A",
            "px": "2.0913",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254431889787,
            "coin": "GAS",
            "side": "A",
            "px": "2.0983",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968953,
            "coin": "KAS",
            "side": "B",
            "px": "0.051445",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432904028,
            "coin": "KAS",
            "side": "A",
            "px": "0.05152",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945282,
            "coin": "KAS",
            "side": "A",
            "px": "0.051534",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945283,
            "coin": "KAS",
            "side": "A",
            "px": "0.051555",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842027,
            "coin": "KAS",
            "side": "A",
            "px": "0.051598",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432842028,
            "coin": "KAS",
            "side": "A",
            "px": "0.051665",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254431808666,
            "coin": "KAS",
            "side": "A",
            "px": "0.052123",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433000983,
            "coin": "ENS",
            "side": "A",
            "px": "10.549",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432995953,
            "coin": "FXS",
            "side": "B",
            "px": "0.77769",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432995954,
            "coin": "FXS",
            "side": "B",
            "px": "0.77724",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433000984,
            "coin": "FXS",
            "side": "B",
            "px": "0.77709",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433000985,
            "coin": "FXS",
            "side": "B",
            "px": "0.77631",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433002221,
            "coin": "FXS",
            "side": "B",
            "px": "0.77754",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432952949,
            "coin": "FXS",
            "side": "A",
            "px": "0.77842",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432915018,
            "coin": "MET",
            "side": "B",
            "px": "0.29813",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432909930,
            "coin": "MET",
            "side": "A",
            "px": "0.30084",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254431782470,
            "coin": "NOT",
            "side": "A",
            "px": "0.000519",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254431870803,
            "coin": "NOT",
            "side": "A",
            "px": "0.00052",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432958264,
            "coin": "FTT",
            "side": "B",
            "px": "0.55282",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432958265,
            "coin": "FTT",
            "side": "B",
            "px": "0.55137",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432958266,
            "coin": "FTT",
            "side": "A",
            "px": "0.5541",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432949649,
            "coin": "BSV",
            "side": "B",
            "px": "19.698",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432949650,
            "coin": "BSV",
            "side": "B",
            "px": "19.679",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001066,
            "coin": "VVV",
            "side": "B",
            "px": "0.9478",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433001067,
            "coin": "VVV",
            "side": "A",
            "px": "0.9491",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432995955,
            "coin": "GMX",
            "side": "B",
            "px": "8.1368",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968432,
            "coin": "GMX",
            "side": "B",
            "px": "8.1334",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432977304,
            "coin": "GMX",
            "side": "A",
            "px": "8.1456",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432977305,
            "coin": "GMX",
            "side": "A",
            "px": "8.1491",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945290,
            "coin": "GMX",
            "side": "A",
            "px": "8.1536",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432996035,
            "coin": "ZEREBRO",
            "side": "A",
            "px": "0.028273",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432364736,
            "coin": "ZEREBRO",
            "side": "B",
            "px": "0.027394",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433002231,
            "coin": "ZEREBRO",
            "side": "B",
            "px": "0.027731",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254433002232,
            "coin": "ZEREBRO",
            "side": "A",
            "px": "0.028313",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432983506,
            "coin": "GRIFFAIN",
            "side": "B",
            "px": "0.019736",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432961191,
            "coin": "GRIFFAIN",
            "side": "A",
            "px": "0.019847",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254432738102,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13679",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432957039,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07082",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432974788,
            "coin": "ARB",
            "side": "A",
            "px": "0.194",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc86242898f4b63883f05628d8cf2c326071797b8",
            "oid": 254432978430,
            "coin": "VINE",
            "side": "B",
            "px": "0.028053",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6e1a6f7bb47aca69aaa511210227d40926ff256",
            "oid": 254432909261,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13594",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6e1a6f7bb47aca69aaa511210227d40926ff256",
            "oid": 254432820251,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13583",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6e1a6f7bb47aca69aaa511210227d40926ff256",
            "oid": 254432820252,
            "coin": "DOGE",
            "side": "A",
            "px": "0.13641",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc029043cd00b80363130fa058818459a521842a1",
            "oid": 254433005619,
            "coin": "XRP",
            "side": "A",
            "px": "2.0275",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc2e137ea6b79a7992dd54d44ee8223f2c7e13e22",
            "oid": 254433001620,
            "coin": "ETC",
            "side": "B",
            "px": "12.895",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x5fb096e3fe7da58aebb47d73cafaf039a1800661",
            "oid": 254432988161,
            "coin": "VINE",
            "side": "B",
            "px": "0.028053",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf36c562be0598b8ce109998b6aa2d53ffcd8316b",
            "oid": 254431759027,
            "coin": "NOT",
            "side": "B",
            "px": "0.000516",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432945493,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07062",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x57dd78cd36e76e2011e8f6dc25cabbaba994494b",
            "oid": 254432998956,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13605",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x727956612a8700627451204a3ae26268bd1a1525",
            "oid": 254432973357,
            "coin": "FIL",
            "side": "B",
            "px": "1.4507",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432973323,
            "coin": "ETC",
            "side": "B",
            "px": "12.888",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432948352,
            "coin": "ETHFI",
            "side": "B",
            "px": "0.73994",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe377e2f11c745c294ad65686635217ec765059f2",
            "oid": 254432963549,
            "coin": "SUI",
            "side": "B",
            "px": "1.3494",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432997808,
            "coin": "OP",
            "side": "B",
            "px": "0.28968",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005769,
            "coin": "@142",
            "side": "A",
            "px": "85478.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7f3d2fcb1d78872135ea9585b354c49b63092110",
            "oid": 254432678261,
            "coin": "vntl:SPACEX",
            "side": "A",
            "px": "515.16",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432970845,
            "coin": "MOODENG",
            "side": "B",
            "px": "0.069416",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x4537e1a5b7bcadd39f6d6211cb757ce431798732",
            "oid": 254433002834,
            "coin": "XRP",
            "side": "A",
            "px": "2.028",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432973578,
            "coin": "WLD",
            "side": "B",
            "px": "0.57324",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432911605,
            "coin": "DOT",
            "side": "B",
            "px": "2.0226",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x92c7d53dc3bac14722b281e74a3feba8f7ba040e",
            "oid": 254432944262,
            "coin": "ARB",
            "side": "B",
            "px": "0.1931",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254433005959,
            "coin": "kFLOKI",
            "side": "B",
            "px": "0.042889",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432948623,
            "coin": "ARB",
            "side": "B",
            "px": "0.19315",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432863227,
            "coin": "DOT",
            "side": "A",
            "px": "2.031",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432825401,
            "coin": "ARB",
            "side": "B",
            "px": "0.19311",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254432969396,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13607",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432866822,
            "coin": "DOT",
            "side": "A",
            "px": "2.033",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432959812,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12021",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432977650,
            "coin": "@151",
            "side": "A",
            "px": "2823.3",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432988397,
            "coin": "IO",
            "side": "B",
            "px": "0.20058",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432988398,
            "coin": "IO",
            "side": "B",
            "px": "0.20049",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432971776,
            "coin": "ETC",
            "side": "B",
            "px": "12.895",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432961745,
            "coin": "ETC",
            "side": "B",
            "px": "12.892",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432936908,
            "coin": "SAGA",
            "side": "B",
            "px": "0.0706",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432936909,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07053",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432965471,
            "coin": "MOODENG",
            "side": "B",
            "px": "0.069545",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005523,
            "coin": "@151",
            "side": "B",
            "px": "2821.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432889768,
            "coin": "kFLOKI",
            "side": "A",
            "px": "0.042954",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432998857,
            "coin": "@151",
            "side": "A",
            "px": "2823.2",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432977439,
            "coin": "@151",
            "side": "A",
            "px": "2823.4",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432961018,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007988",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432957573,
            "coin": "PENDLE",
            "side": "A",
            "px": "2.4289",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb4321b142b2a03ce20fcab2007ff6990b9acba93",
            "oid": 254432072252,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007997",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432866702,
            "coin": "ARB",
            "side": "B",
            "px": "0.19302",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254432994656,
            "coin": "WLD",
            "side": "B",
            "px": "0.57325",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433003267,
            "coin": "ETHFI",
            "side": "B",
            "px": "0.73604",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xaf9f722a676230cc44045efe26fe9a85801ca4fa",
            "oid": 254432973078,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14617",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433009091,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.8262",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432943518,
            "coin": "ARB",
            "side": "B",
            "px": "0.1931",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7469cbaedb847ddac296b6f7c857e3127f239be8",
            "oid": 254433007095,
            "coin": "ETH",
            "side": "B",
            "px": "2821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
            "oid": 254432999964,
            "coin": "ARB",
            "side": "B",
            "px": "0.19311",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc1fce740d83a60de67d039aa927a678ff78c202f",
            "oid": 254432926201,
            "coin": "WLD",
            "side": "B",
            "px": "0.57309",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x399965e15d4e61ec3529cc98b7f7ebb93b733336",
            "oid": 254432992205,
            "coin": "@162",
            "side": "B",
            "px": "0.29761",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe3b6e3443c8f2080704e7421bad9340f13950acb",
            "oid": 254433000037,
            "coin": "ETH",
            "side": "B",
            "px": "2820.8",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xfc8c156428a8e48cb8d0356db16e59bec4c0ecea",
            "oid": 254432371452,
            "coin": "SAND",
            "side": "B",
            "px": "0.14102",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254432494610,
            "coin": "ARB",
            "side": "B",
            "px": "0.19309",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254431352615,
            "coin": "ARB",
            "side": "A",
            "px": "0.19415",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432890397,
            "coin": "WLD",
            "side": "B",
            "px": "0.57297",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432821932,
            "coin": "GRASS",
            "side": "A",
            "px": "0.30364",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254432944118,
            "coin": "ARB",
            "side": "B",
            "px": "0.19313",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xae7a0f9f663bb54bfe95378c7e3de7dd6e28d6bc",
            "oid": 254432973657,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.2978",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432997683,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.296",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432964990,
            "coin": "UNI",
            "side": "B",
            "px": "5.516",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432851594,
            "coin": "UNI",
            "side": "B",
            "px": "5.5191",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433003089,
            "coin": "ETHFI",
            "side": "B",
            "px": "0.73193",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1938cc3babe6cb8812c85a337829425e30de935",
            "oid": 254431269405,
            "coin": "UNI",
            "side": "B",
            "px": "5.513",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
            "oid": 254431331976,
            "coin": "WLD",
            "side": "B",
            "px": "0.57295",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433007752,
            "coin": "UNI",
            "side": "A",
            "px": "5.526",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433008927,
            "coin": "BERA",
            "side": "B",
            "px": "0.94342",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254432487892,
            "coin": "UNI",
            "side": "A",
            "px": "5.5318",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433002736,
            "coin": "kFLOKI",
            "side": "B",
            "px": "0.042902",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x162cc7c861ebd0c06b3d72319201150482518185",
            "oid": 254432999960,
            "coin": "ETH",
            "side": "B",
            "px": "2821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc029043cd00b80363130fa058818459a521842a1",
            "oid": 254433006373,
            "coin": "XRP",
            "side": "A",
            "px": "2.0275",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254432996194,
            "coin": "BTC",
            "side": "A",
            "px": "85428.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254432327810,
            "coin": "BABY",
            "side": "B",
            "px": "0.018339",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432913697,
            "coin": "PENDLE",
            "side": "A",
            "px": "2.4276",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005734,
            "coin": "XLM",
            "side": "B",
            "px": "0.23084",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x727956612a8700627451204a3ae26268bd1a1525",
            "oid": 254433002861,
            "coin": "ASTER",
            "side": "B",
            "px": "0.94081",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x727956612a8700627451204a3ae26268bd1a1525",
            "oid": 254433004794,
            "coin": "SUI",
            "side": "B",
            "px": "1.3507",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433003538,
            "coin": "kFLOKI",
            "side": "B",
            "px": "0.042883",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe3b6e3443c8f2080704e7421bad9340f13950acb",
            "oid": 254432957746,
            "coin": "ETH",
            "side": "B",
            "px": "2821.3",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432888619,
            "coin": "BNB",
            "side": "A",
            "px": "821.52",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007392,
            "coin": "XLM",
            "side": "A",
            "px": "0.23138",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254432764287,
            "coin": "kPEPE",
            "side": "B",
            "px": "0.004123",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254431817907,
            "coin": "WLD",
            "side": "B",
            "px": "0.56975",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432981294,
            "coin": "WLD",
            "side": "B",
            "px": "0.5728",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432849481,
            "coin": "WLD",
            "side": "A",
            "px": "0.57474",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432865747,
            "coin": "kNEIRO",
            "side": "A",
            "px": "0.12067",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432568871,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12018",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007549,
            "coin": "@151",
            "side": "B",
            "px": "2821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254432902167,
            "coin": "BTC",
            "side": "B",
            "px": "85386.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432972163,
            "coin": "SUI",
            "side": "A",
            "px": "1.3513",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007445,
            "coin": "@151",
            "side": "B",
            "px": "2821.9",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432944237,
            "coin": "SOL",
            "side": "B",
            "px": "126.54",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdacb0c5b92766b286db282ea42b3d7ebdc2423f9",
            "oid": 254432465743,
            "coin": "xyz:NVDA",
            "side": "A",
            "px": "173.52",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007391,
            "coin": "XLM",
            "side": "B",
            "px": "0.2307",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432945023,
            "coin": "PYTH",
            "side": "B",
            "px": "0.0674",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe8de5d2ae47fa0a00d1de7a36b2a2d75b63a80f7",
            "oid": 254432984327,
            "coin": "POL",
            "side": "B",
            "px": "0.12281",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254433000372,
            "coin": "STX",
            "side": "B",
            "px": "0.28251",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254432927306,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14622",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x56667d7d8e0e5c529d96f6d92ff2ce2414c5910d",
            "oid": 254432978234,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053304",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254432972349,
            "coin": "SOL",
            "side": "B",
            "px": "126.52",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432951114,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14619",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe71092eeced04572afba4515895758277d0f8b88",
            "oid": 254432360188,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.095023",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xfad17af4900f3d3afa663103f58b41607ee45d70",
            "oid": 254432963519,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14626",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x5fffee2555a15899ad656c1a80f1b35cd0b2c0c1",
            "oid": 254433005419,
            "coin": "@162",
            "side": "B",
            "px": "0.29742",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433008064,
            "coin": "ETH",
            "side": "B",
            "px": "2820.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432939790,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14623",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432929906,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094391",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432929907,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094714",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432929908,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094367",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432929909,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094738",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432929910,
            "coin": "POPCAT",
            "side": "B",
            "px": "0.094344",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432813744,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094751",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf9109ada2f73c62e9889b45453065f0d99260a2d",
            "oid": 254432813746,
            "coin": "POPCAT",
            "side": "A",
            "px": "0.094775",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x3ff54fd26855db3758d0b5ae7aed47440c47f705",
            "oid": 254432962045,
            "coin": "BTC",
            "side": "A",
            "px": "85460.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009477,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14625",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254432927889,
            "coin": "WLFI",
            "side": "A",
            "px": "0.14639",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd158ab29974537b266d37ef88ae71ff4334bab69",
            "oid": 254432986032,
            "coin": "POL",
            "side": "B",
            "px": "0.12281",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432943747,
            "coin": "PYTH",
            "side": "B",
            "px": "0.067413",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xbb475febf78b528f4f0d4ae1a12541406565199c",
            "oid": 254433003466,
            "coin": "BTC",
            "side": "B",
            "px": "85397.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xbb475febf78b528f4f0d4ae1a12541406565199c",
            "oid": 254432961994,
            "coin": "BTC",
            "side": "A",
            "px": "85455.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433005821,
            "coin": "ETH",
            "side": "A",
            "px": "2822.3",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432943835,
            "coin": "PYTH",
            "side": "B",
            "px": "0.06752",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x621c5551678189b9a6c94d929924c225ff1d63ab",
            "oid": 254432481334,
            "coin": "WLD",
            "side": "B",
            "px": "0.57308",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254432897256,
            "coin": "ARB",
            "side": "A",
            "px": "0.19324",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254432927305,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14618",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433004168,
            "coin": "EIGEN",
            "side": "A",
            "px": "0.5054",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009475,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14627",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254432999976,
            "coin": "BTC",
            "side": "A",
            "px": "85429.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254432927307,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14614",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254433008154,
            "coin": "BTC",
            "side": "B",
            "px": "85404.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x14a4914df70e599bdc5baf74b7b51773499dbe21",
            "oid": 254433004807,
            "coin": "AVAX",
            "side": "B",
            "px": "12.911",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xaf9f722a676230cc44045efe26fe9a85801ca4fa",
            "oid": 254432968757,
            "coin": "ETH",
            "side": "A",
            "px": "2824.9",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xfdc7c22506e166206858df11d22acaa5bd195ba6",
            "oid": 254432965880,
            "coin": "ETH",
            "side": "B",
            "px": "2821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254432964271,
            "coin": "PNUT",
            "side": "B",
            "px": "0.08074",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf4e62d9d431a07b9b1efb5dd548e3ad9aadd07ba",
            "oid": 254432933784,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053294",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1c1c270b573d55b68b3d14722b5d5d401511bed0",
            "oid": 254433009476,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14626",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254432614332,
            "coin": "ARB",
            "side": "A",
            "px": "0.1944",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432484130,
            "coin": "ARB",
            "side": "B",
            "px": "0.19305",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433009453,
            "coin": "@162",
            "side": "B",
            "px": "0.29698",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432554453,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12011",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e5c7a2fed36727acf6306e3e5ae12bce86856c7",
            "oid": 254432965524,
            "coin": "TON",
            "side": "B",
            "px": "1.4874",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433001564,
            "coin": "FARTCOIN",
            "side": "A",
            "px": "0.29804",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x38f96bba487a628d2a84afc3fe5c7d77eb07cb04",
            "oid": 254432990380,
            "coin": "ENA",
            "side": "B",
            "px": "0.24193",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x38f96bba487a628d2a84afc3fe5c7d77eb07cb04",
            "oid": 254432991760,
            "coin": "ENA",
            "side": "A",
            "px": "0.24229",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x38f96bba487a628d2a84afc3fe5c7d77eb07cb04",
            "oid": 254432990403,
            "coin": "ENA",
            "side": "A",
            "px": "0.24244",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254433007566,
            "coin": "XRP",
            "side": "B",
            "px": "2.0267",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd02928c445d780ea954fe0b6f0be0f6cb9727678",
            "oid": 254432548460,
            "coin": "ARB",
            "side": "B",
            "px": "0.19303",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
            "oid": 254432426228,
            "coin": "SAGA",
            "side": "A",
            "px": "0.07228",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
            "oid": 254432426233,
            "coin": "SAGA",
            "side": "B",
            "px": "0.06976",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7ca165f354e3260e2f8d5a7508cc9dd2fa009235",
            "oid": 254433006569,
            "coin": "@142",
            "side": "A",
            "px": "85491.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa33a4a057334c7811ad5f45f3c4f0dfa3d081ff8",
            "oid": 254433000071,
            "coin": "LINK",
            "side": "B",
            "px": "12.109",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc2e137ea6b79a7992dd54d44ee8223f2c7e13e22",
            "oid": 254432927288,
            "coin": "SAND",
            "side": "B",
            "px": "0.1412",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432989840,
            "coin": "SOL",
            "side": "A",
            "px": "126.86",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x170a03a53fe05bf4ce6753cb0907d05e827bc3c6",
            "oid": 254433008061,
            "coin": "XRP",
            "side": "B",
            "px": "2.0274",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc1fd39727cd10fe97ccac228888f025309560a93",
            "oid": 254432986082,
            "coin": "POL",
            "side": "B",
            "px": "0.12281",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254432944161,
            "coin": "SOL",
            "side": "A",
            "px": "126.6",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254432975474,
            "coin": "SOL",
            "side": "A",
            "px": "126.58",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb4321b142b2a03ce20fcab2007ff6990b9acba93",
            "oid": 254430852940,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008742",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0a06ec6754b628be489b2c40bba20c8580392a7b",
            "oid": 254432929930,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14623",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf36c562be0598b8ce109998b6aa2d53ffcd8316b",
            "oid": 254432819144,
            "coin": "kSHIB",
            "side": "A",
            "px": "0.008019",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf36c562be0598b8ce109998b6aa2d53ffcd8316b",
            "oid": 254432816328,
            "coin": "kSHIB",
            "side": "A",
            "px": "0.00801",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf36c562be0598b8ce109998b6aa2d53ffcd8316b",
            "oid": 254432494567,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007931",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432846115,
            "coin": "MEW",
            "side": "B",
            "px": "0.00101",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432917456,
            "coin": "MEW",
            "side": "A",
            "px": "0.001015",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf36c562be0598b8ce109998b6aa2d53ffcd8316b",
            "oid": 254432849562,
            "coin": "DOGE",
            "side": "B",
            "px": "0.13502",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x95995f302ad58138d791ce49f9f3b1274e80c60a",
            "oid": 254432943834,
            "coin": "ETH",
            "side": "A",
            "px": "2823.4",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432753510,
            "coin": "SOL",
            "side": "A",
            "px": "126.92",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254432996306,
            "coin": "FIL",
            "side": "A",
            "px": "1.4521",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254433006483,
            "coin": "BTC",
            "side": "B",
            "px": "85414.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xca163b98cbc2ad5b08b01ae7bcb4bef68b9161ad",
            "oid": 254432988290,
            "coin": "BERA",
            "side": "B",
            "px": "0.94391",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433007649,
            "coin": "SOL",
            "side": "A",
            "px": "126.65",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968949,
            "coin": "ZEN",
            "side": "B",
            "px": "9.5256",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945220,
            "coin": "STX",
            "side": "B",
            "px": "0.28276",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432945221,
            "coin": "STX",
            "side": "B",
            "px": "0.28266",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432942188,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14626",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432929040,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14619",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825802,
            "coin": "GRASS",
            "side": "B",
            "px": "0.30322",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432825803,
            "coin": "GRASS",
            "side": "B",
            "px": "0.30311",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432961803,
            "coin": "PENDLE",
            "side": "B",
            "px": "2.4247",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432925946,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053281",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf36c562be0598b8ce109998b6aa2d53ffcd8316b",
            "oid": 254432976924,
            "coin": "XPL",
            "side": "A",
            "px": "0.17763",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433005823,
            "coin": "@142",
            "side": "B",
            "px": "85445.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd63c228ea09e4d999f5fb3ed7d3529756e3a724f",
            "oid": 254432138376,
            "coin": "BRETT",
            "side": "A",
            "px": "0.014719",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254433001560,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.29738",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432881102,
            "coin": "SOL",
            "side": "B",
            "px": "126.33",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x4129c62faf652fea61375dcd9ca8ce24b2bb8b95",
            "oid": 254433002662,
            "coin": "ETH",
            "side": "B",
            "px": "2821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xf27ebb91ea420f73b59b205ac7e0b77a90ec8f3c",
            "oid": 254432846371,
            "coin": "MEW",
            "side": "B",
            "px": "0.00101",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb4321b142b2a03ce20fcab2007ff6990b9acba93",
            "oid": 254430875354,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008742",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x39475d17bcd20adc540e647dae6781b153fbf3b1",
            "oid": 254432972309,
            "coin": "ETH",
            "side": "A",
            "px": "2824.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x59d00a5a603ebe61cdce3f5ff54e9cfabe64ad7c",
            "oid": 254432998431,
            "coin": "CRV",
            "side": "A",
            "px": "0.38307",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x4537e1a5b7bcadd39f6d6211cb757ce431798732",
            "oid": 254433006537,
            "coin": "XRP",
            "side": "B",
            "px": "2.0268",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xdcac85ecae7148886029c20e661d848a4de99ce2",
            "oid": 254432995997,
            "coin": "ADA",
            "side": "A",
            "px": "0.37927",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0fd468a73084daa6ea77a9261e40fdec3e67e0c7",
            "oid": 254433008261,
            "coin": "BTC",
            "side": "B",
            "px": "85411.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x793660578d2628b0ec7e0e2b5c75a4847ae80706",
            "oid": 254432956485,
            "coin": "ETH",
            "side": "B",
            "px": "2821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432956598,
            "coin": "SUI",
            "side": "B",
            "px": "1.3499",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x0fd468a73084daa6ea77a9261e40fdec3e67e0c7",
            "oid": 254433003506,
            "coin": "BTC",
            "side": "A",
            "px": "85428.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432867819,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007997",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432820237,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007996",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432812269,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12021",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432812270,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12014",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7ca165f354e3260e2f8d5a7508cc9dd2fa009235",
            "oid": 254433006545,
            "coin": "BTC",
            "side": "B",
            "px": "85406.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7ca165f354e3260e2f8d5a7508cc9dd2fa009235",
            "oid": 254432963792,
            "coin": "BTC",
            "side": "B",
            "px": "85397.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x57dd78cd36e76e2011e8f6dc25cabbaba994494b",
            "oid": 254433007612,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.2974",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254432980194,
            "coin": "SUI",
            "side": "A",
            "px": "1.3512",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254433006546,
            "coin": "BTC",
            "side": "B",
            "px": "85391.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432826198,
            "coin": "ARB",
            "side": "B",
            "px": "0.19314",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432826199,
            "coin": "ARB",
            "side": "B",
            "px": "0.19306",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432960572,
            "coin": "INJ",
            "side": "B",
            "px": "5.1492",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x010461c14e146ac35fe42271bdc1134ee31c703a",
            "oid": 254432968463,
            "coin": "POL",
            "side": "B",
            "px": "0.1228",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x09bc1cf4d9f0b59e1425a8fde4d4b1f7d3c9410d",
            "oid": 254433008237,
            "coin": "BTC",
            "side": "B",
            "px": "85398.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1edfeccf03f35d2268f5dc4013508a6eb52c702",
            "oid": 254431396005,
            "coin": "BABY",
            "side": "A",
            "px": "0.018484",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xb1edfeccf03f35d2268f5dc4013508a6eb52c702",
            "oid": 254431396007,
            "coin": "BABY",
            "side": "B",
            "px": "0.018256",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa801bb18551016ab0721555f59a27d2b31a8ffbb",
            "oid": 254432945251,
            "coin": "ETH",
            "side": "B",
            "px": "2821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
            "oid": 254431695096,
            "coin": "DOOD",
            "side": "A",
            "px": "0.00372",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
            "oid": 254431695097,
            "coin": "DOOD",
            "side": "B",
            "px": "0.003585",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x298e21ddd7320c393afdbf0c59dc4e4034a4b68d",
            "oid": 254432876215,
            "coin": "UNI",
            "side": "B",
            "px": "5.517",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x92c7d53dc3bac14722b281e74a3feba8f7ba040e",
            "oid": 254432945597,
            "coin": "BCH",
            "side": "B",
            "px": "521.11",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x89a4e7b411cebfcf66347935bb84f406eab4be29",
            "oid": 254432871342,
            "coin": "UNI",
            "side": "B",
            "px": "5.517",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe86b057f5eb764c9738d6b0d38170befd0723664",
            "oid": 254432981043,
            "coin": "TON",
            "side": "A",
            "px": "1.4916",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x621c5551678189b9a6c94d929924c225ff1d63ab",
            "oid": 254432982495,
            "coin": "ETH",
            "side": "A",
            "px": "2824.4",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e7cdb495b4864448a89bdd959e0fc070e47c621",
            "oid": 254433008810,
            "coin": "MON",
            "side": "A",
            "px": "0.023601",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x76bf72092bbfe1c935b6d834f70dd9e75b057fd6",
            "oid": 254432976995,
            "coin": "@151",
            "side": "A",
            "px": "2826.8",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xc06dd151c27dbe535ae95797e8aab1bb8b8cc7c0",
            "oid": 254432974864,
            "coin": "ETH",
            "side": "A",
            "px": "2825.3",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x727956612a8700627451204a3ae26268bd1a1525",
            "oid": 254432967666,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.29769",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254433001329,
            "coin": "@162",
            "side": "A",
            "px": "0.29878",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xaf9f722a676230cc44045efe26fe9a85801ca4fa",
            "oid": 254432949211,
            "coin": "ETH",
            "side": "A",
            "px": "2823.8",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433003430,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.29725",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432975823,
            "coin": "2Z",
            "side": "B",
            "px": "0.10241",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xecb63caa47c7c4e77f60f1ce858cf28dc2b82b00",
            "oid": 254432911697,
            "coin": "2Z",
            "side": "B",
            "px": "0.10234",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432956734,
            "coin": "ETH",
            "side": "B",
            "px": "2818.9",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa775d1bd91ee2cc4cae6113e5fd9c0c17b355277",
            "oid": 254432944093,
            "coin": "BTC",
            "side": "B",
            "px": "85422.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x399965e15d4e61ec3529cc98b7f7ebb93b733336",
            "oid": 254433005590,
            "coin": "@162",
            "side": "A",
            "px": "0.298",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433004699,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82849",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x399965e15d4e61ec3529cc98b7f7ebb93b733336",
            "oid": 254433003825,
            "coin": "@142",
            "side": "B",
            "px": "85457.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x97f20f91c18093c7f738a00ffabc03eea04ec668",
            "oid": 254432877539,
            "coin": "BTC",
            "side": "A",
            "px": "85492.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254433004146,
            "coin": "OP",
            "side": "B",
            "px": "0.28969",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xd6ef5649200101eb44ce8949f4cfb167b0c5a07a",
            "oid": 254432842994,
            "coin": "WIF",
            "side": "B",
            "px": "0.33822",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x023a3d058020fb76cca98f01b3c48c8938a22355",
            "oid": 254432841414,
            "coin": "@151",
            "side": "A",
            "px": "2827.4",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x1dd4af3383fce2ee4a40d9fb6766cbfcce2ddbce",
            "oid": 254430606947,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053246",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xaf9f722a676230cc44045efe26fe9a85801ca4fa",
            "oid": 254433003749,
            "coin": "FARTCOIN",
            "side": "B",
            "px": "0.29719",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xce75a235d1a055e4b3eb88545f7215c6f963055a",
            "oid": 254432949025,
            "coin": "BTC",
            "side": "B",
            "px": "85422.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x348e5365acfa48a26ada7da840ca611e29c950ef",
            "oid": 254433006509,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7717a7a245d9f950e586822b8c9b46863ed7bd7e",
            "oid": 254432947667,
            "coin": "ETH",
            "side": "B",
            "px": "2820.6",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x162cc7c861ebd0c06b3d72319201150482518185",
            "oid": 254433006503,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31dea2516beee92135b96f464eeec3cf292a13f2",
            "oid": 254432981441,
            "coin": "@142",
            "side": "B",
            "px": "85443.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa289ee1e56c0d5d041db762e6123e78af0f7d9ad",
            "oid": 254433006511,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254433009420,
            "coin": "BTC",
            "side": "B",
            "px": "85412.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x9e74a6a1df3c2545ec4e8e54a1502967c7ad15e1",
            "oid": 254433007728,
            "coin": "BTC",
            "side": "B",
            "px": "85401.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xee162a5a60829bd346f0c1ac3514b21fe5f4b290",
            "oid": 254433006507,
            "coin": "BTC",
            "side": "B",
            "px": "76880.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xa33a4a057334c7811ad5f45f3c4f0dfa3d081ff8",
            "oid": 254432996217,
            "coin": "ASTER",
            "side": "B",
            "px": "0.94051",
            "raw_book_diff": "remove"
        },
        {
            "user": "0xe056eebc2c7acfc782d47ebe18ad71735480f72a",
            "oid": 254432945082,
            "coin": "ETH",
            "side": "B",
            "px": "2821.0",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x33ad48ed92e3e91a98278f0f7456eae6eafa516e",
            "oid": 254433006020,
            "coin": "WLFI",
            "side": "A",
            "px": "0.14642",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x223537ac9a856c31f4043e86ced86bb29f06653e",
            "oid": 254432840840,
            "coin": "BNB",
            "side": "A",
            "px": "821.5",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432972007,
            "coin": "KAS",
            "side": "B",
            "px": "0.051471",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433002538,
            "coin": "COMP",
            "side": "B",
            "px": "34.822",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433002539,
            "coin": "COMP",
            "side": "B",
            "px": "34.811",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432996793,
            "coin": "COMP",
            "side": "B",
            "px": "34.802",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432966138,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5407",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432966139,
            "coin": "ORDI",
            "side": "B",
            "px": "3.5395",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824288,
            "coin": "kBONK",
            "side": "B",
            "px": "0.008743",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433003925,
            "coin": "APT",
            "side": "B",
            "px": "1.8571",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432950976,
            "coin": "ME",
            "side": "B",
            "px": "0.30181",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432950977,
            "coin": "ME",
            "side": "B",
            "px": "0.30171",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432988856,
            "coin": "BIO",
            "side": "B",
            "px": "0.050942",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432988857,
            "coin": "BIO",
            "side": "B",
            "px": "0.050928",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432988858,
            "coin": "BIO",
            "side": "B",
            "px": "0.050912",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432759525,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432650317,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432707458,
            "coin": "MEME",
            "side": "B",
            "px": "0.001179",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432829625,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001875",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432443207,
            "coin": "TURBO",
            "side": "B",
            "px": "0.001874",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432996747,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041491",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824431,
            "coin": "BRETT",
            "side": "B",
            "px": "0.014528",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824432,
            "coin": "BRETT",
            "side": "B",
            "px": "0.01452",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433002493,
            "coin": "TIA",
            "side": "B",
            "px": "0.56877",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432988832,
            "coin": "IO",
            "side": "B",
            "px": "0.20058",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432988833,
            "coin": "IO",
            "side": "B",
            "px": "0.20049",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432971941,
            "coin": "ETC",
            "side": "B",
            "px": "12.895",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432969212,
            "coin": "ETC",
            "side": "B",
            "px": "12.892",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433006284,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07066",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432940949,
            "coin": "SAGA",
            "side": "B",
            "px": "0.0706",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432940950,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07053",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432999776,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07066",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433002602,
            "coin": "SAGA",
            "side": "B",
            "px": "0.07066",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824714,
            "coin": "ARB",
            "side": "B",
            "px": "0.19314",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824715,
            "coin": "ARB",
            "side": "B",
            "px": "0.19306",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432859962,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007997",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824281,
            "coin": "kSHIB",
            "side": "B",
            "px": "0.007996",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432814406,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12021",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432814407,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12014",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432814408,
            "coin": "kNEIRO",
            "side": "B",
            "px": "0.12005",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432960642,
            "coin": "INJ",
            "side": "B",
            "px": "5.1492",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432966200,
            "coin": "POL",
            "side": "B",
            "px": "0.1228",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432971969,
            "coin": "ZEN",
            "side": "B",
            "px": "9.5256",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432943097,
            "coin": "STX",
            "side": "B",
            "px": "0.28276",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432943098,
            "coin": "STX",
            "side": "B",
            "px": "0.28266",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432860077,
            "coin": "SAND",
            "side": "B",
            "px": "0.1412",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432933685,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14626",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432927650,
            "coin": "WLFI",
            "side": "B",
            "px": "0.14619",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824416,
            "coin": "GRASS",
            "side": "B",
            "px": "0.30322",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432824417,
            "coin": "GRASS",
            "side": "B",
            "px": "0.30311",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432996748,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041478",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432966185,
            "coin": "AIXBT",
            "side": "B",
            "px": "0.041461",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432925760,
            "coin": "PENDLE",
            "side": "B",
            "px": "2.4247",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432925667,
            "coin": "PROMPT",
            "side": "B",
            "px": "0.053281",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432964040,
            "coin": "MOODENG",
            "side": "B",
            "px": "0.069545",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432957433,
            "coin": "KAITO",
            "side": "B",
            "px": "0.62158",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432921043,
            "coin": "OP",
            "side": "B",
            "px": "0.28964",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432921044,
            "coin": "OP",
            "side": "B",
            "px": "0.28952",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433006232,
            "coin": "VIRTUAL",
            "side": "B",
            "px": "0.82901",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432999788,
            "coin": "EIGEN",
            "side": "B",
            "px": "0.5047",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432994469,
            "coin": "EIGEN",
            "side": "B",
            "px": "0.5046",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432964020,
            "coin": "XPL",
            "side": "B",
            "px": "0.17754",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432964021,
            "coin": "XPL",
            "side": "B",
            "px": "0.17745",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254433006219,
            "coin": "LDO",
            "side": "B",
            "px": "0.57624",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x31ca8395cf837de08b24da3f660e77761dfb974b",
            "oid": 254432864119,
            "coin": "MNT",
            "side": "B",
            "px": "0.97174",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x97f20f91c18093c7f738a00ffabc03eea04ec668",
            "oid": 254433009908,
            "coin": "RESOLV",
            "side": "A",
            "px": "0.076815",
            "raw_book_diff": {
                "new": {
                    "sz": "509.0"
                }
            }
        },
        {
            "user": "0xf5d4d37b5563db58deb809396620a3811c1516ca",
            "oid": 254432452289,
            "coin": "MON",
            "side": "A",
            "px": "0.023585",
            "raw_book_diff": {
                "update": {
                    "origSz": "489.0",
                    "newSz": "0.0"
                }
            }
        },
        {
            "user": "0x479f209db7b652e37a7996ed701963b2076ef221",
            "oid": 254432858514,
            "coin": "MON",
            "side": "A",
            "px": "0.023588",
            "raw_book_diff": {
                "update": {
                    "origSz": "1632.0",
                    "newSz": "0.0"
                }
            }
        },
        {
            "user": "0xf5d4d37b5563db58deb809396620a3811c1516ca",
            "oid": 254432452289,
            "coin": "MON",
            "side": "A",
            "px": "0.023585",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x479f209db7b652e37a7996ed701963b2076ef221",
            "oid": 254432858514,
            "coin": "MON",
            "side": "A",
            "px": "0.023588",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x7083a8e36b44865c8d40379502bc081259a0ba66",
            "oid": 254433009909,
            "coin": "MON",
            "side": "B",
            "px": "0.023588",
            "raw_book_diff": {
                "new": {
                    "sz": "2060857.0"
                }
            }
        },
        {
            "user": "0x97f20f91c18093c7f738a00ffabc03eea04ec668",
            "oid": 254433009911,
            "coin": "WLD",
            "side": "A",
            "px": "0.5743",
            "raw_book_diff": {
                "new": {
                    "sz": "373.2"
                }
            }
        },
        {
            "user": "0xd52e4cb5e4dcb9e1aaffc63aaccc77d9052f62c0",
            "oid": 254433009914,
            "coin": "BTC",
            "side": "B",
            "px": "85420.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.1"
                }
            }
        },
        {
            "user": "0x5ddd3dcb63352eb0750df792b666fcb0d7e9b4ad",
            "oid": 254432984101,
            "coin": "@241",
            "side": "A",
            "px": "0.001216",
            "raw_book_diff": "remove"
        },
        {
            "user": "0x5ddd3dcb63352eb0750df792b666fcb0d7e9b4ad",
            "oid": 254433009915,
            "coin": "@241",
            "side": "A",
            "px": "0.001224",
            "raw_book_diff": {
                "new": {
                    "sz": "23487.0"
                }
            }
        },
        {
            "user": "0x774200d82eb522999459d87fd543068d254242da",
            "oid": 254433009916,
            "coin": "xyz:XYZ100",
            "side": "A",
            "px": "25180.0",
            "raw_book_diff": {
                "new": {
                    "sz": "0.0004"
                }
            }
        },
        {
            "user": "0xe0afe5de6831901db80d7ca9d17475f78b1a870c",
            "oid": 254433009917,
            "coin": "xyz:XYZ100",
            "side": "A",
            "px": "25226.0",
            "raw_book_diff": {
                "new": {
                    "sz": "3.7163"
                }
            }
        },
        {
            "user": "0xe687b524903fea10a648326d90130ffabe0215f9",
            "oid": 254432946258,
            "coin": "kPEPE",
            "side": "A",
            "px": "0.004124",
            "raw_book_diff": {
                "update": {
                    "origSz": "808195.0",
                    "newSz": "803989.0"
                }
            }
        }
    ]
}
`
