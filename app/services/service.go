package services

import "api/app/services/datastruct"

type Service interface {
	HealthCheck() datastruct.HealthCheck
}

func InitService() Service{
	return mainService{}
}

type mainService struct {}

func (mainService) HealthCheck() datastruct.HealthCheck {
	return datastruct.HealthCheck{}
}
