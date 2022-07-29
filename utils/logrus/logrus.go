package logrus

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// NewLogs 级别
func NewLogs() *logrus.Logger {
	log := logrus.New()
	// JSONFormatter将日志格式化为可解析的json
	log.SetFormatter(&logrus.JSONFormatter{
		// 设置用于封送时间戳的格式。
		TimestampFormat: "2006-01-02 15:04:05",
		// 允许在输出中禁用自动时间戳（true--不输出，false -- 输出）
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})
	/**    文件输出   */
	log.SetOutput(os.Stdout)
	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	uploadDir := "utils/logrus/logruslogtest/" + time.Now().Format("2006/01/02/15/04/05/")
	_ = os.MkdirAll(uploadDir, os.ModePerm)
	file, err := os.OpenFile(uploadDir+"logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr", err.Error())
	}
	/** 	设置级别	***/
	log.SetLevel(logrus.DebugLevel)
	/**		记入函数  开启这个模式会增加性能开销。		*/
	log.SetReportCaller(false)
	return log
}
