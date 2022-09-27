package main

import "fmt"

type Beverage interface {
	getDescription() string
	cost() float32
	getBeverage() Beverage
}

type DarkRoast struct{}

func (dr DarkRoast) cost() float32 {
	return 0.99
}

func (dr DarkRoast) getDescription() string {
	return "DarkRoast: 0.99$"
}

func (dr DarkRoast) getBeverage() Beverage {
	return nil
}

type Espresso struct{}

func (e Espresso) cost() float32 {
	return 0.99
}

func (e Espresso) getDescription() string {
	return "Espresso: 0.99$"
}

func (e Espresso) getBeverage() Beverage {
	return nil
}

type Milk struct {
	beverage Beverage
}

func (m *Milk) getBeverage() Beverage {
	return m.beverage
}

func (m *Milk) getDescription() string {
	return m.getBeverage().getDescription() + "\nMilk: 0.2$"
}

func (m *Milk) cost() float32 {
	return m.getBeverage().cost() + 0.2
}

type Soy struct {
	beverage Beverage
}

func (s *Soy) getBeverage() Beverage {
	return s.beverage
}

func (s *Soy) getDescription() string {
	return s.getBeverage().getDescription() + "\nSoy: 0.2$"
}

func (s *Soy) cost() float32 {
	return s.getBeverage().cost() + 0.2
}

func main() {
	order1 := Milk{Espresso{}}
	fmt.Println(order1.getDescription())
	fmt.Println(order1.cost())

	fmt.Println("---------------------")

	order2 := Soy{DarkRoast{}}
	fmt.Println(order2.getDescription())
	fmt.Println(order2.cost())
}
