package main

import (
	"testing"
)


func assertStrings(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}


func assertError(t testing.TB, got, want error) {
  t.Helper()
  if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
  t.Helper()
  got, err := dictionary.Search(word)
  if err != nil {
    t.Fatal("should find added word:", err)
  }
  assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
  dictionary := Dictionary{"test": "this is a test"}
  t.Run("known word", func(t *testing.T) {
    want := "this is a test"
    got, _ := dictionary.Search("test")
    assertStrings(t, want, got)

  })

  t.Run("unknown word", func(t *testing.T) {
    _, err :=  dictionary.Search("unknown")
    assertError(t, err, ErrNotFound)
  })
}

func TestAdd(t *testing.T) {
  dictionary := Dictionary{}
  t.Run("new word", func(t *testing.T) {
    word := "test"
    definition := "this is a test"
    dictionary.Add(word, definition)
    assertDefinition(t, dictionary, word, definition)
  })

  t.Run("existing word", func(t *testing.T) {
    word := "test"
    definition := "this is a test"
    dictionary = Dictionary{word: definition}
    err := dictionary.Add(word, "new test")

    assertError(t, err, ErrWordExists)
    assertDefinition(t, dictionary, word, definition)
  })
}
