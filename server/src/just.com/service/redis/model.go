package redis
import (
	"log"
	"gopkg.in/redis.v3"
)

type RedisService struct {
	conn *redis.Client
	log  *log.Logger
}

func NewRedisService(conn *redis.Client, log *log.Logger) *RedisService {
	if conn == nil || log == nil {
		return
	}
	rs := new(RedisService)
	rs.conn = conn
	rs.log = log
	return rs
}
