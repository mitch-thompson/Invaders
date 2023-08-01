package render

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var PROJECT_ROOT string = "../.."

func TestNewSprite(t *testing.T) {
	cases := []struct {
		name, path string
		want       Sprite
	}{
		{"Create DEFENDER Sprite", filepath.Join(PROJECT_ROOT, DEFENDER), Sprite{filepath.Join(PROJECT_ROOT, DEFENDER)}},
		{"Create UFO Sprite", filepath.Join(PROJECT_ROOT, ALIEN), Sprite{filepath.Join(PROJECT_ROOT, ALIEN)}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := NewSprite(c.path)
			if !reflect.DeepEqual(*got, c.want) {
				t.Errorf("Failed to create sprit: Expected %+v got %+v", c.want, got)
			}
			if err != nil {
				t.Errorf("Unexpected error %+v", err)
			}
		})
	}
}

func TestFailedNewSprite(t *testing.T) {
	path := ""
	_, err := NewSprite(path)
	if err == nil {
		t.Errorf("expected an error but got none")
	} else if !strings.Contains(err.Error(), "Required Sprite : file not found") {
		t.Errorf("unexpected error message: %v", err)
	}
}
