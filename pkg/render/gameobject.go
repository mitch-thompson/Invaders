package render

const (
	DEFENDER = "static/images/defender.png"
	ALIEN    = "static/images/ufo.png"
)

type GameObject struct {
	path     string
	position Position
}

type Defender struct {
	GameObject
}

type Alien struct {
	GameObject
	isBottom bool
}

func NewDefender(position Position) *Defender {
	return &Defender{
		GameObject: GameObject{
			path:     DEFENDER,
			position: position,
		},
	}
}

func NewAlien(position Position, isBottom bool) *Alien {
	return &Alien{
		GameObject: GameObject{
			path:     ALIEN,
			position: position,
		},
		isBottom: isBottom,
	}
}
