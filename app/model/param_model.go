package model

type Param struct {
	Id     int64  `json:"id" gorm:"column:id;type:bigint;comment:用户ID;primaryKey"` // id
	Name   string `json:"name" gorm:"column:name;type:varchar(100);comment:参数名称"`  // 参数名称
	AreaId int64  `json:"area_id" gorm:"column:area_id;type:bigint;comment:参数域id"` // 参数域id
	Area   string `json:"area" gorm:"column:area;type:varchar(100);comment:参数域"`   // 参数域
	Value  string `json:"value" gorm:"column:value;type:varchar(500);comment:参数值"` // 参数值
	Mark   string `json:"mark" gorm:"column:mark;type:varchar(500);comment:备注"`    // 备注
}

func (Param) TableName() string {
	return "param"
}

type ParamArea struct {
	Id   int64  `json:"id" gorm:"column:id;type:bigint;comment:id;primaryKey"` // id
	Name string `json:"name" gorm:"column:id;name:varchar(100);comment:域名称"`   // 域名称
	Mark string `json:"mark" gorm:"column:mark;type:varchar(500);comment:备注"`  // 备注
}

func (ParamArea) TableName() string {
	return "param_area"
}

type ReqCreateParam struct {
	Name   string `json:"name"`    // 参数名称
	AreaId int64  `json:"area_id"` // 参数域id
	Value  string `json:"value"`   // 参数值
	Mark   string `json:"mark"`    // 备注
}

type ReqModifyParam struct {
	Id     int64  `json:"id"`      // id
	Name   string `json:"name"`    // 参数名称
	AreaId int64  `json:"area_id"` // 参数域id
	Area   string `json:"area"`    // 参数域
	Value  string `json:"value"`   // 参数值
	Mark   string `json:"mark"`    // 备注
}

type ReqQueryParam struct {
	Name   string `json:"name" form:"name"`       // 参数名称
	AreaId int64  `json:"area_id" form:"area_id"` // 参数域id
	PageInfo
}

type ReqCreateArea struct {
	Name string `json:"name"` // 域名称
	Mark string `json:"mark"` // 备注
}

type ReqQueryArea struct {
	Name string `json:"name" form:"name"` // 参数名称
	PageInfo
}

type ReqModifyArea struct {
	Id   int64  `json:"id"`   // id
	Name string `json:"name"` // 域名称
	Mark string `json:"mark"` // 备注
}
