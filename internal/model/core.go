package model

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

var WGServer http.Server
var MyDB *xorm.Engine
var RedisDB *redis.Client
