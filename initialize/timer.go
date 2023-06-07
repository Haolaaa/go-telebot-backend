package initialize

import (
	"admin/config"
	"admin/global"
	"fmt"

	"github.com/robfig/cron/v3"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for i := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				fmt.Println(detail.TableName)
			}(global.CONFIG.Timer.Detail[i])
		}
	}
}
