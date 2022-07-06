package single_flight

import "sync"

//防止刺穿的包

type Group struct {
	mutx sync.Mutex
	//记录所有key对应的结果
	mp map[string]*call
}

type call struct {
	wg *sync.WaitGroup //来一组协程锁等待一下

	result interface{} //返回的内容
	err error //返回的错误信息
}

//用于执行函数，参数需要包含同一组的“key”，以及待执行的函数“fc”
func (g *Group) do(key string, fc func()(interface{}, error)) (interface{}, error)  {
	//只允许一个协程操作map
	g.mutx.Lock()
	if c, ok := g.mp[key]; ok {
		//操作完了，释放吧
		g.mutx.Unlock()
		//等结果出来咯
		c.wg.Wait()

		return c.result, c.err
	}

	c := &call{
		wg: &sync.WaitGroup{},
	}
	//只有一个在获取数据，其他协程都要等他哟
	c.wg.Add(1)

	//写到map中提供其他协程读取
	g.mp[key] = c

	//操作完了，释放把
	g.mutx.Unlock()

	//真正执行的地方, 记录返回的情况
	result, err := fc()
	c.result = result
	c.err = err

	//获取到了，可以结束了
	c.wg.Done()

	return result, err
}