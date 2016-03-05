package image_test
import (
	"testing"
	"just.com/service/image"
	"log"
	"just.com/test"
	"os"
)

func TestAdd(t *testing.T) {

	session := test.DataSource.NewSession()
	imageService := image.ImageService{}
	imageService.Session = session
	imageService.Log = log.New(os.Stdout, "mlearing", log.Llongfile)
	imageTable := imageService.Add("name", "url", "user_id", 100, 200)
	session.Commit()
	log.Println(imageTable)
}