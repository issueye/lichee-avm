package service

import (
	"github.com/issueye/lichee/app/common"
	"github.com/issueye/lichee/app/model"
	"github.com/issueye/lichee/utils"
)

type JobService struct{}

func NewJobService() *JobService {
	return new(JobService)
}

// Save
// 写入数据
func (job JobService) Save(data *model.Job) error {
	return common.LocalDb.Create(data).Error
}

// Delete
// 写入数据
func (job JobService) Del(id int64) error {
	return common.LocalDb.Model(&model.Job{}).Where("id = ?", id).Error
}

// Delete
// 写入数据
func (job JobService) GetById(id int64) (*model.Job, error) {
	data := new(model.Job)
	err := common.LocalDb.Model(&model.Job{}).Where("id = ?", id).Find(data).Error
	return data, err
}

// Query
// 写入数据
func (job JobService) Query(req *model.ReqQueryJob) ([]*model.ResQueryJob, error) {
	list := make([]*model.Job, 0)

	err := common.LocalDb.Model(&model.Job{}).Where(&list).Error
	// 数据的条数
	req.Total = int64(len(list))

	resList := make([]*model.ResQueryJob, 0)
	for _, data := range list {
		// 获取参数域
		pa, err := NewParamService().GetAreaById(data.AreaId)
		if err != nil {
			common.Log.Errorf("【%d】未找到参数域信息", data.AreaId)
			continue
		}

		res := new(model.ResQueryJob)
		res.Id = data.Id
		res.Name = data.Name
		res.Expr = data.Expr
		res.Mark = data.Mark
		res.Enable = data.Enable
		res.Path = data.Path
		res.AreaId = pa.Id
		res.Area = pa.Name
		res.CreateTime = data.CreateTime.Format(utils.FormatDateTimeMs)
		resList = append(resList, res)
	}

	return resList, err
}
