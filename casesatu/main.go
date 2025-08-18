package main

import (
	"encoding/json"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

type FruitType string
type Fruits struct {
	ID    int       `json:"fruitId"`
	Name  string    `json:"fruitName"`
	Type  FruitType `json:"fruitType,omitempty"`
	Stock int       `json:"stock"`
}

func andisFruits(f []Fruits) []string {
	result := make([]string, 0)

	for _, fruit := range f {
		name := strings.ToLower(fruit.Name)
		if !slices.Contains(result, name) {
			result = append(result, name)
		}
	}

	return result
}

func fruitBasket(f []Fruits) map[FruitType][]string {
	result := make(map[FruitType][]string)

	for _, fruit := range f {
		result[fruit.Type] = append(result[fruit.Type], fruit.Name)
	}

	return result
}

func countFruitPerBasket(f []Fruits) map[FruitType]int {
	result := make(map[FruitType]int)

	for _, fruit := range f {
		result[fruit.Type] += fruit.Stock
	}

	return result
}

func main() {
	b, err := os.ReadFile("fruits.json")
	if err != nil {
		panic(err)
	}

	var fruits []Fruits
	if err := json.Unmarshal(b, &fruits); err != nil {
		panic(err)
	}

	println("Andi's fruits are:")
	andis := andisFruits(fruits)
	for _, fruit := range andis {
		println(fruit)
	}

	println("------------------------------")
	println("Fruit Basket:")
	basket := fruitBasket(fruits)
	var keys []FruitType
	for k := range basket {
		keys = append(keys, k)
	}

	slices.Sort(keys)
	dupes := make([]string, 0)
	for _, k := range keys {
		println(string(k) + ":")
		for _, fruit := range basket[k] {
			fruit = strings.ToLower(fruit)
			if !slices.Contains(dupes, fruit) {
				dupes = append(dupes, fruit)
				println("  - " + fruit)
			}
		}
	}
	println("------------------------------")

	countBasket := countFruitPerBasket(fruits)
	var countKeys []FruitType
	for k := range countBasket {
		countKeys = append(countKeys, k)
	}

	slices.Sort(countKeys)
	println("Count of fruits per basket:")
	for _, k := range countKeys {
		println(string(k) + ": " + strconv.Itoa(countBasket[k]))
	}
	println("------------------------------")
}
