package main

import "math"

type Rectangle struct {
  Width float64
  Height float64
}

type Circle struct {
  Radius float64
}

type Triangle struct {
  Height float64
  Base float64
}

type Shape interface {
  Area() float64
}

func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
  return 0.5 * t.Base * t.Height
}

func Perimeter(rectangle Rectangle) float64 {
  return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
  return (rectangle.Height * rectangle.Width)
}

func main() {}
