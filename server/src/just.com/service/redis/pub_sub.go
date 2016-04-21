package redis
import "encoding/json"

func (self *RedisService) Publish(topic string, talkMessage *TalkMessage) {
	message, err := json.Marshal(talkMessage)
	if err != nil {
		return
	}
	publishErr := self.client.Publish(topic, string(message)).Err()
	if publishErr != nil {
		self.log.Println(publishErr)
		return
	}
}

func (self *RedisService) Subscribe(topic string) *TalkMessage {
	pubSub, err := self.client.Subscribe(topic)
	if err != nil {
		self.log.Println(err)
	}
	// igonore subscribe/pong message
	message, messageErr := pubSub.ReceiveMessage()
	if messageErr != nil {
		self.log.Println(messageErr)
	}
	talkMessage := new(TalkMessage)
	unmarshalErr := json.Unmarshal([]byte(message.Payload), talkMessage)
	if unmarshalErr != nil {
		self.log.Println(unmarshalErr)
		return
	}
	return talkMessage
}