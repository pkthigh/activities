package errs

import "fmt"

// ERRS 错误定义
type ERRS int

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

// Int to error code
func (errs ERRS) Int() int {
	return int(errs)
}

func (errs ERRS) Error() error {
	return fmt.Errorf("%v", string(errs))
}
