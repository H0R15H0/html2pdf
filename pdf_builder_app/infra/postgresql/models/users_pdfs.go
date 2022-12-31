// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UsersPDF is an object representing the database table.
type UsersPDF struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    null.String `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	S3URL     string      `boil:"s3_url" json:"s3_url" toml:"s3_url" yaml:"s3_url"`
	UnifiedAt null.Time   `boil:"unified_at" json:"unified_at,omitempty" toml:"unified_at" yaml:"unified_at,omitempty"`

	R *usersPDFR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L usersPDFL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UsersPDFColumns = struct {
	ID        string
	UserID    string
	S3URL     string
	UnifiedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	S3URL:     "s3_url",
	UnifiedAt: "unified_at",
}

var UsersPDFTableColumns = struct {
	ID        string
	UserID    string
	S3URL     string
	UnifiedAt string
}{
	ID:        "users_pdfs.id",
	UserID:    "users_pdfs.user_id",
	S3URL:     "users_pdfs.s3_url",
	UnifiedAt: "users_pdfs.unified_at",
}

// Generated where

var UsersPDFWhere = struct {
	ID        whereHelperstring
	UserID    whereHelpernull_String
	S3URL     whereHelperstring
	UnifiedAt whereHelpernull_Time
}{
	ID:        whereHelperstring{field: "\"users_pdfs\".\"id\""},
	UserID:    whereHelpernull_String{field: "\"users_pdfs\".\"user_id\""},
	S3URL:     whereHelperstring{field: "\"users_pdfs\".\"s3_url\""},
	UnifiedAt: whereHelpernull_Time{field: "\"users_pdfs\".\"unified_at\""},
}

// UsersPDFRels is where relationship names are stored.
var UsersPDFRels = struct {
	User                  string
	UnifiedPDFPartialPDFS string
}{
	User:                  "User",
	UnifiedPDFPartialPDFS: "UnifiedPDFPartialPDFS",
}

// usersPDFR is where relationships are stored.
type usersPDFR struct {
	User                  *User           `boil:"User" json:"User" toml:"User" yaml:"User"`
	UnifiedPDFPartialPDFS PartialPDFSlice `boil:"UnifiedPDFPartialPDFS" json:"UnifiedPDFPartialPDFS" toml:"UnifiedPDFPartialPDFS" yaml:"UnifiedPDFPartialPDFS"`
}

// NewStruct creates a new relationship struct
func (*usersPDFR) NewStruct() *usersPDFR {
	return &usersPDFR{}
}

func (r *usersPDFR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *usersPDFR) GetUnifiedPDFPartialPDFS() PartialPDFSlice {
	if r == nil {
		return nil
	}
	return r.UnifiedPDFPartialPDFS
}

// usersPDFL is where Load methods for each relationship are stored.
type usersPDFL struct{}

var (
	usersPDFAllColumns            = []string{"id", "user_id", "s3_url", "unified_at"}
	usersPDFColumnsWithoutDefault = []string{"id", "s3_url"}
	usersPDFColumnsWithDefault    = []string{"user_id", "unified_at"}
	usersPDFPrimaryKeyColumns     = []string{"id"}
	usersPDFGeneratedColumns      = []string{}
)

type (
	// UsersPDFSlice is an alias for a slice of pointers to UsersPDF.
	// This should almost always be used instead of []UsersPDF.
	UsersPDFSlice []*UsersPDF
	// UsersPDFHook is the signature for custom UsersPDF hook methods
	UsersPDFHook func(context.Context, boil.ContextExecutor, *UsersPDF) error

	usersPDFQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	usersPDFType                 = reflect.TypeOf(&UsersPDF{})
	usersPDFMapping              = queries.MakeStructMapping(usersPDFType)
	usersPDFPrimaryKeyMapping, _ = queries.BindMapping(usersPDFType, usersPDFMapping, usersPDFPrimaryKeyColumns)
	usersPDFInsertCacheMut       sync.RWMutex
	usersPDFInsertCache          = make(map[string]insertCache)
	usersPDFUpdateCacheMut       sync.RWMutex
	usersPDFUpdateCache          = make(map[string]updateCache)
	usersPDFUpsertCacheMut       sync.RWMutex
	usersPDFUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var usersPDFAfterSelectHooks []UsersPDFHook

var usersPDFBeforeInsertHooks []UsersPDFHook
var usersPDFAfterInsertHooks []UsersPDFHook

var usersPDFBeforeUpdateHooks []UsersPDFHook
var usersPDFAfterUpdateHooks []UsersPDFHook

var usersPDFBeforeDeleteHooks []UsersPDFHook
var usersPDFAfterDeleteHooks []UsersPDFHook

var usersPDFBeforeUpsertHooks []UsersPDFHook
var usersPDFAfterUpsertHooks []UsersPDFHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UsersPDF) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UsersPDF) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UsersPDF) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UsersPDF) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UsersPDF) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UsersPDF) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UsersPDF) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UsersPDF) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UsersPDF) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range usersPDFAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUsersPDFHook registers your hook function for all future operations.
func AddUsersPDFHook(hookPoint boil.HookPoint, usersPDFHook UsersPDFHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		usersPDFAfterSelectHooks = append(usersPDFAfterSelectHooks, usersPDFHook)
	case boil.BeforeInsertHook:
		usersPDFBeforeInsertHooks = append(usersPDFBeforeInsertHooks, usersPDFHook)
	case boil.AfterInsertHook:
		usersPDFAfterInsertHooks = append(usersPDFAfterInsertHooks, usersPDFHook)
	case boil.BeforeUpdateHook:
		usersPDFBeforeUpdateHooks = append(usersPDFBeforeUpdateHooks, usersPDFHook)
	case boil.AfterUpdateHook:
		usersPDFAfterUpdateHooks = append(usersPDFAfterUpdateHooks, usersPDFHook)
	case boil.BeforeDeleteHook:
		usersPDFBeforeDeleteHooks = append(usersPDFBeforeDeleteHooks, usersPDFHook)
	case boil.AfterDeleteHook:
		usersPDFAfterDeleteHooks = append(usersPDFAfterDeleteHooks, usersPDFHook)
	case boil.BeforeUpsertHook:
		usersPDFBeforeUpsertHooks = append(usersPDFBeforeUpsertHooks, usersPDFHook)
	case boil.AfterUpsertHook:
		usersPDFAfterUpsertHooks = append(usersPDFAfterUpsertHooks, usersPDFHook)
	}
}

// One returns a single usersPDF record from the query.
func (q usersPDFQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UsersPDF, error) {
	o := &UsersPDF{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for users_pdfs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UsersPDF records from the query.
func (q usersPDFQuery) All(ctx context.Context, exec boil.ContextExecutor) (UsersPDFSlice, error) {
	var o []*UsersPDF

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UsersPDF slice")
	}

	if len(usersPDFAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UsersPDF records in the query.
func (q usersPDFQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count users_pdfs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q usersPDFQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if users_pdfs exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UsersPDF) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// UnifiedPDFPartialPDFS retrieves all the partial_pdf's PartialPDFS with an executor via unified_pdf_id column.
func (o *UsersPDF) UnifiedPDFPartialPDFS(mods ...qm.QueryMod) partialPDFQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"partial_pdfs\".\"unified_pdf_id\"=?", o.ID),
	)

	return PartialPDFS(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (usersPDFL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUsersPDF interface{}, mods queries.Applicator) error {
	var slice []*UsersPDF
	var object *UsersPDF

	if singular {
		var ok bool
		object, ok = maybeUsersPDF.(*UsersPDF)
		if !ok {
			object = new(UsersPDF)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUsersPDF)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUsersPDF))
			}
		}
	} else {
		s, ok := maybeUsersPDF.(*[]*UsersPDF)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUsersPDF)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUsersPDF))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &usersPDFR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &usersPDFR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UsersPDFS = append(foreign.R.UsersPDFS, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UsersPDFS = append(foreign.R.UsersPDFS, local)
				break
			}
		}
	}

	return nil
}

// LoadUnifiedPDFPartialPDFS allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (usersPDFL) LoadUnifiedPDFPartialPDFS(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUsersPDF interface{}, mods queries.Applicator) error {
	var slice []*UsersPDF
	var object *UsersPDF

	if singular {
		var ok bool
		object, ok = maybeUsersPDF.(*UsersPDF)
		if !ok {
			object = new(UsersPDF)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUsersPDF)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUsersPDF))
			}
		}
	} else {
		s, ok := maybeUsersPDF.(*[]*UsersPDF)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUsersPDF)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUsersPDF))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &usersPDFR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &usersPDFR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`partial_pdfs`),
		qm.WhereIn(`partial_pdfs.unified_pdf_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load partial_pdfs")
	}

	var resultSlice []*PartialPDF
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice partial_pdfs")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on partial_pdfs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for partial_pdfs")
	}

	if len(partialPDFAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UnifiedPDFPartialPDFS = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &partialPDFR{}
			}
			foreign.R.UnifiedPDF = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.UnifiedPDFID) {
				local.R.UnifiedPDFPartialPDFS = append(local.R.UnifiedPDFPartialPDFS, foreign)
				if foreign.R == nil {
					foreign.R = &partialPDFR{}
				}
				foreign.R.UnifiedPDF = local
				break
			}
		}
	}

	return nil
}

