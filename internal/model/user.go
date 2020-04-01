package model

import (
	"pika/driver"

	"github.com/jinzhu/gorm"
)

//User - User Model struct
type User struct {
	ID   string
	Name string
}

var user *User

func UserModel() *User {
	return user
}

func (u *User) TableName() string {
	return `t_user`
}

func (u *User) findByMap(wheremaps map[string]interface{}) (*User, error) {
	var user User
	if err := driver.DB.Table(u.TableName()).Where(wheremaps).Find(&user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &user, nil
}

//Save - save data
func (u *User) Save() error {
	return nil
}

//FindByID - find data
func (u *User) FindByID(userID uint64) (*User, error) {
	return u.findByMap(map[string]interface{}{
		"id": userID,
	})
}
