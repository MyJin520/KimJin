package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局Logger实例，供项目其他模块调用
var Logger *zap.Logger

// InitZapLogger 初始化zap日志管理器
func InitZapLogger(logRootDir string) {
	// 1. 生成今日的日志目录（年/月/日层级）
	today := time.Now()
	dateDir := filepath.Join(
		logRootDir,
		today.Format("2006"), // 年（Go的参考时间格式，固定值）
		today.Format("01"),   // 月
		today.Format("02"),   // 日
	)

	// 确保目录存在，不存在则递归创建（权限0755：所有者可读可写可执行，其他只读可执行）
	if err := os.MkdirAll(dateDir, 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 2. 定义日志编码配置（JSON格式，易解析；也可改为ConsoleEncoder）
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller", // 显示调用日志的文件和行号
		LineEnding:   zapcore.DefaultLineEnding,
		EncodeLevel:  zapcore.CapitalLevelEncoder, // 级别大写（INFO/ERROR/WARN）
		EncodeTime:   zapcore.ISO8601TimeEncoder,  // 时间格式：2026-02-27T15:04:05Z07:00
		EncodeCaller: zapcore.ShortCallerEncoder,  // 调用者格式：logger/zap_logger.go:50
	}

	// 3. 为不同级别日志配置文件写入器（lumberjack处理文件轮转）
	// 3.1 Info级别日志配置
	infoWriter := &lumberjack.Logger{
		Filename:   filepath.Join(dateDir, "info.log"), // 日志文件路径
		MaxSize:    100,                                // 单个文件最大100MB
		MaxBackups: 30,                                 // 保留30个备份文件
		MaxAge:     7,                                  // 日志文件保留7天
		Compress:   true,                               // 压缩备份文件（节省空间）
	}

	// 3.2 Error级别日志配置
	errorWriter := &lumberjack.Logger{
		Filename:   filepath.Join(dateDir, "error.log"),
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}

	// 3.3 Warn级别日志配置
	warnWriter := &lumberjack.Logger{
		Filename:   filepath.Join(dateDir, "warn.log"),
		MaxSize:    100,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}

	// 4. 构建不同级别的zap Core（日志核心，负责编码和输出）
	// Info Core：只处理Info级别日志，输出到info.log
	infoCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(infoWriter),
		zapcore.InfoLevel,
	)

	// Error Core：只处理Error级别日志，输出到error.log
	errorCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(errorWriter),
		zapcore.ErrorLevel,
	)

	// Warn Core：只处理Warn级别日志，输出到warn.log
	warnCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(warnWriter),
		zapcore.WarnLevel,
	)

	// 5. 控制台输出Core（可选，便于开发调试）
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // 控制台用易读的文本格式
		zapcore.AddSync(os.Stdout),               // 输出到标准输出（控制台）
		zapcore.DebugLevel,                       // 控制台输出所有级别日志
	)

	// 6. 合并所有Core：NewTee会将日志分发到匹配级别的Core
	multiCore := zapcore.NewTee(infoCore, errorCore, warnCore, consoleCore)

	// 7. 创建最终的Logger实例
	Logger = zap.New(
		multiCore,
		zap.AddCaller(),                   // 显示日志调用者（文件+行号）
		zap.AddStacktrace(zap.ErrorLevel), // Error级别日志显示堆栈信息
	)

	// 替换zap的全局Logger（可选，可直接用zap.L()调用）
	zap.ReplaceGlobals(Logger)

}

// 封装便捷的日志调用方法
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// Sync 确保日志刷入文件
func Sync() {
	err := Logger.Sync()
	if err != nil {
		fmt.Println("日志保存失败：" + fmt.Sprintf("%v", err))
	}
}
