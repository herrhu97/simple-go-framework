package session

import (
	"database/sql"
	"github.com/herrhu97/simple-go-framework/orm/clause"
	"strings"

	"github.com/herrhu97/simple-go-framework/log"
	"github.com/herrhu97/simple-go-framework/orm/dialect"
	"github.com/herrhu97/simple-go-framework/orm/schema"
)

// Session 用于实现与数据库的交互,执行sql语句
type Session struct {
	db       *sql.DB         // db.Open()返回的对象指针
	tx       *sql.Tx         // 事务
	dialect  dialect.Dialect // go类型于具体数据库类型转换
	refTable *schema.Schema  // 数据库中具体的表的类型与Go类型的转换
	clause   clause.Clause   // 存储暂时的临时sql子句，拼凑大sql
	sql      strings.Builder // 拼接sql用的sb对象
	sqlVars  []interface{}   // 替代占位符的具体参数
}

// CommonDB is a minimal function set of db
type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

// DB returns tx if a tx begins. otherwise return *sql.DB
func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// Raw 将pre-sql与sql变量填充到Session中
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	//在传递时给可变参数变量后面添加...，这样就可以将切片中的元素进行传递，而不是传递可变参数变量本身。
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
