package main

import "fmt"

type Point struct {
	x, y float64
}

// Value reveiver always pass the COPY of original value
// Print() method could be called on value or pointer type.
// in second case compiler will dereference the the pointer, make the copy and
// pass in method COPY of origianl value.
func (p Point) Print(name string) {
	fmt.Printf("Point: %s(%v, %v)\n", name, p.x, p.y)
}

// Method can receve only reference to Point
func (p *Point) Move(d Point) {
	p.x += d.x
	p.y += d.y
}

func (p Point) Add(d Point) Point {
	p.x += d.x
	p.y += d.y
	return p
}

func main() {
	a := Point{3, 4}
	a.Print("A")

	fmt.Println("Move by (1, 1)")
	a.Move(Point{1, 1})
	a.Print("A")

	fmt.Println("Add to original point(1, 1)")
	c := a.Add(Point{1, 1})
	a.Print("A")
	c.Print("C")

	fmt.Println("Add to referenced data (1, 1)")
	b := &a
	c = b.Add(Point{1, 1})
	a.Print("A")
	b.Print("B")
	c.Print("C")

}
