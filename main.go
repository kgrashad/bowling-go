package bowling

type Game struct {
	score int
}

func (g *Game) Roll(pins int) {
	g.score += pins
}

func (g *Game) Score() (score int) {
	return g.score
}
