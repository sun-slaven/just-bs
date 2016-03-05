package main
import (
	"os"
	"log"
	"encoding/json"
	"just.com/etc"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"just.com/group/user"
	"net/http"
	"time"
	"just.com/middleware"
	"just.com/model/qiniu"
	"just.com/model/db"
)

func main() {
	// 1.config
	path := os.Getenv("JUST_PATH")
	configByte, configErr := ioutil.ReadFile(path + "/etc/config.json")
	if configErr != nil {
		log.Println(configErr)
	}
	config := etc.Config{}
	configUnmarshal := json.Unmarshal(configByte, &config)
	if configUnmarshal != nil {
		log.Println(configUnmarshal)
	}
	// 2.db
	dataSource := db.New(config.DBConfig)
	log.Println(dataSource)
	// 2.qiniu fs
	qiniuFileSystem := qiniu.New(config.QiniuConfig)
	log.Println(qiniuFileSystem)
	// interface
	gin.SetMode("release")
	router := gin.New()
	justGroup := router.Group("/just")
	justGroup.Use(middleware.ContextMiddleWare)
	justGroup.Use(middleware.LogMiddleware)
	justGroup.Use(middleware.TokenMiddleWare)
	user.BuildRouter(justGroup)

	s := &http.Server{
		Addr:":" + config.Port,
		Handler:router,
		ReadTimeout:60 * 60 * time.Second,
		WriteTimeout:60 * 60 * time.Second,
	}
	log.Println("liesten at" + config.Port)
	s.ListenAndServe()
}

