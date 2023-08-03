package render

import (
	"fyne.io/fyne/v2"
	"reflect"
	"testing"
)

func TestNewNullLayout(t *testing.T) {
	want := &NullLayout{}
	got := NewNullLayout()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Failed creating NullLayout %v, got %v", want, got)
	}
}

func TestNullLayout_MinSize(t *testing.T) {
	want := fyne.NewSize(0, 0)
	nullLayout := NullLayout{}
	got := nullLayout.MinSize(nil)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Min Size of Null Layout incorrect: Expected %v, got %v", want, got)
	}
}
