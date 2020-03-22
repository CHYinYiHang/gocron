package logging

import (
	"fmt"
	"github.com/CHYinYiHang/gocron/pkg/config"
	"github.com/CHYinYiHang/gocron/pkg/file"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup initialize the log instance
func LoadLogin() {
	var err error
	filePath := config.Config.Server.LogFilePath
	fileName := config.Config.Server.LogFileName
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	// 设置日志输出到文件
	// 定义多个写入器
	writers := []io.Writer{F, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)

	logger = log.New(fileAndStdoutWriter, DefaultPrefix, log.LstdFlags)

}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info output logs at info level
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error output logs at error level
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

// setPrefix set the prefix of the log output
func setPrefix(level Level) {
	_, files, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[gocron]-%s-[%s:%d] ", levelFlags[level], filepath.Base(files), line)
	} else {
		logPrefix = fmt.Sprintf("[gocron]-%s ", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
