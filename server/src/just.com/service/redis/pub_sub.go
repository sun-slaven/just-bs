package redis
import "github.com/garyburd/redigo/redis"


func (self *RedisService) Pub() {
}

func (self *RedisService) Sub(msg string) error {
	psConn := redis.PubSubConn{
		Conn:self.conn,
	}
	return psConn.Subscribe()
}

func (self *RedisService) UnSub(msg string) error {
	psConn := redis.PubSubConn{
		Conn:self.conn,
	}
	return psConn.Unsubscribe()
}