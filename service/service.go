package service

import "activities/service/framework"

// Service 活动服务实现
type Service struct {
	Activities map[string]framework.Activities
}

// ActivitiesList 活动列表
func (srv *Service) ActivitiesList() {}

// ParticipateActivities 参与活动
func (srv *Service) ParticipateActivities() {}
