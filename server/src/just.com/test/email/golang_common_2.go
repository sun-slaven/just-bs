package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"html/template"
)

type EmailContext struct {
	To  []string
	Sub interface{}
}

func SendMail2() {
	RequestURI := "http://sendcloud.sohu.com/webapi/mail.send_template.json"

	//不同于登录SendCloud站点的帐号，您需要登录后台创建发信子帐号，使用子帐号和密码才可以进行邮件的发送。
	PostParams := url.Values{
		"api_user": {"will19940412_test_TQpPHj"},
		"api_key":  {"A0wbotFTuAZNz8dH"},
		"from":     {"service@sendcloud.im"},
		"fromname": {"王浩"},
		"substitution_vars":{`{"to": ["992444037@qq.com"],"sub":{"%name%": ["小泡子仔"],"%url%":["www.baidu.com"]}}`},
		"template_invoke_name":  {"test_template"},
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
	fmt.Println(string(BodyByte))
}

func main() {
	//	SendMail2()
	fmt.Println(getContent("992444037@qq.com", "小泡子仔", ""))
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
