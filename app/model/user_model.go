package model

import "time"

type User struct {
	Id         int64     `json:"id" gorm:"column:id;type:bigint;comment:用户ID;primaryKey"`              // 用户ID
	Account    string    `json:"account" gorm:"column:account;type:varchar(100);comment:登录名"`          // 登录名
	Name       string    `json:"name" gorm:"column:name;type:varchar(100);comment:用户名"`                // 用户名
	Password   string    `json:"password" gorm:"password:account;type:varchar(100);comment:用户密码"`      // 用户密码
	Mark       string    `json:"mark" gorm:"column:mark;type:varchar(500);comment:备注"`                 // 备注
	Enable     int64     `json:"enable" gorm:"column:enable;type:int;comment:启用"`                      // 启用
	LoginTime  time.Time `json:"login_time" gorm:"column:login_time;type:varchar(100);comment:登录时间"`   // 登录时间
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:varchar(100);comment:创建时间"` // 创建时间
}

func (User) TableName() string {
	return "user"
}

type ReqCreateUser struct {
	Account  string `json:"account"`  // 登录名
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 用户密码
	Mark     string `json:"mark"`     // 备注
}

type ReqModifyUser struct {
	Id       int64  `json:"id"`       // 用户ID
	Account  string `json:"account"`  // 登录名
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 用户密码
	Mark     string `json:"mark"`     // 备注
}

type ReqQueryUser struct {
	Account string `json:"account" form:"account"` // 登录名
	Name    string `json:"name" form:"name"`       // 用户名
	Mark    string `json:"mark" form:"mark"`       // 备注
	PageInfo
}

type ResQueryUser struct {
	Id         int64  `json:"id"`          // 用户ID
	Account    string `json:"account"`     // 登录名
	Name       string `json:"name"`        // 用户名
	Mark       string `json:"mark"`        // 备注
	Enable     int64  `json:"enable"`      // 启用
	LoginTime  string `json:"login_time"`  // 登录时间
	CreateTime string `json:"create_time"` // 创建时间
}

type LoginUser struct {
	Account  string `json:"account"`  // 登录名
	Password string `json:"password"` // 用户密码
}
