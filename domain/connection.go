package domain

import "github.com/urcop/emotionalTracker/domain/repositories"

type Connection interface {
	User() repositories.User
}
