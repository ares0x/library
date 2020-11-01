package zap

import (
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func InitLogger(_logPath string, _maxSize, _maxBackups, _maxAge int)  {
    writeSyncer := getLogWriter(_logPath, _maxSize, _maxBackups, _maxAge)
    encoder := getEncoder()
    core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
    logger := zap.New(core,zap.AddCaller())
    SugarLogger = logger.Sugar()
}

func getLogWriter(_logPath string, _maxSize, _maxBackups, _maxAge int) zapcore.WriteSyncer {
    lumberJackLogger := &lumberjack.Logger{
        Filename:   _logPath,
        MaxSize:    _maxSize,
        MaxBackups: _maxBackups,
        MaxAge:     _maxAge,
        Compress:   false,
    }
    return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.CallerKey = "fileName"            //fileName
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetLog() *zap.SugaredLogger {
    return SugarLogger
}