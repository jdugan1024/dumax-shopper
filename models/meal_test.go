// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMeals(t *testing.T) {
	t.Parallel()

	query := Meals()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMealsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMealsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Meals().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMealsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MealSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMealsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MealExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Meal exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MealExists to return true, but got false.")
	}
}

func testMealsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mealFound, err := FindMeal(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mealFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMealsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Meals().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMealsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Meals().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMealsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mealOne := &Meal{}
	mealTwo := &Meal{}
	if err = randomize.Struct(seed, mealOne, mealDBTypes, false, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}
	if err = randomize.Struct(seed, mealTwo, mealDBTypes, false, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mealOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mealTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Meals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMealsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mealOne := &Meal{}
	mealTwo := &Meal{}
	if err = randomize.Struct(seed, mealOne, mealDBTypes, false, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}
	if err = randomize.Struct(seed, mealTwo, mealDBTypes, false, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mealOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mealTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mealBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func mealAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Meal) error {
	*o = Meal{}
	return nil
}

func testMealsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Meal{}
	o := &Meal{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mealDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Meal object: %s", err)
	}

	AddMealHook(boil.BeforeInsertHook, mealBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mealBeforeInsertHooks = []MealHook{}

	AddMealHook(boil.AfterInsertHook, mealAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mealAfterInsertHooks = []MealHook{}

	AddMealHook(boil.AfterSelectHook, mealAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mealAfterSelectHooks = []MealHook{}

	AddMealHook(boil.BeforeUpdateHook, mealBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mealBeforeUpdateHooks = []MealHook{}

	AddMealHook(boil.AfterUpdateHook, mealAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mealAfterUpdateHooks = []MealHook{}

	AddMealHook(boil.BeforeDeleteHook, mealBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mealBeforeDeleteHooks = []MealHook{}

	AddMealHook(boil.AfterDeleteHook, mealAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mealAfterDeleteHooks = []MealHook{}

	AddMealHook(boil.BeforeUpsertHook, mealBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mealBeforeUpsertHooks = []MealHook{}

	AddMealHook(boil.AfterUpsertHook, mealAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mealAfterUpsertHooks = []MealHook{}
}

func testMealsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMealsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mealColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMealToManyMealIngredients(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Meal
	var b, c MealIngredient

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, mealIngredientDBTypes, false, mealIngredientColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mealIngredientDBTypes, false, mealIngredientColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.MealID = a.ID
	c.MealID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.MealIngredients().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.MealID == b.MealID {
			bFound = true
		}
		if v.MealID == c.MealID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MealSlice{&a}
	if err = a.L.LoadMealIngredients(ctx, tx, false, (*[]*Meal)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MealIngredients); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MealIngredients = nil
	if err = a.L.LoadMealIngredients(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MealIngredients); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMealToManyAddOpMealIngredients(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Meal
	var b, c, d, e MealIngredient

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mealDBTypes, false, strmangle.SetComplement(mealPrimaryKeyColumns, mealColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MealIngredient{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, mealIngredientDBTypes, false, strmangle.SetComplement(mealIngredientPrimaryKeyColumns, mealIngredientColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*MealIngredient{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMealIngredients(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.MealID {
			t.Error("foreign key was wrong value", a.ID, first.MealID)
		}
		if a.ID != second.MealID {
			t.Error("foreign key was wrong value", a.ID, second.MealID)
		}

		if first.R.Meal != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Meal != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MealIngredients[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MealIngredients[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MealIngredients().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMealsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMealsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MealSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMealsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Meals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mealDBTypes = map[string]string{`ID`: `INT`, `Name`: `TEXT`, `MealType`: `TEXT`}
	_           = bytes.MinRead
)

func testMealsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mealPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mealAllColumns) == len(mealPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mealDBTypes, true, mealPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMealsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mealAllColumns) == len(mealPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Meal{}
	if err = randomize.Struct(seed, o, mealDBTypes, true, mealColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mealDBTypes, true, mealPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mealAllColumns, mealPrimaryKeyColumns) {
		fields = mealAllColumns
	} else {
		fields = strmangle.SetComplement(
			mealAllColumns,
			mealPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MealSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMealsUpsert(t *testing.T) {
	t.Parallel()
	if len(mealAllColumns) == len(mealPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Meal{}
	if err = randomize.Struct(seed, &o, mealDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Meal: %s", err)
	}

	count, err := Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mealDBTypes, false, mealPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Meal struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Meal: %s", err)
	}

	count, err = Meals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}