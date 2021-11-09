package readers

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"

	"modules/requests"
)

type config = requests.AppConfig

func ReadSettings(path string) config {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	byteData, _ := ioutil.ReadAll(f)

	var cfg config
	json.Unmarshal(byteData, &cfg)

	return cfg
}
