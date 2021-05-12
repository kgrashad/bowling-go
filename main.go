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

		if isSpare(g.rolls, i) {
			score += g.rolls[i+2]
		}

		score += g.rolls[i]
	}

	return
}

func isSpare(rolls []int, i int) bool {
	return i < len(rolls)-2 && i%2 == 0 && rolls[i]+rolls[i+1] == 10
}
