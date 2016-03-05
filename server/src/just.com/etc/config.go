package etc

type Config struct {
	DBConfig `json:"db"`
	QiniuConfig `json:"qiniu"`
	Port string `json:"port"`
}

type DBConfig struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	MaxIdleConns int `json:"maxIdleConns"`
	MaxOpenConns int `json:"maxOpenConns"`
}

type QiniuConfig struct {
	AK     string `json:"ak"`
	SK     string `json:"sk"`
	Bucket string    `json:"bucket"`
	Domin  string `json:"domain"`
}
