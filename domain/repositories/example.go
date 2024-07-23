package repositories

import "github.com/FoodMoodOTG/examplearch/domain/models"

type Example interface {
	Insert(example *models.Example) (string, error)
	GetExample(id string) (*models.Example, error)
	All() ([]*models.Example, error)
}
