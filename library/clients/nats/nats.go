package nats

import (
	"activities/common"
	"activities/library/clients/nats/message"
	"activities/library/config"
	"activities/library/logger"
	"activities/library/storage"
	"encoding/json"
	"time"

	stan "github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
)

var client *Nats

// Nats nats-io client
type Nats struct {
	time string
	conn stan.Conn
	msgs chan *stan.Msg
}

func init() {
	conf := config.GetNatsConf()
	conn, err := stan.Connect(
		conf.Cluster,
		conf.Client,
		stan.NatsURL(conf.URLs),
		stan.ConnectWait(15*time.Second),
		stan.Pings(3, 5),
		stan.SetConnectionLostHandler(
			func(c stan.Conn, err error) {
				logger.FatalF("nats connect interrupt: %v", err)
			},
		),
	)
	if err != nil {
		logger.ErrorF("nats connect error: %v", err)
		panic(err)
	}
	client = &Nats{time: time.Now().Format("2006-01-02"), conn: conn, msgs: make(chan *stan.Msg, 1024)}
	// 订阅
	client.Sub(common.ItemSubject)
	client.Sub(common.PkcHandOverNewSubject)
	// 单线程消费消息
	go func() {
		for {
			msg := <-client.msgs
			switch msg.Subject {
			case "DailyStatistics":
				client.DailyStatistics(msg)
			case common.ItemSubject.String():
				message.ItemRecordHandler(msg)
			case common.PkcHandOverNewSubject.String():
				message.HandOverRecordHandler(msg)
			}

		}
	}()
	// 每日12点定时任务(统计今日数据并清空当日数据)
	go func() {
		for {
			nowtime := time.Now()
			next := nowtime.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(nowtime))
			<-t.C
			client.msgs <- &stan.Msg{pb.MsgProto{Subject: "DailyStatistics", Data: []byte(nowtime.Format("2006-01-02"))}, nil}
		}
	}()
}

// Sub 订阅
func (nats *Nats) Sub(subject common.Subject) {
	_, err := nats.conn.Subscribe(subject.String(), func(msg *stan.Msg) {
		nats.msgs <- msg
	})
	if err != nil {
		logger.ErrorF("Nats Sub Subject %v fail: %v", subject, err)
	}
}

// Close 连接关闭
func (nats *Nats) Close() error {
	if err := nats.conn.Close(); err != nil {
		logger.ErrorF("nats conn close error: %v", err)
		return err
	}
	return nil
}

// Statistics 统计
type Statistics struct {
	Item      map[string]string `json:"item"`
	Hand      map[string]string `json:"hand"`
	Insurance map[string]string `json:"insurance"`
}

// JSONFormatString Json格式字符串
func (statistics *Statistics) JSONFormatString() string {
	bytes, _ := json.Marshal(statistics)
	return string(bytes[:])
}

// DailyStatistics 每日统计
func (nats *Nats) DailyStatistics(msg *stan.Msg) {
	statistics := &Statistics{
		Item:      make(map[string]string),
		Hand:      make(map[string]string),
		Insurance: make(map[string]string),
	}

	itemStore, _ := storage.GetRdsDB(common.ItemRecordStore)
	for _, uid := range itemStore.Keys("*").Val() {
		statistics.Item[uid] = itemStore.Get(uid).Val()
	}
	itemStore.FlushDb()

	handStore, _ := storage.GetRdsDB(common.HandOverRecordStore)
	for _, uid := range handStore.Keys("*").Val() {
		statistics.Hand[uid] = itemStore.Get(uid).Val()
	}
	handStore.FlushDb()

	insuranceStore, _ := storage.GetRdsDB(common.InsuranceRecordStore)
	for _, uid := range insuranceStore.Keys("*").Val() {
		statistics.Insurance[uid] = itemStore.Get(uid).Val()
	}
	insuranceStore.FlushDb()

	dailyStore, _ := storage.GetRdsDB(common.DailyStatistics)
	if err := dailyStore.Set(nats.time, statistics.JSONFormatString(), time.Hour*24*30).Err(); err != nil {
		logger.ErrorF("DailyStatistics redis set Time: %s Data: %s Error: %s", nats.time, statistics.JSONFormatString(), err)
	} else {
		logger.InfoF("DailyStatistics Time: %v Data:%v", nats.time, statistics.JSONFormatString())
	}
	nats.time = string(msg.Data[:])
}
