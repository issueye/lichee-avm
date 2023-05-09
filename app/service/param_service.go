package service

import (
	"github.com/issueye/lichee/app/common"
	"github.com/issueye/lichee/app/model"
)

type ParamService struct{}

func NewParamService() *ParamService {
	return new(ParamService)
}

// Save
// 写入数据
func (p ParamService) Save(data *model.Param) error {
	// 如果存在数据则更新数据，没有则写入数据
	count := int64(0)
	err := common.LocalDb.
		Model(&model.Param{}).
		Where("id = ?", data.Id).
		Count(&count).
		Error

	if err != nil {
		return err
	}

	if count > 0 {
		err = common.LocalDb.
			Model(&model.Param{}).
			Updates(data).
			Error

		return err
	} else {
		err = common.LocalDb.
			Create(data).
			Error

		return err
	}
}

func (p ParamService) GetAreaById(id int64) (*model.ParamArea, error) {
	data := new(model.ParamArea)
	err := common.LocalDb.
		Model(&model.ParamArea{}).
		Where("id = ?", id).
		Find(data).
		Error

	return data, err
}

func (p ParamService) GetAreaList(req *model.ReqQueryArea) ([]*model.ParamArea, error) {
	list := make([]*model.ParamArea, 0)

	query := common.LocalDb.
		Model(&model.ParamArea{})

	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}

	err := query.
		Find(&list).
		Error

	req.Total = int64(len(list))
	return list, err
}

// DelArea
// 写入数据
func (p ParamService) DelArea(id int64) error {
	return common.LocalDb.
		Where("id = ?", id).
		Delete(&model.ParamArea{}).
		Error
}

// Delete
// 写入数据
func (p ParamService) GetById(id int64) (*model.Param, error) {
	data := new(model.Param)
	err := common.LocalDb.
		Where("id = ?", id).
		Find(data).
		Error
	return data, err
}

// ModifyArea
// 写入数据
func (p ParamService) ModifyArea(data *model.ParamArea) error {
	return common.LocalDb.
		Model(&model.ParamArea{}).
		Where("id = ?", data.Id).
		Updates(data).Error
}

// DelParam
// 写入数据
func (p ParamService) DelParam(area, id int64) error {
	return common.LocalDb.
		Where("area_id = ?", area).
		Where("id = ?", id).
		Delete(&model.Param{}).
		Error
}

// ModifyParam
// 修改数据
func (p ParamService) ModifyParam(data *model.Param) error {
	return common.LocalDb.
		Model(&model.Param{}).
		Where("area = ?", data.Area).
		Where("id = ?", data.Id).
		Updates(data).Error
}

// Query
// 查询数据
func (p ParamService) Query(req *model.ReqQueryParam) ([]*model.Param, error) {
	list := make([]*model.Param, 0)

	query := common.LocalDb.
		Model(&model.Param{})

	if req.Name != "" {
		query = query.
			Where("name = ?", req.Name)
	}

	err := query.Find(&list).Error

	// 数据的条数
	req.Total = int64(len(list))
	return list, err
}

func (p ParamService) CreateArea(area *model.ParamArea) error {
	return common.LocalDb.
		Create(area).
		Error
}