// SetUser of the usersPDF to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UsersPDFS.
func (o *UsersPDF) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"users_pdfs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, usersPDFPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &usersPDFR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UsersPDFS: UsersPDFSlice{o},
		}
	} else {
		related.R.UsersPDFS = append(related.R.UsersPDFS, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct.
func (o *UsersPDF) RemoveUser(ctx context.Context, exec boil.ContextExecutor, related *User) error {
	var err error

	queries.SetScanner(&o.UserID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("user_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.User = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.UsersPDFS {
		if queries.Equal(o.UserID, ri.UserID) {
			continue
		}

		ln := len(related.R.UsersPDFS)
		if ln > 1 && i < ln-1 {
			related.R.UsersPDFS[i] = related.R.UsersPDFS[ln-1]
		}
		related.R.UsersPDFS = related.R.UsersPDFS[:ln-1]
		break
	}
	return nil
}

// AddUnifiedPDFPartialPDFS adds the given related objects to the existing relationships
// of the users_pdf, optionally inserting them as new records.
// Appends related to o.R.UnifiedPDFPartialPDFS.
// Sets related.R.UnifiedPDF appropriately.
func (o *UsersPDF) AddUnifiedPDFPartialPDFS(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PartialPDF) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.UnifiedPDFID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"partial_pdfs\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"unified_pdf_id"}),
				strmangle.WhereClause("\"", "\"", 2, partialPDFPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.UnifiedPDFID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &usersPDFR{
			UnifiedPDFPartialPDFS: related,
		}
	} else {
		o.R.UnifiedPDFPartialPDFS = append(o.R.UnifiedPDFPartialPDFS, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &partialPDFR{
				UnifiedPDF: o,
			}
		} else {
			rel.R.UnifiedPDF = o
		}
	}
	return nil
}

// SetUnifiedPDFPartialPDFS removes all previously related items of the
// users_pdf replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.UnifiedPDF's UnifiedPDFPartialPDFS accordingly.
// Replaces o.R.UnifiedPDFPartialPDFS with related.
// Sets related.R.UnifiedPDF's UnifiedPDFPartialPDFS accordingly.
func (o *UsersPDF) SetUnifiedPDFPartialPDFS(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PartialPDF) error {
	query := "update \"partial_pdfs\" set \"unified_pdf_id\" = null where \"unified_pdf_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.UnifiedPDFPartialPDFS {
			queries.SetScanner(&rel.UnifiedPDFID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.UnifiedPDF = nil
		}
		o.R.UnifiedPDFPartialPDFS = nil
	}

	return o.AddUnifiedPDFPartialPDFS(ctx, exec, insert, related...)
}

// RemoveUnifiedPDFPartialPDFS relationships from objects passed in.
// Removes related items from R.UnifiedPDFPartialPDFS (uses pointer comparison, removal does not keep order)
// Sets related.R.UnifiedPDF.
func (o *UsersPDF) RemoveUnifiedPDFPartialPDFS(ctx context.Context, exec boil.ContextExecutor, related ...*PartialPDF) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.UnifiedPDFID, nil)
		if rel.R != nil {
			rel.R.UnifiedPDF = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("unified_pdf_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.UnifiedPDFPartialPDFS {
			if rel != ri {
				continue
			}

			ln := len(o.R.UnifiedPDFPartialPDFS)
			if ln > 1 && i < ln-1 {
				o.R.UnifiedPDFPartialPDFS[i] = o.R.UnifiedPDFPartialPDFS[ln-1]
			}
			o.R.UnifiedPDFPartialPDFS = o.R.UnifiedPDFPartialPDFS[:ln-1]
			break
		}
	}

	return nil
}

