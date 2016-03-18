package service

import (
	"testing"
	"just.com/service/image"
	"log"
	"just.com/test"
	"os"
)

var imageService image.ImageService

func init() {
	session := test.DataSource.NewSession()
	imageService = image.ImageService{}
	imageService.Session = session
	imageService.Log = log.New(os.Stdout, "mlearing", log.Llongfile)
}

func TestAdd(t *testing.T) {
	err := imageService.Add("name", "123", "user_id", 100, 200)
	if err != nil {
		t.Fail()
	}
	imageService.Session.Commit()
}

func TestFindById(t *testing.T) {
	imageTable, err := imageService.FindById("1bf178e0-3402-475b-a4b2-3740c286bb3d")
	log.Println(imageTable)
	if err != nil {
		t.Fail()
	}
}

func TestFindByUrl(t *testing.T) {
	imageTable, err := imageService.FindByUrl("123")
	if err != nil {
		t.Fail()
	}
	log.Println(imageTable)
}