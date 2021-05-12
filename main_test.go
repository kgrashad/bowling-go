package bowling

import "testing"

func TestBowling(t *testing.T) {

	t.Run("Test zero rolls should return zero score", func(t *testing.T) {
		game := Game{}

		game.Roll(0)

		got := game.Score()
		want := 0

		assertScore(want, got, t)
	})
}

func assertScore(want, got int, t *testing.T) {
	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}
}
