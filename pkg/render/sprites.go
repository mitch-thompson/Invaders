package render

import (
	"fmt"
	"os"
)

const (
	DEFENDER = "static/images/defender.png"
	ALIEN    = "static/images/ufo.png"
)

// Sprite represents an object on the canvas
type Sprite struct {
	path string
}

// NewSprite creates a new instances of the Sprite
func NewSprite(path string) (*Sprite, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("Required Sprite " + path + ": file not found")
	}
	return &Sprite{path}, nil
}
