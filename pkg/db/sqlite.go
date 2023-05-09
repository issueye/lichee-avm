package db

import (
	"time"

	sqlite "github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func InitSqlite(cfg *Config, log *zap.SugaredLogger) (*gorm.DB, error) {
	newLogger := glogger.New(
		Writer{
			log:    log,
			BPrint: cfg.LogMode,
		},
		glogger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  glogger.Info,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	var l glogger.Interface

	if cfg.LogMode {
		l = newLogger.LogMode(glogger.Info)
	} else {
		l = newLogger.LogMode(glogger.Silent)
	}

	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   l,
	})

	if err != nil {
		log.Panicf("连接数据库异常: %v", err)
		return nil, err
	}

	db.Debug()
	log.Infof("初始化sqlite数据库完成! dsn: %s", cfg.Path)

	return db, nil
}
