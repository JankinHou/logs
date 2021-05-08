package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestNewLogs(t *testing.T) {
	logsConfig := &LogsConfig{
		LogLevel:     1,
		LogsType:     "logs",
		LogsRootPath: "runtime/",
		LogSaveName:  "log",
		LogsFileExt:  "log",
		LogsFormat:   "string",
	}
	start := time.Now().UnixNano()
	StartLogs(logsConfig)
	// wg := sync.WaitGroup{}
	for i := 0; i < 50000; i++ {
		Error(i, time.Now().UnixNano())
		time.Sleep(time.Second / 2)
		// wg.Add(1)
		// go func(i int) {
		// 	// defer wg.Done()

		// }(i)
	}
	// wg.Wait()
	time.Sleep(time.Minute)
	fmt.Println("程序共执行：", (time.Now().UnixNano()-start)/1e6, "毫秒")
	// time.Sleep(time.Minute)
	t.Error(1111)
}
