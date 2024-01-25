package main

// mostly written by ChatGPT

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Meal struct {
	Name        string
	Ingredients []Ingredient
}

type Ingredient struct {
	Name     string
	Quantity float64
	Measure  string
}

func Parse(data []string) []Meal {
	var meals []Meal
	var meal Meal

	for _, line := range data {
		fmt.Println("> ", line)
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Check if it is a meal line or an ingredient line
		if !strings.HasPrefix(line, "-") {
			// If it is a meal, save the previous one and start a new one
			if meal.Name != "" {
				meals = append(meals, meal)
			}
			meal = Meal{Name: strings.Split(line, ",")[0]}
		} else {
			// If it is an ingredient, add it to the current meal
			line = strings.TrimPrefix(line, "- ")
			parts := strings.Split(line, ",")
			qty, _ := strconv.ParseFloat(strings.Trim(parts[1], " "), 64)
			name := strings.Trim(parts[0], " ")
			measure := strings.Trim(parts[2], " ")
			ingredient := Ingredient{Name: name, Quantity: qty, Measure: measure}
			meal.Ingredients = append(meal.Ingredients, ingredient)
		}
	}

	// Add the last meal
	meals = append(meals, meal)

	return meals
}

func main() {
	// The data source could be a text file or any other sources
	file, err := os.Open("meals.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	meals := Parse(data)
	fmt.Printf("Got %d meals\n", len(meals))

	// Prints the parsed meals and ingredients
	for _, meal := range meals {
		fmt.Println(meal.Name)
		for _, ingredient := range meal.Ingredients {
			fmt.Printf("- %s, %0.1f, %s \n", ingredient.Name, ingredient.Quantity, ingredient.Measure)
		}
	}
}
