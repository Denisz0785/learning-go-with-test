package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "checking how program work is",
	}
	t.Run("known word", func(t *testing.T) {
		want := "checking how program work is"

		got, _ := dictionary.Search("test")
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		want := "unknown word"
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatalf("expected error but got %v", err)
		}
		assertStrings(t, err.Error(), want)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want: %s got: %s", want, got)
	}

}
