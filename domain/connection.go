package domain

import "github.com/FoodMoodOTG/examplearch/domain/repositories"

type Connection interface {
	Example() repositories.Example
}
