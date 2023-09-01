package handler

import (
	"reader/cache"
	"reader/storage"
)

type Server struct {
	Server *storage.Storage
	Cache  *cache.Cache
}
