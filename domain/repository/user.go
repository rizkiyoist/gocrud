package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rizkiyoist/gocrud/domain/model"
)

type userDB struct {
	db *gorm.DB
}

type User interface {
	Get(users *model.User) (*model.User, error)
}

func (u userDB) Find(ID int64) (*model.User, error) {
	m := new(model.User)
	m.ID = ID
	res := u.db.Find(&m).Scan(&m)

	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}
