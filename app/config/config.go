package config

import (
	"github.com/issueye/lichee/app/common"
	"github.com/issueye/lichee/app/model"
	"github.com/issueye/lichee/app/service"
	"github.com/issueye/lichee/utils"
	"github.com/spf13/cast"
)

type Result struct {
	value  any    // 参数值
	name   string // 参数名称
	mark   string // 参数说明
	areaId int64  // 参数域
}

func (r *Result) Name() string {
	return r.name
}

func (r *Result) String() string {
	return cast.ToString(r.value)
}

func (r *Result) Int64() int64 {
	return cast.ToInt64(r.value)
}

func (r *Result) Int() int {
	return cast.ToInt(r.value)
}

func (r *Result) Float() float64 {
	return cast.ToFloat64(r.value)
}

func (r *Result) Bool() bool {
	return cast.ToBool(r.value)
}

// 获取参数
func GetSysParam(name string) *Result {
	r := new(Result)
	data := new(model.Param)
	err := common.LocalDb.Model(&model.Param{}).Where("name = ?", name).Find(data).Error
	if err != nil {
		r.areaId = 0
		r.mark = ""
		r.name = ""
		r.value = ""
		return r
	}

	r.name = data.Name
	r.value = data.Value
	r.areaId = data.AreaId
	r.mark = data.Mark
	return r
}

// 获取参数
func SetSysParam(name, value, mark string) error {
	data := new(model.Param)
	data.Id = utils.Lid{}.GenID()
	data.Name = name
	data.Mark = mark
	data.Value = value
	data.AreaId = common.SYS_AREA
	data.Area = common.SYS_AREA_NAME
	return service.NewParamService().Save(data)
}
