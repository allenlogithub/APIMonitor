package readers

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

func ReadJson(path string) map[string]interface{} {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	byteData, _ := ioutil.ReadAll(f)
	
	var result map[string]interface{}
	json.Unmarshal([]byte(byteData), &result)

	return result
}
