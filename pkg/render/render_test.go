package render

import (
	"fyne.io/fyne/v2"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas()

	if c == nil {
		t.Error("Expected New Canvas to return a non-nil value, but got nil")
	}

	_, isCanvasObject := c.Canvas.(fyne.CanvasObject)
	if !isCanvasObject {
		t.Error("Expected Canvas to be of type fyne.CanvasObject, but got a different type")
	}
}
