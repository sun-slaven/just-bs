package test
import (
	"just.com/model/db"
	"os"
	"io/ioutil"
	"log"
	"just.com/etc"
	"encoding/json"
)

var DataSource *db.DataSource
var Log *log.Logger

func init() {
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
	DataSource = db.NewDatSource(config.DBConfig)
	Log = log.New(os.Stdout, "[mlearing]", log.Ltime | log.Lshortfile)
}