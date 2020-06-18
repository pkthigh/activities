package common

// REDIS redis area
type REDIS int

const (
	// Default 默认
	Default REDIS = iota
	// ItemRecordStore 道具记录存储
	ItemRecordStore
	// HandOverRecordStore 每手记录存储
	HandOverRecordStore
	// InsuranceRecordStore 保险记录存储
	InsuranceRecordStore
)

// Int to int
func (redis REDIS) Int() int {
	return int(redis)
}
