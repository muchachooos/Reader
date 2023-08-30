package client

import "github.com/gin-gonic/gin"

func HtmlHandler(context *gin.Context) {
	context.File("./resources/html/index.html")
}
