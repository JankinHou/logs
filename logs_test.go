package logs

import (
	"fmt"
	"sync"
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
		LogsFormat:   "json",
	}
	start := time.Now().UnixNano()
	StartLogs(logsConfig)
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Error(i, time.Now().UnixNano())
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Minute)
	fmt.Println("程序共执行：", (time.Now().UnixNano()-start)/1e6, "毫秒")
	// time.Sleep(time.Minute)
	t.Error(1111)
}
