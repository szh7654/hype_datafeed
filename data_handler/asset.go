package datahandler

import (
	"context"
	dr "indexer/data_retriver"
	util "indexer/util"
	"sort"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	AssetNames    []string          = make([]string, 0)
	Assets        []*Asset          = make([]*Asset, 0)
	NameToAssetId map[string]uint16 = make(map[string]uint16)
)

var (
	assetCollection = util.MongoClient.Database("quant").Collection("asset")
	newAssetChan    = make(chan AssetData, 10)
)

type AssetData struct {
	Name    string `json:"name"`
	AssetId uint16 `json:"asset_id"`
}

func init() {
	go func() {
		ctx := context.TODO()
		for {
			asset := <-newAssetChan
			log.Info().Any("asset", asset).Msg("insert asset")
			_, err := util.MongoClient.Database("quant").Collection("asset").InsertOne(ctx, AssetData{
				Name:    asset.Name,
				AssetId: asset.AssetId,
			})
			if err != nil {
				log.Error().Err(err).Any("asset", asset).Msg("insert asset failed")
			}
		}
	}()

	ctx := context.TODO()

	cursor, err := assetCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var results []AssetData
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].AssetId < results[j].AssetId
	})
	AssetNames = make([]string, len(results))
	Assets = make([]*Asset, len(results))
	log.Info().Msgf("fetch %d assets from mongo", len(results))
	for i, result := range results {
		AssetNames[i] = result.Name
		Assets[i] = &Asset{
			SzDecimals:  uint8(result.AssetId),
			MaxLeverage: uint16(result.AssetId),
			OraclePx:    0,
			MarkPx:      0,
			MidPx:       0,
			DayNtlVlm:   0,
		}
		NameToAssetId[result.Name] = result.AssetId
	}

	metaAndAssetCtxs, err := dr.FetchMetaAndAssetCtxs()
	if err != nil {
		panic(err)
	}
	applyMetaAndAssetCtxs(metaAndAssetCtxs)
}

func applyMetaAndAssetCtxs(metaAndAssetCtxs *dr.MetaAndAssetCtxs) {
	for i, assetInfo := range metaAndAssetCtxs.Meta.Universe {
		assetCtx := metaAndAssetCtxs.Ctxs[i]
		asset, _ := GetAsset(assetInfo.Name)
		asset.OraclePx = assetCtx.OraclePx
		asset.MarkPx = assetCtx.MarkPx
		asset.MidPx = assetCtx.MidPx
		asset.DayNtlVlm = assetCtx.DayNtlVlm
		asset.OpenInterest = assetCtx.OpenInterest
		asset.Funding = assetCtx.Funding
		asset.SzDecimals = uint8(assetInfo.SzDecimals)
		asset.MaxLeverage = uint16(assetInfo.MaxLeverage)
	}
}

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

func GetAsset(name string) (*Asset, uint16) {
	if assetId, ok := NameToAssetId[name]; !ok {
		NameToAssetId[name] = uint16(len(Assets))
		AssetNames = append(AssetNames, name)
		asset := &Asset{}
		Assets = append(Assets, asset)
		addBook()
		newAssetChan <- AssetData{
			Name:    name,
			AssetId: NameToAssetId[name],
		}
		return asset, NameToAssetId[name]
	} else {
		return Assets[assetId], assetId
	}
}
