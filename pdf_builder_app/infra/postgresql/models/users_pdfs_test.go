// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testUsersPDFS(t *testing.T) {
	t.Parallel()

	query := UsersPDFS()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUsersPDFSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
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

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersPDFSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UsersPDFS().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersPDFSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UsersPDFSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUsersPDFSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UsersPDFExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if UsersPDF exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UsersPDFExists to return true, but got false.")
	}
}

func testUsersPDFSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	usersPDFFound, err := FindUsersPDF(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if usersPDFFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUsersPDFSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UsersPDFS().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUsersPDFSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UsersPDFS().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUsersPDFSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	usersPDFOne := &UsersPDF{}
	usersPDFTwo := &UsersPDF{}
	if err = randomize.Struct(seed, usersPDFOne, usersPDFDBTypes, false, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}
	if err = randomize.Struct(seed, usersPDFTwo, usersPDFDBTypes, false, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = usersPDFOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = usersPDFTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UsersPDFS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUsersPDFSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	usersPDFOne := &UsersPDF{}
	usersPDFTwo := &UsersPDF{}
	if err = randomize.Struct(seed, usersPDFOne, usersPDFDBTypes, false, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}
	if err = randomize.Struct(seed, usersPDFTwo, usersPDFDBTypes, false, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = usersPDFOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = usersPDFTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func usersPDFBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func usersPDFAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UsersPDF) error {
	*o = UsersPDF{}
	return nil
}

func testUsersPDFSHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UsersPDF{}
	o := &UsersPDF{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, usersPDFDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UsersPDF object: %s", err)
	}

	AddUsersPDFHook(boil.BeforeInsertHook, usersPDFBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	usersPDFBeforeInsertHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.AfterInsertHook, usersPDFAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	usersPDFAfterInsertHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.AfterSelectHook, usersPDFAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	usersPDFAfterSelectHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.BeforeUpdateHook, usersPDFBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	usersPDFBeforeUpdateHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.AfterUpdateHook, usersPDFAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	usersPDFAfterUpdateHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.BeforeDeleteHook, usersPDFBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	usersPDFBeforeDeleteHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.AfterDeleteHook, usersPDFAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	usersPDFAfterDeleteHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.BeforeUpsertHook, usersPDFBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	usersPDFBeforeUpsertHooks = []UsersPDFHook{}

	AddUsersPDFHook(boil.AfterUpsertHook, usersPDFAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	usersPDFAfterUpsertHooks = []UsersPDFHook{}
}

func testUsersPDFSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsersPDFSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(usersPDFColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUsersPDFToManyUnifiedPDFPartialPDFS(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UsersPDF
	var b, c PartialPDF

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, partialPDFDBTypes, false, partialPDFColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, partialPDFDBTypes, false, partialPDFColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UnifiedPDFID = a.ID
	c.UnifiedPDFID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UnifiedPDFPartialPDFS().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UnifiedPDFID == b.UnifiedPDFID {
			bFound = true
		}
		if v.UnifiedPDFID == c.UnifiedPDFID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := UsersPDFSlice{&a}
	if err = a.L.LoadUnifiedPDFPartialPDFS(ctx, tx, false, (*[]*UsersPDF)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UnifiedPDFPartialPDFS); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UnifiedPDFPartialPDFS = nil
	if err = a.L.LoadUnifiedPDFPartialPDFS(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UnifiedPDFPartialPDFS); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testUsersPDFToManyAddOpUnifiedPDFPartialPDFS(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UsersPDF
	var b, c, d, e PartialPDF

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, usersPDFDBTypes, false, strmangle.SetComplement(usersPDFPrimaryKeyColumns, usersPDFColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*PartialPDF{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, partialPDFDBTypes, false, strmangle.SetComplement(partialPDFPrimaryKeyColumns, partialPDFColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*PartialPDF{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUnifiedPDFPartialPDFS(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UnifiedPDFID {
			t.Error("foreign key was wrong value", a.ID, first.UnifiedPDFID)
		}
		if a.ID != second.UnifiedPDFID {
			t.Error("foreign key was wrong value", a.ID, second.UnifiedPDFID)
		}

		if first.R.UnifiedPDF != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.UnifiedPDF != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UnifiedPDFPartialPDFS[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UnifiedPDFPartialPDFS[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UnifiedPDFPartialPDFS().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testUsersPDFToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UsersPDF
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, usersPDFDBTypes, false, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := UsersPDFSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UsersPDF)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testUsersPDFToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UsersPDF
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, usersPDFDBTypes, false, strmangle.SetComplement(usersPDFPrimaryKeyColumns, usersPDFColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UsersPDFS[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testUsersPDFSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
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

func testUsersPDFSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UsersPDFSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUsersPDFSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UsersPDFS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	usersPDFDBTypes = map[string]string{`ID`: `uuid`, `UserID`: `uuid`, `S3URL`: `text`, `UnifiedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testUsersPDFSUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(usersPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(usersPDFAllColumns) == len(usersPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUsersPDFSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(usersPDFAllColumns) == len(usersPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UsersPDF{}
	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, usersPDFDBTypes, true, usersPDFPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(usersPDFAllColumns, usersPDFPrimaryKeyColumns) {
		fields = usersPDFAllColumns
	} else {
		fields = strmangle.SetComplement(
			usersPDFAllColumns,
			usersPDFPrimaryKeyColumns,
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

	slice := UsersPDFSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUsersPDFSUpsert(t *testing.T) {
	t.Parallel()

	if len(usersPDFAllColumns) == len(usersPDFPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UsersPDF{}
	if err = randomize.Struct(seed, &o, usersPDFDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UsersPDF: %s", err)
	}

	count, err := UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, usersPDFDBTypes, false, usersPDFPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UsersPDF struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UsersPDF: %s", err)
	}

	count, err = UsersPDFS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
