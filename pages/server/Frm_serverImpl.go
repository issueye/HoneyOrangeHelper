package server

import (
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/internal/gow32"
	"fmt"
	"strings"
	"time"

	"github.com/ying32/govcl/vcl"
)

// ::private::
type TFrm_serverFields struct {
	ShowLogCount int32 // 显示日志条数
}

func (f *TFrm_server) OnFormCreate(sender vcl.IObject) {
	f.Monitor()
}

// 添加日志到主窗口中
func (f *TFrm_server) addLog(msg string) {
	vcl.ThreadSync(func() {
		now := time.Now().Format("2006-01-02 15:04:05")
		f.Mmo_run_log.Lines().Add(fmt.Sprintf("%s %s", now, msg))

		if f.Mmo_run_log.Lines().Count() > f.ShowLogCount {
			f.Mmo_run_log.Lines().Delete(0)
		}

		// 移动到最后一行
		gow32.SendMessageW(f.Mmo_run_log.Handle(), gow32.WM_VSCROLL, gow32.SB_BOTTOM, 0)
	})
}

func (f *TFrm_server) Monitor() {
	go func() {
		for msg := range global.MsgChannel {
			f.addLog(strings.TrimSpace(msg))
			global.Sugared.Info(strings.TrimSpace(msg))
		}
	}()
}
