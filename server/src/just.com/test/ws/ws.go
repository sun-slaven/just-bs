package main
import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"github.com/garyburd/redigo/redis"
	"os"
)

var upgreder = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool { return true },
}
var Logger *log.Logger = log.New(os.Stdout, "[test]:", log.Lshortfile | log.Ltime)

func main() {
	router := gin.Default()
	test := router.Group("/test")
	test.Any("/ws", func(c *gin.Context) {
		// 1.http --> ws conn
		wsConn, err := upgreder.Upgrade(c.Writer, c.Request, nil)
		defer wsConn.Close()
		if err != nil {
			Logger.Println(err)
			return
		}
		// 2. dial-->conn
		redisConn, connErr := redis.DialURL("redis://127.0.0.1:6379")
		defer redisConn.Close()
		if connErr != nil {
			Logger.Println(connErr)
		}
		// 3.conn-->pubSubConn
		pubConn := new(redis.PubSubConn)
		pubConn.Conn = redisConn
		defer pubConn.Close()
		subScribeErr := pubConn.Subscribe("NEWS")
		if subScribeErr != nil {
			Logger.Println(subScribeErr)
		}
		pubConn.Receive()
		//		//  订阅成功的信息
		//		for {
		//			// 订阅信息
		//			reply := pubConn.Receive()
		//			replyMess := reply.(redis.Message)
		//			log.Println(replyMess)
		//			conn.WriteMessage(websocket.TextMessage, replyMess.Data)
		//			//			replyStr, replyStrEr := redis.Values(redisConn.Receive())
		//			//			if replyStrEr != nil{
		//			//				log.Println(replyStrEr)
		//			//			}
		//			//			log.Println(replyStr)
		//			//			conn.WriteMessage(websocket.TextMessage, []byte("123"))
		//		}
		// channel
		var browserMessage chan []byte = make(chan []byte,4)
		var topicMessage chan []byte = make(chan []byte,4)
		go BrowserMessage(browserMessage, wsConn)
		go TopicMessage(topicMessage, pubConn)
		for {
			select {
			case m := <-browserMessage:
				log.Println("browser data is", string(m))
				replay,replayErr:=  redisConn.Do("PUBLISH", "NEWS", string(m))
				if replayErr != nil{
					Logger.Println(replayErr)
				}
				Logger.Println(replay)
			case m := <-topicMessage:
				Logger.Println("get topic send is:", string(m))
				// TextMessage or BinaryMessage
				writeErr:= wsConn.WriteMessage(websocket.TextMessage, m)
				if writeErr!= nil{
					Logger.Println(writeErr)
				}
			}
		}
	})
	test.Static("/client", "D:project/just/just_bs/server/src/just.com/test/ws")
	router.Run(":8080")
}

/*接收来自浏览器的数据*/
func BrowserMessage(bMessage chan []byte, conn *websocket.Conn) {
	for {
		_, data, _ := conn.ReadMessage()
		bMessage<-data
	}
}

/*接收其他人发来的关于这个topic的消息*/
func TopicMessage(topicMessage chan []byte, psc *redis.PubSubConn) {
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			Logger.Printf("%s: message: %s\n", v.Channel, v.Data)
			topicMessage <- v.Data
		case redis.Subscription:
			Logger.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case redis.PMessage:
			Logger.Printf("%s: message: %s\n",v.Channel,v.Data)
		case error:
			Logger.Println(v.Error())
		}
	}
}