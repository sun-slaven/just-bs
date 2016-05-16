package etc
import (
	"os"
	"log"
	"io/ioutil"
)

func PrintBanner(path string, logger *log.Logger) {
	banner, err := os.Open(path + "/etc/banner.txt")
	if err != nil {
		logger.Println(err)
	}
	bytes, readErr := ioutil.ReadAll(banner)
	if readErr != nil {
		logger.Println(readErr)
	}
	log.Println(string(bytes))
}
