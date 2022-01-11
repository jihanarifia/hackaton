package dao

type DB interface {
	Health() error
}
