package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"reader/cache"
	"reader/client"
	"reader/configuration"
	"reader/handler"
	"reader/model"
	"reader/stanreader"
	"reader/storage"
	"strconv"
)

const configPath = "configuration.json"

func main() {
	config := configuration.GetConfig(configPath)

	// Подключаемся к базе данных
	db, err := sqlx.Open("postgres", getDSN(config.DBConf))
	if err != nil {
		panic(err)
	}

	s := storage.Storage{DB: *db}
	c := cache.NewCache()

	reader := stanreader.Reader{Storage: &s, Cache: c}
	reader.Run()

	server := handler.Server{Server: &s, Cache: c}
	router := gin.Default()

	// Возвращаяет HtmlHandler клиента
	router.GET("/", client.HtmlHandler)

	// Возвращает инфорацию о заказе
	router.GET("/order", server.GetOrderHandler)

	port := ":" + strconv.Itoa(8080)
	err = router.Run(port)
	if err != nil {
		panic(err)
	}
}

func getDSN(cfg model.DBConf) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Sslmode)
}
