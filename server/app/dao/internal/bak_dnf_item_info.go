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

// BakDnfItemInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type BakDnfItemInfoDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns bakDnfItemInfoColumns
}

// BakDnfItemInfoColumns defines and stores column names for table bak_dnf_item_info.
type bakDnfItemInfoColumns struct {
	ItNo                 string //
	ItName               string //
	ItEngName            string //
	ItExplain            string //
	MasterType           string //
	SubType              string //
	Job                  string //
	Class                string //
	Revert               string //
	Level                string //
	Skill                string //
	CreateRatio          string //
	Rarity               string //
	Weight               string //
	Price                string //
	Cash                 string //
	Medal                string //
	Durability           string //
	Cooltime             string //
	HpMax                string //
	MpMax                string //
	PhyAtt               string //
	PhyDef               string //
	MagAtt               string //
	MagDef               string //
	EquipPhyAtt          string //
	EquipPhyDef          string //
	EquipMagAtt          string //
	EquipMagDef          string //
	RefFire              string //
	RefWater             string //
	RefDark              string //
	RefLight             string //
	RefAll               string //
	RefSlow              string //
	RefFreeze            string //
	RefPoison            string //
	RefStun              string //
	RefCus               string //
	RefBlind             string //
	RefLite              string //
	RefSton              string //
	RefSleep             string //
	RefDeekement         string //
	RefDeadlystrike      string //
	RefBleeding          string //
	RefConfuse           string //
	RefHold              string //
	RefAllStat           string //
	RefPierce            string //
	RefStuck             string //
	InvenMax             string //
	HpRegenrate          string //
	MpRegenrate          string //
	MovSpeed             string //
	AttSpeed             string //
	Quest                string //
	HitRecovery          string //
	Jump                 string //
	AttElement           string //
	AttActiveStatus      string //
	AttActiveStatusRatio string //
	AttActiveStatusPow   string //
	AttBackforce         string //
	AttUpforce           string //
	AttHpDrain           string //
	AttMpDrain           string //
	CriticalhitRate      string //
	StuckRate            string //
	AttDefenseIgnore     string //
	SkillLevelup         string //
	SetType              string //
	Url                  string //
	JewelType            string //
}

var (
	// BakDnfItemInfo is globally public accessible object for table bak_dnf_item_info operations.
	BakDnfItemInfo = BakDnfItemInfoDao{
		M:     g.DB("default").Model("bak_dnf_item_info").Safe(),
		DB:    g.DB("default"),
		Table: "bak_dnf_item_info",
		Columns: bakDnfItemInfoColumns{
			ItNo:                 "it_no",
			ItName:               "it_name",
			ItEngName:            "it_eng_name",
			ItExplain:            "it_explain",
			MasterType:           "master_type",
			SubType:              "sub_type",
			Job:                  "job",
			Class:                "class",
			Revert:               "revert",
			Level:                "level",
			Skill:                "skill",
			CreateRatio:          "create_ratio",
			Rarity:               "rarity",
			Weight:               "weight",
			Price:                "price",
			Cash:                 "cash",
			Medal:                "medal",
			Durability:           "durability",
			Cooltime:             "cooltime",
			HpMax:                "hp_max",
			MpMax:                "mp_max",
			PhyAtt:               "phy_att",
			PhyDef:               "phy_def",
			MagAtt:               "mag_att",
			MagDef:               "mag_def",
			EquipPhyAtt:          "equip_phy_att",
			EquipPhyDef:          "equip_phy_def",
			EquipMagAtt:          "equip_mag_att",
			EquipMagDef:          "equip_mag_def",
			RefFire:              "ref_fire",
			RefWater:             "ref_water",
			RefDark:              "ref_dark",
			RefLight:             "ref_light",
			RefAll:               "ref_all",
			RefSlow:              "ref_slow",
			RefFreeze:            "ref_freeze",
			RefPoison:            "ref_poison",
			RefStun:              "ref_stun",
			RefCus:               "ref_cus",
			RefBlind:             "ref_blind",
			RefLite:              "ref_lite",
			RefSton:              "ref_ston",
			RefSleep:             "ref_sleep",
			RefDeekement:         "ref_deekement",
			RefDeadlystrike:      "ref_deadlystrike",
			RefBleeding:          "ref_bleeding",
			RefConfuse:           "ref_confuse",
			RefHold:              "ref_hold",
			RefAllStat:           "ref_all_stat",
			RefPierce:            "ref_pierce",
			RefStuck:             "ref_stuck",
			InvenMax:             "inven_max",
			HpRegenrate:          "hp_regenrate",
			MpRegenrate:          "mp_regenrate",
			MovSpeed:             "mov_speed",
			AttSpeed:             "att_speed",
			Quest:                "quest",
			HitRecovery:          "hit_recovery",
			Jump:                 "jump",
			AttElement:           "att_element",
			AttActiveStatus:      "att_active_status",
			AttActiveStatusRatio: "att_active_status_ratio",
			AttActiveStatusPow:   "att_active_status_pow",
			AttBackforce:         "att_backforce",
			AttUpforce:           "att_upforce",
			AttHpDrain:           "att_hp_drain",
			AttMpDrain:           "att_mp_drain",
			CriticalhitRate:      "criticalhit_rate",
			StuckRate:            "stuck_rate",
			AttDefenseIgnore:     "att_defenseIgnore",
			SkillLevelup:         "skill_levelup",
			SetType:              "set_type",
			Url:                  "url",
			JewelType:            "jewel_type",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *BakDnfItemInfoDao) Ctx(ctx context.Context) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *BakDnfItemInfoDao) As(as string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *BakDnfItemInfoDao) TX(tx *gdb.TX) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *BakDnfItemInfoDao) Master() *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *BakDnfItemInfoDao) Slave() *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *BakDnfItemInfoDao) Args(args ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *BakDnfItemInfoDao) LeftJoin(table ...string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *BakDnfItemInfoDao) RightJoin(table ...string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *BakDnfItemInfoDao) InnerJoin(table ...string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *BakDnfItemInfoDao) Fields(fieldNamesOrMapStruct ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *BakDnfItemInfoDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *BakDnfItemInfoDao) Option(option int) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *BakDnfItemInfoDao) OmitEmpty() *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *BakDnfItemInfoDao) Filter() *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Filter()}
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
func (d *BakDnfItemInfoDao) Where(where interface{}, args ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Where(where, args...)}
}

