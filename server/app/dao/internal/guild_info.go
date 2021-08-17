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

// GuildInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type GuildInfoDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns guildInfoColumns
}

// GuildInfoColumns defines and stores column names for table guild_info.
type guildInfoColumns struct {
	GuildId          string //
	ServerId         string //
	GuildName        string //
	MasterId         string //
	MasterNo         string //
	MasterName       string //
	GuildUrl         string //
	GuildIcon        string //
	CreateTime       string //
	Lev              string //
	Ability          string //
	ExpireFlag       string //
	ExpireTime       string //
	MemberSecedeTime string //
	MemberCount      string //
	RecommendFlag    string //
	RecommendTime    string //
	GuildPoint       string //
	GuildPointAcc    string //
	GuildPointPrev   string //
	GuildRank        string //
	GuildWarPoint    string //
	FinalEntry       string //
	FinalWin         string //
	GuildIconAuth    string //
	GuildExp         string //
}

var (
	// GuildInfo is globally public accessible object for table guild_info operations.
	GuildInfo = GuildInfoDao{
		M:     g.DB("default").Model("guild_info").Safe(),
		DB:    g.DB("default"),
		Table: "guild_info",
		Columns: guildInfoColumns{
			GuildId:          "guild_id",
			ServerId:         "server_id",
			GuildName:        "guild_name",
			MasterId:         "master_id",
			MasterNo:         "master_no",
			MasterName:       "master_name",
			GuildUrl:         "guild_url",
			GuildIcon:        "guild_icon",
			CreateTime:       "create_time",
			Lev:              "lev",
			Ability:          "ability",
			ExpireFlag:       "expire_flag",
			ExpireTime:       "expire_time",
			MemberSecedeTime: "member_secede_time",
			MemberCount:      "member_count",
			RecommendFlag:    "recommend_flag",
			RecommendTime:    "recommend_time",
			GuildPoint:       "guild_point",
			GuildPointAcc:    "guild_point_acc",
			GuildPointPrev:   "guild_point_prev",
			GuildRank:        "guild_rank",
			GuildWarPoint:    "guild_war_point",
			FinalEntry:       "final_entry",
			FinalWin:         "final_win",
			GuildIconAuth:    "guild_icon_auth",
			GuildExp:         "guild_exp",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *GuildInfoDao) Ctx(ctx context.Context) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *GuildInfoDao) As(as string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *GuildInfoDao) TX(tx *gdb.TX) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *GuildInfoDao) Master() *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *GuildInfoDao) Slave() *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *GuildInfoDao) Args(args ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *GuildInfoDao) LeftJoin(table ...string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *GuildInfoDao) RightJoin(table ...string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *GuildInfoDao) InnerJoin(table ...string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *GuildInfoDao) Fields(fieldNamesOrMapStruct ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *GuildInfoDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *GuildInfoDao) Option(option int) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *GuildInfoDao) OmitEmpty() *GuildInfoDao {
	return &GuildInfoDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *GuildInfoDao) Filter() *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Filter()}
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
func (d *GuildInfoDao) Where(where interface{}, args ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *GuildInfoDao) WhereIn(column string, in interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *GuildInfoDao) WhereNotIn(column string, in interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *GuildInfoDao) WhereNull(columns ...string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *GuildInfoDao) WhereNotNull(columns ...string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *GuildInfoDao) WhereBetween(column string, min, max interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *GuildInfoDao) WhereOr(where interface{}, args ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *GuildInfoDao) WherePri(where interface{}, args ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *GuildInfoDao) And(where interface{}, args ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *GuildInfoDao) Or(where interface{}, args ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *GuildInfoDao) Group(groupBy string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *GuildInfoDao) Order(orderBy ...string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *GuildInfoDao) Limit(limit ...int) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *GuildInfoDao) Offset(offset int) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *GuildInfoDao) Page(page, limit int) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *GuildInfoDao) Batch(batch int) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Batch(batch)}
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
func (d *GuildInfoDao) Cache(duration time.Duration, name ...string) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *GuildInfoDao) Data(data ...interface{}) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.GuildInfo.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *GuildInfoDao) All(where ...interface{}) ([]*model.GuildInfo, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.GuildInfo
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.GuildInfo.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *GuildInfoDao) One(where ...interface{}) (*model.GuildInfo, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.GuildInfo
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *GuildInfoDao) FindOne(where ...interface{}) (*model.GuildInfo, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.GuildInfo
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *GuildInfoDao) FindAll(where ...interface{}) ([]*model.GuildInfo, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.GuildInfo
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
func (d *GuildInfoDao) Struct(pointer interface{}, where ...interface{}) error {
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
func (d *GuildInfoDao) Structs(pointer interface{}, where ...interface{}) error {
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
func (d *GuildInfoDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *GuildInfoDao) Chunk(limit int, callback func(entities []*model.GuildInfo, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.GuildInfo
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *GuildInfoDao) LockUpdate() *GuildInfoDao {
	return &GuildInfoDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *GuildInfoDao) LockShared() *GuildInfoDao {
	return &GuildInfoDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *GuildInfoDao) Unscoped() *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *GuildInfoDao) Union(unions ...*gdb.Model) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *GuildInfoDao) UnionAll(unions ...*gdb.Model) *GuildInfoDao {
	return &GuildInfoDao{M: d.M.Union(unions...)}
}
