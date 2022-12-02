package spacify

import "testing"

func TestSpacify(t *testing.T) {
	got := Spacify("hello")
	expected := "h e l l o"

	if got != expected {
		t.Errorf("got %s expected %s", got, expected)
	}
}
