package model

import "sync"

// 此文件用于记录当前系统中的一些参数，可以用于优雅的退出等等...

// All_GOROUTINE_STATUS_MAP 用于控制所有goroutine开关，当程序关闭是进行关闭
var (
	All_GOROUTINE_STATUS_MAP = make(map[string]bool, 8)
	SYN_WAIT_GROUP           sync.WaitGroup
	PROCESS_SWTICH           = true
)

const (
	HOT_RELOAD_GOROUTINE = "HOT_RELOAD_GOROUTINE"
	MISSION_GOTOUTINE    = "MISSION_GOTOUTINE" // 此常量用于标记启动的迁移表mission goroutine
	CONFIGS              = "/configs/migration.ini"
)
