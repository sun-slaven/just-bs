package etc

type Config struct {
	DBConfig `json:"db"`
	QiniuConfig `json:"qiniu"`
	Port string `json:"port"`
	RedisConfig `json:"redis"`
	SendCloudConfig `json:"send_cloud"`
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

type SendCloudConfig struct {
	ApiUser    string `json:"api_user"`
	ApiKey     string `json:"api_key"`
	From       string `json:"from"`
	FromName   string `json:"from_name"`
	Subject    string `json:"subject"`
	RequestUrl string `json:"request_url"`
	ActiveUrl  string `json:"active_url"`
}
