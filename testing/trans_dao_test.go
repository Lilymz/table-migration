package testing

import (
	"github.com/Lilymz/table-migration/v2/pkg/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGetPrimary(t *testing.T) {
	dbConfig := mysql.Config{
		DSN:                       "root:hstest@2014@tcp(192.168.113.45:3306)/cluster_server_cxy?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                                  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                // 根据当前 MySQL 版本自动配置
	}
	singletonOperator := dao.SingletonOperator
	connect := singletonOperator.Connect(dbConfig, &gorm.Config{})
	if connect {
		t.Log("数据库连接成功...")
		primary, err := dao.GetPrimary("submit_message_send_history")
		if err != nil {
			t.Error("主键获取失败")
		} else {
			t.Log(primary)
		}
	}

}
