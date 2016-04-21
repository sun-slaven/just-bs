package redis
import (
	"log"
	"gopkg.in/redis.v3"
	"just.com/etc"
	"just.com/query/vo/user"
	"just.com/common"
	"time"
)

type RedisService struct {
	client *redis.Client
	log    *log.Logger
}

func NewRedisService(config *etc.RedisConfig, log *log.Logger) *RedisService {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Url,
		Password: "", // no password set
		DB:       0, // use default DB
	})
	redisService := new(RedisService)
	redisService.client = client
	redisService.log = log
	return redisService
}

/*聊天室model*/
type TalkMessage struct {
	Topic string    `json:"topic"`
	UserVo *user.UserVo    `json:"user_vo"`
	CreateTime string    `json:"craete_time"`
}

func NewTalkMessage(topic string,userVo *user.UserVo) *TalkMessage {
	talkMessage := new(TalkMessage)
	talkMessage.Topic = topic
	talkMessage.UserVo = userVo
	talkMessage.CreateTime = common.TimeFormat(time.Now())
	return talkMessage
}


