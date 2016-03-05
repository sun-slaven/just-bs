package service
import (
	"github.com/go-xorm/xorm"
	"log"
	"time"
	"code.google.com/p/go-uuid/uuid"
	"errors"
	"just.com/model/db/table"
)

type TokenService  struct {
	Session *xorm.Session
	Log     *log.Logger
	data    table.TokenTable
}

func (self *TokenService) Make(userId string) (string, error) {
	// 1.check
	querySql := `SELECT COUNT("UUID") FROM "TOKEN" WHERE "USER_ID" = ? AND "FROZEN_TIME" > ?`
	count, _ := self.Session.Sql(querySql, userId, time.Now()).Count(&TokenService{})
	if count > 0 {
		return "", errors.New(SERVICE_TOKEN_COUNT_ERR)
	}
	// 2.insert
	tokenTable := table.TokenTable{}
	tokenTable.UUID = uuid.New()
	tokenTable.UserId = userId
	tokenTable.CreateTime = time.Now()
	tokenTable.DeadTime = time.Now().Add(2 * 24 * time.Hour)
	insertNum, insertErr := self.Session.InsertOne(&tokenTable)
	if insertNum == 0 {
		if insertErr != nil {
			log.Println(insertErr)
			return "", insertErr
		}
		return "", errors.New(SERVICE_TOKEN_INSERT_ERR)
	}
	return tokenTable.UUID, nil
}

func (self *TokenService) Check() bool {
	sql := `SELECT * FROM "TOKEN"
		WHERE "UUID" = ?
		AND "USER_ID" = ?
	 	AND "FROZEN_TIME" < ?`
	tokenTable := table.TokenTable{}
	getFlag, _ := self.Session.Sql(sql, self.data.UUID, self.data.UserId, time.Now()).Get(&tokenTable)
	return getFlag
}