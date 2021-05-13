package bowling

const (
	MaxRolls       int = 22
	NumberOfFrames int = 10
	MaxRollScore   int = 10
)

const (
	ErrInvalidPinCount = BowlingErr("invalid pin count, should be 0-10")
	ErrTooManyRolls    = BowlingErr("you can't roll more than 12 times per game")
)

type BowlingErr string

type BowlingRolls [MaxRolls]int

type Game struct {
	rolls     BowlingRolls
	rollCount int
}

func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > 10 {
		return ErrInvalidPinCount
	}

	if g.rollCount == 12 {
		return ErrTooManyRolls
	}

	g.rolls[g.rollCount] = pins
	g.rollCount++
	return nil
}

func (g *Game) Score() (score int) {
	rollIndex := 0
	for i := 0; i < NumberOfFrames; i++ {

		if g.isSpare(rollIndex) {
			score += MaxRollScore + g.rolls[rollIndex+2]
			rollIndex += 2
		} else if g.isStrike(rollIndex) {
			score += MaxRollScore + g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
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

func (e BowlingErr) Error() string {
	return string(e)
}
