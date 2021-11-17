/*
to test it, do the following:
	move the "modules" folder under the GOROOT
	go clean -testcache && \
	go test -v -cover modules/readers
*/

package readers

import (
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"modules/requests"
)

func TestGetLocatedFolder(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"/a/b/", "/a/b/"},
		{"/a/b/c", "/a/b/"},
	}

	for _, c := range cases {
		rt := getLocatedFolder(c.input)
		if rt != c.want {
			t.Errorf("readers.getLocatedFolder of (%s) was incorrect, got: %s, want: %s", c.input, rt, c.want)
		}
	}
}

func TestReadSettings(t *testing.T) {
	filePath := "/source/test.json"
	jsonString := `{
	"domain": "http://172.17.0.2:80",
	"rounds": 2,
	"workers": 2,
	"cases": [
		{
			"route": "/route",
			"request_type": "GET",
			"est_elapse": 200
		}
	]
}
	`
	// assign json data into a defined struct
	var cfg requests.AppConfig
	json.Unmarshal([]byte(jsonString), &cfg)

	// create file and write json data into it
	f, err := os.Create(filePath)
	if err != nil {
		t.Errorf("Create test.json file failed.")
	}
	defer f.Close()
	buf := bytes.NewBufferString(jsonString)
	_, err = buf.WriteTo(f)
	if err != nil {
		t.Errorf("Error in writing json string to test.json")
	}

	// load the json data from the generated json file
	r := ReadSettings(filePath)
	if !reflect.DeepEqual(cfg, r) {
		t.Errorf("readers.ReadSettings of (%#v) was incorrect, got: %#v, want: %#v", "/source/docs/simpleExample.json", r, cfg)
	}

	// remove the json file
	err = os.Remove(filePath)
	if err != nil {
		t.Errorf("Error in removing json file")
	}
}
