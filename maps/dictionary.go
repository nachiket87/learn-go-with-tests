package main

import "errors"

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
  return string(e)
}

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, description string) error {
  _, err := d.Search(word)

  switch err {
  case ErrNotFound:
    d[word] = description
  case nil:
    return ErrWordExists
  default:
    return err
  }
  return nil
}
