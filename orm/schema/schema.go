package schema

import (
	"go/ast"
	"reflect"

	"github.com/herrhu97/simple-go-framework/orm/dialect"
)

// Field represents a column of database
type Field struct {
	Name string // 值名
	Type string // 值类型
	Tag  string // 额外的约束条件，比如"primary key"
}

// Schema represents a table of database，用于实现对象(object)和表(table)的转换
type Schema struct {
	Model      interface{}       // 具体对象
	Name       string            // 结构体的名称做表名
	Fields     []*Field          // 对象的所有Fields
	FieldNames []string          // Fields的所有name
	fieldMap   map[string]*Field // Field的name到Field的映射
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("orm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
