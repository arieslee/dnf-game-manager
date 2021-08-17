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

// GmManifestNotuseDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type GmManifestNotuseDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns gmManifestNotuseColumns
}

// GmManifestNotuseColumns defines and stores column names for table gm_manifest_notuse.
type gmManifestNotuseColumns struct {
	MId   string //
	Level string //
}

var (
	// GmManifestNotuse is globally public accessible object for table gm_manifest_notuse operations.
	GmManifestNotuse = GmManifestNotuseDao{
		M:     g.DB("taiwan").Model("gm_manifest_notuse").Safe(),
		DB:    g.DB("taiwan"),
		Table: "gm_manifest_notuse",
		Columns: gmManifestNotuseColumns{
			MId:   "m_id",
			Level: "level",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *GmManifestNotuseDao) Ctx(ctx context.Context) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *GmManifestNotuseDao) As(as string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *GmManifestNotuseDao) TX(tx *gdb.TX) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *GmManifestNotuseDao) Master() *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *GmManifestNotuseDao) Slave() *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *GmManifestNotuseDao) Args(args ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *GmManifestNotuseDao) LeftJoin(table ...string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *GmManifestNotuseDao) RightJoin(table ...string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *GmManifestNotuseDao) InnerJoin(table ...string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *GmManifestNotuseDao) Fields(fieldNamesOrMapStruct ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *GmManifestNotuseDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *GmManifestNotuseDao) Option(option int) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *GmManifestNotuseDao) OmitEmpty() *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *GmManifestNotuseDao) Filter() *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Filter()}
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
func (d *GmManifestNotuseDao) Where(where interface{}, args ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *GmManifestNotuseDao) WhereIn(column string, in interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *GmManifestNotuseDao) WhereNotIn(column string, in interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *GmManifestNotuseDao) WhereNull(columns ...string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *GmManifestNotuseDao) WhereNotNull(columns ...string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *GmManifestNotuseDao) WhereBetween(column string, min, max interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *GmManifestNotuseDao) WhereOr(where interface{}, args ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *GmManifestNotuseDao) WherePri(where interface{}, args ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *GmManifestNotuseDao) And(where interface{}, args ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *GmManifestNotuseDao) Or(where interface{}, args ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *GmManifestNotuseDao) Group(groupBy string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *GmManifestNotuseDao) Order(orderBy ...string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *GmManifestNotuseDao) Limit(limit ...int) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *GmManifestNotuseDao) Offset(offset int) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *GmManifestNotuseDao) Page(page, limit int) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *GmManifestNotuseDao) Batch(batch int) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Batch(batch)}
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
func (d *GmManifestNotuseDao) Cache(duration time.Duration, name ...string) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *GmManifestNotuseDao) Data(data ...interface{}) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.GmManifestNotuse.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *GmManifestNotuseDao) All(where ...interface{}) ([]*model.GmManifestNotuse, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.GmManifestNotuse
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.GmManifestNotuse.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *GmManifestNotuseDao) One(where ...interface{}) (*model.GmManifestNotuse, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.GmManifestNotuse
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *GmManifestNotuseDao) FindOne(where ...interface{}) (*model.GmManifestNotuse, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.GmManifestNotuse
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *GmManifestNotuseDao) FindAll(where ...interface{}) ([]*model.GmManifestNotuse, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.GmManifestNotuse
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
func (d *GmManifestNotuseDao) Struct(pointer interface{}, where ...interface{}) error {
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
func (d *GmManifestNotuseDao) Structs(pointer interface{}, where ...interface{}) error {
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
func (d *GmManifestNotuseDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *GmManifestNotuseDao) Chunk(limit int, callback func(entities []*model.GmManifestNotuse, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.GmManifestNotuse
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *GmManifestNotuseDao) LockUpdate() *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *GmManifestNotuseDao) LockShared() *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *GmManifestNotuseDao) Unscoped() *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *GmManifestNotuseDao) Union(unions ...*gdb.Model) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *GmManifestNotuseDao) UnionAll(unions ...*gdb.Model) *GmManifestNotuseDao {
	return &GmManifestNotuseDao{M: d.M.Union(unions...)}
}
