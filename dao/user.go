package dao

import (
	"gin_demo/pkg/gorm"
	"time"
)

type User struct {
	Id        int64     `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Account   string    `json:"account" gorm:"column:account"`
	Password  string    `json:"password" gorm:"column:password"`
	Age       uint8     `json:"age" gorm:"column:age"`
	AvatarUrl string    `json:"avatar_url" gorm:"column:avatar_url"`
	Status    uint8     `json:"status" gorm:"column:status"`
	Birth     string    `json:"birth" gorm:"column:birth"`
	Sex       int       `json:"sex" gorm:"column:sex"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Create(user *User) error {
	if err := gorm.Client("main").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindOne(condition map[string]interface{}) (*User, error) {
	result := &User{}
	if err := gorm.Client("main").Where(condition).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u *User) Find(condition map[string]interface{}) ([]*User, error) {
	var result []*User
	if err := gorm.Client("main").Where(condition).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
