package user
import (
	"just.com/middleware"
)

func UserDelete(c *middleware.Context, userId string) (bool, error) {
	c.Log.Println(userId)
	return false, nil
}