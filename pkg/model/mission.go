package model

import (
	"os"
	"time"
)

type Mission struct {
	DataBase    string    `json:"dataBase"`    // 数据库名称
	SourceTable string    `json:"sourceTable"` // 待迁移表
	TargetTable string    `json:"targetTable"` // 迁移目的表,若为空则以sourceTable按照年月
	Condition   string    `json:"condition"`
	Step        string    `json:"step"`       //每次移动的数据大小
	PrimaryKey  string    `json:"primaryKey"` //主键名称
	Status      int       `json:"status"`     // 此迁移配置的状态（比如迁移开始，停止，暂停）
	UpdateTime  time.Time `json:"updateTime"` // 默认创建会记录,之后会在每次的迁移都会更新（包含成功或者失败）
}

const (
	RUNNING = iota
	PAUSE
	STOP
)

// New 迁移构构造器
func (mission *Mission) New(dataBase, sourceTable, targetTable, condition, step, primaryKey string, status int) *Mission {
	return &Mission{
		DataBase:    dataBase,
		SourceTable: sourceTable,
		TargetTable: targetTable,
		Condition:   condition,
		Step:        step,
		PrimaryKey:  primaryKey,
		Status:      status,
	}
}

// MissionHolder 迁移配置持有容器
var MissionHolder map[string]Mission

func GetMigrationPath() string {
	path, _ := os.Getwd()
	migrationPath := path + CONFIGS
	return migrationPath
}
