package service

import (
	"fmt"

	"github.com/issueye/lichee/app/common"
	"github.com/issueye/lichee/app/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return new(UserService)
}

// Save
// 写入数据
func (user UserService) Create(data *model.User) error {
	return common.LocalDb.Create(data).Error
}

// 修改用户信息
func (user UserService) Modify(data *model.User) error {
	return common.LocalDb.Model(&model.User{}).Where("id = ?", data.Id).Updates(data).Error
}

func (user UserService) FindUser(lu *model.LoginUser) (*model.User, error) {
	data := new(model.User)
	err := common.LocalDb.Model(&model.User{}).Where("account = ?", lu.Account).Find(data).Error
	return data, err
}

// Delete
// 写入数据
func (user UserService) Del(id int64) error {
	return common.LocalDb.Where("id = ?", id).Delete(&model.User{}).Error
}

// Delete
// 写入数据
func (user UserService) GetById(id int64) (*model.User, error) {
	data := new(model.User)
	err := common.LocalDb.Model(data).Where("id = ?", id).Find(data).Error
	return data, err
}

// Query
// 写入数据
func (user UserService) Query(req *model.ReqQueryUser) ([]*model.ResQueryUser, error) {
	list := make([]*model.User, 0)

	query := common.LocalDb.Model(&model.User{})

	if req.Name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%%%s%%", req.Name))
	}

	if req.Account != "" {
		query = query.Where("account = ?", req.Account)
	}

	if req.Mark != "" {
		query = query.Where("mark like ?", fmt.Sprintf("%%%s%%", req.Mark))
	}

	err := query.Find(&list).Error
	// 数据的条数
	req.Total = int64(len(list))

	resList := make([]*model.ResQueryUser, 0)

	for _, data := range list {
		res := new(model.ResQueryUser)
		res.Id = data.Id
		res.Name = data.Name
		res.Account = data.Account
		res.Mark = data.Mark
		res.Enable = data.Enable
		res.CreateTime = data.CreateTime
		res.LoginTime = data.LoginTime
		resList = append(resList, res)
	}

	return resList, err
}
