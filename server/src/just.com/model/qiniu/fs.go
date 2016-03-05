package qiniu
import (
	"qiniupkg.com/api.v7/kodo"
	"just.com/etc"
)

type QiniuFileSystem  struct {
	ak     string
	sk     string
	bucket string
	domain string
}

func New(config etc.QiniuConfig) *QiniuFileSystem {
	q := new(QiniuFileSystem)
	q.ak = config.AK
	q.sk = config.SK
	q.bucket = config.Bucket
	q.domain = config.Domin
	return q
}

func (self *QiniuFileSystem) MakeToken(key string) string {
	kodo.SetMac(self.ak, self.sk)
	zone := 0
	c := kodo.New(zone, nil)
	policy := &kodo.PutPolicy{
		Scope:self.bucket + ":" + key,
		Expires:3600, // one hour
	}
	return c.MakeUptoken(policy)
}