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

// MemberSecurityGradeDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type MemberSecurityGradeDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns memberSecurityGradeColumns
}

// MemberSecurityGradeColumns defines and stores column names for table member_security_grade.
type memberSecurityGradeColumns struct {
	MId                      string //
	LastVisitTime            string //
	PassFailCnt              string //
	LastVaccineDate          string //
	LastWindowDate           string //
	GoblinPassMod            string //
	GoblinFailCnt            string //
	SecurityCardReg          string //
	SecurityCardFailCnt      string //
	MOptReg                  string //
	PcOptReg                 string //
	BlackIpTryTime           string //
	LinearPassFailCnt        string //
	LastPassFailTime         string //
	LastCheckTime            string //
	PassModifyCheck          string //
	MemberPcReg              string //
	GatekeeperOtpReg         string //
	GoblinValidityTime       string //
	SecurityCardValidityTime string //
	ValidityIp               string //
	CargopadStatus           string //
	CargopadMod              string //
	CargopadValidityTime     string //
}

var (
	// MemberSecurityGrade is globally public accessible object for table member_security_grade operations.
	MemberSecurityGrade = MemberSecurityGradeDao{
		M:     g.DB("taiwan").Model("member_security_grade").Safe(),
		DB:    g.DB("taiwan"),
		Table: "member_security_grade",
		Columns: memberSecurityGradeColumns{
			MId:                      "m_id",
			LastVisitTime:            "last_visit_time",
			PassFailCnt:              "pass_fail_cnt",
			LastVaccineDate:          "last_vaccine_date",
			LastWindowDate:           "last_window_date",
			GoblinPassMod:            "goblin_pass_mod",
			GoblinFailCnt:            "goblin_fail_cnt",
			SecurityCardReg:          "security_card_reg",
			SecurityCardFailCnt:      "security_card_fail_cnt",
			MOptReg:                  "m_opt_reg",
			PcOptReg:                 "pc_opt_reg",
			BlackIpTryTime:           "black_ip_try_time",
			LinearPassFailCnt:        "linear_pass_fail_cnt",
			LastPassFailTime:         "last_pass_fail_time",
			LastCheckTime:            "last_check_time",
			PassModifyCheck:          "pass_modify_check",
			MemberPcReg:              "member_pc_reg",
			GatekeeperOtpReg:         "gatekeeper_otp_reg",
			GoblinValidityTime:       "goblin_validity_time",
			SecurityCardValidityTime: "security_card_validity_time",
			ValidityIp:               "validity_ip",
			CargopadStatus:           "cargopad_status",
			CargopadMod:              "cargopad_mod",
			CargopadValidityTime:     "cargopad_validity_time",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *MemberSecurityGradeDao) Ctx(ctx context.Context) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *MemberSecurityGradeDao) As(as string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *MemberSecurityGradeDao) TX(tx *gdb.TX) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *MemberSecurityGradeDao) Master() *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *MemberSecurityGradeDao) Slave() *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *MemberSecurityGradeDao) Args(args ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MemberSecurityGradeDao) LeftJoin(table ...string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MemberSecurityGradeDao) RightJoin(table ...string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *MemberSecurityGradeDao) InnerJoin(table ...string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *MemberSecurityGradeDao) Fields(fieldNamesOrMapStruct ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *MemberSecurityGradeDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *MemberSecurityGradeDao) Option(option int) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *MemberSecurityGradeDao) OmitEmpty() *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *MemberSecurityGradeDao) Filter() *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Filter()}
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
func (d *MemberSecurityGradeDao) Where(where interface{}, args ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *MemberSecurityGradeDao) WhereIn(column string, in interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *MemberSecurityGradeDao) WhereNotIn(column string, in interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *MemberSecurityGradeDao) WhereNull(columns ...string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *MemberSecurityGradeDao) WhereNotNull(columns ...string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *MemberSecurityGradeDao) WhereBetween(column string, min, max interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *MemberSecurityGradeDao) WhereOr(where interface{}, args ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *MemberSecurityGradeDao) WherePri(where interface{}, args ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *MemberSecurityGradeDao) And(where interface{}, args ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *MemberSecurityGradeDao) Or(where interface{}, args ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *MemberSecurityGradeDao) Group(groupBy string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *MemberSecurityGradeDao) Order(orderBy ...string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *MemberSecurityGradeDao) Limit(limit ...int) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *MemberSecurityGradeDao) Offset(offset int) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *MemberSecurityGradeDao) Page(page, limit int) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *MemberSecurityGradeDao) Batch(batch int) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Batch(batch)}
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
func (d *MemberSecurityGradeDao) Cache(duration time.Duration, name ...string) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *MemberSecurityGradeDao) Data(data ...interface{}) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.MemberSecurityGrade.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *MemberSecurityGradeDao) All(where ...interface{}) ([]*model.MemberSecurityGrade, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.MemberSecurityGrade
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.MemberSecurityGrade.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *MemberSecurityGradeDao) One(where ...interface{}) (*model.MemberSecurityGrade, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.MemberSecurityGrade
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *MemberSecurityGradeDao) FindOne(where ...interface{}) (*model.MemberSecurityGrade, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.MemberSecurityGrade
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *MemberSecurityGradeDao) FindAll(where ...interface{}) ([]*model.MemberSecurityGrade, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.MemberSecurityGrade
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
func (d *MemberSecurityGradeDao) Struct(pointer interface{}, where ...interface{}) error {
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
func (d *MemberSecurityGradeDao) Structs(pointer interface{}, where ...interface{}) error {
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
func (d *MemberSecurityGradeDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *MemberSecurityGradeDao) Chunk(limit int, callback func(entities []*model.MemberSecurityGrade, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.MemberSecurityGrade
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *MemberSecurityGradeDao) LockUpdate() *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *MemberSecurityGradeDao) LockShared() *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *MemberSecurityGradeDao) Unscoped() *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *MemberSecurityGradeDao) Union(unions ...*gdb.Model) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *MemberSecurityGradeDao) UnionAll(unions ...*gdb.Model) *MemberSecurityGradeDao {
	return &MemberSecurityGradeDao{M: d.M.Union(unions...)}
}
