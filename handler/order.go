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

	message, err := s.Server.GetOrder(orderUid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	context.JSON(200, message)
}
