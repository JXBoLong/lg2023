package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/lg"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for i := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.GVA_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.GVA_CONFIG.Timer.Detail[i])
		}
	}
	_, err := global.GVA_Timer.AddTaskByFunc("testOneMinute", "* * * * *", func() {
		fmt.Println("每分钟一次")
	})
	if err != nil {
		fmt.Println("add timer error:", err)
	}
	_, err = global.GVA_Timer.AddTaskByFunc("testFiveMinute", "*/5 * * * *", func() {
		fmt.Println("每五分钟一次")
	})
	if err != nil {
		fmt.Println("add timer error:", err)
	}
	_, err = global.GVA_Timer.AddTaskByFunc("AutoMaticUnEnableProject", "0 * * * *", lg.AutoMaticUnEnableProject)
	if err != nil {
		fmt.Println("设置每小时自动下架项目失败", err)
	}
}
