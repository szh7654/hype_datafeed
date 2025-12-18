package datahandler

import (
	dr "indexer/data_retriver"
	"indexer/util"
)

var (
	AssetNames    []string          = make([]string, 0)
	Assets        []*Asset          = make([]*Asset, 0)
	NameToAssetId map[string]uint16 = make(map[string]uint16)
)

type Asset struct {
	OraclePx     float64
	MarkPx       float64
	MidPx        float64
	DayNtlVlm    float64
	OpenInterest float64
	Funding      float64
	SzDecimals   uint8
	MaxLeverage  uint16
}

func AddAsset(asset Asset, name string) {
	AssetNames = append(AssetNames, name)
	Assets = append(Assets, &asset)
	NameToAssetId[name] = uint16(len(Assets))
}

func GetAsset(name string) *Asset {
	if assetId, ok := NameToAssetId[name]; !ok {
		return nil
	} else {
		return Assets[assetId]
	}
}

func InitAssets(assetData dr.MetaAndAssetCtxs) {
	for i, assetInfo := range assetData.Meta.Universe {
		assetCtx := assetData.Ctxs[i]
		AddAsset(Asset{
			SzDecimals:   uint8(assetInfo.SzDecimals),
			MaxLeverage:  uint16(assetInfo.MaxLeverage),
			OraclePx:     util.ToF64(assetCtx.OraclePx),
			MarkPx:       util.ToF64(assetCtx.MarkPx),
			MidPx:        util.ToF64(assetCtx.MidPx),
			DayNtlVlm:    util.ToF64(assetCtx.DayNtlVlm),
			OpenInterest: util.ToF64(assetCtx.OpenInterest),
			Funding:      util.ToF64(assetCtx.Funding),
		}, assetInfo.Name)
	}
}
