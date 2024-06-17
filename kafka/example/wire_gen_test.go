package example

import (
	"github.com/IBM/sarama"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func Test_SyncProducer(t *testing.T) {
	app, err := wireApp()
	if err != nil {
		t.Fatal(err)
	}

	p, o, err := app.SyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "ccccaa",
		Value: sarama.StringEncoder("axxx"),
	})
	if err != nil {
		app.Logger.Errorf("send message error: %v", err)
		return
	}
	app.Logger.Infof("partition: %d, offset: %d", p, o)

}

func Test_AsyncProducer(t *testing.T) {
	app, err := wireApp()
	if err != nil {
		t.Fatal(err)
	}
	defer app.AsyncProducer.AsyncClose()
	go func() {
		// [!important] 异步生产者发送后必须把返回值从 Errors 或者 Successes 中读出来 不然会阻塞 sarama 内部处理逻辑 导致只能发出去一条消息
		select {
		case _ = <-app.AsyncProducer.Successes():
		case e := <-app.AsyncProducer.Errors():
			if e != nil {
				app.Logger.Errorf("[Producer] err:%v msg:%+v \n", e.Msg, e.Err)
			}
		}
	}()
	var limit = 100
	var count int64 = 0
	// 异步发送
	for i := 0; i < 100; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{Topic: "12xxx", Key: nil, Value: sarama.StringEncoder(str)}
		// 异步发送只是写入内存了就返回了，并没有真正发送出去
		// sarama 库中用的是一个 channel 来接收，后台 goroutine 异步从该 channel 中取出消息并真正发送
		app.AsyncProducer.Input() <- msg
		atomic.AddInt64(&count, 1)
		if atomic.LoadInt64(&count)%2 == 0 {
			app.Logger.Errorf("已发送消息数:%v\n", count)
		}

	}
	app.Logger.Errorf("发送完毕 总发送消息数:%v\n", limit)
}
