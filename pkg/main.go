package main

import (
	. "github.com/Lilymz/table-migration/v2/pkg/config"
	. "github.com/Lilymz/table-migration/v2/pkg/model"
	"github.com/Lilymz/table-migration/v2/pkg/service"
)

func main() {
	// 获取当前带迁移的表配置，用于决定开启多少个goroutine
	func() {
		if goroutineNum := len(MissionHolder); goroutineNum > 0 {
			// 此处的2是热更新的mission和程序总开关
			SYN_WAIT_GROUP.Add(goroutineNum + 2)
		} else {
			DaoLog.Fatal("mission goroutine create failed ,mission size:", goroutineNum)
		}
	}()
	// 开启热加载
	service.StartUpReload()

}
