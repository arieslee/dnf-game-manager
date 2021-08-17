package internal

import (
	"context"
	"database/sql"
	"dnf-game-manager/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"
)

// EventNewmember0709EntryDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type EventNewmember0709EntryDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns eventNewmember0709EntryColumns
}

// EventNewmember0709EntryColumns defines and stores column names for table event_newmember0709_entry.
type eventNewmember0709EntryColumns struct {
	MId        string //
	OccDate    string //
	ServerId   string //
	CharacNo   string //
	Item1No    string //
	Item1Check string //
	Item2No    string //
	Item2Check string //
}

var (
	// EventNewmember0709Entry is globally public accessible object for table event_newmember0709_entry operations.
	EventNewmember0709Entry = EventNewmember0709EntryDao{
		M:     g.DB("default").Model("event_newmember0709_entry").Safe(),
		DB:    g.DB("default"),
		Table: "event_newmember0709_entry",
		Columns: eventNewmember0709EntryColumns{
			MId:        "m_id",
			OccDate:    "occ_date",
			ServerId:   "server_id",
			CharacNo:   "charac_no",
			Item1No:    "item1_no",
			Item1Check: "item1_check",
			Item2No:    "item2_no",
			Item2Check: "item2_check",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *EventNewmember0709EntryDao) Ctx(ctx context.Context) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *EventNewmember0709EntryDao) As(as string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *EventNewmember0709EntryDao) TX(tx *gdb.TX) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *EventNewmember0709EntryDao) Master() *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *EventNewmember0709EntryDao) Slave() *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *EventNewmember0709EntryDao) Args(args ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *EventNewmember0709EntryDao) LeftJoin(table ...string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *EventNewmember0709EntryDao) RightJoin(table ...string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *EventNewmember0709EntryDao) InnerJoin(table ...string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *EventNewmember0709EntryDao) Fields(fieldNamesOrMapStruct ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *EventNewmember0709EntryDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *EventNewmember0709EntryDao) Option(option int) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *EventNewmember0709EntryDao) OmitEmpty() *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *EventNewmember0709EntryDao) Filter() *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *EventNewmember0709EntryDao) Where(where interface{}, args ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *EventNewmember0709EntryDao) WhereIn(column string, in interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *EventNewmember0709EntryDao) WhereNotIn(column string, in interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *EventNewmember0709EntryDao) WhereNull(columns ...string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *EventNewmember0709EntryDao) WhereNotNull(columns ...string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *EventNewmember0709EntryDao) WhereBetween(column string, min, max interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *EventNewmember0709EntryDao) WhereOr(where interface{}, args ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *EventNewmember0709EntryDao) WherePri(where interface{}, args ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *EventNewmember0709EntryDao) And(where interface{}, args ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *EventNewmember0709EntryDao) Or(where interface{}, args ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *EventNewmember0709EntryDao) Group(groupBy string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *EventNewmember0709EntryDao) Order(orderBy ...string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *EventNewmember0709EntryDao) Limit(limit ...int) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *EventNewmember0709EntryDao) Offset(offset int) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *EventNewmember0709EntryDao) Page(page, limit int) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *EventNewmember0709EntryDao) Batch(batch int) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *EventNewmember0709EntryDao) Cache(duration time.Duration, name ...string) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *EventNewmember0709EntryDao) Data(data ...interface{}) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.EventNewmember0709Entry.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *EventNewmember0709EntryDao) All(where ...interface{}) ([]*model.EventNewmember0709Entry, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.EventNewmember0709Entry
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.EventNewmember0709Entry.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *EventNewmember0709EntryDao) One(where ...interface{}) (*model.EventNewmember0709Entry, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.EventNewmember0709Entry
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *EventNewmember0709EntryDao) FindOne(where ...interface{}) (*model.EventNewmember0709Entry, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.EventNewmember0709Entry
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *EventNewmember0709EntryDao) FindAll(where ...interface{}) ([]*model.EventNewmember0709Entry, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.EventNewmember0709Entry
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *EventNewmember0709EntryDao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *EventNewmember0709EntryDao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *EventNewmember0709EntryDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *EventNewmember0709EntryDao) Chunk(limit int, callback func(entities []*model.EventNewmember0709Entry, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.EventNewmember0709Entry
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *EventNewmember0709EntryDao) LockUpdate() *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *EventNewmember0709EntryDao) LockShared() *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *EventNewmember0709EntryDao) Unscoped() *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *EventNewmember0709EntryDao) Union(unions ...*gdb.Model) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *EventNewmember0709EntryDao) UnionAll(unions ...*gdb.Model) *EventNewmember0709EntryDao {
	return &EventNewmember0709EntryDao{M: d.M.Union(unions...)}
}
