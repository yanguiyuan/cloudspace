// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	Username   string    `gorm:"column:username;not null" json:"username"`
	Password   string    `gorm:"column:password;not null" json:"password"`
	Gender     string    `gorm:"column:gender;not null;default:unknown" json:"gender"`
	Email      string    `gorm:"column:email;not null" json:"email"`
	Role       string    `gorm:"column:role;not null;default:user" json:"role"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
