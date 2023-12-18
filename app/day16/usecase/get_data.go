package usecase

import (
	"fmt"
	"in_memory_cache_demo/adapters"
	"in_memory_cache_demo/interfaces"
	"time"
)

func GetData() string {
	memory_cache_adapter := adapters.NewMemoryCacheAdapter()
	redis_adapter := adapters.NewRedisAdapter()
	totalTime := 0
	for {
		value := interfaces.GetData(memory_cache_adapter, "key")
		if value == "" {
			fmt.Println("インメモリキャッシュに無いのでRedisから取得します")
			value = interfaces.GetData(redis_adapter, "key")
			interfaces.SetData(memory_cache_adapter, "key", value)
		} else {
			fmt.Println("インメモリキャッシュにありました")
		}
		fmt.Println(value)
		// 3秒待機
		time.Sleep(3 * time.Second)
		// 開始から何秒経過したかを表示
		totalTime += 3
		fmt.Println(totalTime, "秒経過")
	}
}