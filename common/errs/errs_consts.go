package errs

import "fmt"

// ERRS 错误定义
type ERRS int

// 公告错误
const (
	// Default 没有错误
	Default ERRS = iota
	// MissingParameter 缺失参数
	MissingParameter
	// IllegalParameter 非法参数
	IllegalParameter
	// DatabaseError 数据库错误
	DatabaseError
)

// 业务错误
const (
	// UnknownActivityType 未知活动类型
	UnknownActivityType ERRS = iota + 1000
	// ThisActivityTypeNotCurrentlySet 当前未设置这个活动类型
	ThisActivityTypeNotCurrentlySet
	// ActivityIDNotMatch 活动ID不匹配
	ActivityIDNotMatch
	// ActivityStatusNotMatch 活动状态不匹配
	ActivityStatusNotMatch
)

// Int to error code
func (errs ERRS) Int() int {
	return int(errs)
}

func (errs ERRS) Error() error {
	return fmt.Errorf("%v", string(errs))
}
