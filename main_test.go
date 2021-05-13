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

	t.Run("Strike roll should return 10 plus next two rolls", func(t *testing.T) {
		got := rollAndGetScore(10, 2, 3)
		want := 20

		assertScore(want, got, t)
	})

	t.Run("Strike roll and spare roll should return correct score", func(t *testing.T) {
		got := rollAndGetScore(10, 6, 4, 3, 0)
		want := 20 + 13 + 3

		assertScore(want, got, t)
	})

	t.Run("Perfect game returns 300 score", func(t *testing.T) {
		got := rollAndGetScore(10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10)
		want := 300

		assertScore(want, got, t)
	})

	t.Run("Negative roll pins returns an error", func(t *testing.T) {
		game := Game{}
		err := game.Roll(-1)
		assertError(ErrInvalidPinCount, err, t)
	})

	t.Run("Above 10 roll pins returns an error", func(t *testing.T) {
		game := Game{}
		err := game.Roll(11)
		assertError(ErrInvalidPinCount, err, t)
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

func assertError(want, got error, t *testing.T) {
	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}
