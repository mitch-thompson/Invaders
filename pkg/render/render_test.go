package render

import (
	"fyne.io/fyne/v2"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(640, 480)

	if c == nil {
		t.Error("Expected New Canvas to return a non-nil value, but got nil")
	}

	_, isCanvasObject := c.Canvas.(fyne.CanvasObject)
	if !isCanvasObject {
		t.Error("Expected Canvas to be of type fyne.CanvasObject, but got a different type")
	}
}

func TestCalculateButtonSize(t *testing.T) {
	cases := []struct {
		name string
		size fyne.Size
		want fyne.Size
	}{
		{"Greater than 10", fyne.NewSize(640, 480), fyne.NewSize(10, 10)},
		{"Less than 10", fyne.NewSize(90, 90), fyne.NewSize(9, 9)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := calculateButtonSize(c.size)
			if got != c.want {
				t.Errorf("Expected button size to be %v, but got %v", c.want, got)
			}
		})
	}
}
