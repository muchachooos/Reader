package storage

import (
	"database/sql"
	"errors"
	"reader/model"
)

func (s *Storage) GetOrder(orderUid string) (model.Message, error) {
	// Делаем запрос в базу данных на получение заказа по его UID
	var order model.Order

	err := s.DB.Get(&order, "SELECT * FROM new_order WHERE order_uid = $1", orderUid)
	if err != nil {
		// Если заказа с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrOrderNotFound
		}
		return model.Message{}, err
	}

	// Делаем запрос в базу данных на получение данных о товарах в заказе по его UID
	items := make([]model.Item, 0)

	err = s.DB.Select(&items, "SELECT * FROM item WHERE order_uid = $1", orderUid)
	if err != nil {
		// Если итемов принадлежащих заказу с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrItemsNotFound
		}
		return model.Message{}, err
	}

	// Делаем запрос в базу данных на получение данных о доставке по UID заказа
	var delivery model.Delivery

	err = s.DB.Get(&delivery, "SELECT * FROM delivery WHERE order_uid = $1", orderUid)
	if err != nil {
		// Если данных о доставке заказа с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrDeliveryNotFound
		}
		return model.Message{}, err
	}

	// Делаем запрос в базу данных на получение данных о оплате по UID заказа
	var payment model.Payment

	err = s.DB.Get(&payment, "SELECT * FROM payment WHERE order_uid = $1", orderUid)
	if err != nil {
		// Если данных об оплате заказа с таким UID нет, возвращаем ошибку
		if errors.Is(err, sql.ErrNoRows) {
			return model.Message{}, model.ErrPaymentNotFound
		}
		return model.Message{}, err
	}

	return model.MapMessage(order, items, delivery, payment), err
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
