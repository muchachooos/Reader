package model

func MapOrder(msg Message) (Order, []Item, Delivery, Payment) {
	order := Order{
		OrderUid:          msg.OrderUid,
		TrackNumber:       msg.TrackNumber,
		Entry:             msg.Entry,
		Locale:            msg.Locale,
		InternalSignature: msg.InternalSignature,
		CustomerId:        msg.CustomerId,
		DeliveryService:   msg.DeliveryService,
		Shardkey:          msg.Shardkey,
		SmId:              msg.SmId,
		DateCreated:       msg.DateCreated,
		OofShard:          msg.OofShard,
	}

	items := make([]Item, 0, len(msg.Items))

	for i := range msg.Items {
		items = append(items, Item{
			OrderUid:    msg.OrderUid,
			ChrtId:      msg.Items[i].ChrtId,
			TrackNumber: msg.Items[i].TrackNumber,
			Price:       msg.Items[i].Price,
			Rid:         msg.Items[i].Rid,
			Name:        msg.Items[i].Name,
			Sale:        msg.Items[i].Sale,
			Size:        msg.Items[i].Size,
			TotalPrice:  msg.Items[i].TotalPrice,
			NmId:        msg.Items[i].NmId,
			Brand:       msg.Items[i].Brand,
			Status:      msg.Items[i].Status,
		})
	}

	delivery := Delivery{
		OrderUid: msg.OrderUid,
		Name:     msg.Delivery.Name,
		Phone:    msg.Delivery.Phone,
		Zip:      msg.Delivery.Zip,
		City:     msg.Delivery.City,
		Address:  msg.Delivery.Address,
		Region:   msg.Delivery.Region,
		Email:    msg.Delivery.Email,
	}

	payment := Payment{
		OrderUid:     msg.OrderUid,
		Transaction:  msg.Payment.Transaction,
		RequestId:    msg.Payment.RequestId,
		Currency:     msg.Payment.Currency,
		Provider:     msg.Payment.Provider,
		Amount:       msg.Payment.Amount,
		PaymentDt:    msg.Payment.PaymentDt,
		Bank:         msg.Payment.Bank,
		DeliveryCost: msg.Payment.DeliveryCost,
		GoodsTotal:   msg.Payment.GoodsTotal,
		CustomFee:    msg.Payment.CustomFee,
	}

	return order, items, delivery, payment
}
