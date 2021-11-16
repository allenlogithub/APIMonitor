/*
to test it, do the following:
	move the "modules" folder under the GOROOT
	go test -v -cover modules/readers
*/

package readers

import (
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
	var cfg requests.AppConfig
	cfg.Domain = "http://172.17.0.2:80"
	cfg.Rounds = 2
	cfg.Workers = 2
	var case1 requests.RequestConfig
	case1.Route = "/route"
	case1.RequestType = "GET"
	case1.EstElapse = 200
	cfg.Cases = append(cfg.Cases, case1)

	r := ReadSettings("/source/docs/simpleExample.json")
	if !reflect.DeepEqual(cfg, r) {
		t.Errorf("readers.ReadSettings of (%#v) was incorrect, got: %#v, want: %#v", "/source/docs/simpleExample.json", r, cfg)
	}
}
