// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package schema

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ReceivedMessage is an object representing the database table.
type ReceivedMessage struct {
	MessageID  string    `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	AttendeeID string    `boil:"attendee_id" json:"attendee_id" toml:"attendee_id" yaml:"attendee_id"`
	ReadAt     time.Time `boil:"read_at" json:"read_at" toml:"read_at" yaml:"read_at"`

	R *receivedMessageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L receivedMessageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReceivedMessageColumns = struct {
	MessageID  string
	AttendeeID string
	ReadAt     string
}{
	MessageID:  "message_id",
	AttendeeID: "attendee_id",
	ReadAt:     "read_at",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ReceivedMessageWhere = struct {
	MessageID  whereHelperstring
	AttendeeID whereHelperstring
	ReadAt     whereHelpertime_Time
}{
	MessageID:  whereHelperstring{field: "\"received_messages\".\"message_id\""},
	AttendeeID: whereHelperstring{field: "\"received_messages\".\"attendee_id\""},
	ReadAt:     whereHelpertime_Time{field: "\"received_messages\".\"read_at\""},
}

// ReceivedMessageRels is where relationship names are stored.
var ReceivedMessageRels = struct {
	Attendee string
	Message  string
}{
	Attendee: "Attendee",
	Message:  "Message",
}

// receivedMessageR is where relationships are stored.
type receivedMessageR struct {
	Attendee *Attendee
	Message  *Message
}

// NewStruct creates a new relationship struct
func (*receivedMessageR) NewStruct() *receivedMessageR {
	return &receivedMessageR{}
}

// receivedMessageL is where Load methods for each relationship are stored.
type receivedMessageL struct{}

var (
	receivedMessageAllColumns            = []string{"message_id", "attendee_id", "read_at"}
	receivedMessageColumnsWithoutDefault = []string{"message_id", "attendee_id"}
	receivedMessageColumnsWithDefault    = []string{"read_at"}
	receivedMessagePrimaryKeyColumns     = []string{"message_id", "attendee_id"}
)

type (
	// ReceivedMessageSlice is an alias for a slice of pointers to ReceivedMessage.
	// This should generally be used opposed to []ReceivedMessage.
	ReceivedMessageSlice []*ReceivedMessage

	receivedMessageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	receivedMessageType                 = reflect.TypeOf(&ReceivedMessage{})
	receivedMessageMapping              = queries.MakeStructMapping(receivedMessageType)
	receivedMessagePrimaryKeyMapping, _ = queries.BindMapping(receivedMessageType, receivedMessageMapping, receivedMessagePrimaryKeyColumns)
	receivedMessageInsertCacheMut       sync.RWMutex
	receivedMessageInsertCache          = make(map[string]insertCache)
	receivedMessageUpdateCacheMut       sync.RWMutex
	receivedMessageUpdateCache          = make(map[string]updateCache)
	receivedMessageUpsertCacheMut       sync.RWMutex
	receivedMessageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single receivedMessage record from the query.
func (q receivedMessageQuery) One(exec boil.Executor) (*ReceivedMessage, error) {
	o := &ReceivedMessage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: failed to execute a one query for received_messages")
	}

	return o, nil
}

// All returns all ReceivedMessage records from the query.
func (q receivedMessageQuery) All(exec boil.Executor) (ReceivedMessageSlice, error) {
	var o []*ReceivedMessage

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "schema: failed to assign all query results to ReceivedMessage slice")
	}

	return o, nil
}

// Count returns the count of all ReceivedMessage records in the query.
func (q receivedMessageQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to count received_messages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q receivedMessageQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "schema: failed to check if received_messages exists")
	}

	return count > 0, nil
}

// Attendee pointed to by the foreign key.
func (o *ReceivedMessage) Attendee(mods ...qm.QueryMod) attendeeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AttendeeID),
	}

	queryMods = append(queryMods, mods...)

	query := Attendees(queryMods...)
	queries.SetFrom(query.Query, "\"attendees\"")

	return query
}

// Message pointed to by the foreign key.
func (o *ReceivedMessage) Message(mods ...qm.QueryMod) messageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MessageID),
	}

	queryMods = append(queryMods, mods...)

	query := Messages(queryMods...)
	queries.SetFrom(query.Query, "\"messages\"")

	return query
}

// LoadAttendee allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (receivedMessageL) LoadAttendee(e boil.Executor, singular bool, maybeReceivedMessage interface{}, mods queries.Applicator) error {
	var slice []*ReceivedMessage
	var object *ReceivedMessage

	if singular {
		object = maybeReceivedMessage.(*ReceivedMessage)
	} else {
		slice = *maybeReceivedMessage.(*[]*ReceivedMessage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &receivedMessageR{}
		}
		args = append(args, object.AttendeeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &receivedMessageR{}
			}

			for _, a := range args {
				if a == obj.AttendeeID {
					continue Outer
				}
			}

			args = append(args, obj.AttendeeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`attendees`), qm.WhereIn(`attendees.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Attendee")
	}

	var resultSlice []*Attendee
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Attendee")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for attendees")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for attendees")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Attendee = foreign
		if foreign.R == nil {
			foreign.R = &attendeeR{}
		}
		foreign.R.ReceivedMessages = append(foreign.R.ReceivedMessages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AttendeeID == foreign.ID {
				local.R.Attendee = foreign
				if foreign.R == nil {
					foreign.R = &attendeeR{}
				}
				foreign.R.ReceivedMessages = append(foreign.R.ReceivedMessages, local)
				break
			}
		}
	}

	return nil
}

