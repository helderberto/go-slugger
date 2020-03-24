package slugger

import "testing"

func TestSlugger(t *testing.T) {
	input := "Lorem Ipsum is simply dummy"
	inputSpecial := "Lorem && ** Ipsum is simply dummy"
	want := "lorem-ipsum-is-simply-dummy"

	got := slugger(Params{text: input})
	gotSpecial := slugger(Params{text: inputSpecial})

	t.Run("slug text without special characters", func(t *testing.T) {
		if got != want {
			t.Errorf("got %v; want %v", got, want)
		}
	})

	t.Run("slug text with special characters", func(t *testing.T) {
		if gotSpecial != want {
			t.Errorf("got %v; want %v", gotSpecial, want)
		}
	})
}
