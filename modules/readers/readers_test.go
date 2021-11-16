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
	paths := []struct {
		x string
		y string
	}{
		{"/a/b/", "/a/b/"},
		{"/a/b/c", "/a/b/"},
	}

	for _, path := range paths {
		rt := getLocatedFolder(path.x)
		if rt != path.y {
			t.Errorf("readers.getLocatedFolder of (%s) was incorrect, got: %s, want: %s", path.x, rt, path.y)
		}
	}
}
