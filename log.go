package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

// var lock sync.Mutex
var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 3 // 打印的深度，set是第0级，outlogs是第1级，5个方法是第2级，使用的地方是第4级
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"INFO", "DEBUG", "WARN", "ERROR", "FATAL"}
)

type logData struct {
	Level    string        `json:"level"`
	Package  string        `json:"package"`
	FileInfo string        `json:"file_info"`
	FilePath string        `json:"file_path"` // 文件路径
	Content  []interface{} `json:"content"`
}

const (
	INFO Level = iota
	DEBUG
	WARN
	ERROR
	FATAL
)

type LogsConfig struct {
	LogLevel     Level
	LogsType     string
	LogsRootPath string
	LogSaveName  string
	LogsFileExt  string
	LogsFormat   string // 数据格式，当前支持string和json

}

var logsConfig *LogsConfig

func StartLogs(config *LogsConfig) {
	var err error
	logsConfig = config
	filePath := getLogFilePath()
	fileName := getLogFileName()

	F, err := openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)

}

func Debug(v ...interface{}) {
	if DEBUG >= logsConfig.LogLevel {
		outLogs(DEBUG, v...)
	}
	// return
}

func Info(v ...interface{}) {
	if INFO >= logsConfig.LogLevel {
		outLogs(INFO, v...)
	}
	// return
}

func Warn(v ...interface{}) {
	if WARN >= logsConfig.LogLevel {
		outLogs(WARN, v...)
	}
	// return
}

func Error(v ...interface{}) {
	if ERROR >= logsConfig.LogLevel {
		outLogs(ERROR, v...)
	}
	// return
}

func Fatal(v ...interface{}) {
	if FATAL >= logsConfig.LogLevel {
		outLogs(FATAL, v...)
	}
}

//
func setPrefix(level Level) string {
	// 得到的file是打印错误的文件名
	// 得到的line是打印错误的行号
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	logger.SetFlags(0)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]",
			levelFlags[level],
			filepath.Base(file),
			line)
	} else {
		logPrefix = fmt.Sprintf("[%s]",
			levelFlags[level])
	}
	return logPrefix
	// logger.SetPrefix(logPrefix)
}

func setLogData(level Level, v ...interface{}) {
	var logInfo logData
	logInfo.Level = levelFlags[level]
	// 得到的file是打印错误的文件名
	// 得到的line是打印错误的行号
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	if ok {
		// 设置文件名和行号
		logInfo.FileInfo = fmt.Sprintf("%s:%d", filepath.Base(file), line)
		// 设置包名
		logInfo.Package = filepath.Base(filepath.Dir(file))
		// 设置文件的绝对路径
		logInfo.FilePath = fmt.Sprintf("%s:%d", file, line)
	}
	logInfo.Content = append(logInfo.Content, v...)

	var logString string

	switch logsConfig.LogsFormat {
	case "json":
		logbyte, err := json.Marshal(&logInfo)
		if err != nil {
			return
		}
		logString = string(logbyte)
	case "string":
		// 输出key value 格式
		logString = fmt.Sprintf("%+v", logInfo)
	}

	// 删除默认的时间戳前缀
	// logger.SetFlags(0)
	logger.Println(logString)
}

// 主要用来判断日志类型。并非记录的格式
// 当前计划支持 logs(写本地log文件) redis  elk 等
func outLogs(level Level, v ...interface{}) {
	// lock.Lock()
	// defer lock.Unlock()
	// json和控制台输出
	switch logsConfig.LogsType {
	case "logs":
		setLogData(level, v...)
	case "redis":
		// todo
	case "elk":
		// todo
	default:
		// log.Println(levelFlags[level], v)
	}
	
	log.Println(setPrefix(level), v)
	// logger.Println(levelFlags[level], v)
	// setPrefix(level)
	// logger.Println(v)
}

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:03")
}
