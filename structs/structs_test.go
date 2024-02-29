package main

import "testing"

func TestPerimeter(t *testing.T) {
  want := 40.0
  got := Perimeter(10.0, 10.0)

  if got != want {
    t.Errorf("got %.2f want %.2f", got, want)
  }

}

func TestArea(t *testing.T) {
  want := 100.0
  got := Area(10.0, 10.0)

  if got != want {
    t.Errorf("got %.2f want %.2f", got, want)
  }

}
