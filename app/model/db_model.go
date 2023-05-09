package model

type DbSource struct {
	Id       int64  `json:"id" gorm:"column:id;type:bigint;comment:ID;primaryKey"`                                // ID
	Name     string `json:"name" gorm:"column:name;type:varchar(100);comment:名称"`                                 // 名称
	Host     string `json:"host" gorm:"column:host;type:varchar(100);comment:地址"`                                 // 地址
	Port     int64  `json:"port" gorm:"column:port;type:bigint;comment:端口号"`                                      // 端口号
	Database string `json:"database" gorm:"column:database;type:varchar(50);comment:数据库"`                         // 数据库
	User     string `json:"user" gorm:"column:user;type:varchar(100);comment:账户"`                                 // 账号
	Password string `json:"password" gorm:"column:password;type:varchar(100);comment:密码"`                         // 密码
	Type     int64  `json:"type" gorm:"column:type;type:int;comment:数据库类型 0 sqlserver 1 mysql 2 oracle 3 sqlite"` // 类型 0 sqlserver 1 mysql 2 oracle 3 sqlite
	Path     string `json:"path" gorm:"column:path;type:varchar(500);comment:数据库路径"`                              // 数据库路径
	Mark     string `json:"mark" gorm:"column:mark;type:varchar(500);comment:备注"`                                 // 备注
}

func (DbSource) TableName() string {
	return "db_source"
}

// ReqCreateDbSource
// 创建
type ReqCreateDbSource struct {
	Name     string `json:"name" binding:"required" label:"名称" example:"测试SQLSERVER_001"` // 名称
	Host     string `json:"host" binding:"required" label:"服务地址" example:"."`             // 地址
	Port     int64  `json:"port" binding:"required" label:"端口号" example:"1433"`           // 端口号
	Database string `json:"database" binding:"required" label:"数据库名称" example:"TEST"`     // 数据库
	User     string `json:"user" binding:"required" label:"账户" example:"sa"`              // 账号
	Password string `json:"password" binding:"required" label:"密码" example:"123456"`      // 密码
	Type     int64  `json:"type" label:"数据库类型" example:"0"`                               // 类型
	Path     string `json:"path" label:"数据库路径 SQLITE时需要" example:"db\test.db"`            // 类型
	Mark     string `json:"mark"`                                                         // 备注
}

// 修改
type ReqModifyDbSource struct {
	Id       int64  `json:"id"  binding:"required" label:"id" example:"10000"`            // ID
	Name     string `json:"name" binding:"required" label:"名称" example:"测试SQLSERVER_001"` // 名称
	Host     string `json:"host" binding:"required" label:"服务地址" example:"."`             // 地址
	Port     int64  `json:"port" binding:"required" label:"端口号" example:"1433"`           // 端口号
	Database string `json:"database" binding:"required" label:"数据库名称" example:"TEST"`     // 数据库
	User     string `json:"user" binding:"required" label:"账户" example:"sa"`              // 账号
	Password string `json:"password" binding:"required" label:"密码" example:"123456"`      // 密码
	Type     int64  `json:"type" label:"数据库类型" example:"0"`                               // 类型
	Path     string `json:"path" label:"数据库路径 SQLITE时需要" example:"db\test.db"`            // 类型
	Mark     string `json:"mark"`                                                         // 备注
}

// ReqQueryDbSource
// 请求 查询
type ReqQueryDbSource struct {
	Name string `json:"name" form:"name"` // 名称
	PageInfo
}
