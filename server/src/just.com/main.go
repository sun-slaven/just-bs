package main
import (
	"os"
	"log"
	"encoding/json"
	"just.com/etc"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"just.com/model/qiniu"
	"just.com/model/db"
	"just.com/action/course"
	"just.com/action/user"
	"just.com/action/token"
	"just.com/middleware"
	"just.com/action/college"
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
	logger := log.New(os.Stdout, "[MLearing]", log.Ltime | log.Llongfile)
	// 2.db
	dataSource := db.NewDatSource(config.DBConfig)
	logger.Println(dataSource)
	// 3.redis
	//	rds := db.NewRedisDataSource(config.RedisConfig)
	// 3.qiniu fs
	qiniuFileSystem := qiniu.New(config.QiniuConfig)
	logger.Println(qiniuFileSystem)
	// email
	sendConfig := config.SendCloudConfig
	// interface
	//	gin.SetMode("release")
	router := gin.Default()
	router.Static("/res", path + "/res")
	router.Static("/web", path + "/..")
	router.StaticFile("/favicon.ico", path + "/res/favicon.ico")
	mLearingGroup := router.Group("/api/v1")
	mLearingGroup.Use(middleware.ContextMiddleWare(dataSource, logger))
	//	justGroup.Use(middleware.LogMiddleware)
	mLearingGroup.Use(middleware.TokenTest)
	mLearingGroup.Use(middleware.EmailMiddleware(sendConfig))
	//	mLearingGroup.Use(middleware.TokenMiddleWare)
	college.BuildRouter(mLearingGroup.Group("/colleges"))
	course.BuildRouter(mLearingGroup.Group("/courses"))
	user.BuildRouter(mLearingGroup.Group("/users"))
	token.BuildRouter(mLearingGroup.Group("/tokens"))
	s := &http.Server{
		Addr: config.Port,
		Handler:router,
		ReadTimeout:60 * 60 * time.Second,
		WriteTimeout:60 * 60 * time.Second,
	}
	log.Println("liesten at" + config.Port)
	s.ListenAndServe()
}

