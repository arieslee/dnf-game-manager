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

// TmeCharacDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type TmeCharacDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns tmeCharacColumns
}

// TmeCharacColumns defines and stores column names for table tme_charac.
type tmeCharacColumns struct {
	MId               string //
	CharacNo          string //
	CharacName        string //
	Village           string //
	Job               string //
	Lev               string //
	Exp               string //
	GrowType          string //
	HP                string //
	MaxHP             string //
	MaxMP             string //
	PhyAttack         string //
	PhyDefense        string //
	MagAttack         string //
	MagDefense        string //
	ElementResist     string //
	SpecProperty      string //
	InvenWeight       string //
	HpRegen           string //
	MpRegen           string //
	MoveSpeed         string //
	AttackSpeed       string //
	CastSpeed         string //
	HitRecovery       string //
	Jump              string //
	CharacWeight      string //
	Fatigue           string //
	MaxFatigue        string //
	PremiumFatigue    string //
	MaxPremiumFatigue string //
	CreateTime        string //
	LastPlayTime      string //
	DungeonClearPoint string //
	DeleteTime        string //
	DeleteFlag        string //
	GuildId           string //
	GuildRight        string //
	MemberFlag        string //
}

var (
	// TmeCharac is globally public accessible object for table tme_charac operations.
	TmeCharac = TmeCharacDao{
		M:     g.DB("default").Model("tme_charac").Safe(),
		DB:    g.DB("default"),
		Table: "tme_charac",
		Columns: tmeCharacColumns{
			MId:               "m_id",
			CharacNo:          "charac_no",
			CharacName:        "charac_name",
			Village:           "village",
			Job:               "job",
			Lev:               "lev",
			Exp:               "exp",
			GrowType:          "grow_type",
			HP:                "HP",
			MaxHP:             "maxHP",
			MaxMP:             "maxMP",
			PhyAttack:         "phy_attack",
			PhyDefense:        "phy_defense",
			MagAttack:         "mag_attack",
			MagDefense:        "mag_defense",
			ElementResist:     "element_resist",
			SpecProperty:      "spec_property",
			InvenWeight:       "inven_weight",
			HpRegen:           "hp_regen",
			MpRegen:           "mp_regen",
			MoveSpeed:         "move_speed",
			AttackSpeed:       "attack_speed",
			CastSpeed:         "cast_speed",
			HitRecovery:       "hit_recovery",
			Jump:              "jump",
			CharacWeight:      "charac_weight",
			Fatigue:           "fatigue",
			MaxFatigue:        "max_fatigue",
			PremiumFatigue:    "premium_fatigue",
			MaxPremiumFatigue: "max_premium_fatigue",
			CreateTime:        "create_time",
			LastPlayTime:      "last_play_time",
			DungeonClearPoint: "dungeon_clear_point",
			DeleteTime:        "delete_time",
			DeleteFlag:        "delete_flag",
			GuildId:           "guild_id",
			GuildRight:        "guild_right",
			MemberFlag:        "member_flag",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *TmeCharacDao) Ctx(ctx context.Context) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *TmeCharacDao) As(as string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *TmeCharacDao) TX(tx *gdb.TX) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *TmeCharacDao) Master() *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *TmeCharacDao) Slave() *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *TmeCharacDao) Args(args ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *TmeCharacDao) LeftJoin(table ...string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *TmeCharacDao) RightJoin(table ...string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *TmeCharacDao) InnerJoin(table ...string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *TmeCharacDao) Fields(fieldNamesOrMapStruct ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *TmeCharacDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *TmeCharacDao) Option(option int) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *TmeCharacDao) OmitEmpty() *TmeCharacDao {
	return &TmeCharacDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *TmeCharacDao) Filter() *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Filter()}
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
func (d *TmeCharacDao) Where(where interface{}, args ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *TmeCharacDao) WhereIn(column string, in interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *TmeCharacDao) WhereNotIn(column string, in interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *TmeCharacDao) WhereNull(columns ...string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *TmeCharacDao) WhereNotNull(columns ...string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *TmeCharacDao) WhereBetween(column string, min, max interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *TmeCharacDao) WhereOr(where interface{}, args ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *TmeCharacDao) WherePri(where interface{}, args ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *TmeCharacDao) And(where interface{}, args ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *TmeCharacDao) Or(where interface{}, args ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *TmeCharacDao) Group(groupBy string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *TmeCharacDao) Order(orderBy ...string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *TmeCharacDao) Limit(limit ...int) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *TmeCharacDao) Offset(offset int) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *TmeCharacDao) Page(page, limit int) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *TmeCharacDao) Batch(batch int) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Batch(batch)}
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
func (d *TmeCharacDao) Cache(duration time.Duration, name ...string) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *TmeCharacDao) Data(data ...interface{}) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.TmeCharac.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *TmeCharacDao) All(where ...interface{}) ([]*model.TmeCharac, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.TmeCharac
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.TmeCharac.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *TmeCharacDao) One(where ...interface{}) (*model.TmeCharac, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.TmeCharac
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *TmeCharacDao) FindOne(where ...interface{}) (*model.TmeCharac, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.TmeCharac
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *TmeCharacDao) FindAll(where ...interface{}) ([]*model.TmeCharac, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.TmeCharac
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
func (d *TmeCharacDao) Struct(pointer interface{}, where ...interface{}) error {
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
func (d *TmeCharacDao) Structs(pointer interface{}, where ...interface{}) error {
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
func (d *TmeCharacDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *TmeCharacDao) Chunk(limit int, callback func(entities []*model.TmeCharac, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.TmeCharac
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *TmeCharacDao) LockUpdate() *TmeCharacDao {
	return &TmeCharacDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *TmeCharacDao) LockShared() *TmeCharacDao {
	return &TmeCharacDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *TmeCharacDao) Unscoped() *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *TmeCharacDao) Union(unions ...*gdb.Model) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *TmeCharacDao) UnionAll(unions ...*gdb.Model) *TmeCharacDao {
	return &TmeCharacDao{M: d.M.Union(unions...)}
}
