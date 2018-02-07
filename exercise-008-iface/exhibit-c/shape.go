package main

import (
	"fmt"
	"math"
)

////////////
// Square //
////////////

type Square struct {
	side float64
}

func (s *Square) Name() string {
	return "Square"
}

func (s *Square) Perimeter() float64 {
	return 4 * s.side
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

////////////
// Circle //
////////////

type Circle struct {
	radius float64
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//////////////
// Triangle //
//////////////

type RightTriangle struct{
  Leg1 float64
  Leg2 float64 
}

func (t *RightTriangle) Name() string{
  return "Right Triangle"
}

func (t *RightTriangle) Perimeter() float64{
  return t.Leg1 + t.Leg2 + math.Sqrt(t.Leg1 * t.Leg1 + t.Leg2 * t.Leg2)
}

func (t *RightTriangle) Area() float64{
  return (t.Leg1 * t.Leg2) / 2.0  
}

////////////////
// Efficiency //
////////////////

type Shape interface {
	Name() string
	Perimeter() float64
	Area() float64
}

func Efficiency(s Shape) {
	name := s.Name()
	area := s.Area()
	rope := s.Perimeter()

	efficiency := 100.0 * area / (rope * rope)
	fmt.Printf("Efficiency of a %s is %f\n", name, efficiency)
}

func main() {

  t := RightTriangle{
    Leg1: 10.0,
    Leg2: 10.0,
  } 
  Efficiency(&t) 
  
	s := Square{side: 10.0}
	Efficiency(&s)

	c := Circle{radius: 10.0}
	Efficiency(&c)
}
