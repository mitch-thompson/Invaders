package render

import (
	"fyne.io/fyne/v2"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewEmptyGameBoard(t *testing.T) {
	want := &GameBoard{
		gameObjects: []*GameObject{},
	}

	got := NewEmptyGameBoard()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Failed creating GameBoard: Expected %v, got %v", want, got)
	}
}

func TestNewGameObject(t *testing.T) {
	sprite := &Sprite{filepath.Join(PROJECT_ROOT, ALIEN)}
	canvasObject := fyne.CanvasObject(nil)
	want := GameObject{
		object:       *sprite,
		x:            0,
		y:            0,
		isAlien:      false,
		isBottom:     false,
		canvasObject: canvasObject,
	}
	got := NewGameObject(sprite, 0, 0, false, false)

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("Failed creating GameObject: Expected %v, got %v", want, got)
	}
}
