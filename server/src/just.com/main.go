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
	"just.com/middleware"
	action_router "just.com/action/router"
	"just.com/action/public"
)

func main() {
	// 1.config
	path := os.Getenv("JUST_PATH")
	configByte, configErr := ioutil.ReadFile(path + "/etc/config.json")
	if configErr != nil {
		log.Println(configErr)
	}
	config := new(etc.Config)
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
	qiniuFileSystem := qiniu.NewQiniuFileSystem(config.QiniuConfig)
	//	gin.SetMode("release")
	router := gin.Default()
	router.Static("/web", path + "/..")
	router.Static("/res", path + "/res")
	router.StaticFile("/favicon.ico", path + "/res/favicon.ico")
	mainGroup := router.Group("/api/v1")

	// middleware
	mainGroup.Use(middleware.ResponseMiddleware())
	mainGroup.Use(middleware.ContextMiddleWare(dataSource, logger))
	//	justGroup.Use(middleware.LogMiddleware)
	mainGroup.Use(middleware.FileSystemMiddlaware(qiniuFileSystem))
	mainGroup.Use(middleware.EmailMiddleware(config.SendCloudConfig))
	mainGroup.Use(middleware.ApiMiddleware())
	mainGroup.Use(middleware.TokenMiddleWare(config.WhiteListConfig))
//	mainGroup.Use(middleware.RoleMiddleware())
	// router
	action_router.BuildRouter(mainGroup)

	// public
	publicGroup := router.Group("/api/v1")
	publicGroup.Use(middleware.ContextMiddleWare(dataSource, logger))
	public.BuildRouter(publicGroup.Group("/public"))

	server := &http.Server{
		Addr: config.Port,
		Handler:router,
		ReadTimeout:60 * 60 * time.Second,
		WriteTimeout:60 * 60 * time.Second,
	}
	log.Println("liesten at" + config.Port)
	server.ListenAndServe()
}

