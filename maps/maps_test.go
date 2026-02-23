package main

import "testing"

func TestMaps(t *testing.T) {
	t.Run("fetch value using key", func(t *testing.T) {
		dictonary := Dictionary{"test": "this is just a test"}

		got, _ := dictonary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		dictonary := Dictionary{"test": "this is just a test"}

		_, err := dictonary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new key and value", func(t *testing.T) {
		dictonary := Dictionary{}

		dictonary.Add("test", "this is just a test")
		definition := "this is just a test"

		assertDefinition(t, dictonary, "test", definition)

	})

	t.Run("add existing key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "this is just another test")
		definition := "this is just a test"

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertDefinition(t, dictionary, "test", definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		dictionary.Update("test", "this is just an updated test")
		definition := "this is just an updated test"

		assertDefinition(t, dictionary, "test", definition)
	})

	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "this is just a test")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Delete("test")

		assertError(t, err, nil)

		_, err = dictionary.Search("test")
		assertError(t, err, ErrNotFound)

	})

	t.Run("delete non existing key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Delete("unknown")

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func assertDefinition(t testing.TB, d Dictionary, key, definition string) {
	t.Helper()
	got, err := d.Search(key)

	if err != nil {
		t.Fatal("expected to get no error")
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
