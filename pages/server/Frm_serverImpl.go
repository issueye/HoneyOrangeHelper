package server

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/internal/gow32"
	"HoneyOrangeHelper/internal/helper_cmd"
	"HoneyOrangeHelper/pkg/logger"
	"HoneyOrangeHelper/pkg/utils"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewServerForm(owner vcl.IComponent, parent vcl.IWinControl, data *config.Server) *TFrm_server {
	page := NewFrm_server(owner)
	page.data = data
	page.SetParent(parent)
	page.SetAlign(types.AlClient)
	page.SetBorderStyle(types.BsNone)
	page.Pnl_actions.SetAlignment(types.TaLeftJustify)
	page.SetData(data)
	return page
}

// ::private::
type TFrm_serverFields struct {
	ShowLogCount int32 // 显示日志条数
	data         *config.Server
	IsRunning    bool
	ctx          context.Context
	cancel       context.CancelFunc
	log          *zap.Logger
	sugar        *zap.SugaredLogger
	PluginBtns   []*PluginBtn
}

func (f *TFrm_server) OnFormCreate(sender vcl.IObject) {
	f.InitData()
	f.SetEvents()
}

func (f *TFrm_server) InitLogger() {
	cfg := new(logger.Config)
	cfg.Level = config.GetParam(config.LOG, "level", 0).Int()
	path := fmt.Sprintf("%s/logs/%d", global.ROOT_PATH, f.data.Code)
	utils.IsExistsCreatePath(fmt.Sprintf("%s/logs", global.ROOT_PATH), fmt.Sprintf("%d", f.data.Code))
	cfg.Path = path
	cfg.MaxSize = config.GetParam(config.LOG, "max-size", 100).Int() // MB
	cfg.MaxBackups = config.GetParam(config.LOG, "max-backups", 100).Int()
	cfg.MaxAge = config.GetParam(config.LOG, "max-age", 100).Int() // days
	cfg.Compress = config.GetParam(config.LOG, "compress", true).Bool()
	f.sugar, f.log = logger.InitLogger(cfg)
}

func (f *TFrm_server) CloseWindow() {
	f.log.Info("关闭窗口")
	f.log.Sync()
	f.log = nil
	f.sugar = nil

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
		f.IsRunning = false
		f.Btn_server_run.SetImageIndex(0)
		f.Btn_server_run.SetCaption("启动")
		f.StopServer()
	} else {
		f.IsRunning = true
		f.Btn_server_run.SetImageIndex(1)
		f.Btn_server_run.SetCaption("停止")
		f.RunServer()
	}
}

func (f *TFrm_server) InitLogger() {
	path := fmt.Sprintf("%s/logs/%d", global.ROOT_PATH, f.data.Code)
	utils.IsExistsCreatePath(fmt.Sprintf("%s/logs", global.ROOT_PATH), fmt.Sprintf("%d", f.data.Code))

	fp := path + "/server.log"

	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   fp,
				MaxSize:    100, // megabytes
				MaxBackups: 10,
				MaxAge:     28, // days
				Compress:   true,
			}),
			zapcore.DebugLevel,
		),
	)
}

func (f *TFrm_server) StopServer() {
	f.cancel()

	f.log.Info("关闭窗口")
	f.log.Sync()
	f.log = nil
	f.sugar = nil
}

func (f *TFrm_server) RunServer() {
	f.Mmo_run_log.Clear()

	f.ctx, f.cancel = context.WithCancel(context.Background())
	params := make([]string, 0)
	for _, param := range f.data.Params {
		if param.Type == "运行参数" {
			params = append(params, fmt.Sprintf("%s=%s", param.Key, param.Value))
		}
	}

	f.InitLogger()

	msgChan, err := helper_cmd.Run(f.ctx, true, f.data.Path, params...)
	if err != nil {
		f.addLog(err.Error())
	}

	f.Monitor(msgChan)
}

func (f *TFrm_server) InitData() {
	f.PluginBtns = make([]*PluginBtn, 0)
	f.IsRunning = false
	f.ShowLogCount = 1000
	f.Btn_server_run.SetImageIndex(0)
	f.Btn_server_run.SetCaption("启动")

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

func (f *TFrm_server) Monitor(msgChan chan string) {
	go func() {
		for msg := range msgChan {
			f.addLog(strings.TrimSpace(msg))
			f.sugar.Info(strings.TrimSpace(msg))
		}
	}()
}