// WhereIn builds `column IN (in)` statement.
func (d *BakDnfItemInfoDao) WhereIn(column string, in interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.WhereIn(column, in)}
}

// WhereNotIn builds `column NOT IN (in)` statement.
func (d *BakDnfItemInfoDao) WhereNotIn(column string, in interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.WhereNotIn(column, in)}
}

// WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.
func (d *BakDnfItemInfoDao) WhereNull(columns ...string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.WhereNull(columns...)}
}

// WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.
func (d *BakDnfItemInfoDao) WhereNotNull(columns ...string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.WhereNotNull(columns...)}
}

// WhereBetween builds `column BETWEEN min AND max` statement.
func (d *BakDnfItemInfoDao) WhereBetween(column string, min, max interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.WhereBetween(column, min, max)}
}

// WhereOr adds "OR" condition to the where statement.
func (d *BakDnfItemInfoDao) WhereOr(where interface{}, args ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.WhereOr(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *BakDnfItemInfoDao) WherePri(where interface{}, args ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *BakDnfItemInfoDao) And(where interface{}, args ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *BakDnfItemInfoDao) Or(where interface{}, args ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *BakDnfItemInfoDao) Group(groupBy string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *BakDnfItemInfoDao) Order(orderBy ...string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *BakDnfItemInfoDao) Limit(limit ...int) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *BakDnfItemInfoDao) Offset(offset int) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *BakDnfItemInfoDao) Page(page, limit int) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *BakDnfItemInfoDao) Batch(batch int) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Batch(batch)}
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
func (d *BakDnfItemInfoDao) Cache(duration time.Duration, name ...string) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *BakDnfItemInfoDao) Data(data ...interface{}) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.BakDnfItemInfo.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *BakDnfItemInfoDao) All(where ...interface{}) ([]*model.BakDnfItemInfo, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.BakDnfItemInfo
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.BakDnfItemInfo.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *BakDnfItemInfoDao) One(where ...interface{}) (*model.BakDnfItemInfo, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.BakDnfItemInfo
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *BakDnfItemInfoDao) FindOne(where ...interface{}) (*model.BakDnfItemInfo, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.BakDnfItemInfo
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *BakDnfItemInfoDao) FindAll(where ...interface{}) ([]*model.BakDnfItemInfo, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.BakDnfItemInfo
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
func (d *BakDnfItemInfoDao) Struct(pointer interface{}, where ...interface{}) error {
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
func (d *BakDnfItemInfoDao) Structs(pointer interface{}, where ...interface{}) error {
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
func (d *BakDnfItemInfoDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *BakDnfItemInfoDao) Chunk(limit int, callback func(entities []*model.BakDnfItemInfo, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.BakDnfItemInfo
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *BakDnfItemInfoDao) LockUpdate() *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *BakDnfItemInfoDao) LockShared() *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *BakDnfItemInfoDao) Unscoped() *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Unscoped()}
}

// Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.
func (d *BakDnfItemInfoDao) Union(unions ...*gdb.Model) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Union(unions...)}
}

// UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.
func (d *BakDnfItemInfoDao) UnionAll(unions ...*gdb.Model) *BakDnfItemInfoDao {
	return &BakDnfItemInfoDao{M: d.M.Union(unions...)}
}
