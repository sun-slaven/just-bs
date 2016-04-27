package qiniu
import (
	"qiniupkg.com/api.v7/kodo"
	"just.com/etc"
)

type QiniuFileSystem  struct {
	ak          string
	sk          string
	bucket      string
	domain      string
	callBackUrl string
}

func New(config etc.QiniuConfig) *QiniuFileSystem {
	q := new(QiniuFileSystem)
	q.ak = config.AK
	q.sk = config.SK
	q.bucket = config.Bucket
	q.domain = config.Domin
	q.callBackUrl = config.CallbackUrl
	return q
}

func (self *QiniuFileSystem) MakeToken(key string, fileType string) string {
	kodo.SetMac(self.ak, self.sk)
	zone := 0
	c := kodo.New(zone, nil)
	policy := &kodo.PutPolicy{
		Scope:self.bucket + ":" + key,
		Expires:3600, // one hour
		CallbackUrl: self.callBackUrl,
	}
	switch fileType {
	case "image":
		policy.CallbackUrl = "key=$(key)&hash=$(etag)&w=$(imageInfo.width)&h=$(imageInfo.height)&fsize$(fsize)"
	case "file":
		policy.CallbackUrl = "key=$(key)&hash=$(etag)&fsize$(fsize)"
	case "video":
		policy.CallbackUrl = "key=$(key)&hash=$(etag)&w=$(imageInfo.width)&h=$(imageInfo.height)"
		policy.PersistentOps = "avthumb/mp4/s/640x360/vb/1.25m"
	}

	return c.MakeUptoken(policy)
}