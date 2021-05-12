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

	t.Run("Spare roll should return sum of pins and next roll", func(t *testing.T) {
		got := rollAndGetScore(4, 6, 3)
		want := 16

		assertScore(want, got, t)
	})

	t.Run("Two spare rolls return 25", func(t *testing.T) {
		got := rollAndGetScore(4, 6, 5, 5)
		want := 25

		assertScore(want, got, t)
	})

	t.Run("10 pins roll across two frame should not add next roll", func(t *testing.T) {
		got := rollAndGetScore(0, 6, 4, 2)
		want := 12

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
