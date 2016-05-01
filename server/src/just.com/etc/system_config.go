package etc
import (
	"strings"
)

type Config struct {
	DBConfig `json:"db"`
	QiniuConfig `json:"qiniu"`
	DeploymentConfig DeploymentConfig `json:"deployment"`
	RedisConfig `json:"redis"`
	SendCloudConfig `json:"send_cloud"`
	WhiteListConfig  []White `json:"white_list"`
}

type DBConfig struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	MaxIdleConns int `json:"maxIdleConns"`
	MaxOpenConns int `json:"maxOpenConns"`
}

type QiniuConfig struct {
	AK          string `json:"ak"`
	SK          string `json:"sk"`
	Bucket      string    `json:"bucket"`
	Domin       string `json:"domain"`
	CallbackUrl string `json:"callback_url"`
}

type RedisConfig  struct {
	Url string `json:"url"`
}

type DeploymentConfig struct {
	Dev        DeploymentItemConfig        `json:"dev"`
	Production DeploymentItemConfig `json:"production"`
}

type DeploymentItemConfig struct {
	Port        string `json:"port"`
	SwaggerHost string `json:"swagger_host"`
}

type SendCloudConfig struct {
	ApiUser    string `json:"api_user"`
	ApiKey     string `json:"api_key"`
	From       string `json:"from"`
	FromName   string `json:"from_name"`
	Subject    string `json:"subject"`
	RequestUrl string `json:"request_url"`
	ActiveUrl  string `json:"active_url"`
}

type White struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (self *White) Match(method, path string) bool {
	path = strings.TrimSuffix(path, "/")
	if (strings.ToLower(self.Method) == strings.ToLower(method)) &&( self.Path == path) {
		return true
	}
	return false
}