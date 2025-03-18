package domain

import "github.com/urcop/emotionalTracker/domain/repositories"

type Connection interface {
	Example() repositories.Example
}
