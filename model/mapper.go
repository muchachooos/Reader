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

func MapMessage(order Order, items []Item, delivery Delivery, payment Payment) Message {
	orderJSON := Order{
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerId:        order.CustomerId,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.Shardkey,
		SmId:              order.SmId,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
	}

	itemsJSON := make([]ItemJSON, 0, len(items)-1)

	for i := range items {
		itemsJSON = append(itemsJSON, ItemJSON{
			ChrtId:      items[i].ChrtId,
			TrackNumber: items[i].TrackNumber,
			Price:       items[i].Price,
			Rid:         items[i].Rid,
			Name:        items[i].Name,
			Sale:        items[i].Sale,
			Size:        items[i].Size,
			TotalPrice:  items[i].TotalPrice,
			NmId:        items[i].NmId,
			Brand:       items[i].Brand,
			Status:      items[i].Status,
		})
	}

	deliveryJSON := DeliveryJSON{
		Name:    delivery.Name,
		Phone:   delivery.Phone,
		Zip:     delivery.Zip,
		City:    delivery.City,
		Address: delivery.Address,
		Region:  delivery.Region,
		Email:   delivery.Email,
	}

	paymentJSON := PaymentJSON{
		Transaction:  payment.Transaction,
		RequestId:    payment.RequestId,
		Currency:     payment.Currency,
		Provider:     payment.Provider,
		Amount:       payment.Amount,
		PaymentDt:    payment.PaymentDt,
		Bank:         payment.Bank,
		DeliveryCost: payment.DeliveryCost,
		GoodsTotal:   payment.GoodsTotal,
		CustomFee:    payment.CustomFee,
	}

	result := Message{
		OrderUid:          orderJSON.OrderUid,
		TrackNumber:       orderJSON.TrackNumber,
		Entry:             orderJSON.Entry,
		Delivery:          deliveryJSON,
		Payment:           paymentJSON,
		Items:             itemsJSON,
		Locale:            orderJSON.Locale,
		InternalSignature: orderJSON.InternalSignature,
		CustomerId:        orderJSON.CustomerId,
		DeliveryService:   orderJSON.DeliveryService,
		Shardkey:          orderJSON.Shardkey,
		SmId:              orderJSON.SmId,
		DateCreated:       orderJSON.DateCreated,
		OofShard:          orderJSON.OofShard,
	}

	return result
}
