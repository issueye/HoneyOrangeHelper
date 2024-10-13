package home

import (
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/pages/about"
	"HoneyOrangeHelper/pages/server_mana"
	"context"
	"strings"
	"time"

	"github.com/robfig/cron"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TaskScheduler struct {
	cron    *cron.Cron
	running bool
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		cron:    cron.New(),
		running: false,
	}
}

func (ts *TaskScheduler) AddTask(spec string, cmd func()) {
	ts.cron.AddFunc(spec, cmd)
}

func (ts *TaskScheduler) Start() {
	ts.cron.Start()
	ts.running = true
}

func (ts *TaskScheduler) Stop() {
	ts.cron.Stop()
	ts.running = false
}

func (ts *TaskScheduler) IsRunning() bool {
	return ts.running
}

var scheduler *TaskScheduler

// ::private::
type TFrmHomeFields struct {
	IsServerRun    bool
	ShowLogCount   int32
	IsTrueClose    bool
	ctx            context.Context // 上下文
	cancel         context.CancelFunc
	LineServerPath string // 排队叫号服务路径
	LineServerPort string // 排队叫号服务端口
}

func (f *TFrmHome) OnFormCreate(sender vcl.IObject) {
	f.IsServerRun = false

	f.TForm.SetPosition(types.PoScreenCenter)
	f.Tray_icon.SetVisible(true)
	f.Tray_icon.SetHint("蜜桔工具")
	f.Tray_icon.SetOnClick(f.OnTrayIconClick)

	f.Timer.SetOnTimer(f.OnTimer)
	f.Meu_item_about.SetOnClick(f.OnAboutClick)
	f.TForm.SetOnClose(f.OnFormClose)
	f.TForm.SetOnDestroy(f.OnFormDestroy)

	f.Meu_add_server.SetOnClick(f.OnServerAddClick)
	f.PM_close.SetOnClick(f.OnAppCloseClick)
	f.PM_show.SetOnClick(f.OnAppShowClick)

	panels := f.StatusBar.Panels()
	panels.Items(1).SetText("停止")

	f.OnStartScheduler()
	f.Monitor()
}

func (f *TFrmHome) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	f.Hide()
	if !f.IsTrueClose {
		*action = types.CaHide
	}
}

func (f *TFrmHome) OnFormDestroy(sender vcl.IObject) {
	if f.cancel != nil {
		f.cancel()
	} else {
		if scheduler != nil {
			scheduler.Stop()
			scheduler = nil
		}
	}
}

func (f *TFrmHome) OnAppCloseClick(sender vcl.IObject) {
	f.IsTrueClose = true
	f.Close()
}

func (f *TFrmHome) OnAppShowClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnAboutClick(sender vcl.IObject) {
	about := about.NewFrmAbout(f)
	about.SetPosition(types.PoOwnerFormCenter)
	about.ShowModal()
}

func (f *TFrmHome) OnServerAddClick(sender vcl.IObject) {
	frm := server_mana.NewFrm_server_mana(f)
	frm.SetPosition(types.PoOwnerFormCenter)
	frm.ShowModal()
}

func (f *TFrmHome) OnTrayIconClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnTimer(sender vcl.IObject) {
	item := f.StatusBar.Panels().Items(3)
	now := time.Now().Format("2006-01-02 15:04:05")
	item.SetText(now)
}

// 添加日志到主窗口中
func (f *TFrmHome) addLog(msg string) {
	vcl.ThreadSync(func() {
		// now := time.Now().Format("2006-01-02 15:04:05")
		// f.MmoRunLog.Lines().Add(fmt.Sprintf("%s %s", now, msg))

		// if f.MmoRunLog.Lines().Count() > f.ShowLogCount {
		// 	f.MmoRunLog.Lines().Delete(0)
		// }

		//移动到最后一行
		// gow32.SendMessageW(f.MmoRunLog.Handle(), gow32.WM_VSCROLL, gow32.SB_BOTTOM, 0)
	})
}

func (f *TFrmHome) OnStartScheduler() {
	// _AutoRestart :=
	// if _AutoRestart {
	// 	if scheduler == nil {
	// 		scheduler = NewTaskScheduler()
	// 	} else {
	// 		scheduler.Stop()
	// 		scheduler = NewTaskScheduler()
	// 	}
	// 	scronString := initialize.ReadStringFromIni("System", "AutoRestartCron", "0 0 0,13,18,21 * * ?")
	// 	scheduler.AddTask(scronString, func() {
	// 		f.OnRunServerClick(nil)
	// 	})
	// 	scheduler.Start()
	// } else {
	// 	if scheduler != nil {
	// 		scheduler.Stop()
	// 		scheduler = nil
	// 	}
	// }
}

func (f *TFrmHome) Monitor() {
	go func() {
		for {
			select {
			case msg := <-global.MsgChannel:
				f.addLog(strings.TrimSpace(msg))
				global.Sugared.Info(strings.TrimSpace(msg))
			}
		}
	}()
}
