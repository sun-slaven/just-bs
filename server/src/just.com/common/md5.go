package common
import "crypto/md5"

func Md5(value string) string {
	// new
	md5Ctx := md5.New()
	// write
	md5Ctx.Write([]byte(value))
	// sum
	result := md5Ctx.Sum(nil)
	return string(result)
}