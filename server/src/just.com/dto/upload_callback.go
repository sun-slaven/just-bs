package dto

type UploadCallbackRequest struct {
	Key  string `json:"key" form:"key"`
	Hash string `json:"hash" form:"hash"`
	W    int64 `json:"w" form:"w"`
	H    int64 `json:"h" form:"h"`
}
