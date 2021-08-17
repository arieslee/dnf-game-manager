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

// MemberPunishInfoHistory2013Dao is the manager for logic model data accessing
// and custom defined data operations functions management.
type MemberPunishInfoHistory2013Dao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns memberPunishInfoHistory2013Columns
}

// MemberPunishInfoHistory2013Columns defines and stores column names for table member_punish_info_history_2013.
type memberPunishInfoHistory2013Columns struct {
	No          string //
	MId         string //
	PunishType  string //
	OccTime     string //
	PunishValue string //
	ApplyFlag   string //
	StartTime   string //
	EndTime     string //
	AdminId     string //
	Reason      string //
	IsKicked    string //
	FirstSsn    string //
	SecondSsn   string //
}

var (
	// MemberPunishInfoHistory2013 is globally public accessible object for table member_punish_info_history_2013 operations.
	MemberPunishInfoHistory2013 = MemberPunishInfoHistory2013Dao{
		M:     g.DB("default").Model("member_punish_info_history_2013").Safe(),
		DB:    g.DB("default"),
		Table: "member_punish_info_history_2013",
		Columns: memberPunishInfoHistory2013Columns{
			No:          "no",
			MId:         "m_id",
			PunishType:  "punish_type",
			OccTime:     "occ_time",
			PunishValue: "punish_value",
			ApplyFlag:   "apply_flag",
			StartTime:   "start_time",
			EndTime:     "end_time",
			AdminId:     "admin_id",
			Reason:      "reason",
			IsKicked:    "is_kicked",
			FirstSsn:    "first_ssn",
			SecondSsn:   "second_ssn",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *MemberPunishInfoHistory2013Dao) Ctx(ctx context.Context) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *MemberPunishInfoHistory2013Dao) As(as string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *MemberPunishInfoHistory2013Dao) TX(tx *gdb.TX) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *MemberPunishInfoHistory2013Dao) Master() *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *MemberPunishInfoHistory2013Dao) Slave() *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *MemberPunishInfoHistory2013Dao) Args(args ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MemberPunishInfoHistory2013Dao) LeftJoin(table ...string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MemberPunishInfoHistory2013Dao) RightJoin(table ...string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MemberPunishInfoHistory2013Dao) InnerJoin(table ...string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *MemberPunishInfoHistory2013Dao) Fields(fieldNamesOrMapStruct ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *MemberPunishInfoHistory2013Dao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *MemberPunishInfoHistory2013Dao) Option(option int) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *MemberPunishInfoHistory2013Dao) OmitEmpty() *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *MemberPunishInfoHistory2013Dao) Filter() *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Filter()}
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
func (d *MemberPunishInfoHistory2013Dao) Where(where interface{}, args ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *MemberPunishInfoHistory2013Dao) WhereIn(column string, in interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *MemberPunishInfoHistory2013Dao) WhereNotIn(column string, in interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *MemberPunishInfoHistory2013Dao) WhereNull(columns ...string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *MemberPunishInfoHistory2013Dao) WhereNotNull(columns ...string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *MemberPunishInfoHistory2013Dao) WhereBetween(column string, min, max interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *MemberPunishInfoHistory2013Dao) WhereOr(where interface{}, args ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *MemberPunishInfoHistory2013Dao) WherePri(where interface{}, args ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *MemberPunishInfoHistory2013Dao) And(where interface{}, args ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *MemberPunishInfoHistory2013Dao) Or(where interface{}, args ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *MemberPunishInfoHistory2013Dao) Group(groupBy string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *MemberPunishInfoHistory2013Dao) Order(orderBy ...string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *MemberPunishInfoHistory2013Dao) Limit(limit ...int) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *MemberPunishInfoHistory2013Dao) Offset(offset int) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *MemberPunishInfoHistory2013Dao) Page(page, limit int) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *MemberPunishInfoHistory2013Dao) Batch(batch int) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Batch(batch)}
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
func (d *MemberPunishInfoHistory2013Dao) Cache(duration time.Duration, name ...string) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *MemberPunishInfoHistory2013Dao) Data(data ...interface{}) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.MemberPunishInfoHistory2013.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *MemberPunishInfoHistory2013Dao) All(where ...interface{}) ([]*model.MemberPunishInfoHistory2013, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.MemberPunishInfoHistory2013
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.MemberPunishInfoHistory2013.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *MemberPunishInfoHistory2013Dao) One(where ...interface{}) (*model.MemberPunishInfoHistory2013, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.MemberPunishInfoHistory2013
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *MemberPunishInfoHistory2013Dao) FindOne(where ...interface{}) (*model.MemberPunishInfoHistory2013, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.MemberPunishInfoHistory2013
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *MemberPunishInfoHistory2013Dao) FindAll(where ...interface{}) ([]*model.MemberPunishInfoHistory2013, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.MemberPunishInfoHistory2013
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
func (d *MemberPunishInfoHistory2013Dao) Struct(pointer interface{}, where ...interface{}) error {
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
func (d *MemberPunishInfoHistory2013Dao) Structs(pointer interface{}, where ...interface{}) error {
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
func (d *MemberPunishInfoHistory2013Dao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *MemberPunishInfoHistory2013Dao) Chunk(limit int, callback func(entities []*model.MemberPunishInfoHistory2013, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.MemberPunishInfoHistory2013
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *MemberPunishInfoHistory2013Dao) LockUpdate() *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *MemberPunishInfoHistory2013Dao) LockShared() *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *MemberPunishInfoHistory2013Dao) Unscoped() *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *MemberPunishInfoHistory2013Dao) Union(unions ...*gdb.Model) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *MemberPunishInfoHistory2013Dao) UnionAll(unions ...*gdb.Model) *MemberPunishInfoHistory2013Dao {
	return &MemberPunishInfoHistory2013Dao{M: d.M.Union(unions...)}
}
