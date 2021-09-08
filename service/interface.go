package service

import "github.com/vembry/web-api-go/repository"

type HealthServiceParam struct {
	IHealthRepo *repository.IHealthRepo
}

type IHealthService interface {
	Check() bool
	Validate(id int) bool
}
