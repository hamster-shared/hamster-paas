package models

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	in := []int64{5, 2, 6, 3, 1, 4}
	want := []int64{4, 1, 3, 6, 2, 5}
	got := reverse(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("reverse(%v)=%v,want %v", in, got, want)
	}
}
