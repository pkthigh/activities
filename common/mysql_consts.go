package common

// MYSQL 数据库名常量
type MYSQL string

const (
	// ActivityDsn 活动数据库
	ActivityDsn MYSQL = "activity_dsn"
)

func (mysql MYSQL) String() string {
	return string(mysql)
}
