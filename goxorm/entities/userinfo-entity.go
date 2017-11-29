package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int        `xorm:"int(10) not null pk autoincr 'uid'"` //语义标签
	UserName   string     `xorm:"varchar(64) null default null 'username'"`
	DepartName string     `xorm:"varchar(64) null default null 'departname'"`
	CreateAt   *time.Time `xorm:"date null default null 'created'"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}
