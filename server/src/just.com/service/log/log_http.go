package log
import (
	"log"
	"just.com/model/db"
	"just.com/model/db/table"
	"code.google.com/p/go-uuid/uuid"
	"time"
	"encoding/json"
)

type LogHttpService struct {
	ds  *db.DataSource
	log *log.Logger
}

func NewLogHttpService(ds *db.DataSource, log *log.Logger) *LogHttpService {
	return &LogHttpService{
		ds:ds,
		log:log,
	}
}

//单独的事务
func (self *LogHttpService)Log(method, url, body, userId, error string, status int, responseBody interface{}) {
	session := self.ds.NewSession()
	beginErr := session.Begin()
	if beginErr != nil {
		self.log.Println(beginErr)
		return
	}
	dataByte, marshalErr := json.Marshal(responseBody)
	if marshalErr != nil {
		self.log.Println(marshalErr)
		return
	}
	logTable := &table.LogHttp{
		UUID:uuid.New(),
		Method:method,
		Url:url,
		Body:body,
		UserId:userId,
		Status:status,
		ErrorMessage:error,
		ResponseBody:string(dataByte),
		CreateTime:time.Now(),
	}
	insertNum, insertErr := session.InsertOne(logTable)
	if insertNum == 0 {
		if insertErr != nil {
			self.log.Println(insertErr)
		}
		rollbackErr := session.Rollback()
		if rollbackErr != nil {
			self.log.Println(rollbackErr)
		}
		return
	}
	commitErr := session.Commit()
	if commitErr != nil {
		rollbackErr := session.Rollback()
		if rollbackErr != nil {
			self.log.Println(rollbackErr)
		}
		self.log.Println(commitErr)
		return
	}
}
