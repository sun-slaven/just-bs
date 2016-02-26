package model

type Config struct {
	DBConfig `json:"db"`
	QiniuConfig `json:"qiniu"`
	Port string `json:"port"`
}

type DBConfig struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	MaxIdleConns int64 `json:"maxIdleConns"`
	MaxOpenConns int64 `json:"maxOpenConns"`
}

type QiniuConfig struct {
	Bucket string    `json:"bucket"`
	AK     string `json:"ak"`
	SK     string `json:"sk"`
	Domin  string `json:"domain"`
}
