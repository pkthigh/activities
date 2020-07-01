package common

// REDIS redis area
type REDIS int

const (
	// ItemRecordStore 道具记录存储
	ItemRecordStore REDIS = 1
	// HandOverRecordStore 每手记录存储
	HandOverRecordStore REDIS = 2
	// InsuranceRecordStore 保险记录存储
	InsuranceRecordStore REDIS = 3

	// DailyStatistics 每日统计
	DailyStatistics REDIS = 15
)

// Int to int
func (redis REDIS) Int() int {
	return int(redis)
}
