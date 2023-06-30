package dao

import (
	"errors"
	log "github.com/Lilymz/table-migration/v2/pkg/config"
)

// GetPrimary 查询主键
func GetPrimary(table string) (string, error) {
	var (
		primary string
		err     error
	)
	defer func() {
		message := recover()
		if nil != message {
			log.DaoLog.Error("table:"+table+" get primary key failed msg:", message)
			err = errors.Unwrap(message.(error))
		}
	}()
	db := SingletonOperator.GetDB()
	db.Raw("SELECT COLUMN_NAME AS Key_name FROM INFORMATION_SCHEMA.COLUMNS WHERE   TABLE_NAME = '" + table + "' AND COLUMN_KEY = 'PRI' LIMIT 1").Scan(&primary)
	return primary, err
}

// 查询更新的列

// 获取某张表所有列名
