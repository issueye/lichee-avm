package service

import (
	"fmt"

	"github.com/issueye/lichee/app/common"
	"github.com/issueye/lichee/app/model"
)

type DbSourceService struct{}

func NewDbSourceService() *DbSourceService {
	return new(DbSourceService)
}

// Create
// 写入数据
func (dbSource DbSourceService) Create(data *model.DbSource) error {
	return common.LocalDb.Create(data).Error
}

// Delete
// 写入数据
func (dbSource DbSourceService) Del(id int64) error {
	return common.LocalDb.Where("id = ?", id).Delete(&model.DbSource{}).Error
}

// Delete
// 写入数据
func (dbSource DbSourceService) GetById(id int64) (*model.DbSource, error) {
	data := new(model.DbSource)
	err := common.LocalDb.Model(&model.DbSource{}).Where("id = ?", id).Find(data).Error
	return data, err
}

// Query
// 写入数据
func (dbSource DbSourceService) Query(req *model.ReqQueryDbSource) ([]*model.DbSource, error) {
	list := make([]*model.DbSource, 0)
	query := common.LocalDb.Model(&model.DbSource{})
	if req.Name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%%%s%%", req.Name))
	}
	err := query.Find(&list).Error
	if err != nil {
		return nil, err
	}
	req.Total = int64(len(list))
	return list, nil
}
