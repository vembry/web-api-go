package service

import (
	"fmt"
)

func InitService(param HealthServiceParam) IHealthService {
	fmt.Println("initializing health service")
	return &param
}

func (serv HealthServiceParam) Check() bool {
	return (*serv.IHealthRepo).CheckHealth()
}

func (serv HealthServiceParam) Validate(id int) bool {
	return (*serv.IHealthRepo).ValidateHealth(id)
}
