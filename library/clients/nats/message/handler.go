package message

import (
	"activities/common"
	"activities/library/logger"
	"activities/library/storage"
	"encoding/json"
	"strconv"
	"time"

	"github.com/nats-io/stan.go"
)

// ItemRecordHandler 道具消息处理
func ItemRecordHandler(msg *stan.Msg) {
	var record ItemRecord
	if err := json.Unmarshal(msg.Data, &record); err != nil {
		logger.ErrorF("ItemRecordHandler JSON Unmarshal error: %v", err)
		return
	}
	// 忽略时间币
	if record.CoinType != COIN_TYPE_PKC {
		return
	}
	store, _ := storage.GetRdsDB(common.ItemRecordStore)

	userid := strconv.Itoa(record.Uid)
	if !store.Exists(userid).Val() {
		if err := store.Set(userid, record.Cost, time.Hour*24*8).Err(); err != nil {
			logger.ErrorF("ItemRecordHandler redis set %v error: %v", userid, err)
		}
	} else {
		if err := store.IncrByFloat(userid, record.Cost).Err(); err != nil {
			logger.ErrorF("ItemRecordHandler redis IncrByFloat %v error: %v", userid, err)
		}
	}
	logger.InfoF("[道具]Subject: %v, UserID: %v", msg.Subject, userid)
}

// HandOverRecordHandler 每手消息处理
func HandOverRecordHandler(msg *stan.Msg) {
	var record HandOverRecord
	if err := json.Unmarshal(msg.Data, &record); err != nil {
		logger.ErrorF("HandOverRecordHandler JSON Unmarshal error: %v", err)
		return
	}
	if record.CoinType != COIN_TYPE_PKC {
		return
	}

	// 忽略非联盟局消息
	if record.UnionId == 0 {
		return
	}

	hors, _ := storage.GetRdsDB(common.HandOverRecordStore)
	irs, _ := storage.GetRdsDB(common.InsuranceRecordStore)

	// 记录手数
	for key := range record.Process.Players {
		userid := strconv.Itoa(int(key))
		if !hors.Exists(userid).Val() {
			if err := hors.Set(userid, 1, time.Hour*24*8).Err(); err != nil {
				logger.ErrorF("HandOverRecordStore redis Set %v error: %v", userid, err)
			}
		} else {
			if err := hors.Incr(userid).Err(); err != nil {
				logger.ErrorF("HandOverRecordStore redis Incr %v error: %v", userid, err)
			}
		}
		logger.InfoF("[手数]Subject: %v, UserID: %v", msg.Subject)
	}

	// 记录保险
	for _, ins := range record.InsDetails {
		userid := strconv.Itoa(ins.UserId)
		if !irs.Exists(userid).Val() {
			if err := irs.Set(userid, ins.Buy, time.Hour*24*8).Err(); err != nil {
				logger.ErrorF("InsuranceRecordStore redis Set %v error: %v", userid, err)
			}
		} else {
			if err := irs.IncrBy(userid, ins.Buy).Err(); err != nil {
				logger.ErrorF("InsuranceRecordStore redis IncrBy %v error: %v", userid, err)
			}
		}
		logger.InfoF("[保险]Subject: %v, UserID: %v", msg.Subject, userid)
	}
}
