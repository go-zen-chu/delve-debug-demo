package count

import (
	"sync"
)

// Parallel1000count should count upto 1000 and return it
func Parallel1000count() int {
	count := 0
	// 並列処理がすべて完了するのを wait する
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// goroutine で並列化
		go func() {
			defer wg.Done()
			count++ // 1 を足すだけ、簡単!
		}()
	}
	// 処理がすべて終わるまで待つ。これで 1000 回インクリメントされる
	wg.Wait()
	return count
}
