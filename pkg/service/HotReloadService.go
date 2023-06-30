package service

import (
	"github.com/Lilymz/table-migration/v2/pkg/config"
	"github.com/Lilymz/table-migration/v2/pkg/model"
	"gopkg.in/ini.v1"
	"strconv"
	"time"
)

const ITEM = "item"
const DATA_BASE = ".dataBase"
const SOURCE_TABLE = ".sourceTable"
const TARGET_TABLE = ".targetTable"
const CONDITION = ".condition"
const STEP = ".step"
const PRIMARY_KEY = ".primaryKey"
const STATUS = ".status"
const DEFAULT = "default"
const MSSION = "mission"

// 这个用于处理热加载migration.ini 当文件发生变动时触发实时更新
// 1.首次初始化所定义的配置，初始化数据放入到MissionHolder中
// 2.开启一个协程进行热更新迁移配置，热更新的间隔通过system.ini中的reload.interval，若未配置默认30s加载一次

func StartUpReload() {
	LoadIni()
	go func() {
		for {
			load, _ := ini.Load("E:\\goland\\table-migration\\configs\\system.ini")
			section := load.Section("system")
			isRunning, _ := section.Key("process.switch").Bool()
			model.PROCESS_SWTICH = isRunning
			if !isRunning {
				break
			}
			time.Sleep(time.Minute)
		}
		model.SYN_WAIT_GROUP.Done()
		config.DaoLog.Info("程序总控制协程开关关闭...")
	}()
	go func() {
		for model.PROCESS_SWTICH {
			load, err := ini.Load("E:\\goland\\table-migration\\configs\\system.ini")
			section := load.Section("system")
			// 读取系统中的配置间隔配置更新热加载
			internalTime := 30
			if err == nil {
				timeStep, err := section.Key("reload.interval").Int()
				if err != nil {
					return
				}
				internalTime = timeStep
			}
			time.Sleep(time.Duration(internalTime) * time.Second)
			LoadIni()
		}
		model.SYN_WAIT_GROUP.Done()
		config.DaoLog.Info("迁移表配置热加载关闭...")
	}()
}

func LoadIni() {
	// 解析配置文件
	// todo 后续打包需要换回来 migrationPath := model.GetMigrationPath()
	holderMap := make(map[string]*model.Mission, 8)
	if missionConfig, err := ini.Load("E:\\goland\\table-migration\\configs\\migration.ini"); err == nil {
		defaultSection := missionConfig.Section(DEFAULT)
		defaultDataBase := defaultSection.Key("dataBase").Value()
		missionSection := missionConfig.Section(MSSION)
		var index int
		for {
			sourceTableKey := ITEM + strconv.Itoa(index) + SOURCE_TABLE
			if !missionSection.HasKey(sourceTableKey) {
				break
			}
			mission, result, key := productMissionByConfig(missionSection, index, defaultDataBase)
			if !result {
				config.DaoLog.Warn(ITEM + strconv.Itoa(index) + " product mission config failed")
				index++
				continue
			}
			holderMap[key] = mission
			index++
		}
		model.MissionHolder = holderMap
		for key, value := range model.MissionHolder {
			config.DaoLog.Info("待迁移表："+key, value)
		}
	} else {
		config.DaoLog.Fatal("migration.ini load failed,err:", err.Error())
	}
}
func productMissionByConfig(section *ini.Section, index int, defaultDataBase string) (*model.Mission, bool, string) {
	var (
		dataBase, sourceTable, targetTable, condition, step, primaryKey, status string
	)
	dataBase = section.Key(ITEM + strconv.Itoa(index) + DATA_BASE).Value()
	if "" == dataBase {
		dataBase = defaultDataBase
	}
	sourceTable = section.Key(ITEM + strconv.Itoa(index) + SOURCE_TABLE).Value()
	if "" == sourceTable {
		config.DaoLog.Warn(ITEM + strconv.Itoa(index) + "product failed missing sourceTable!")
		return nil, true, ""
	}
	targetTable = section.Key(ITEM + strconv.Itoa(index) + TARGET_TABLE).Value()
	if "" == targetTable {
		targetTable = sourceTable + "_" + time.Now().Format("20060102")
	}
	condition = section.Key(ITEM + strconv.Itoa(index) + CONDITION).Value()
	step = section.Key(ITEM + strconv.Itoa(index) + STEP).Value()
	primaryKey = section.Key(ITEM + strconv.Itoa(index) + PRIMARY_KEY).Value()
	// todo 此处如果primaryKey为空需要通过数据库获取某张表得主键
	status = section.Key(ITEM + strconv.Itoa(index) + STATUS).Value()
	if "" == status {
		status = "0"
	}
	return model.Mission.New(model.Mission{}, dataBase, sourceTable, targetTable, condition, step, primaryKey, 0), true, sourceTable
}
