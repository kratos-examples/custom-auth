package data

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/pkg/models"
	"github.com/yylego/must"
	"github.com/yylego/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	loggergorm "gorm.io/gorm/logger"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	db *gorm.DB // GORM database connection instance // GORM 数据库连接实例
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	dsn := must.Nice(c.Database.Source)                     // Get database DSN from config // 从配置获取数据库DSN
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{ // Open SQLite database connection // 打开SQLite数据库连接
		Logger: loggergorm.Default.LogMode(loggergorm.Info), // Enable GORM info logging // 启用GORM信息日志
	}))

	must.Done(db.AutoMigrate(&models.Admin{}, &models.Token{})) // Auto migrate database schema // 自动迁移数据库表结构

	mockAdmin1(db) // Create mock admin1 and token data // 创建模拟管理员1及其令牌数据
	mockAdmin2(db) // Create mock admin2 and token data // 创建模拟管理员2及其令牌数据

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		must.Done(rese.P1(db.DB()).Close()) // Close database connection on cleanup // 清理时关闭数据库连接
	}
	return &Data{
		db: db,
	}, cleanup, nil
}

// DB returns the database instance
//
// DB 返回数据库实例
func (d *Data) DB() *gorm.DB {
	return d.db
}

// mockAdmin1 creates test admin with username "abc" and access token
//
// mockAdmin1 创建测试管理员（用户名 "abc"）及其访问令牌
func mockAdmin1(db *gorm.DB) {
	must.Done(db.Create(&models.Admin{
		Username: "abc",
		Password: "123",
		Mailbox:  "",
		Status:   0,
	}).Error)
	must.Done(db.Create(&models.Token{
		AdminID:   1,
		Token:     "95d9fda7f675444d9acc3c8225dbf7de",
		Type:      "access",
		ExpiresAt: time.Now().UTC().Add(30 * time.Minute), // Token expires in 30 minutes // 令牌30分钟后过期
	}).Error)
}

// mockAdmin2 creates test admin with username "xyz" and access token
//
// mockAdmin2 创建测试管理员（用户名 "xyz"）及其访问令牌
func mockAdmin2(db *gorm.DB) {
	must.Done(db.Create(&models.Admin{
		Username: "xyz",
		Password: "456",
		Mailbox:  "",
		Status:   0,
	}).Error)
	must.Done(db.Create(&models.Token{
		AdminID:   2,
		Token:     "46421ed7de4a4fcc888ff84541defbc3",
		Type:      "access",
		ExpiresAt: time.Now().UTC().Add(30 * time.Minute), // Token expires in 30 minutes // 令牌30分钟后过期
	}).Error)
}
