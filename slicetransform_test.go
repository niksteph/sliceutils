package slicetransform_test

import (
	"fmt"
	"github.com/niksteph/slicetransform"
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	test := []int{2, 3, 4, 7, 8}
	want := []int{2, 4, 8}
	got := slicetransform.Filter(test, func(n int) bool {
		return n%2 == 0
	})
	if !slices.Equal(want, got) {
		t.Errorf("want:%v, but got %v", want, got)
	}
}

func TestMap(t *testing.T) {
	test := []int{2, 3, 4, 7, 8}
	want := []string{"10", "11", "100", "111", "1000"}
	got := slicetransform.Map(test, func(n int) string {
		return fmt.Sprintf("%b", n)
	})
	if !slices.Equal(want, got) {
		t.Errorf("want:%v, but got %v", want, got)
	}
}

func TestReduce(t *testing.T) {
	test := []int{2, 3, 4, 7, 8}
	want := 2 + 3 + 4 + 7 + 8
	got := slicetransform.Reduce(test, 0, func(intermed, current int) int {
		return intermed + current
	})
	if want != got {
		t.Errorf("want:%v, but got %v", want, got)
	}
}

func TestZipEqualLength(t *testing.T) {
	test1 := []int{4, 3, 2, 1}
	test2 := []int{1, 2, 3, 4}
	want := []float64{4.0 / 1.0, 3.0 / 2.0, 2.0 / 3.0, 1.0 / 4.0}
	got := slicetransform.Zip(test1, 0, test2, 1, func(numerator, denominator int) float64 {
		return float64(numerator) / float64(denominator)
	})
	if !slices.Equal(want, got) {
		t.Errorf("want:%v, but got %v", want, got)
	}
}

func TestZipShorterFirst(t *testing.T) {
	test1 := []int{4, 3, 2}
	test2 := []int{1, 2, 3, 4}
	want := []float64{4.0 / 1.0, 3.0 / 2.0, 2.0 / 3.0, 0.0}
	got := slicetransform.Zip(test1, 0, test2, 1, func(numerator, denominator int) float64 {
		return float64(numerator) / float64(denominator)
	})
	if !slices.Equal(want, got) {
		t.Errorf("want:%v, but got %v", want, got)
	}
}

func TestZipLongerFirst(t *testing.T) {
	test1 := []int{4, 3, 2, 1}
	test2 := []int{1, 2, 3}
	want := []float64{4.0 / 1.0, 3.0 / 2.0, 2.0 / 3.0, 1.0}
	got := slicetransform.Zip(test1, 0, test2, 1, func(numerator, denominator int) float64 {
		return float64(numerator) / float64(denominator)
	})
	if !slices.Equal(want, got) {
		t.Errorf("want:%v, but got %v", want, got)
	}
}
