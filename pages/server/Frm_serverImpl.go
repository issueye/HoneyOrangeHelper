package server

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/internal/gow32"
	"HoneyOrangeHelper/internal/helper_cmd"
	"HoneyOrangeHelper/internal/utils"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/issueye/ipc_grpc/grpc/pb"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewServerForm(ctx context.Context, owner vcl.IComponent, parent vcl.IWinControl, data *config.Server) *TFrm_server {
	page := NewFrm_server(owner)
	page.data = data
	page.SetParent(parent)
	page.SetAlign(types.AlClient)
	page.SetBorderStyle(types.BsNone)
	page.Pnl_actions.SetAlignment(types.TaLeftJustify)
	page.SetData(data)
	page.ServerMessage = make(chan string, 100)
	page.Monitor(ctx)
	return page
}

// ::private::
type TFrm_serverFields struct {
	ShowLogCount     int32 // 显示日志条数
	data             *config.Server
	IsRunning        bool
	ctx              context.Context
	cancel           context.CancelFunc
	log              *zap.Logger
	sugar            *zap.SugaredLogger
	PluginBtns       []*PluginBtn
	lumberJackLogger *lumberjack.Logger
	runResult        *helper_cmd.RunResult
	ServerMessage    chan string
}

func (f *TFrm_server) OnFormCreate(sender vcl.IObject) {
	f.InitData()
	f.SetEvents()
}

func (f *TFrm_server) CloseWindow() {
	for _, btn := range f.PluginBtns {
		btn.Btn.Free()
	}

	f.Close()
}

func (f *TFrm_server) SetEvents() {
	f.Btn_remove_server.SetOnClick(f.OnBtn_remove_serverClick)
	f.Btn_server_run.SetOnClick(f.OnBtn_server_runClick)
	f.Btn_server_settings.SetOnClick(f.OnBtn_server_settingsClick)
	f.Btn_clear_log.SetOnClick(f.OnBtn_clear_logClick)
}

func (f *TFrm_server) OnBtn_clear_logClick(sender vcl.IObject) {
	f.Mmo_run_log.Clear()
}

func (f *TFrm_server) OnBtn_server_settingsClick(sender vcl.IObject) {
	if f.IsRunning {
		vcl.ShowMessage("请先停止服务")
		return
	}

	NewServerManaForm(f, 1, f.data).ShowModal()
	info, err := config.GetServer(f.data.ID)
	if err != nil {
		vcl.ShowMessage(err.Error())
		return
	}

	f.SetData(info)
}

func (f *TFrm_server) SetData(data *config.Server) {
	f.Lbl_id.SetCaption(fmt.Sprintf("ID: %d", data.Code))
	f.Lbl_name.SetCaption(fmt.Sprintf("名称: %s", data.Name))
	f.Lbl_path.SetCaption(fmt.Sprintf("路径: %s", data.Path))

	code := f.data.Code
	f.data = data
	f.data.Code = code

	// 释放所有插件按钮
	for _, btn := range f.PluginBtns {
		btn.Btn.Free()
	}

	// 处理插件
	f.PluginBtns = make([]*PluginBtn, 0)
	for _, p := range f.data.Plugins {
		btn := NewPluginBtn(f, f.Pnl_actions, *f.Btn_server_run.BorderSpacing(), p, func() bool {
			return f.IsRunning
		})
		f.PluginBtns = append(f.PluginBtns, btn)
	}
}

func (f *TFrm_server) OnBtn_remove_serverClick(sender vcl.IObject) {
	if f.IsRunning {
		vcl.ShowMessage("请先停止服务")
		return
	}

	msg := message.NewMessage(watermill.NewUUID(), message.Payload(f.data.ToJson()))
	global.PubSub.Publish(global.TOPIC_SERVER_REMOVE, msg)
}

func (f *TFrm_server) OnBtn_server_runClick(sender vcl.IObject) {
	if f.IsRunning {
		f.StopServer()
	} else {
		f.RunServer()
	}
}

