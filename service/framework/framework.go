package framework

import "activities/common"

// Activity 活动抽象类
type Activity interface {
	ID() int64
	Type() common.ActivityType
	Status() common.ActivityStatus

	// Judge 判断玩家是否满足参与活动的条件
	Judge(userid string) bool

	// Participate 参与活动
	Participate(userid string)
}
