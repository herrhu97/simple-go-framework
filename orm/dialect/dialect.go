package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

// Dialect 实现sql方言（类型转换、原生判别表是否存在sql）
type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    // 将 Go 语言的类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) // 返回某个表是否存在的 SQL 语句，参数是表名(table)
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
