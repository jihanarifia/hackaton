package postgres

import (
	"hackaton/pkg/dao"

	"github.com/jinzhu/gorm"
)

type DB struct {
	conn *gorm.DB
}

func NewDB(conn *gorm.DB) dao.DB {
	return &DB{conn: conn}
}

// Health check ping postgres
func (db *DB) Health() error {
	return db.conn.DB().Ping()
}
