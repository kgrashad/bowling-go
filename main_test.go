package bowling

import "testing"

func TestBowling(t *testing.T) {

	t.Run("Zero rolls should return zero score", func(t *testing.T) {
		got := rollAndGetScore(0)
		want := 0

		assertScore(want, got, t)
	})

	t.Run("Two rolls should return sum of pins", func(t *testing.T) {
		got := rollAndGetScore(2, 3)
		want := 5

		assertScore(want, got, t)
	})
}

func rollAndGetScore(rolls ...int) (score int) {
	game := Game{}

	for _, roll := range rolls {
		game.Roll(roll)
	}

	return game.Score()
}

func assertScore(want, got int, t *testing.T) {
	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}
