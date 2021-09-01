package main

import (
	"fmt"
	"math"
)

// -----------------
// https://gobyexample.com/interfaces
// -----------------

type GeometryCalculator interface {
	area() float64
	perimeter() float64
}

type Square struct {
	edge float64
}

func (s *Square) area() float64 {
	return s.edge * s.edge
}

func (s *Square) perimeter() float64 {
	return 4 * s.edge
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(gc GeometryCalculator) {
	fmt.Printf("%+v - %T\n", gc, gc)
	fmt.Println("Area     :", gc.area())
	fmt.Println("Perimeter:", gc.perimeter())
	fmt.Println()
}

// func main() {
// 	s := &Square{edge: 2}
// 	r := &Rectangle{width: 3, height: 4}
// 	c := &Circle{radius: 5}

// 	Measure(s)
// 	Measure(r)
// 	Measure(c)
// }

// -----------------
// https://zetcode.com/golang/interface/
// -----------------

type Animal interface {
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Meow!"
}

type Duck struct{}

func (d Duck) Sound() string {
	return "Quack!"
}

// func main() {
// 	animals := []Animal{Dog{}, Cat{}, Duck{}}

// 	for _, a := range animals {
// 		fmt.Println(a.Sound())
// 	}
// }

// -----------------
// https://golangbot.com/interfaces-part-2
// -----------------

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p *Person) Describe() {
	fmt.Printf("%s is %d old years old.\n", p.name, p.age)
}

type Address struct {
	province string
	country  string
}

func (a *Address) Describe() {
	fmt.Printf("Province: %s - Country: %s", a.province, a.country)
}

// func main() {
// 	var d Describer

// 	p := &Person{"Ronaldo", 36}
// 	d = p
// 	d.Describe()

// 	a := &Address{province: "Madeira", country: "Portugal"}
// 	d = a
// 	d.Describe()
// }

// -----------------

const BonusFee = 325.58

type SalaryCalculator interface {
	DisplaySalary()
}

type BonusCalculator interface {
	CalculateBonus() float64
}

type Employee struct {
	name       string
	salary     float64
	bonusCount int
	bonusFee   float64
}

func (e *Employee) DisplaySalary() {
	fmt.Printf("%s has salary %.2f\n", e.name, (e.salary + e.CalculateBonus()))
}

func (e *Employee) CalculateBonus() float64 {
	return float64(e.bonusCount) * e.bonusFee
}

// func main() {
// 	e := &Employee{
// 		name:       "Ashley",
// 		salary:     12403.27,
// 		bonusCount: 2,
// 		bonusFee:   BonusFee,
// 	}

// 	var sc SalaryCalculator = e
// 	sc.DisplaySalary()

// 	var bc BonusCalculator = e
// 	fmt.Println("Bonus:", bc.CalculateBonus())
// }

type EmployeeOperations interface {
	SalaryCalculator
	BonusCalculator
}

func main() {
	e := &Employee{
		name:       "James",
		salary:     9621.48,
		bonusCount: 4,
		bonusFee:   BonusFee,
	}

	var eops EmployeeOperations = e

	eops.DisplaySalary()
	fmt.Println("Bonus:", eops.CalculateBonus())
}
