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
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want: %s got: %s", want, got)
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("want: %s got: %s", want, got)
	}

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "dance"
		definition := "is happy moving"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "dance"
		definition := "is happy moving"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "alala")

		assertError(t, err, ErrWordAlreadyExist)
		assertDefinition(t, dictionary, word, definition)

	})
}

func assertDefinition(t *testing.T, dict Dictionary, word, def string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatalf("error of getting word: %v", err)
	}
	assertStrings(t, got, def)
}

func TestUpdate(t *testing.T) {
	t.Run("word exist", func(t *testing.T) {
		word := "dance"
		definition := "is happy moving"
		dictionary := Dictionary{word: definition}

		newDefinition := "a la la ho"

		err := dictionary.Update(word, newDefinition)
		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("word not exist", func(t *testing.T) {
		word := "dance"
		dictionary := Dictionary{}
		newDefinition := "a la la ho"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordNotExist)

	})
}

func TestDelete(t *testing.T) {
	word := "dance"
	definition := "is happy moving"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)

}
