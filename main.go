package main

import "fmt"

type Bird interface {
	Fly()
}

type Human interface {
	GoToWork()
}

type Combined interface {
	Bird
	Human
}

type SuperHero struct{}

func (s *SuperHero) Fly() {
	fmt.Println("Superhero can fly")
}

func (s *SuperHero) GoToWork() {
	fmt.Println("Superhero can go to work")
}

func (s *SuperHero) FightCrime() {
	fmt.Println("Superhero can fight crime")
}

type Superman struct {
	SuperHero
}

func (s *Superman) Fly() {
	fmt.Println("Superman flies and catches Lex Luthor")
}

func Summariser(c Combined) {
	c.Fly()
	c.GoToWork()
}

func main() {
	var superHero *SuperHero = &SuperHero{}
	Summariser(superHero)

	fmt.Println("--------------------")

	var superman *Superman = &Superman{}
	Summariser(superman)
}
