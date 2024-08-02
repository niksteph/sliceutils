package slicetransform_test

import(
	"testing"
	"github.com/niksteph/slicetransform"
	"slices"
)

func TestFilter(t *testing.T) {
	test := []int{2,3,4,7,8}
	want := []int{2,4,8}
	got := slicetransform.Filter(test, func(n int) bool {
		return n%2 == 0
	})
	if !slices.Equal(want,got) {
		t.Errorf("want:%v, but got %v", want, got)
	}
}
