package service

import "github.com/vembry/web-api-template-go/repository"

type HealthServiceParam struct {
	IHealthRepo *repository.IHealthRepo
}

type IHealthService interface {
	Check() bool
	Validate(id int) bool
}
