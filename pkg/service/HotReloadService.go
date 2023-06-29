package service

import (
	"fmt"
	"github.com/Lilymz/table-migration/v2/pkg/config"
	"github.com/Lilymz/table-migration/v2/pkg/model"
	"gopkg.in/ini.v1"
	"strconv"
	"time"
)

// 这个用于处理热加载migration.ini 当文件发生变动时触发实时更新
// 1.首次初始化所定义的配置，初始化数据放入到MissionHolder中
// 2.开启一个协程进行热更新迁移配置，热更新的间隔通过system.ini中的reload.interval，若未配置默认30s加载一次

func StartUpReload() {
	LoadIni()
	//
	go func() {
		model.SYN_WAIT_GROUP.Done()
	}()
	go func() {
		for model.PROCESS_SWTICH {

		}
		model.SYN_WAIT_GROUP.Done()
	}()

}

func LoadIni() {
	// 解析配置文件
	//migrationPath := model.GetMigrationPath()
	if migrationIni, err := ini.Load("E:\\goland\\table-migration\\configs\\migration.ini"); err == nil {
		mission := migrationIni.Section("mission")
		sectionDefault := migrationIni.Section("default")
		defaultDataBase := sectionDefault.Key("dataBase").Value()
		var index int
		for {
			sourceTableKey := "item" + strconv.Itoa(index) + ".sourceTable"
			if !mission.HasKey(sourceTableKey) {
				break
			}
			var dataBase, sourceTable, targetTable, condition, step, primary, status string
			dataBaseKey := "item" + strconv.Itoa(index) + ".dataBase"
			targetTableKey := "item" + strconv.Itoa(index) + ".targetTable"
			conditionKey := "item" + strconv.Itoa(index) + ".condition"
			stepKey := "item" + strconv.Itoa(index) + ".step"
			primaryKey := "item" + strconv.Itoa(index) + ".primaryKey"
			statusKey := "item" + strconv.Itoa(index) + ".status"
			if mission.HasKey(dataBaseKey) {
				dataBase = mission.Key(dataBaseKey).Value()
				if dataBase == "" {
					dataBase = defaultDataBase
				}
			}
			if mission.HasKey(sourceTableKey) {
				sourceTable = mission.Key(sourceTableKey).Value()
			}
			if mission.HasKey(targetTableKey) {
				targetTable = mission.Key(targetTableKey).Value()
				if targetTable == "" {
					targetTable = targetTableKey + "_" + time.Now().Format("20060102")
				}
			}
			if mission.HasKey(conditionKey) {
				condition = mission.Key(conditionKey).Value()
			}
			if mission.HasKey(stepKey) {
				step = mission.Key(stepKey).Value()
				if step == "" {
					step = "1000"
				}
			}
			if mission.HasKey(primaryKey) {
				primary = mission.Key(primaryKey).Value()
				if primary == "" {
					primary = "sn"
				}
			}
			if mission.HasKey(dataBaseKey) {
				dataBase = mission.Key(dataBaseKey).Value()
				if dataBase == "" {
					dataBase = defaultDataBase
				}
			}
			index++
			//model.Mission.New()
		}

	} else {
		fmt.Println(err)
		config.DaoLog.Fatal("migration.ini load failed,err:", err.Error())
	}

}
