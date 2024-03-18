package main

import "fmt"


type Rect struct {
  height int
  width int
}

type What struct {
  height int
  width int
}

func (x Rect) Area() int {
  return x.height * x.width
}

func (x What) Area() int {
  return x.height * x.width
}

type I interface {
  Area() int
}

func a(x I) int {
  return x.Area()
}

func main() {
  x := Rect{40, 40}
  y := What{40, 60}
  fmt.Println(a(x))
  fmt.Println(a(y))
}
