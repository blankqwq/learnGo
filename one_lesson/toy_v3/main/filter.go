package main

type Filter func(c *Context) error

type FilterBuild func(next Filter) Filter

type FilterInterface interface {
	// 在加闭包
	Handle2(c *Context,next Filter) Filter
	Handle(next Filter) Filter
}

type TestFilter struct {

}

func (f *TestFilter)Handle(next Filter) Filter  {
	return func(c *Context) error {
		// 只能说是简单的aop编程，无法做到更加复杂的业务逻辑,以为这个中间件是静态的，而不是动态执行的
		println("start:...")
		// 这里也是可以修改的
		err :=next(c)
		println("end: ...")
		return err
	}
}

func (f *TestFilter)Handle2(c *Context,next Filter) Filter  {
	// 循环生成
	// 外部先执行,
	//	然后执行next
		next(c)
	for i := 0; i < 5 ; i++ {

	}

	return func(c *Context) error {
		return nil
	}
}
