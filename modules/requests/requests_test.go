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
	cases := []struct{
		input Input
		want string
	}{
		{
			Input{
				"A": "1",
				"B": "2",
			},
			"A=1&B=2",
		},{
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
