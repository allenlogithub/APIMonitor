/*
to test it, do the following:
	move the "modules" folder under the GOROOT
	go test -v -cover modules/readers
*/

package readers

import (
	"testing"
)

func TestGetLocatedFolder(t *testing.T) {
	cases := []struct {
		input string
		want string
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
