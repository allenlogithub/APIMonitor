/*
to test it, do the following:
	move the "modules" folder under the GOROOT
	go test -v -cover modules/export
*/

package export

import (
	"testing"
)

func equalSlice(s1 []interface{}, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}

func stringSliceToInterfaceSlice(s []string) []interface{} {
	var rt []interface{}
	for _, v := range s {
		rt = append(rt, v)
	}

	return rt
}

func TestGetHeaders(t *testing.T) {
	type Input struct {
		X string
		Y string
	}
	type Want []interface{}
	cases := []struct {
		input Input
		want  Want
	}{
		{
			Input{
				X: "abc",
				Y: "abc",
			},
			Want{
				"X",
				"Y",
			},
		},
	}

	for _, c := range cases {
		rt := stringSliceToInterfaceSlice(getHeaders(c.input))
		if !equalSlice(rt, c.want) {
			t.Errorf("export.getHeaders of (%s) was incorrect, got: %s, want: %s", c.input, rt, c.want)
		}
	}
}

func TestGetValues(t *testing.T) {
	type Input struct {
		X string
		Y int
		Z interface{}
	}
	type Want []interface{}
	cases := []struct {
		input Input
		want  Want
	}{
		{
			Input{
				X: "abc",
				Y: 123,
				Z: 123,
			},
			Want{
				"abc",
				123,
				123,
			},
		},
		{
			Input{
				X: "abc",
				Y: 123,
				Z: "cba",
			},
			Want{
				"abc",
				123,
				"cba",
			},
		},
	}

	for _, c := range cases {
		rt := getValues(c.input)
		if !equalSlice(rt, c.want) {
			t.Errorf("export.getValues of (%#v) was incorrect, got: %#v, want: %#v", c.input, rt, c.want)
		}
	}
}
