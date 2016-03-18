package common_test
import (
	"just.com/common"
	"log"
	"testing"
	"html/template"
	"os"
)

func TestMd5(t *testing.T) {
	log.Println(common.Md5("123456"))
	Template()
}

func Template() {
	t := template.New("text")
	t,_= t.Parse(`点击以下连接，即可完成验证 <a href="{{.link}}">{{.link}}<a>`)
	linkMap := make(map[string]string)
	linkMap["link"] = "123"
	t.Execute(os.Stdout,linkMap)
}
