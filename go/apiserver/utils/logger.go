package utils

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
	"path"
	"runtime"
	"strings"
	"sync"
)

var singletonLog *logrus.Logger
var once sync.Once

func GetLogger() *logrus.Logger {
	once.Do(func() {
		singletonLog = logrus.New()
		fmt.Println("kubeyard logger initiated. This should be called only once.")
		var exampleFormatter = &zt_formatter.ZtFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				return fmt.Sprintf("%s()", strings.ReplaceAll(f.Function, "kubeyard-dashboard", "")), fmt.Sprintf("%s:%d", filename, f.Line)
			},
			Formatter: nested.Formatter{
				HideKeys:    true,
				FieldsOrder: []string{"component", "category"},
			},
		}
		//singletonLog.set
		//singletonLog.SetReportCaller(true)
		singletonLog.SetLevel(logrus.DebugLevel)
		singletonLog.SetReportCaller(true)
		singletonLog.SetFormatter(exampleFormatter)
	})

	return singletonLog
}
