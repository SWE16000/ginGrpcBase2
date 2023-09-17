package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)


// 首字母大写，公开这两个logger实例
var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

// 对三个条件定义为一个结构体
type LoggerConfig struct {
	encoder      zapcore.Encoder
	writerSyncer zapcore.WriteSyncer
	levelEnable  zapcore.Level
}

func init()  {
	InitLogger()
}

// logger初始化方法
func InitLogger() {
	//// 实例化结构体
	//config := LoggerConfig{
	//	// 使用zap自带的NewJSONEncoder定义为一个JSON编码器作为输出
	//	encoder: zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
	//	// 匿名函数，返回文件保存位置
	//	// 注意这里需要使用AddSync对输出流的类型执行转换
	//	writerSyncer: func() zapcore.WriteSyncer {
	//		file, _ := os.OpenFile("./logs/logger.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//		return zapcore.AddSync(file)
	//	}(),
	//	// 日志等级配置，使用最低级的debuglevel
	//	levelEnable: zapcore.DebugLevel,
	//}

	//// 将所有参数依次填入
	//core := zapcore.NewCore(
	//	config.encoder,
	//	//config.writerSyncer,
	//
	//	// 第一个是文件输出流，第二个为控制台输出（注意这里仍然需要AddSync仅需转型）
	//	zapcore.NewMultiWriteSyncer(
	//		config.writerSyncer,
	//		zapcore.AddSync(os.Stdout),
	//	),
	//	config.levelEnable,
	//)
	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(
			getWriterSyncer(),
			zapcore.AddSync(os.Stdout)),
		getLevelEnable(),
	)

	// 使用new方法实例化logger
	Logger = zap.New(core)
	// 变换为sugarlogger
	SugarLogger = Logger.Sugar()
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()       // 获取配置实例
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // 规定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 全部大写
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSyncer() zapcore.WriteSyncer {

	file, _ := os.OpenFile("./logs/"+time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	return zapcore.AddSync(file)
}

func getLevelEnable() zapcore.Level {
	return zapcore.DebugLevel
}