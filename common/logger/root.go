package logger

import (
	"os"
	"gcoresys/common"
)

var (
	root          = &logger{[]interface{}{}, new(swapHandler)}
	StdoutHandler = StreamHandler(os.Stdout, LogfmtFormat())
	StderrHandler = StreamHandler(os.Stderr, LogfmtFormat())
)

func init() {
	root.SetHandler(DiscardHandler())
}

// New returns a new logger with the given context.
// New is a convenient alias for Root().New
func New(ctx ...interface{}) Logger {
	return root.New(ctx...)
}

// Root returns the root logger
func Root() Logger {
	return root
}

// The following functions bypass the exported logger methods (logger.Debug,
// etc.) to keep the call depth the same for all paths to logger.write so
// runtime.Caller(2) always refers to the call site in client code.

// Trace is a convenient alias for Root().Trace
func Trace(msg string, ctx ...interface{}) {
	root.write(msg, LvlTrace, ctx)
}

// Debug is a convenient alias for Root().Debug
func Debug(msg string, ctx ...interface{}) {
	root.write(msg, LvlDebug, ctx)
}

// Info is a convenient alias for Root().Info
func Info(msg string, ctx ...interface{}) {
	root.write(msg, LvlInfo, ctx)
}

// Warn is a convenient alias for Root().Warn
func Warn(msg string, ctx ...interface{}) {
	root.write(msg, LvlWarn, ctx)
}

// Error is a convenient alias for Root().Error
func Error(msg string, ctx ...interface{}) {
	root.write(msg, LvlError, ctx)
}

// Crit is a convenient alias for Root().Crit
func Crit(msg string, ctx ...interface{}) {
	root.write(msg, LvlCrit, ctx)
	os.Exit(1)
}


// 占时废弃
// 初始化log如果handler在外边设置了则使用外边传进来的handler
func InitLoggerBak(logLevel Lvl, handler *GlogHandler) {
	logHandler := handler
	if handler == nil {
		logHandler = NewGlogHandler(StreamHandler(os.Stdout, TerminalFormat(true)))
		logHandler.Verbosity(Lvl(logLevel))
	}
	Root().SetHandler(logHandler)
	// 每个微服务启动时都会调用该方法，因此在这里输出docker环境最合适不过了
	common.PrintCurDockerEnv()
}

// 初始化log如果handler在外边设置了则使用外边传进来的handler
func InitLogger(logLevel Lvl, configPath interface{}) {
	var logHandler *GlogHandler
	if configPath == nil {
		logHandler = NewGlogHandler(StreamHandler(os.Stdout, TerminalFormat(true)))
	} else if _, ok := configPath.(string); ok {
		targetDir := "/var/log/qy"
		if !PathExists(targetDir) {
			os.MkdirAll(targetDir, os.ModePerm)
		}
		fileHandler, err := FileHandler(targetDir + "/" + configPath.(string), TerminalFormat(false))
		if err != nil {
			panic(err.Error())
		}
		// gscore日志不打印到终端
		if configPath.(string) == "gscore.log" {
			logHandler = NewGlogHandler(MultiHandler(
				fileHandler,
			))
		} else {
			logHandler = NewGlogHandler(MultiHandler(
				StreamHandler(os.Stdout, TerminalFormat(true)),
				fileHandler,
			))
		}
	} else {
		logHandler = configPath.(*GlogHandler)
	}
	logHandler.Verbosity(Lvl(logLevel))

	Root().SetHandler(logHandler)
	// 每个微服务启动时都会调用该方法，因此在这里输出docker环境最合适不过了
	common.PrintCurDockerEnv()
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
