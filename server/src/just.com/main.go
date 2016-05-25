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
	"flag"
	"just.com/action/token"
)

func main() {
	logger := log.New(os.Stdout, "[MLearing]", log.Ltime | log.Llongfile)
	flag.Parse()
	deployment := flag.Arg(0)
	// 1.config
	path := os.Getenv("JUST_PATH")
	configByte, configErr := ioutil.ReadFile(path + "/etc/config.json")
	if configErr != nil {
		log.Println(configErr)
	}
	config := new(etc.Config)
	configUnmarshal := json.Unmarshal(configByte, config)
	if configUnmarshal != nil {
		log.Println(configUnmarshal)
	}
	// 2.db
	sqlLog, sqlLogErr := os.OpenFile(path + "/log/sql_log", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if sqlLogErr != nil {
		logger.Println(sqlLogErr)
		return
	}
	defer sqlLog.Close()
	dataSource := db.NewDatSource(config.DBConfig, sqlLog)
	defer dataSource.Close()
	// 3.redis
	//	rds := db.NewRedisDataSource(config.RedisConfig)
	// 3.qiniu fs
	qiniuFileSystem := qiniu.NewQiniuFileSystem(config.QiniuConfig)
	//	gin.SetMode("release")
	router := gin.Default()
	router.Static("/web", path + "/..")
	router.Static("/res", path + "/res")
	mainGroup := router.Group("/api/v1")

	// middleware
	mainGroup.Use(middleware.ResponseMiddleware())
	mainGroup.Use(middleware.ContextMiddleWare(dataSource, logger))
	mainGroup.Use(middleware.FileSystemMiddleware(qiniuFileSystem))
	mainGroup.Use(middleware.EmailMiddleware(config.SendCloudConfig))
	mainGroup.Use(middleware.ApiMiddleware())
	mainGroup.Use(middleware.TokenMiddleWare(config.WhiteListConfig))
	mainGroup.Use(middleware.LogMiddleware())
	//	mainGroup.Use(middleware.RoleMiddleware())
	// router
	action_router.BuildRouter(mainGroup)

	// public
	publicGroup := router.Group("/api/v1")
	publicGroup.Use(middleware.ResponseMiddleware())
	publicGroup.Use(middleware.ContextMiddleWare(dataSource, logger))

	tokenGroup := router.Group("/api/v1")
	tokenGroup.Use(middleware.ResponseMiddleware())
	tokenGroup.Use(middleware.ContextMiddleWare(dataSource, logger))
	tokenGroup.OPTIONS("/tokens", token.TokenOptionHandle)
	// deployment
	var deploymentItem etc.DeploymentItemConfig
	if deployment == "production" {
		deploymentItem = config.DeploymentConfig.Production
	}else {
		deploymentItem = config.DeploymentConfig.Dev
	}
	public.BuildRouter(publicGroup.Group("/public"), path, deploymentItem.SwaggerHost)

	server := &http.Server{
		Addr: deploymentItem.Port,
		Handler:router,
		ReadTimeout:60 * 60 * time.Second,
		WriteTimeout:60 * 60 * time.Second,
	}
	etc.PrintBanner(path, logger)
	listenErr := server.ListenAndServe()
	if listenErr != nil {
		log.Println(listenErr)
		return
	}
	log.Println("liesten at" + deploymentItem.Port)
}