package bowling

const MaxRolls int = 22
const NumberOfFrames int = 10
const MaxRollScore int = 10

type BowlingRolls [MaxRolls]int

type Game struct {
	rolls      BowlingRolls
	rollsCount int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.rollsCount] = pins
	g.rollsCount++
}

func (g *Game) Score() (score int) {
	rollIndex := 0
	for i := 0; i < NumberOfFrames; i++ {

		if isSpare(g.rolls, rollIndex) {
			score += 10 + g.rolls[rollIndex+2]
			rollIndex += 2
		} else if isStrike(g.rolls, rollIndex) {
			score += 10 + g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
			rollIndex++
		} else {
			score += g.rolls[rollIndex] + g.rolls[rollIndex+1]
			rollIndex += 2
		}
	}

	return
}

func isSpare(rolls BowlingRolls, i int) bool {
	return rolls[i]+rolls[i+1] == MaxRollScore
}

func isStrike(rolls BowlingRolls, i int) bool {
	return rolls[i] == MaxRollScore
}
