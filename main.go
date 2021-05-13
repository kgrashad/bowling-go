package bowling

const MaxRolls int = 22
const NumberOfFrames int = 10
const MaxRollScore int = 10

type BowlingRolls [MaxRolls]int

type Game struct {
	rolls     BowlingRolls
	rollCount int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.rollCount] = pins
	g.rollCount++
}

func (g *Game) Score() (score int) {
	rollIndex := 0
	for i := 0; i < NumberOfFrames; i++ {

		if g.isSpare(rollIndex) {
			score += 10 + g.rolls[rollIndex+2]
			rollIndex += 2
		} else if g.isStrike(rollIndex) {
			score += 10 + g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
			rollIndex++
		} else {
			score += g.rolls[rollIndex] + g.rolls[rollIndex+1]
			rollIndex += 2
		}
	}

	return
}

func (g *Game) isSpare(i int) bool {
	return g.rolls[i]+g.rolls[i+1] == MaxRollScore
}

func (g *Game) isStrike(i int) bool {
	return g.rolls[i] == MaxRollScore
}
