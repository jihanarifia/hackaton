package postgres

import "hackaton/pkg/model"

func (db *DB) CreateUser(user model.User) error {
	return db.conn.Create(&user).Error
}

func (db *DB) GetUsers() (users []model.User, err error) {
	err = db.conn.Find(&users).Error
	return
}
