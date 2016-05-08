package file
import "strings"

const BASE_URL = "http://7xnz7k.com1.z0.glb.clouddn.com/"
const DEFAULT_IMAGE = "default.png"
const DEFAULT_FILE = "98620d12-1a5f-47ae-877a-296944f30b75.mp4"

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
