package service
import (
	"log"
	"time"
	"code.google.com/p/go-uuid/uuid"
	"just.com/model/db/table"
	"just.com/common"
)

/*return id*/
func (self *TokenService) Make(userId string) (*XToken, error) {
	// 1.delete
	tokenTable := new(table.TokenTable)
	tokenTable.UserId = userId
	self.Session.Delete(tokenTable)
	// 2.insert
	tokenTable.UUID = uuid.New()
	//	tokenTable.UserId = userId
	tokenTable.CreateTime = time.Now()
	tokenTable.DeadTime = time.Now().Add(24 * time.Hour)
	insertNum, insertErr := self.Session.InsertOne(tokenTable)
	if insertNum == 0 {
		if insertErr != nil {
			log.Println(insertErr)
			return nil, TOKEN_MAKE_ERR
		}
		return nil, TOKEN_MAKE_ERR
	}
	xToken := new(XToken)
	xToken.Id = tokenTable.UUID
	xToken.UserId = tokenTable.UserId
	return xToken, nil
}

/*each check add dead time*/
func (self *TokenService) Check(xToken *XToken) bool {
	if common.IsEmpty(xToken.Id,xToken.UserId) == true{
		return false
	}
	sql := `SELECT * FROM "TOKEN"
		WHERE "UUID" = ?
		AND "USER_ID" = ?
	 	AND "DEAD_TIME" > ?`
	tokenTable := new(table.TokenTable)
	getFlag, getErr := self.Session.Sql(sql, xToken.Id, xToken.UserId, time.Now()).Get(tokenTable)
	if getErr != nil {
		self.Log.Println(getErr)
	}
	// update the dead time
	if getFlag == true {
		u := new(table.TokenTable)
		u.DeadTime = time.Now().Add(24 * time.Hour)
		_, updateErr := self.Session.Id(tokenTable.UUID).Update(u)
		if updateErr != nil {
			self.Log.Println(updateErr)
		}
	}
	return getFlag
}