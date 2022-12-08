package mascot

import "testing"

func TestBestMascot(t *testing.T) {
	got := BestMascot()
	want := "tux"

	assertEquals(got, want, t)
}

func assertEquals(got, want string, t *testing.T) {
	if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}