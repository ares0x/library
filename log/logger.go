package logs

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var (
	MaxSize    = 200 // 每个日志文件最大尺寸200M
	MaxBackups = 20  // 日志文件最多保存20个备份
)

func NewLoggers(logPtah string, maxAge int, compress bool) *logrus.Logger {
	logger := logrus.New()
	lumberLogger := &lumberjack.Logger{
		Filename:   logPtah,
		MaxSize:    MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
		LocalTime:  true,       // 备份文件名本地/UTC时间
	}
	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	logMultWriter := io.MultiWriter(os.Stdout, lumberLogger)

	logger.SetOutput(logMultWriter)
	return logger
}
