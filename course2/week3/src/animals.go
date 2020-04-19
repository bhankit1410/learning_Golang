package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Printf("%s \n", a.food)
}
func (a Animal) Move() {
	fmt.Printf("%s \n", a.locomotion)
}
func (a Animal) Speak() {
	fmt.Printf("%s \n", a.noise)
}

func main() {
	var cow Animal
	var bird Animal
	var snake Animal
	cow.food = "grass"
	bird.food = "worms"
	snake.food = "mice"
	cow.locomotion = "walk"
	bird.locomotion = "fly"
	snake.locomotion = "slither"
	cow.noise = "moo"
	bird.noise = "peep"
	snake.noise = "hsss"

	i := 2
	for i > 0 {
		var inputLine string
		fmt.Printf(">")
		// fmt.Scan(&inputLine)
		reader := bufio.NewReader(os.Stdin)
		inputLine, errName := reader.ReadString('\n')
		var animal, feature string
		if errName == nil && len(strings.SplitAfter(inputLine, " ")) > 1{
			animal, feature = strings.ToLower(strings.SplitAfter(inputLine, " ")[0]), strings.ToLower(strings.SplitAfter(inputLine, " ")[1])
		}
		if (strings.TrimSpace(animal) == "cow") && (strings.TrimSpace(feature) == "eat") {
			cow.Eat()
		} else if (strings.TrimSpace(animal) == "cow") && (strings.TrimSpace(feature) == "move") {
			cow.Move()
		} else if (strings.TrimSpace(animal) == "cow") && (strings.TrimSpace(feature) == "speak") {
			cow.Speak()
		} else if (strings.TrimSpace(animal) == "bird") && (strings.TrimSpace(feature) == "eat") {
			bird.Eat()
		} else if (strings.TrimSpace(animal) == "bird") && (strings.TrimSpace(feature) == "move") {
			bird.Move()
		} else if (strings.TrimSpace(animal) == "bird") && (strings.TrimSpace(feature) == "speak") {
			bird.Speak()
		} else if (strings.TrimSpace(animal) == "snake") && (strings.TrimSpace(feature) == "eat") {
			snake.Eat()
		} else if (strings.TrimSpace(animal) == "snake") && (strings.TrimSpace(feature) == "move") {
			snake.Move()
		} else if (strings.TrimSpace(animal) == "snake") && (strings.TrimSpace(feature) == "speak") {
			snake.Speak()
		} else {
			fmt.Printf("format incorrect: correct format is > animal feature \n")
		}

	}
}
