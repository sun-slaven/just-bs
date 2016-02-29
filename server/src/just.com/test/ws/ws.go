package main
import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"github.com/garyburd/redigo/redis"
)

var upgreder = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool { return true },
}

func main() {
	router := gin.Default()
	test := router.Group("/test")
	test.Any("/ws", func(c *gin.Context) {
		log.Println("ws")
		conn, err := upgreder.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print(err)
			return
		}
		defer conn.Close()
		//		conn.WriteMessage(websocket.TextMessage, []byte("Hello World"))
		redisConn, connErr := redis.DialURL("redis://127.0.0.1:6379")
		if connErr != nil {
			log.Println(connErr)
		}
		defer redisConn.Close()
		pubConn := redis.PubSubConn{}
		pubConn.Conn = redisConn
		defer pubConn.Close()
		subScribeErr := pubConn.Subscribe("NEWS")
		if subScribeErr != nil {
			log.Println(subScribeErr)
		}
		pubConn.Receive()
		for {
			reply := pubConn.Receive()
			replyMess := reply.(redis.Message)
			conn.WriteMessage(websocket.TextMessage, replyMess.Data)
			//			replyStr, replyStrEr := redis.Values(redisConn.Receive())
			//			if replyStrEr != nil{
			//				log.Println(replyStrEr)
			//			}
			//			log.Println(replyStr)
			//			conn.WriteMessage(websocket.TextMessage, []byte("123"))
		}
	})

	test.Static("/cli", "D:/project/just/just_bs/server/src/just.com/test/ws")

	router.Run(":8080")
}