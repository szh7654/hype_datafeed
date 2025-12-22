package datahandler

import (
	dr "indexer/data_retriver"

	"github.com/emirpasic/gods/v2/maps/treemap"
	"github.com/rs/zerolog/log"
)

var (
	books = []Book{}
)

type Book struct {
	LimitAsks   Levels
	LimitBids   Levels
	TriggerAsks Levels
	TriggerBids Levels
}
type OidToSize = map[uint64]float64 // key为oid，value为sz

type Levels struct{ *treemap.Map[float64, Level] } // key为price

type Level = map[uint32]OidToSize // key为 userId

func addBook() {
	books = append(books, Book{
		LimitAsks:   Levels{treemap.New[float64, Level]()},
		LimitBids:   Levels{treemap.New[float64, Level]()},
		TriggerAsks: Levels{treemap.New[float64, Level]()},
		TriggerBids: Levels{treemap.New[float64, Level]()},
	})
}

func (t Book) applyBids(addrOrders []dr.AddrOrder) {
	t.LimitBids.applyaAddrOrders(addrOrders)
}

func (t Book) applyAsks(addrOrders []dr.AddrOrder) {
	t.LimitAsks.applyaAddrOrders(addrOrders)
}

func (t Levels) applyaAddrOrders(addrOrders []dr.AddrOrder) {
	for _, addrOrder := range addrOrders {
		_, userId := GetUser(addrOrder.Addr)
		setActive(userId, addrOrder.Addr)
		px := addrOrder.Order.LimitPx
		sz := addrOrder.Order.Sz
		oid := addrOrder.Order.Oid
		t.add(px, sz, oid, userId)
	}
}

func (t Levels) applyOrderBookDiff(userId uint32, orderBookDiff dr.OrderBookDiff) {
	t.applyLevelAction(LevelAction{Px: orderBookDiff.Px, Sz: orderBookDiff.OrderBookDiffAction.NewSz, Oid: orderBookDiff.Oid, UserId: userId, ActionType: orderBookDiff.OrderBookDiffAction.ActionType})
}


type LevelAction struct {
	Px         float64
	Sz         float64
	Oid        uint64
	UserId     uint32
	ActionType dr.ActionType
}

func (t Levels) applyLevelAction(levelAction LevelAction) {
	switch levelAction.ActionType {
	case dr.ActionTypeNew:
		t.add(levelAction.Px, levelAction.Sz, levelAction.Oid, levelAction.UserId)
	case dr.ActionTypeUpdate:
		t.update(levelAction.Px, levelAction.Sz, levelAction.Oid, levelAction.UserId)
	case dr.ActionTypeRemove:
		t.remove(levelAction.Px, levelAction.Sz, levelAction.Oid, levelAction.UserId)
	}
}

func (t Levels) add(px float64, sz float64, oid uint64, userId uint32) {
	if level, ok := t.Get(px); ok {
		oidToSz := level[userId]
		if oidToSz == nil {
			oidToSz = map[uint64]float64{oid: sz}
			level[userId] = oidToSz
		} else {
			oidToSz[oid] = sz
			level[userId] = oidToSz
		}
	} else { // new level
		level := map[uint32]OidToSize{userId: map[uint64]float64{oid: sz}}
		t.Put(px, level)
	}
}

func (t Levels) remove(px float64, sz float64, oid uint64, userId uint32) {
	if level, ok := t.Get(px); ok {
		oidToSz := level[userId]
		if oidToSz == nil {
			log.Warn().Msgf("user id %v at price %f not found", userId, px)
		} else {
			oidToSz[oid] = sz
			level[userId] = oidToSz
		}
	} else {
		log.Warn().Msgf("level not found %v", px)
	}
}

func (t Levels) update(px float64, sz float64, oid uint64, userId uint32) {
	if level, ok := t.Get(px); ok {
		oidToSz := level[userId]
		if oidToSz == nil {
			log.Warn().Msgf("user id %v at price %f not found, continue update", userId, px)
			oidToSz[oid] = sz
			level[userId] = oidToSz
		} else {
			oidToSz[oid] = sz
			level[userId] = oidToSz
		}
	} else {
		level := map[uint32]OidToSize{userId: map[uint64]float64{oid: sz}}
		t.Put(px, level)
		log.Warn().Msgf("level not found %v for user id %v during update, add this order", px, userId)
	}
}
