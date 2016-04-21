package file
import "just.com/model/db/table"

type ImageVo struct {
	Url    string `json:"url"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

func (self *ImageVo) NewImageVo(t *table.ImageTable)  {
	self.Url = t.Url
	self.Width = t.Width
	self.Height = t.Height
}

func NewImageVo(url string,width,height int64)  *ImageVo{
	image := new(ImageVo)
	image.Url = url
	image.Width = width
	image.Height = height
	return image
}
