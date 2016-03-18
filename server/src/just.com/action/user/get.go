package user
import (
	"just.com/middleware"
)

func UserGet(c *middleware.Context, userId string) (bool, error) {
	c.Ds.NewSession()
	c.Log.Println(userId)
	return false, nil
}