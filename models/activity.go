package models

// AcActivity of the model
type AcActivity struct {
	ID        int64  `json:"id"`
	CreateAt  int    `json:"create_at"`
	UpdateAt  int    `json:"update_at"`
	NameZh    string `json:"name_zh"`
	ContentZh string `json:"content_zh"`
	NameEn    string `json:"name_en"`
	ContentEn string `json:"content_en"`
	AcType    int    `json:"ac_type"` // [0:其他 / 1:手数活动
	Status    int    `json:"status"`  // [0:未开启 / 1:进行中 / 2:提前下线 3: 已结束]
	PageURL   string `json:"page_url"`
	PicURL    string `json:"pic_url"`
	HandNum   int    `json:"hand_num"`
	Remark    string `json:"remark"`
	CreatedBy int64  `json:"created_by"`
	UpdatedBy int64  `json:"updated_by"`
}

// TableName ...
func (AcActivity) TableName() string {
	return "ac_activity"
}

// AcActivityDetail 活动条件
type AcActivityDetail struct {
	ID           int64   `json:"id"`
	CreateAt     int     `json:"create_at"`
	UpdateAt     int     `json:"update_at"`
	CreatedBy    int64   `json:"created_by"`
	UpdatedBy    int64   `json:"updated_by"`
	ActivityID   int64   `json:"activity_id"`
	PropFeeSeven int64   `json:"prop_fee_seven"` // 近7日道具费
	Insurance    int64   `json:"insurance"`      // 保险
	Bonus        int64   `json:"bonus"`          // 礼包奖金
	BonusNum     int     `json:"bonus_num"`      // 礼包份数
	Odds         float32 `json:"odds"`           // 中奖概率
	RealOdds     float32 `json:"real_odds"`      // 实际中奖概率
	BonusAll     int64   `json:"bonus_all"`      // 总奖金
}

// TableName ...
func (AcActivityDetail) TableName() string {
	return "ac_activity_detail"
}
