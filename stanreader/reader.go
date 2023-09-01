package stanreader

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"reader/cache"
	"reader/model"
	"reader/storage"
)

type Reader struct {
	Storage *storage.Storage
	Cache   *cache.Cache
}

func (r *Reader) Run() {
	// Подключение к NATS
	nc, err := nats.Connect(stan.DefaultNatsURL)
	if err != nil {
		panic(err)
	}

	sc, err := stan.Connect("test-cluster", "MY_TEST_READER", stan.NatsConn(nc))
	if err != nil {
		panic(err)
	}

	// Настройка подписки(получить всё, получить новое и др)
	startOpt := stan.StartAt(pb.StartPosition_NewOnly)

	// Оформление подписки на канал
	_, err = sc.Subscribe("TEST_CHANNEL", r.getMsgHandler, startOpt)
	if err != nil {
		panic(err)
	}
}

func (r *Reader) getMsgHandler(msg *stan.Msg) {
	var message model.Message

	err := json.Unmarshal(msg.Data, &message)
	if err != nil {
		panic(err)
	}

	order, items, delivery, payment := model.MapOrder(message)
	if err != nil {
		panic(err)
	}

	err = r.Storage.CreateOrder(order)
	if err != nil {
		panic(err)
	}

	err = r.Storage.CreateItem(items)
	if err != nil {
		panic(err)
	}

	err = r.Storage.CreateDelivery(delivery)
	if err != nil {
		panic(err)
	}

	err = r.Storage.CreatePayment(payment)
	if err != nil {
		panic(err)
	}

	// Кладём заказ в кэш
	r.Cache.Set(message)
}
