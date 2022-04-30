package dbx

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"reflect"
	"time"
)

var globalDB *gorm.DB

func getDefaultGormConfig() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: false,
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
	}
}

//初始化gorm链接
func InitDb(cfg DbConfig, plugin ...gorm.Plugin) {
	//dsn := "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.Dsn, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}), getDefaultGormConfig())

	if err != nil {
		log.Fatal("mysql 链接错误", err)
	}
	for _, p := range plugin {
		db.Use(p)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("mysql 链接错误", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	globalDB = db
}

//事务key
type ctxTransactionKey struct {
}

func GetDb(ctx context.Context) *gorm.DB {
	iFace := ctx.Value(ctxTransactionKey{})

	if iFace != nil {
		tx, ok := iFace.(*gorm.DB)
		if !ok {
			log.Panicf("unexpect context value type: %s", reflect.TypeOf(tx))
			return nil
		}

		return tx
	}

	return globalDB.WithContext(ctx)
}

func Transaction(ctx context.Context, fc func(ctx context.Context) error) error {
	return globalDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		newCtx := context.WithValue(ctx, ctxTransactionKey{}, tx)
		return fc(newCtx)
	})
}
