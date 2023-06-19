package main

import "testing"

func TestFav(t *testing.T) {
	got := fav("kasargod")
	want := "My fav place is:kasargod"

	if got != want {
		t.Errorf("Error!! got: %s and want is %s", got, want)
	}
}
