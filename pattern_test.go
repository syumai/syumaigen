package syumaigen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parsePattern(t *testing.T) {
	const value = "012\n345"
	want := [][]int{{0, 1, 2}, {3, 4, 5}}
	got := parsePattern(value)
	if d := cmp.Diff(want, got); d != "" {
		t.Errorf(d)
	}
}
