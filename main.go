package bowling

type Game struct {
	rolls []int
}

func (g *Game) Roll(pins int) {
	g.rolls = append(g.rolls, pins)
}

func (g *Game) Score() (score int) {
	count := len(g.rolls)

	for i := 0; i < count; i++ {

		if i < count-2 && g.rolls[i]+g.rolls[i+1] == 10 {
			score += g.rolls[i+2]
		}

		score += g.rolls[i]
	}

	return
}
