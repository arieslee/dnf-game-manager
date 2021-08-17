package snowflake

// 位数更短的snowflake id
import (
	"github.com/yitter/idgenerator-go/idgen"
	"sync"
)

type GenId struct {
}

var singleton sync.Once
var instance *GenId

// GetInstance 单例
func GetInstance() *GenId {
	singleton.Do(func() {
		var options = idgen.NewIdGeneratorOptions(0)
		idgen.SetIdGenerator(options)
		instance = &GenId{}
	})
	return instance
}
func (g *GenId) NextId() uint64 {
	return idgen.NextId()
}
