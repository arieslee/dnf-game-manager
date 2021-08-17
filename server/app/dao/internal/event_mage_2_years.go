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

// EventMage2YearsDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type EventMage2YearsDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns eventMage2YearsColumns
}

// EventMage2YearsColumns defines and stores column names for table event_mage_2years.
type eventMage2YearsColumns struct {
	MId        string //
	ServerInfo string //
	CharacNo   string //
	CharacName string //
	CreateTime string //
	DeleteTime string //
	DeleteFlag string //
}

var (
	// EventMage2Years is globally public accessible object for table event_mage_2years operations.
	EventMage2Years = EventMage2YearsDao{
		M:     g.DB("taiwan").Model("event_mage_2years").Safe(),
		DB:    g.DB("taiwan"),
		Table: "event_mage_2years",
		Columns: eventMage2YearsColumns{
			MId:        "m_id",
			ServerInfo: "server_info",
			CharacNo:   "charac_no",
			CharacName: "charac_name",
			CreateTime: "create_time",
			DeleteTime: "delete_time",
			DeleteFlag: "delete_flag",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *EventMage2YearsDao) Ctx(ctx context.Context) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *EventMage2YearsDao) As(as string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *EventMage2YearsDao) TX(tx *gdb.TX) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *EventMage2YearsDao) Master() *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *EventMage2YearsDao) Slave() *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *EventMage2YearsDao) Args(args ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *EventMage2YearsDao) LeftJoin(table ...string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *EventMage2YearsDao) RightJoin(table ...string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *EventMage2YearsDao) InnerJoin(table ...string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *EventMage2YearsDao) Fields(fieldNamesOrMapStruct ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *EventMage2YearsDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *EventMage2YearsDao) Option(option int) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *EventMage2YearsDao) OmitEmpty() *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *EventMage2YearsDao) Filter() *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Filter()}
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
func (d *EventMage2YearsDao) Where(where interface{}, args ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *EventMage2YearsDao) WhereIn(column string, in interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *EventMage2YearsDao) WhereNotIn(column string, in interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *EventMage2YearsDao) WhereNull(columns ...string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *EventMage2YearsDao) WhereNotNull(columns ...string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *EventMage2YearsDao) WhereBetween(column string, min, max interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *EventMage2YearsDao) WhereOr(where interface{}, args ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *EventMage2YearsDao) WherePri(where interface{}, args ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *EventMage2YearsDao) And(where interface{}, args ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *EventMage2YearsDao) Or(where interface{}, args ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *EventMage2YearsDao) Group(groupBy string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *EventMage2YearsDao) Order(orderBy ...string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *EventMage2YearsDao) Limit(limit ...int) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *EventMage2YearsDao) Offset(offset int) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *EventMage2YearsDao) Page(page, limit int) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *EventMage2YearsDao) Batch(batch int) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Batch(batch)}
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
func (d *EventMage2YearsDao) Cache(duration time.Duration, name ...string) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *EventMage2YearsDao) Data(data ...interface{}) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.EventMage2Years.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *EventMage2YearsDao) All(where ...interface{}) ([]*model.EventMage2Years, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.EventMage2Years
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.EventMage2Years.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *EventMage2YearsDao) One(where ...interface{}) (*model.EventMage2Years, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.EventMage2Years
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *EventMage2YearsDao) FindOne(where ...interface{}) (*model.EventMage2Years, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.EventMage2Years
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *EventMage2YearsDao) FindAll(where ...interface{}) ([]*model.EventMage2Years, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.EventMage2Years
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
func (d *EventMage2YearsDao) Struct(pointer interface{}, where ...interface{}) error {
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
func (d *EventMage2YearsDao) Structs(pointer interface{}, where ...interface{}) error {
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
func (d *EventMage2YearsDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *EventMage2YearsDao) Chunk(limit int, callback func(entities []*model.EventMage2Years, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.EventMage2Years
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *EventMage2YearsDao) LockUpdate() *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *EventMage2YearsDao) LockShared() *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *EventMage2YearsDao) Unscoped() *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *EventMage2YearsDao) Union(unions ...*gdb.Model) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *EventMage2YearsDao) UnionAll(unions ...*gdb.Model) *EventMage2YearsDao {
	return &EventMage2YearsDao{M: d.M.Union(unions...)}
}
