package logs

import "testing"

func testConsoleCalls(bl *BeeLogger) {
	bl.Emergency("emergency")
	bl.Alert("alert")
	bl.Critical("critical")
	bl.Error("error")
	bl.Warning("warning")
	bl.Notice("notice")
	bl.Informational("informational")
	bl.Debug("debug")
}

func TestConsole(t *testing.T) {
	log1 := NewLogger(100000)
	log1.EnableFuncCallDepth(true)
	log1.SetLogger("console", nil)
	testConsoleCalls(log1)
}
