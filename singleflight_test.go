package single_flight

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//检查缓存是否存在

var g = &Group{
	mp: map[string]*call{},
	mutx: sync.Mutex{},
}

//var g Group
func TestDo(t *testing.T) {

	nums := 10
	wg := sync.WaitGroup{}
	wg.Add(nums)

	for i := 0; i < nums; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, "进入")
			_, _ = g.do("key", func() (interface{}, error) {
				time.Sleep(1*time.Second)
				fmt.Println(123)
				return nil, nil
			})
			//testSingle1()
			fmt.Println(i, "获取成功")
		}(i)
	}

	wg.Wait()


}
