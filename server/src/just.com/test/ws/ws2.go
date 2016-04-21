package main
import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"gopkg.in/redis.v3"
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
		client:= redis.NewClient(&redis.Options{
			Addr:  "localhost:6379",
		})
		defer client.Close()
		// 3.conn-->pubSub
		pubSub,pubSubErr:= client.Subscribe("NEWS")
		defer pubSub.Close()
		if pubSubErr != nil {
			Logger.Println(pubSubErr)
		}
		// channel
		var browserMessage chan []byte = make(chan []byte)
		var topicMessage chan string = make(chan string)
		go BrowserMessage2(browserMessage, wsConn)
		go TopicMessage2(topicMessage, pubSub)
		for {
			select {
			case m := <-browserMessage:
				log.Println("browser data is", string(m))
				if err:= client.Publish("NEWS",string(m)).Err();err!= nil{
					Logger.Println(err)
				}
			case m := <-topicMessage:
				Logger.Println("get topic send is:", m)
				// TextMessage or BinaryMessage
				writeErr:= wsConn.WriteMessage(websocket.TextMessage, []byte(m))
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
func BrowserMessage2(bMessage chan []byte, conn *websocket.Conn) {
	for {
		_, data, _ := conn.ReadMessage()
		bMessage<-data
	}
}

/*接收其他人发来的关于这个topic的消息*/
func TopicMessage2(topicMessage chan string, pubSub *redis.PubSub) {
	for {
		v,_ := pubSub.Receive()
		switch vTemp:= v.(type) {
		case *redis.Message:
			Logger.Printf("%s: message: %s\n", vTemp.Channel, vTemp.Payload)
			topicMessage <- vTemp.Payload
		case *redis.Subscription:
			Logger.Printf("%s: %s %d\n", vTemp.Channel, vTemp.Kind, vTemp.Count)
		case *redis.PMessage:
			Logger.Printf("%s: message: %s\n",vTemp.Channel,vTemp.Payload)
		default:
			Logger.Println(vTemp)
		}
	}
}