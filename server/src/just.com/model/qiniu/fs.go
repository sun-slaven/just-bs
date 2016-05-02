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

type FileToken struct {
	Token string        `json:"token"`
}

func NewQiniuFileSystem(config etc.QiniuConfig) *QiniuFileSystem {
	q := new(QiniuFileSystem)
	q.ak = config.AK
	q.sk = config.SK
	q.bucket = config.Bucket
	q.domain = config.Domin
	q.callBackUrl = config.CallbackUrl
	return q
}

// 生成上传token
func (self *QiniuFileSystem) MakeToken(fileType string) *FileToken {
	kodo.SetMac(self.ak, self.sk)
	zone := 0
	c := kodo.New(zone, nil)
	policy := &kodo.PutPolicy{
		Scope:self.bucket,
		Expires:1200, // 20 min
		CallbackUrl: self.callBackUrl,
	}
	switch fileType {
	case "image":
		policy.CallbackBody = "ey=$(key)&hash=$(etag)&w=$(imageInfo.width)&h=$(imageInfo.height)"
	case "file":
		policy.CallbackBody = "key=$(key)&fsize$(fsize)"
	case "video":
		policy.CallbackBody = "key=$(key)&fsize$(fsize)"
		policy.PersistentOps = "avthumb/mp4/s/640x360/vb/1.25m"
	}
	return &FileToken{Token:c.MakeUptoken(policy)}
}