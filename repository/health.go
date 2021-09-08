package repository

import "fmt"

func InitRepository(param HealthRepoParam) IHealthRepo {
	fmt.Println("initializing health repository")
	return &param
}

func (repo HealthRepoParam) CheckHealth() bool {
	return true
}

func (repo HealthRepoParam) ValidateHealth(id int) bool {
	return (id%2 == 0) || (id%3 == 0) || (id%5 == 0) || (id%7 == 0)
}
