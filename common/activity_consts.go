package common

// ActivityType 活动类型常量
type ActivityType int

const (
	// Other 其他活动
	Other ActivityType = 0
	// HandNum 手数活动
	HandNum ActivityType = 1
)

// Int to int
func (ty ActivityType) Int() int {
	return int(ty)
}

// ActivityStatus 活动状态常量
type ActivityStatus int

const (
	// Not 未开始
	Not ActivityStatus = 0
	// Ongoing 进行中
	Ongoing ActivityStatus = 1
	// Offline 提前下线
	Offline ActivityStatus = 2
	// Closed 已关闭
	Closed ActivityStatus = 3
)

// Int to int
func (status ActivityStatus) Int() int {
	return int(status)
}
