package domain

type Context interface {
	Make() Context

	Services() Services
	Connection() Connection
}
