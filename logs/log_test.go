//	log := NewLogger(10000)
//	log.SetLogger("console", "")
//
//	> the first params stand for how many channel
//
// Use it like this:
//
//	log.Debug("debug")
//	log.Informational("info")
//	log.Notice("notice")
//	log.Warning("warning")
//	log.Error("error")
//	log.Critical("critical")
//	log.Alert("alert")
//	log.Emergency("emergency")

package logs

import (
	"testing"
)

func TestLogger(t *testing.T) {
	// 控制台
	log1 := NewLogger(10000)
	log1.EnableFuncCallDepth(true)
	log1.SetLogger("console", nil)
	log1.Error("false")

	// 输出文件
	log := NewLogger(10000)
	log.SetLogger("file", map[string]interface{}{"filename": "test.log"})
	log.Debug("debug")
	log.Informational("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
}
