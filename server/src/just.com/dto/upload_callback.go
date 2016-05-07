package dto

type UploadCallbackRequest struct {
	Type string `json:"type" form:"type"`
	Key  string `json:"key" form:"key"`
	W    int64 `json:"w" form:"w"`
	H    int64 `json:"h" form:"h"`
}
