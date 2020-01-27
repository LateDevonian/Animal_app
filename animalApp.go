package main

import (
	"fmt"
)

type Animal interface {
	//lists methods that will perform when defined
	Species() string
	Called() string
	Speak() string
	Eat() string
	Interact() string
}

//create a struct and define the methods
type Mammal struct {
	species string
	called  string
	noise   string
	food    string
	action  string
}

//define what the methods take and return
func (m Mammal) Species() string {
	return m.species
}
func (m Mammal) Speak() string {
	return m.noise
}
func (m Mammal) Eat() string {
	return m.food
}
func (m Mammal) Interact() string {
	return m.action
}
func (m Mammal) Called() string {
	return m.called
}

func describeAnimal() {
	//
}

func main() {
	//
	fido := Mammal{"dog", "Fido", "woof", "chum", "runaround"}
	fluffy := Mammal{"cat", "Fluffy", "meow", "fish", "sleep"}
	willy := Mammal{"whale", "Willy", "squeak", "krill", "swim"}
	//print individually
	fmt.Println(fido.Species(), "makes this noise", fido.Speak())
	fmt.Println("and prefers to eat", fido.Eat())
	fmt.Println(fluffy.Called(), "is, however a ", fluffy.Species())
	//send through an interface as a slice
	//extract the names of the animals
	animals := []Animal{fido, fluffy, willy}
	fmt.Println("Animal names are:")

	for _, animal := range animals {
		fmt.Println(animal.Called())
	}
}
