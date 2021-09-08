package repository

type HealthRepoParam struct {
}
type IHealthRepo interface {
	CheckHealth() bool
	ValidateHealth(id int) bool
}
