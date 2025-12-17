package dataretriver

import eth "github.com/ethereum/go-ethereum/common"

var UserStateChan = make(chan UserStateWithId, 1000)
var BlockFillChan = make(chan *BlockFill, 1000)
var BlockOrderStatusChan = make(chan *BlockOrderStatus, 1000)
var BlockOrderBookDiffChan = make(chan *BlockOrderBookDiff, 1000)

var UserStateReqChan = make(chan map[uint32]eth.Address, 1000)

var UserAndPositionSnapshotChan = make(chan map[uint64]Position, 1000)

type UserStateWithId struct {
	UserId    uint32
	UserState *UserState
}
