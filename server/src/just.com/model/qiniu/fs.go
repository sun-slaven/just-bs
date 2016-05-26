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

const UPLOAD_TYPE_ICON = "icon"
const UPLOAD_TYPE_ATTACHMENT = "attachment"
const UPLOAD_TYPE_VIDEO = "video"

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
	case UPLOAD_TYPE_ICON:
		policy.CallbackBody = "key=$(key)&w=$(imageInfo.width)&h=$(imageInfo.height)&type=icon"
	case UPLOAD_TYPE_ATTACHMENT:
		policy.CallbackBody = "key=$(key)&type=attachment"
	case UPLOAD_TYPE_VIDEO:
		policy.CallbackBody = "key=$(key)&type=video"
//		policy.PersistentOps = "avthumb/mp4/s/640x360/vb/1.25m"
		policy.PersistentOps = "avthumb/mp4"
	}
	return &FileToken{Token:c.MakeUptoken(policy)}
}