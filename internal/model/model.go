package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"

	"github.com/pudongping/gin-blog-service/global"
	"github.com/pudongping/gin-blog-service/pkg/setting"
)

type Model struct {
	// id  int(10) unsigned is_nullable NO
	Id uint32 `gorm:"primary_key" json:"id"`
	// 创建时间  int(10) unsigned is_nullable YES
	CreatedOn uint32 `json:"created_on"`
	// 创建人  varchar(100) is_nullable YES
	CreatedBy string `json:"created_by"`
	// 修改时间  int(10) unsigned is_nullable YES
	ModifiedOn uint32 `json:"modified_on"`
	// 修改人  varchar(100) is_nullable YES
	ModifiedBy string `json:"modified_by"`
	// 删除时间  int(10) unsigned is_nullable YES
	DeletedOn uint32 `json:"deleted_on"`
	// 是否删除 0 为未删除、1 为已删除  tinyint(3) unsigned is_nullable YES
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	config := mysql.New(mysql.Config{
		DSN: dsn,
	})

	var level logger.LogLevel
	// 允许我们在命令行里查看请求的 sql 信息
	// Silent —— 静默模式，不打印任何信息
	// Error —— 发生错误了才打印
	// Warn —— 发生警告级别以上的错误才打印
	// Info —— 打印所有信息，包括 SQL 语句
	if global.ServerSetting.RunMode == "debug" {
		level = logger.Info
	} else {
		level = logger.Error
	}

	db, err := gorm.Open(config, &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		return nil, err
	}

	// 命令行打印数据库请求的信息
	// *gorm.DB 对象的 DB() 方法，可以直接获取到 database/sql 包里的 *sql.DB 对象
	sqlDB, _ := db.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(databaseSetting.ConnMaxLifetime) * time.Second)

	// 创建和维护数据表结构
	migration(db)

	return db, nil
}

// migration 自动迁移
func migration(db *gorm.DB) {
	// db.AutoMigrate()
}
