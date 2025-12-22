package datahandler

import (
	"time"

	"github.com/apache/arrow/go/v18/arrow"
	"github.com/apache/arrow/go/v18/arrow/array"
	eth "github.com/ethereum/go-ethereum/common"
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

type User struct {
	AccountValue      float64
	TotalNtlPos       float64
	CrossAccountValue float64
	Pnl               float64
	UnrealizedPnl     float64
}

func InitUsers(users []string) {
	for _, user := range users {
		_, _ = GetUser(eth.HexToAddress(user))
	}
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
