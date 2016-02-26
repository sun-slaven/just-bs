package model

type Response struct {
	Ok   int64 `json:"ok"`
	Data interface{} `json:"data"`
	Err  string `json:"err"`
}