// UsersPDFS retrieves all the records using an executor.
func UsersPDFS(mods ...qm.QueryMod) usersPDFQuery {
	mods = append(mods, qm.From("\"users_pdfs\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"users_pdfs\".*"})
	}

	return usersPDFQuery{q}
}

// FindUsersPDF retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUsersPDF(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*UsersPDF, error) {
	usersPDFObj := &UsersPDF{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"users_pdfs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, usersPDFObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from users_pdfs")
	}

	if err = usersPDFObj.doAfterSelectHooks(ctx, exec); err != nil {
		return usersPDFObj, err
	}

	return usersPDFObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UsersPDF) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no users_pdfs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(usersPDFColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	usersPDFInsertCacheMut.RLock()
	cache, cached := usersPDFInsertCache[key]
	usersPDFInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			usersPDFAllColumns,
			usersPDFColumnsWithDefault,
			usersPDFColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(usersPDFType, usersPDFMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(usersPDFType, usersPDFMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"users_pdfs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"users_pdfs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into users_pdfs")
	}

	if !cached {
		usersPDFInsertCacheMut.Lock()
		usersPDFInsertCache[key] = cache
		usersPDFInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UsersPDF.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UsersPDF) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	usersPDFUpdateCacheMut.RLock()
	cache, cached := usersPDFUpdateCache[key]
	usersPDFUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			usersPDFAllColumns,
			usersPDFPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update users_pdfs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"users_pdfs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, usersPDFPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(usersPDFType, usersPDFMapping, append(wl, usersPDFPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update users_pdfs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for users_pdfs")
	}

	if !cached {
		usersPDFUpdateCacheMut.Lock()
		usersPDFUpdateCache[key] = cache
		usersPDFUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q usersPDFQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for users_pdfs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for users_pdfs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UsersPDFSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usersPDFPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"users_pdfs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, usersPDFPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in usersPDF slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all usersPDF")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UsersPDF) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no users_pdfs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(usersPDFColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	usersPDFUpsertCacheMut.RLock()
	cache, cached := usersPDFUpsertCache[key]
	usersPDFUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			usersPDFAllColumns,
			usersPDFColumnsWithDefault,
			usersPDFColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			usersPDFAllColumns,
			usersPDFPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert users_pdfs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(usersPDFPrimaryKeyColumns))
			copy(conflict, usersPDFPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"users_pdfs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(usersPDFType, usersPDFMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(usersPDFType, usersPDFMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert users_pdfs")
	}

	if !cached {
		usersPDFUpsertCacheMut.Lock()
		usersPDFUpsertCache[key] = cache
		usersPDFUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UsersPDF record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UsersPDF) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UsersPDF provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), usersPDFPrimaryKeyMapping)
	sql := "DELETE FROM \"users_pdfs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from users_pdfs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for users_pdfs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q usersPDFQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no usersPDFQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from users_pdfs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for users_pdfs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UsersPDFSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(usersPDFBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usersPDFPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"users_pdfs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, usersPDFPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from usersPDF slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for users_pdfs")
	}

	if len(usersPDFAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UsersPDF) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUsersPDF(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UsersPDFSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UsersPDFSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), usersPDFPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"users_pdfs\".* FROM \"users_pdfs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, usersPDFPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UsersPDFSlice")
	}

	*o = slice

	return nil
}

// UsersPDFExists checks if the UsersPDF row exists.
func UsersPDFExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"users_pdfs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if users_pdfs exists")
	}

	return exists, nil
}

// Exists checks if the UsersPDF row exists.
func (o *UsersPDF) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UsersPDFExists(ctx, exec, o.ID)
}