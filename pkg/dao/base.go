package dao

import (
	log "github.com/Lilymz/table-migration/v2/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var connectErr error

type DbOperator struct {
}

var SingletonOperator = &DbOperator{}

// Connect 数据库配置初始化
// todo 如果error 可以考虑重连 goroutine +1
func (DbOperator) Connect(config mysql.Config, gormConfig *gorm.Config) bool {
	db, connectErr = gorm.Open(mysql.New(config), gormConfig)
	if connectErr != nil {
		log.DaoLog.Error("mysql connect error err", connectErr)
		return false
	}
	return true
}
func (DbOperator) GetDB() *gorm.DB {
	return db
}
