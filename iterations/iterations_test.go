package iterations

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat a character 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeat characters based on repeatCount", func(t *testing.T) {
		repeated := Repeat("b", 2)
		expected := "bb"
		assertCorrectMessage(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
