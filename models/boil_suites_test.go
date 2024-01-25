// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Ingredients", testIngredients)
	t.Run("Meals", testMeals)
	t.Run("MealIngredients", testMealIngredients)
}

func TestDelete(t *testing.T) {
	t.Run("Ingredients", testIngredientsDelete)
	t.Run("Meals", testMealsDelete)
	t.Run("MealIngredients", testMealIngredientsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Ingredients", testIngredientsQueryDeleteAll)
	t.Run("Meals", testMealsQueryDeleteAll)
	t.Run("MealIngredients", testMealIngredientsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Ingredients", testIngredientsSliceDeleteAll)
	t.Run("Meals", testMealsSliceDeleteAll)
	t.Run("MealIngredients", testMealIngredientsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Ingredients", testIngredientsExists)
	t.Run("Meals", testMealsExists)
	t.Run("MealIngredients", testMealIngredientsExists)
}

func TestFind(t *testing.T) {
	t.Run("Ingredients", testIngredientsFind)
	t.Run("Meals", testMealsFind)
	t.Run("MealIngredients", testMealIngredientsFind)
}

func TestBind(t *testing.T) {
	t.Run("Ingredients", testIngredientsBind)
	t.Run("Meals", testMealsBind)
	t.Run("MealIngredients", testMealIngredientsBind)
}

func TestOne(t *testing.T) {
	t.Run("Ingredients", testIngredientsOne)
	t.Run("Meals", testMealsOne)
	t.Run("MealIngredients", testMealIngredientsOne)
}

func TestAll(t *testing.T) {
	t.Run("Ingredients", testIngredientsAll)
	t.Run("Meals", testMealsAll)
	t.Run("MealIngredients", testMealIngredientsAll)
}

func TestCount(t *testing.T) {
	t.Run("Ingredients", testIngredientsCount)
	t.Run("Meals", testMealsCount)
	t.Run("MealIngredients", testMealIngredientsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Ingredients", testIngredientsHooks)
	t.Run("Meals", testMealsHooks)
	t.Run("MealIngredients", testMealIngredientsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Ingredients", testIngredientsInsert)
	t.Run("Ingredients", testIngredientsInsertWhitelist)
	t.Run("Meals", testMealsInsert)
	t.Run("Meals", testMealsInsertWhitelist)
	t.Run("MealIngredients", testMealIngredientsInsert)
	t.Run("MealIngredients", testMealIngredientsInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Ingredients", testIngredientsReload)
	t.Run("Meals", testMealsReload)
	t.Run("MealIngredients", testMealIngredientsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Ingredients", testIngredientsReloadAll)
	t.Run("Meals", testMealsReloadAll)
	t.Run("MealIngredients", testMealIngredientsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Ingredients", testIngredientsSelect)
	t.Run("Meals", testMealsSelect)
	t.Run("MealIngredients", testMealIngredientsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Ingredients", testIngredientsUpdate)
	t.Run("Meals", testMealsUpdate)
	t.Run("MealIngredients", testMealIngredientsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Ingredients", testIngredientsSliceUpdateAll)
	t.Run("Meals", testMealsSliceUpdateAll)
	t.Run("MealIngredients", testMealIngredientsSliceUpdateAll)
}
