package main
import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
		conn.WriteMessage(websocket.TextMessage, []byte("Hello World"))
	})

	test.Static("/cli", "D:/project/just/just_bs/server/src/just.com/test/ws")

	router.Run(":8080")
}
