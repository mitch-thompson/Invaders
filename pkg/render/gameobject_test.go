package render

import (
	"reflect"
	"testing"
)

func TestNewDefender(t *testing.T) {
	want := Defender{
		GameObject{
			path:     DEFENDER,
			position: Position{0, 0},
		}}
	testPosition := Position{0, 0}
	got := NewDefender(testPosition)
	if reflect.DeepEqual(got, want) {
		t.Errorf("Failed creating Defender: Expected %v, got %v", want, got)
	}
}

func TestNewAlien(t *testing.T) {
	want := Alien{
		GameObject: GameObject{
			path:     ALIEN,
			position: Position{0, 0},
		},
		isBottom: false,
	}
	testPosition := Position{0, 0}
	got := NewAlien(testPosition, false)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Failed to create Alien: Expected %v, got %v", want, got)
	}
}
