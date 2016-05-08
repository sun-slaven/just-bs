package file
import (
	"strings"
	"just.com/value"
)

type ImageVo struct {
	Url    string `json:"url"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

func NewImageVo(url string, width, height int64) *ImageVo {
	image := new(ImageVo)
	if strings.TrimSpace(url) == "" {
		url = value.DEFAULT_IMAGE
	}
	image.Url = value.BASE_URL + url
	image.Width = width
	image.Height = height
	return image
}
