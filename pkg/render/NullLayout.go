package render

import "fyne.io/fyne/v2"

type NullLayout struct {
}

func NewNullLayout() *NullLayout {
	return &NullLayout{}
}

func (n *NullLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}

func (n *NullLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
}