func (f *TFrm_server) StopServer() {
	f.IsRunning = false
	f.Btn_server_run.SetImageIndex(0)
	f.Btn_server_run.SetCaption("启动")
	f.ShowLogCount = int32(config.GetParam(config.LOG, "show-log-count", 100).Int())

	defer func() {
		defer f.cancel()

		code := strconv.FormatInt(f.data.Code, 10)
		global.PluginSrv.Events.PushEvent(code, pb.EventType_ET_STOP, &pb.ServerInfo{
			Id:          code,
			Name:        f.data.Name,
			ProcessName: f.data.Path,
			State:       pb.StateType_ST_STOP,
		})

		if f.log != nil {
			// f.addLog("服务已停止...")
			f.log.Info("服务已停止")
			f.log.Sync()
			f.log = nil
			f.sugar = nil

			f.lumberJackLogger.Close()
		}
	}()

	f.ServerMessage <- "停止服务中..."

	if f.data.CloseScript != "" {
		f.ServerMessage <- "执行关闭脚本..."
		f.ServerMessage <- fmt.Sprintf("脚本路径: %s", f.data.CloseScript)

		_, err := helper_cmd.Run(20)(f.ctx, true, f.data.CloseScript)
		if err != nil {
			f.ServerMessage <- fmt.Sprintf("执行关闭脚本失败: %s", err.Error())
			return
		}
	}

	p := new(utils.Process)
	p.Msg = f.ServerMessage
	err := p.KillProcessAndChildren(f.runResult.Pid)
	if err != nil {
		f.ServerMessage <- fmt.Sprintf("停止服务失败: %s", err.Error())
		return
	}
}

func (f *TFrm_server) RunServer() error {
	f.Mmo_run_log.Clear()

	var err error
	defer func() {
		f.IsRunning = true
		f.Btn_server_run.SetImageIndex(1)
		f.Btn_server_run.SetCaption("停止")

		code := strconv.FormatInt(f.data.Code, 10)
		global.PluginSrv.Events.PushEvent(code, pb.EventType_ET_START, &pb.ServerInfo{
			Id:          code,
			Name:        f.data.Name,
			ProcessName: f.data.Path,
			State:       pb.StateType_ST_START,
		})

		if err != nil {
			f.addLog("启动失败：" + err.Error())
			f.IsRunning = false
			f.Btn_server_run.SetImageIndex(0)
			f.Btn_server_run.SetCaption("启动")
		}
	}()

	f.ctx, f.cancel = context.WithCancel(context.Background())
	params := make([]string, 0)
	for _, param := range f.data.Params {
		if param.Type == "运行参数" {
			params = append(params, fmt.Sprintf("%s=%s", param.Key, param.Value))
		}
	}

	f.InitLogger()

	f.runResult, err = helper_cmd.Run(200)(f.ctx, true, f.data.Path, params...)
	if err != nil {
		return err

	}

	f.RunMonitor(f.ctx, f.runResult.Msg)
	return nil
}

func (f *TFrm_server) InitData() {
	f.SetShowHint(true)
	f.Btn_server_run.SetShowCaption(false)
	f.PluginBtns = make([]*PluginBtn, 0)
	f.IsRunning = false
	f.ShowLogCount = 1000
	f.Btn_server_run.SetImageIndex(0)
	f.Btn_server_run.SetCaption("启动")

}

// 添加日志到主窗口中
func (f *TFrm_server) addLog(msg string) {
	vcl.ThreadSync(func() {
		// now := time.Now().Format("2006-01-02 15:04:05")
		f.Mmo_run_log.Lines().Add(msg)

		if f.Mmo_run_log.Lines().Count() > f.ShowLogCount {
			f.Mmo_run_log.Lines().Delete(0)
		}

		// 移动到最后一行
		gow32.SendMessageW(f.Mmo_run_log.Handle(), gow32.WM_VSCROLL, gow32.SB_BOTTOM, 0)
	})
}

func (f *TFrm_server) RunMonitor(ctx context.Context, msgChan chan string) {
	go func(_ context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-msgChan:
				if msg != "" {
					datas := strings.Split(msg, "\n\t")
					for _, data := range datas {
						f.addLog(data)
						if f.sugar != nil {
							f.sugar.Info(data)
						}
					}
				}
			}
		}
	}(ctx)
}

func (f *TFrm_server) Monitor(ctx context.Context) {
	go func(c context.Context) {
		for {
			select {
			case msg := <-f.ServerMessage:
				f.addLog(msg)
			case <-c.Done():
				return
			}
		}
	}(ctx)
}
