/*
如果开发没有自定义的日志对象定义，那么KisFlow会提供一个默认的日志对象kisDefaultLogger，
这个类实现了KisLogger的全部接口，且都是默认打印到标准输出的形式来打印日志

这里在init()初始化方法中，会判断目前是否已经有设置全局的Logger对象，
如果没有，KisFlow会默认选择kisDefaultLog 作为全局Logger日志对象
*/
package log

import (
	"context"
	"fmt"
)

// kisDefaultLog 默认提供的日志对象
type kisDefaultLog struct{}

func (log *kisDefaultLog) InfoF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (log *kisDefaultLog) ErrorF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (log *kisDefaultLog) DebugF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (log *kisDefaultLog) InfoFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func (log *kisDefaultLog) ErrorFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func (log *kisDefaultLog) DebugFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func init() {
	// 如果没有设置Logger, 则启动时使用默认的kisDefaultLog对象
	if Logger() == nil {
		SetLogger(&kisDefaultLog{})
	}
}
