package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logs *zap.Logger

type config struct {
	ServerName string
	FileName   string
	FileDir    string
	Level      int
	Encoding   string
}

// read log config
func readConfig() config {
	conf := config{}
	conf.ServerName = viper.Get("log.server_name").(string)
	conf.FileName = viper.Get("log.file_name").(string)
	conf.Level = viper.Get("log.level").(int)
	return conf
}

func InitConfig() {
	conf := readConfig()
	if conf.ServerName == "" {
		conf.ServerName = "default"
	}
	if conf.FileDir == "" {
		conf.FileDir = "./logs"
	}
	if conf.FileName == "" {
		conf.FileName = "app"
	}
	if conf.Encoding == "" {
		conf.Encoding = "json"
	}

	//获得当前日期文件名，传入后已有追加，没有新建
	nowTime := time.Now().Format("20060102")
	LogName := conf.FileDir + "/" + conf.FileName + "_" + nowTime + ".log"

	//	是否关闭堆栈信息
	disableStacktrace := false
	// 设置日志级别
	outLogLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	switch {
	case conf.Level == -1:
		//	调试模式
		outLogLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
		disableStacktrace = false
	case conf.Level == 0:
		//	打印模式
		outLogLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
		disableStacktrace = false
	case conf.Level == 1:
		//	警告
		outLogLevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	case conf.Level == 2:
		//	错误
		outLogLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case conf.Level == 3:
		//	中断
		outLogLevel = zap.NewAtomicLevelAt(zap.PanicLevel)
	case conf.Level == 4:
		//	致命错误
		outLogLevel = zap.NewAtomicLevelAt(zap.FatalLevel)
	}
	// 输出选项编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "logger",
		//CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "Stacktrace", //Stacktrace（堆栈跟踪）
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     TimeEncoder,                   // UTC+8 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	// 发生错误打印行号，路由日志不打印行号
	if conf.Level >= 1 {
		encoderConfig.CallerKey = "caller"
	}

	config := zap.Config{
		Level:             outLogLevel,                                            // 日志级别
		Encoding:          conf.Encoding,                                          // 输出编码格式为 console 或 json
		EncoderConfig:     encoderConfig,                                          // 编码器配置
		DisableStacktrace: disableStacktrace,                                      // 堆栈跟踪  bool
		InitialFields:     map[string]interface{}{"serviceName": conf.ServerName}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:       []string{"stdout", LogName},                            // 输出到指定文件 stdout
		ErrorOutputPaths:  []string{"stderr"},                                     // 将系统内的error记录到控制台或者文件
	}
	// 构建日志
	var err error
	Logs, err = config.Build()
	if err != nil {
		panic(err)
	}
}

// TimeEncoder 格式化时间
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Local().Format("2006-01-02 15:04:05"))
}
