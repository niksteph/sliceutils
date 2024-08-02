package slicetransform_test

import(
	"testing"
	"github.com/niksteph/slicetransform"
	"slices"
	"fmt"
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

func TestMap(t *testing.T) {
	test := []int{2,3,4,7,8}
	want := []string{"10","11","100","111","1000"}
	got := slicetransform.Map(test, func(n int) string {
		return fmt.Sprintf("%b", n)
	})
	if !slices.Equal(want,got) {
		t.Errorf("want:%v, but got %v", want, got)
	}
}

func TestReduce(t *testing.T) {
	test := []int{2,3,4,7,8}
	want := 2+3+4+7+8
	got := slicetransform.Reduce(test, 0, func(intermed, current int) int {
		return intermed+current
	})
	if want != got {
		t.Errorf("want:%v, but got %v", want, got)
	}
}
