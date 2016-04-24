package common
import (
	"crypto/md5"
	"encoding/hex"
)


func Md5(value string) string {
	// new
	md5Ctx := md5.New()
	// write
	md5Ctx.Write([]byte(value))
	// sum
	return hex.EncodeToString(md5Ctx.Sum(nil))
}