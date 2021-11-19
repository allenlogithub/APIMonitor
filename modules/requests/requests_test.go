/*
to test it, do the following:
	move the "modules" folder under the GOROOT
	go clean -testcache && \
	go test -v -cover modules/requests
*/

package requests

import (
	"testing"
)

func TestGetUrlParams(t *testing.T) {
	type Input map[string]string
	cases := []struct {
		input Input
		want  string
	}{
		{
			Input{
				"A": "1",
				"B": "2",
			},
			"A=1&B=2",
		}, {
			Input{},
			"",
		},
	}
	for _, c := range cases {
		rt := getUrlParams(c.input)
		if rt != c.want {
			t.Errorf("requests.getUrlParams of (%#v) was incorrect, got: %#v, want: %#v", c.input, rt, c.want)
		}
	}
}

func TestGetFormParams(t *testing.T) {
	type Input map[string]interface{}
	cases := []struct {
		input Input
		want  string
	}{
		{
			Input{
				"A": "1",
				"B": "2",
			},
			"A=1&B=2",
		}, {
			Input{},
			"",
		}, {
			Input{
				"A": 1,
				"B": true,
			},
			"A=1&B=true",
		}, {
			Input{
				"A": 1.1,
				"B": nil,
			},
			"A=1.1&B=%3Cnil%3E",
		},
	}
	for _, c := range cases {
		rt := getFormParams(c.input)
		if rt != c.want {
			t.Errorf("requests.getFormParams of (%#v) was incorrect, got: %#v, want: %#v", c.input, rt, c.want)
		}
	}
}
