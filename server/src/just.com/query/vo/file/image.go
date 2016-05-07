package file
import "strings"

const BASE_URL = "http://7xnz7k.com1.z0.glb.clouddn.com/"
const DEFAULT_IMAGE = "default.png"

type ImageVo struct {
	Url    string `json:"url"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

func NewImageVo(url string, width, height int64) *ImageVo {
	image := new(ImageVo)
	if strings.TrimSpace(url) == "" {
		url = DEFAULT_IMAGE
	}
	image.Url = BASE_URL + url
	image.Width = width
	image.Height = height
	return image
}
