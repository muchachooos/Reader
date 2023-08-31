package model

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrItemsNotFound = errors.New("items not found")
var ErrDeliveryNotFound = errors.New("delivery not found")
var ErrPaymentNotFound = errors.New("payment not found")

type Err struct {
	Error string `json:"error"`
}