// LoadMessage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (receivedMessageL) LoadMessage(e boil.Executor, singular bool, maybeReceivedMessage interface{}, mods queries.Applicator) error {
	var slice []*ReceivedMessage
	var object *ReceivedMessage

	if singular {
		object = maybeReceivedMessage.(*ReceivedMessage)
	} else {
		slice = *maybeReceivedMessage.(*[]*ReceivedMessage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &receivedMessageR{}
		}
		args = append(args, object.MessageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &receivedMessageR{}
			}

			for _, a := range args {
				if a == obj.MessageID {
					continue Outer
				}
			}

			args = append(args, obj.MessageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`messages`), qm.WhereIn(`messages.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Message")
	}

	var resultSlice []*Message
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Message")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for messages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for messages")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Message = foreign
		if foreign.R == nil {
			foreign.R = &messageR{}
		}
		foreign.R.ReceivedMessages = append(foreign.R.ReceivedMessages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MessageID == foreign.ID {
				local.R.Message = foreign
				if foreign.R == nil {
					foreign.R = &messageR{}
				}
				foreign.R.ReceivedMessages = append(foreign.R.ReceivedMessages, local)
				break
			}
		}
	}

	return nil
}

// SetAttendee of the receivedMessage to the related item.
// Sets o.R.Attendee to related.
// Adds o to related.R.ReceivedMessages.
func (o *ReceivedMessage) SetAttendee(exec boil.Executor, insert bool, related *Attendee) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"received_messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"attendee_id"}),
		strmangle.WhereClause("\"", "\"", 2, receivedMessagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.MessageID, o.AttendeeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AttendeeID = related.ID
	if o.R == nil {
		o.R = &receivedMessageR{
			Attendee: related,
		}
	} else {
		o.R.Attendee = related
	}

	if related.R == nil {
		related.R = &attendeeR{
			ReceivedMessages: ReceivedMessageSlice{o},
		}
	} else {
		related.R.ReceivedMessages = append(related.R.ReceivedMessages, o)
	}

	return nil
}

// SetMessage of the receivedMessage to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.ReceivedMessages.
func (o *ReceivedMessage) SetMessage(exec boil.Executor, insert bool, related *Message) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"received_messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
		strmangle.WhereClause("\"", "\"", 2, receivedMessagePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.MessageID, o.AttendeeID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MessageID = related.ID
	if o.R == nil {
		o.R = &receivedMessageR{
			Message: related,
		}
	} else {
		o.R.Message = related
	}

	if related.R == nil {
		related.R = &messageR{
			ReceivedMessages: ReceivedMessageSlice{o},
		}
	} else {
		related.R.ReceivedMessages = append(related.R.ReceivedMessages, o)
	}

	return nil
}

// ReceivedMessages retrieves all the records using an executor.
func ReceivedMessages(mods ...qm.QueryMod) receivedMessageQuery {
	mods = append(mods, qm.From("\"received_messages\""))
	return receivedMessageQuery{NewQuery(mods...)}
}

// FindReceivedMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReceivedMessage(exec boil.Executor, messageID string, attendeeID string, selectCols ...string) (*ReceivedMessage, error) {
	receivedMessageObj := &ReceivedMessage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"received_messages\" where \"message_id\"=$1 AND \"attendee_id\"=$2", sel,
	)

	q := queries.Raw(query, messageID, attendeeID)

	err := q.Bind(nil, exec, receivedMessageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "schema: unable to select from received_messages")
	}

	return receivedMessageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ReceivedMessage) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no received_messages provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(receivedMessageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	receivedMessageInsertCacheMut.RLock()
	cache, cached := receivedMessageInsertCache[key]
	receivedMessageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			receivedMessageAllColumns,
			receivedMessageColumnsWithDefault,
			receivedMessageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(receivedMessageType, receivedMessageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(receivedMessageType, receivedMessageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"received_messages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"received_messages\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "schema: unable to insert into received_messages")
	}

	if !cached {
		receivedMessageInsertCacheMut.Lock()
		receivedMessageInsertCache[key] = cache
		receivedMessageInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the ReceivedMessage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ReceivedMessage) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	receivedMessageUpdateCacheMut.RLock()
	cache, cached := receivedMessageUpdateCache[key]
	receivedMessageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			receivedMessageAllColumns,
			receivedMessagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("schema: unable to update received_messages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"received_messages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, receivedMessagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(receivedMessageType, receivedMessageMapping, append(wl, receivedMessagePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update received_messages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by update for received_messages")
	}

	if !cached {
		receivedMessageUpdateCacheMut.Lock()
		receivedMessageUpdateCache[key] = cache
		receivedMessageUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q receivedMessageQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all for received_messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected for received_messages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReceivedMessageSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("schema: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), receivedMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"received_messages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, receivedMessagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to update all in receivedMessage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to retrieve rows affected all in update all receivedMessage")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ReceivedMessage) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("schema: no received_messages provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(receivedMessageColumnsWithDefault, o)

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

	receivedMessageUpsertCacheMut.RLock()
	cache, cached := receivedMessageUpsertCache[key]
	receivedMessageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			receivedMessageAllColumns,
			receivedMessageColumnsWithDefault,
			receivedMessageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			receivedMessageAllColumns,
			receivedMessagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("schema: unable to upsert received_messages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(receivedMessagePrimaryKeyColumns))
			copy(conflict, receivedMessagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"received_messages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(receivedMessageType, receivedMessageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(receivedMessageType, receivedMessageMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "schema: unable to upsert received_messages")
	}

	if !cached {
		receivedMessageUpsertCacheMut.Lock()
		receivedMessageUpsertCache[key] = cache
		receivedMessageUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single ReceivedMessage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReceivedMessage) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("schema: no ReceivedMessage provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), receivedMessagePrimaryKeyMapping)
	sql := "DELETE FROM \"received_messages\" WHERE \"message_id\"=$1 AND \"attendee_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete from received_messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by delete for received_messages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q receivedMessageQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("schema: no receivedMessageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from received_messages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for received_messages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReceivedMessageSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), receivedMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"received_messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, receivedMessagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "schema: unable to delete all from receivedMessage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "schema: failed to get rows affected by deleteall for received_messages")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ReceivedMessage) Reload(exec boil.Executor) error {
	ret, err := FindReceivedMessage(exec, o.MessageID, o.AttendeeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReceivedMessageSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReceivedMessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), receivedMessagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"received_messages\".* FROM \"received_messages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, receivedMessagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "schema: unable to reload all in ReceivedMessageSlice")
	}

	*o = slice

	return nil
}

// ReceivedMessageExists checks if the ReceivedMessage row exists.
func ReceivedMessageExists(exec boil.Executor, messageID string, attendeeID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"received_messages\" where \"message_id\"=$1 AND \"attendee_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, messageID, attendeeID)
	}
	row := exec.QueryRow(sql, messageID, attendeeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "schema: unable to check if received_messages exists")
	}

	return exists, nil
}