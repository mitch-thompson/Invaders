package render

import (
	"reflect"
	"testing"
)

type mockRenderer struct {
}

func (m *mockRenderer) DrawMenu() {
}

func (m *mockRenderer) DrawDefender(position Position) {
}

func (m *mockRenderer) DrawAlien(position Position, isAlien bool) {
}

func (m *mockRenderer) ClearScreen() {
}

func (m *mockRenderer) DrawScore(score int) {
}

func TestNewPosition(t *testing.T) {
	want := &Position{0, 0}
	got := NewPosition(0, 0)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Failed creating New Position: Expected %v, got %v", want, got)
	}
}

func TestNewGame(t *testing.T) {
	renderer := &mockRenderer{}
	want := renderer
	game := NewGame(renderer)
	if game.renderer != want {
		t.Errorf("Failed creating New Game: Expected %v, got %v", want, game.renderer)
	}
}
