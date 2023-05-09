package initialize

import (
	"fmt"

	"github.com/issueye/lichee/app/common"
	"github.com/issueye/lichee/pkg/logger"
	"github.com/issueye/lichee/utils"
)

func InitLogger() {
	cfg := new(logger.Config)
	cfg.Path = "log"
	cfg.Level = -1
	cfg.MaxAge = 10
	cfg.MaxBackups = 10
	cfg.MaxSize = 5
	cfg.Compress = true
	common.Log, common.Logger = logger.InitLogger(cfg)
	fmt.Printf("【%s】初始化日志完成...\n", utils.Ltime{}.GetNowStr())
}
