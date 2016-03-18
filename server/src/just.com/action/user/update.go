package user
import (
	"just.com/middleware"
	"just.com/dto"
)

func UserUpdate(c *middleware.Context, userId string, userDto *dto.UserDto) (bool, error) {
	return false, nil;
}