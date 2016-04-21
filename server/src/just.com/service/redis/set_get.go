package redis
import (
	"time"
	"encoding/json"
)

/*序列化以后的对象*/
func (self *RedisService) Set(cacheKey *CacheKey, value interface{}) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		self.log.Println(err)
		return err
	}
	cmd := self.client.Set(cacheKey.string(), []byte(valueByte), time.Hour * 24)
	return cmd.Err()
}

func (self *RedisService) Get(cacheKey *CacheKey) (string, error) {
	return self.client.Get(cacheKey.string()).Result()
}