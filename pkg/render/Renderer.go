package render

type Position struct {
	x, y float32
}

type Renderer interface {
	DrawMenu()
	DrawDefender(position Position)
	DrawAlien(position Position, isAlien bool)
	ClearScreen()
	DrawScore(score int)
}

type Game struct {
	renderer Renderer
}

func NewPosition(x, y float32) *Position {
	return &Position{
		x: x,
		y: y,
	}
}

func NewGame(renderer Renderer) *Game {
	return &Game{
		renderer: renderer,
	}
}

func (g *Game) MovePlayer(newPosition Position) {
	g.renderer.DrawDefender(newPosition)
}

func (g *Game) MoveAlien(newPosition Position, alien Alien) {
	g.renderer.DrawAlien(newPosition, alien.isBottom)
}
