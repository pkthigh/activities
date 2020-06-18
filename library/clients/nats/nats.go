package nats

import (
	"activities/common"
	"activities/library/clients/nats/message"
	"activities/library/config"
	"activities/library/logger"
	"time"

	stan "github.com/nats-io/stan.go"
)

var client *Nats

// Nats nats-io client
type Nats struct {
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
	client = &Nats{conn: conn, msgs: make(chan *stan.Msg, 1024)}
	// 订阅
	client.Sub(common.ItemSubject)
	client.Sub(common.HandOverNewSubject)
	client.Sub(common.PkcHandOverNewSubject)
	// 单线程消费消息
	go func() {
		for {
			msg := <-client.msgs
			switch msg.Subject {
			case common.ItemSubject.String():
				message.ItemRecordHandler(msg)
			}
		}
	}()
	// 每日12点定时任务
	go func() {
		for {
			// TODO: 每日统计并清空当日
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
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
