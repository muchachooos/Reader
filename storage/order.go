package storage

import (
	"database/sql"
	"errors"
	"reader/model"
)

func (s *Storage) GetOrder(orderUid string) (model.Message, error) {
	// Делаем запрос в базу данных на получение заказа по его UID
	rowOrder := s.DB.QueryRow("SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM new_order WHERE order_uid = $1", orderUid)

	// Заполняем модель данными из результата запроса
	var order model.Order
	err := rowOrder.Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.Shardkey, &order.SmId, &order.DateCreated, &order.OofShard)
	if err != nil {
		// Если данных о заказе с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrOrderNotFound
		}
		return model.Message{}, err
	}

	// Делаем запрос в базу данных на получение данных о товарах в заказе по его UID
	rowsItems, err := s.DB.Query("SELECT chrt_id, track_number, price, rid, name, sale, 'size', total_price, nm_id, brand, Status FROM item WHERE order_uid = $1", orderUid)
	if err != nil {
		// Если итемов принадлежащих заказу с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrItemsNotFound
		}
		return model.Message{}, err
	}

	// Заполняем модель данными из результата запроса
	items := make([]model.ItemJSON, 0)
	for rowsItems.Next() {
		var item model.ItemJSON
		err = rowsItems.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
		if err != nil {
			return model.Message{}, err
		}

		items = append(items, item)
	}

	// Делаем запрос в базу данных на получение данных о доставке по UID заказа
	rowDelivery := s.DB.QueryRow("SELECT name, phone, zip, city, city, region, email FROM delivery WHERE order_uid = $1", orderUid)

	// Заполняем модель данными из результата запроса
	var delivery model.DeliveryJSON
	err = rowDelivery.Scan(&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
	if err != nil {
		// Если доставки с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrDeliveryNotFound
		}
		return model.Message{}, err
	}

	// Делаем запрос в базу данных на получение данных о оплате по UID заказа
	rowPayment := s.DB.QueryRow("SELECT 'transaction', request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment WHERE order_uid = $1", orderUid)

	// Заполняем модель данными из результата запроса
	var Payment model.PaymentJSON
	err = rowPayment.Scan(&Payment.Transaction, &Payment.RequestId, &Payment.Currency, &Payment.Provider, &Payment.Amount, &Payment.PaymentDt, &Payment.Bank, &Payment.DeliveryCost, &Payment.GoodsTotal, &Payment.CustomFee)
	if err != nil {
		// Если оплаты с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrPaymentNotFound
		}
		return model.Message{}, err
	}

	res := model.Message{
		OrderUid:          orderUid,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Delivery:          delivery,
		Payment:           Payment,
		Items:             items,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerId:        order.CustomerId,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmId:              order.SmId,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
	}

	return res, err
}

func (s *Storage) CreateOrder(order model.Order) error {
	result, err := s.DB.Exec("INSERT INTO new_order (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if countOfModifiedRows == 0 {
		return nil
	}

	return nil
}

func (s *Storage) CreateItem(items []model.Item) error {
	for i := range items {
		result, err := s.DB.Exec("INSERT INTO item (order_uid,chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)", items[i].OrderUid, items[i].ChrtId, items[i].TrackNumber, items[i].Price, items[i].Rid, items[i].Name, items[i].Sale, items[i].Size, items[i].TotalPrice, items[i].NmId, items[i].Brand, items[i].Status)
		if err != nil {
			return err
		}

		countOfModifiedRows, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if countOfModifiedRows == 0 {
			return nil
		}
	}

	return nil
}

func (s *Storage) CreateDelivery(delivery model.Delivery) error {
	result, err := s.DB.Exec("INSERT INTO delivery (order_uid,name,phone,zip,city,address,region,email) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)", delivery.OrderUid, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err != nil {
		return err
	}

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if countOfModifiedRows == 0 {
		return nil
	}

	return nil
}

func (s *Storage) CreatePayment(payment model.Payment) error {
	result, err := s.DB.Exec("INSERT INTO payment (order_uid,transaction,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", payment.OrderUid, payment.Transaction, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	if err != nil {
		return err
	}

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if countOfModifiedRows == 0 {
		return nil
	}

	return nil
}
