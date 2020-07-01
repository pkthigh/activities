package common

// Subject 订阅常量
type Subject string

const (
	// PkcHandOverNewSubject 微币每手结束信息
	PkcHandOverNewSubject Subject = "pkc_hand_over"
	// ItemSubject 道具使用
	ItemSubject Subject = "item"
)

func (subject Subject) String() string {
	return string(subject)
}
