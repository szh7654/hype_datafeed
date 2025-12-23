package datahandler

import (
	"context"
	"sort"
	"time"

	util "indexer/util"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/apache/arrow/go/v18/arrow/array"
	eth "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	UserAddrs       []eth.Address          = make([]eth.Address, 0)
	Users           []*User                = make([]*User, 0)
	AddressToUserId map[eth.Address]uint32 = make(map[eth.Address]uint32)

	userBuilder = array.NewRecordBuilder(pool, arrow.NewSchema(
		[]arrow.Field{
			{Name: "time", Type: arrow.FixedWidthTypes.Timestamp_s},
			{Name: "user_id", Type: arrow.PrimitiveTypes.Uint32},
			{Name: "account_value", Type: arrow.PrimitiveTypes.Float64},
			{Name: "total_ntl_pos", Type: arrow.PrimitiveTypes.Float64},
			{Name: "cross_account_value", Type: arrow.PrimitiveTypes.Float64},
		},
		nil,
	))
	activeUsers map[uint32]eth.Address = make(map[uint32]eth.Address)
)

var (
	userCollection = util.MongoClient.Database("quant").Collection("user")
	newUserChan    = make(chan UserData, 1000)
)

type UserData struct {
	Name   string `json:"name"`
	UserId uint32 `json:"user_id"`
}

func init() {
	go func() {
		ctx := context.TODO()
		for {
			user := <-newUserChan
			log.Info().Any("user", user).Msg("insert user")
			_, err := util.MongoClient.Database("quant").Collection("user").InsertOne(ctx, UserData{
				Name:   user.Name,
				UserId: user.UserId,
			})
			if err != nil {
				log.Error().Err(err).Any("user", user).Msg("insert user failed")
			}
		}
	}()

	ctx := context.TODO()

	cursor, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var results []UserData
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].UserId < results[j].UserId
	})
	UserAddrs = make([]eth.Address, len(results))
	Users = make([]*User, len(results))
	for i, result := range results {
		UserAddrs[i] = eth.HexToAddress(result.Name)
		Users[i] = &User{}
		AddressToUserId[eth.HexToAddress(result.Name)] = result.UserId
	}
	users := util.ExtractUser()
	for _, user := range users {
		_, _ = GetUser(eth.HexToAddress(user))
	}

}

type User struct {
	AccountValue      float64
	TotalNtlPos       float64
	CrossAccountValue float64
	Pnl               float64
	UnrealizedPnl     float64
}

func setActive(userId uint32, address eth.Address) {
	activeUsers[userId] = address
}

func GetUser(address eth.Address) (*User, uint32) {
	if userId, ok := AddressToUserId[address]; !ok {
		userId = uint32(len(Users))
		user := &User{}
		UserAddrs = append(UserAddrs, address)
		Users = append(Users, user)
		AddressToUserId[address] = userId
		newUserChan <- UserData{
			Name:   address.Hex(),
			UserId: userId,
		}
		return user, userId
	} else {
		return Users[userId], userId
	}
}

func UserRecord() arrow.Record {
	userBuilder.Release()
	t := time.Now().Unix()
	for userId, user := range Users {
		userBuilder.Field(0).(*array.TimestampBuilder).Append(arrow.Timestamp(t))
		userBuilder.Field(1).(*array.Uint32Builder).Append(uint32(userId))
		userBuilder.Field(2).(*array.Float64Builder).Append(user.AccountValue)
		userBuilder.Field(3).(*array.Float64Builder).Append(user.TotalNtlPos)
		userBuilder.Field(4).(*array.Float64Builder).Append(user.CrossAccountValue)
	}
	return userBuilder.NewRecord()
}
