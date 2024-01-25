package main

//go:generate sqlboiler --wipe sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jdugan1024/dumax-shopper/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "meals.db")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	tacos, err := models.Meals(qm.Where("name = ?", "Tacos")).One(ctx, db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %s\n", tacos.Name, tacos.MealType)
	ingredients, err := tacos.MealIngredients(qm.Load("Ingredient")).All(ctx, db)
	if err != nil {
		panic(err)
	}
	for _, i := range ingredients {
		parts := strings.Split(i.Quantity.String, " ")
		qty := parts[0]
		units := "pieces"
		if len(parts) > 1 {
			units = parts[1]
		}
		fmt.Printf("- %s, %s, %s\n", i.R.Ingredient.Name, qty, units)
	}

}
