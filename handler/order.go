package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reader/model"
)

func (s *Server) GetOrderHandler(context *gin.Context) {
	orderUid, ok := context.GetQuery("order_uid")
	if orderUid == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "order_uid is missing"})
		return
	}

	msg, ok := s.Cache.Get(orderUid)
	if ok {
		context.JSON(200, msg)
		return
	}

	message, err := s.Server.GetOrder(orderUid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	s.Cache.Set(message)

	context.JSON(200, message)
}
