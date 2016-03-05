package db
import (
	"github.com/go-xorm/xorm"
	"just.com/etc"
	"log"
	_ "github.com/lib/pq"
//	"github.com/go-xorm/core"
)

type DataSource struct {
	*xorm.Engine
}

func New(config etc.DBConfig) *DataSource {
	ds := new(DataSource)
	engine, err := xorm.NewEngine(config.Name, config.Url)
	if err != nil {
		log.Println(err)
		return nil
	}
	engine.SetMaxOpenConns(config.MaxOpenConns)
	engine.SetMaxIdleConns(config.MaxIdleConns)
//	engine.SetTableMapper(core.SameMapper{})
//	engine.SetColumnMapper(core.SameMapper{})
	//	engine.SetLogger()
	engine.ShowSQL(true)
	engine.Charset("UTF-8")
	ds.Engine = engine
	return ds
}
