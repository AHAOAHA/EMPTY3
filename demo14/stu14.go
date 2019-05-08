//context 使用

package main

import (
	"context"
)

func DoSomething(ctx context.Context, arg Arg) {

}

func SubTask(ctx context.Context) {
	var name string
	var ok bool
	//获取name
	if name, ok = ctx.Value("name").(string); !ok {
		name = "world"
	}

	for {
		select {
		case <-time.After(5 * time.Second):
			fmt
		}
	}
}

func main() {
	ctx := context.Background()	//从根创建新的context
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Minute)	//设置超时时间为1min
	defer cancel()	//当函数退出时，取消子任务
	ctx = context.WithValue(ctx, "name", "Hao.IO")
}