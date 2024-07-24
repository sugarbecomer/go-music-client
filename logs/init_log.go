package logs

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"path"
	"runtime"
	"time"
)

func LogInit() {
	// 输出到命令行
	customFormatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceColors:     true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			_, l := f.Func.FileLine(f.PC)
			return fmt.Sprintf(" %s:%d", f.File, l), ""
		},
	}
	log.SetFormatter(customFormatter)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	//输出到文件
	rotateLogs, err := rotatelogs.New(path.Join("logs", "%Y-%m-%d.log"),
		rotatelogs.WithRotationTime(time.Hour*24), // 每天一个新文件
		rotatelogs.WithMaxAge(time.Hour*24*7),     // 日志保留3天
	)
	if err != nil {
		log.Info(err)
		return
	}
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			log.InfoLevel:  rotateLogs,
			log.WarnLevel:  rotateLogs,
			log.ErrorLevel: rotateLogs,
			log.FatalLevel: rotateLogs,
			log.PanicLevel: rotateLogs,
		},
		&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%time%]    [%lvl%]: %msg% \r\n",
		},
	))
}
