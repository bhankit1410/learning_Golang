package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Printf("%s \n", a.food)
}
func (a Cow) Move() {
	fmt.Printf("%s \n", a.locomotion)
}
func (a Cow) Speak() {
	fmt.Printf("%s \n", a.noise)
}
func (a Bird) Eat() {
	fmt.Printf("%s \n", a.food)
}
func (a Bird) Move() {
	fmt.Printf("%s \n", a.locomotion)
}
func (a Bird) Speak() {
	fmt.Printf("%s \n", a.noise)
}
func (a Snake) Eat() {
	fmt.Printf("%s \n", a.food)
}
func (a Snake) Move() {
	fmt.Printf("%s \n", a.locomotion)
}
func (a Snake) Speak() {
	fmt.Printf("%s \n", a.noise)
}

func main() {
	i := 2
	for i > 0 {
		var inputLine string
		fmt.Printf(">")
		// fmt.Scan(&inputLine)
		reader := bufio.NewReader(os.Stdin)
		inputLine, errName := reader.ReadString('\n')
		var requestType, name, animal string
		if errName == nil && len(strings.SplitAfter(inputLine, " ")) > 2 {
			requestType, name, animal = strings.ToLower(strings.SplitAfter(inputLine, " ")[0]), strings.ToLower(strings.SplitAfter(inputLine, " ")[1]), strings.ToLower(strings.SplitAfter(inputLine, " ")[2])
		}
		if strings.TrimSpace(requestType) == "newanimal" {
			if strings.TrimSpace(animal) == "cow" {
				var a1 Animal
				var c1 Cow
				c1.name = name
				a1 = c1
				fmt.Printf("created it")

			} else if strings.TrimSpace(animal) == "bird" {
				var a1 Animal
				var c1 Bird
				c1.name = name
				a1 = c1
				fmt.Printf("created it")

			} else if strings.TrimSpace(animal) == "snake" {
				var a1 Animal
				var c1 Snake
				c1.name = name
				a1 = c1
				fmt.Printf("created it")

			}

		} else if strings.TrimSpace(requestType) == "query" {
			if strings.TrimSpace(animal) == "eat" {
				fmt.Printf(Animal.Eat())
			}

		} else {
			fmt.Printf("format incorrect: correct format is >query/newanimal name animal \n")
		}

	}
}
