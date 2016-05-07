package service
import (
	"just.com/model/db/table"
	"just.com/err"
	"just.com/common"
)

func (self *TokenService)Delete(userId string) *err.HttpError {
	if common.IsEmpty(userId) {
		return err.NO_REQUIRED_PARAM_FOUND
	}
	deleteNum, deleteErr := self.Session.Delete(&table.TokenTable{UserId:userId})
	if deleteNum == 0 {
		if deleteErr != nil {
			self.Log.Println(deleteErr)
		}
		return err.TOKEN_DELETE_ERR
	}
	return nil
}
