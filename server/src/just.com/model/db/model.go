package db
import (
	"github.com/go-xorm/xorm"
	"just.com/etc"
	"log"
	_ "github.com/lib/pq"
	"io"
)

type DataSource struct {
	*xorm.Engine
}

func NewDatSource(config etc.DBConfig, out io.Writer) *DataSource {
	ds := new(DataSource)
	engine, err := xorm.NewEngine(config.Name, config.Url)
	if err != nil {
		log.Println(err)
		return nil
	}
	engine.SetMaxOpenConns(config.MaxOpenConns)
	engine.SetMaxIdleConns(config.MaxIdleConns)
	engine.ShowSQL = true        // 打开 sql
	engine.SetLogger(xorm.NewSimpleLogger(out))
	engine.Charset("UTF-8")
	ds.Engine = engine
	return ds
}

/*
func NewRedisDataSource(config etc.RedisConfig) *RedisDataSource {
	redisCon, redisErr := redis.DialURL(config.Url)
	if redisErr != nil {
		log.Println(redisErr)
	}
	rds := new(RedisDataSource)
	rds.Conn = redisCon
	return rds
}*/
