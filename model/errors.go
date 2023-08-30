package model

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrItemsNotFound = errors.New("order not found")
var ErrDeliveryNotFound = errors.New("order not found")
var ErrPaymentNotFound = errors.New("order not found")

type Err struct {
	Error string `json:"error"`
}
