package model

type Job struct {
	Id         int64  `json:"id" gorm:"column:id;type:bigint;comment:任务ID;primaryKey;"`             // 任务ID
	Name       string `json:"name" gorm:"column:name;type:varchar(200);comment:任务名称"`               // 名称
	Expr       string `json:"expr" gorm:"column:expr;type:varchar(150);comment:时间表达式"`              // 时间表达式
	Mark       string `json:"mark" gorm:"column:mark;type:varchar(500);comment:备注"`                 // 备注
	Enable     bool   `json:"enable" gorm:"column:enable;type:int;comment:启用"`                      // 状态
	Path       string `json:"path" gorm:"column:path;type:varchar(300);comment:脚本路径"`               // 脚本路径
	AreaId     int64  `json:"area_id" gorm:"column:area_id;type:bigint;comment:参数域ID"`              // 参数域ID
	CreateTime string `json:"create_time" gorm:"column:create_time;type:varchar(100);comment:创建时间"` // 创建时间
}

func (Job) TableName() string {
	return "job"
}

// ReqCreateJob
// 请求 创建定时任务结构体
type ReqCreateJob struct {
	Name   string `json:"name" binding:"required" label:"任务名称" example:"测试定时任务"`         // 名称
	Expr   string `json:"expr" binding:"required" label:"时间表达式" example:"0/5 * * * * ?"` // 时间表达式
	Mark   string `json:"mark" label:"任务说明" example:"每五秒执行一次脚本"`                         // 备注
	Path   string `json:"path" binding:"required" label:"脚本路径" example:"test.js"`        // 脚本路径
	AreaId int64  `json:"area_id" binding:"required" label:"参数域" example:"10000"`        // 参数域ID
}

// ReqModifyJob
// 请求 修改定时任务结构体
type ReqModifyJob struct {
	Id     int64  `json:"id" binding:"required" label:"任务ID"`                     // 任务ID
	Name   string `json:"name" binding:"required" label:"任务名称"`                   // 名称
	Expr   string `json:"expr" binding:"required" label:"时间表达式"`                  // 时间表达式
	Mark   string `json:"mark" label:"任务说明"`                                      // 备注
	Path   string `json:"path" binding:"required" label:"脚本路径"`                   // 脚本路径
	AreaId int64  `json:"area_id" binding:"required" label:"参数域" example:"10000"` // 参数域ID
}

// ReqQueryJob
// 请求 查询定时任务结构体
type ReqQueryJob struct {
	Name string `json:"name" form:"name"` // 名称
	PageInfo
}

type ResQueryJob struct {
	Id         int64  `json:"id"`          // 任务ID
	Name       string `json:"name"`        // 名称
	Expr       string `json:"expr"`        // 时间表达式
	Mark       string `json:"mark"`        // 备注
	Enable     bool   `json:"enable"`      // 状态
	Path       string `json:"path"`        // 脚本路径
	Area       string `json:"area"`        // 参数域
	AreaId     int64  `json:"area_id"`     //
	CreateTime string `json:"create_time"` // 创建时间
}
