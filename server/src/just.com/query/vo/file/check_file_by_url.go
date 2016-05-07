package file
import (
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
)

func checkFileByUrl(url string, session *xorm.Session, log *log.Logger) bool {
	getFlag, getErr := session.Get(&table.FileTable{Url:url})
	if getErr != nil {
		log.Println(getErr)
	}
	return getFlag
}
