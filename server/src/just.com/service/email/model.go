package service

import (
	"bytes"
	"io/ioutil"
	"just.com/etc"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"just.com/service/token"
	"qiniupkg.com/x/log.v7"
)

type EmailService struct {
	Config etc.SendCloudConfig
}

func NewEamilService(config etc.SendCloudConfig) *EmailService {
	es := new(EmailService)
	es.Config = config
	return es
}

func (self *EmailService) SendMail(email, username string, userToken *service.UserToken) {
	RequestURI := self.Config.RequestUrl
	//不同于登录SendCloud站点的帐号，您需要登录后台创建发信子帐号，使用子帐号和密码才可以进行邮件的发送。
	activeUrl := getUrl(userToken)
	PostParams := url.Values{
		"api_user": {self.Config.ApiUser},
		"api_key":  {self.Config.ApiKey},
		"from":     {self.Config.From},
		"fromname": {self.Config.FromName},
		"substitution_vars":{getContent(email, username, activeUrl)},
		"template_invoke_name":     {"test_template_active"},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		panic(err)
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(BodyByte))
}

func getContent(email string, username string, url string) string {
	t := template.New("text")
	t, _ = t.Parse(`{"to": ["{{.to}}"],"sub":{"%name%": ["{{.name}}"],"%url%":["{{.url}}"]}}`)
	emailContent := make(map[string]string)
	emailContent["to"] = email
	emailContent["name"] = username
	emailContent["url"] = url
	buffer := new(bytes.Buffer)
	t.Execute(buffer, emailContent)
	return buffer.String()
}

func getUrl(userToken *service.UserToken) string {
	return "http://the-world-wang.top/user/active?id=" + userToken.Id + "&user_id=" + userToken.UserId
}

func Code(n int) {
	if n < 4 {
		n = 4
	}
	str := "0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z"
	strSlice := strings.Split(str, ",")
	length := len(strSlice)
	result := make([]string, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, strSlice[rand.Intn(length)])
	}
}
