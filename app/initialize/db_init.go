package initialize

import (
	"fmt"

	licheeJs "github.com/issueye/lichee-js"
	"github.com/issueye/lichee/app/common"
	"github.com/issueye/lichee/app/model"
	"github.com/issueye/lichee/app/service"
	"github.com/issueye/lichee/global"
	"github.com/issueye/lichee/pkg/db"
	"github.com/issueye/lichee/utils"
)

func InitDB() {
	req := new(model.ReqQueryDbSource)
	list, err := service.NewDbSourceService().Query(req)
	if err != nil {
		common.Log.Errorf("获取数据库源列表失败，失败原因：%s", err.Error())
		return
	}

	for _, data := range list {

		cfg := &db.Config{
			Username: data.User,
			Password: data.Password,
			Host:     data.Host,
			Database: data.Database,
			Port:     int(data.Port),
			LogMode:  true,
		}

		d, err := common.GetDb(cfg, int(data.Type))
		if err != nil {
			common.Log.Errorf("连接数据库【%s】失败，失败原因：%s", data.Name, err.Error())
			continue
		}

		info := &global.DbInfo{
			Cfg:  cfg,
			Name: data.Name,
			DB:   d,
		}

		// 将数据库注册到JS虚拟机
		licheeJs.RegisterDB(fmt.Sprintf("db/%s", info.Name), info.DB)
		global.GdbMap[data.Name] = info
	}

	fmt.Printf("【%s】初始化数据库完成...\n", utils.Ltime{}.GetNowStr())
}

func InitLocalDb() {
	var err error
	common.LocalDb, err = db.InitSqlite(&db.Config{
		Path:    "data.db",
		LogMode: true,
	}, common.Log)
	if err != nil {
		fmt.Printf("初始化本地数据库失败，失败原因：%s", err.Error())
		return
	}

	// 自定检查、生成表
	common.LocalDb.AutoMigrate(
		&model.Job{},
		&model.Param{},
		&model.ParamArea{},
		&model.DbSource{},
		&model.User{},
	)
	fmt.Printf("【%s】初始化库表完成...\n", utils.Ltime{}.GetNowStr())

	InitAdminUser()
	InitSysArea()
}

func InitAdminUser() {
	u, err := service.NewUserService().GetById(10000)
	if err != nil {
		fmt.Printf("获取管理员信息失败，失败原因:%s\n", err.Error())
		return
	}

	if u.Id == 0 {
		err = service.NewUserService().Create(&model.User{
			Id:         10000,
			Account:    "admin",
			Name:       "管理员",
			Password:   "",
			Mark:       "系统生成的管理员账号",
			Enable:     1,
			CreateTime: utils.Ltime{}.GetNowStr(),
		})

		if err != nil {
			fmt.Printf("生成管理员信息失败，失败原因：%s\n", err.Error())
			return
		}
	}
}

func InitSysArea() {
	pa, err := service.NewParamService().GetAreaById(10000)
	if err != nil {
		fmt.Printf("查询参数域失败，失败原因：%s", err.Error())
		return
	}

	if pa.Id == 0 {
		service.NewParamService().CreateArea(&model.ParamArea{
			Id:         10000,
			Name:       "sys",
			Mark:       "系统参数域",
			CreateTime: utils.Ltime{}.GetNowStr(),
		})

		if err != nil {
			fmt.Printf("生成系统参数域失败，失败原因：%s\n", err.Error())
			return
		}
	}
}
