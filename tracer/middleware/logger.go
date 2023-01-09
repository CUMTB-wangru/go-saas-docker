package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Logger 日志配置
func Logger() gin.HandlerFunc {
	filePath := "log/log"
	//linkName := "latest_log.log"

	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	// logrus 实例化对象
	logger := logrus.New()
	// 日志输出
	logger.Out = scr
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// retalog 日志拆分中间件
	logWriter, _ := retalog.New(
		// 分割后的文件名称
		filePath+"%Y%m%d.log",
		// 设置最大保存时间(7天)
		retalog.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		retalog.WithRotationTime(24*time.Hour),
		// 生成软链，指向最新日志文件
		//retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 执行时间，Since从上一个time到现在的时间以毫秒计算
		stopTime := time.Since(startTime).Milliseconds()
		// 结束时间
		spendTime := fmt.Sprintf("%d ms", stopTime)
		// 主机名
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "Unkonwn"
		}
		// 状态码
		statusCode := c.Writer.Status()
		// 客户端IP
		clientIP := c.ClientIP()
		// userAgent
		userAgent := c.Request.UserAgent()
		// 字节数
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		// 请求方式
		method := c.Request.Method
		// 请求路由
		path := c.Request.RequestURI
		// 日志格式
		entry := logger.WithFields(logrus.Fields{
			"HostName": hostName,
			"status": statusCode,
			"SpendTime": spendTime,
			"IP":clientIP,
			"Method":method,
			"Path": path,
			"DataSize": dataSize,
			"Agent": userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
