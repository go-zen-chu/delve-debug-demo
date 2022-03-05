package count

import "testing"

func TestParallel1000count(t *testing.T) {
	t.Run("when nothing given, it should return 1000", func(t *testing.T) {
		want := 1000
		if actual := Parallel1000count(); actual != want {
			t.Errorf("parallel1000count() = %d, want %d", actual, want)
		}
	})
}
