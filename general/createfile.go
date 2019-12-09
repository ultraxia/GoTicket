package general

import (
	"encoding/json"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"path/filepath"
)

func Writefile(cookie selenium.Cookie) {
	if fileName, err := filepath.Abs("cookies.pkl"); err != nil {
		log.Fatalln(err)
	} else {

		file, _ := os.Open(fileName)
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cookie)
	}
	return
}
