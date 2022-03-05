package count

import (
	"fmt"
	"sync"
)

func Parallel1000countDebug() int {
	count := 0
	working := 0
	mutex := &sync.Mutex{}
	// 並列処理がすべて完了するのを wait する
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// goroutine で並列化
		go func() {
			defer wg.Done()
			count++ // 1 を足すだけ、簡単!

			// スレッドセーフな値を比較して、実際にどうなっているのか比較する
			mutex.Lock()
			working++
			mutex.Unlock()
			// 一致してなかったらおかしい
			if working != count {
				fmt.Printf("一致してない! おかしい! %d, %d\n", working, count)
			}
		}()
	}
	// 処理がすべて終わるまで待つ。これで 1000 回インクリメントされる
	wg.Wait()
	return count
}
