package server

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/internal/helper_cmd"
	"context"
	"path/filepath"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type PluginBtn struct {
	Btn            *vcl.TBitBtn
	Info           *config.Plugin
	parent         vcl.IWinControl
	bs             vcl.TControlBorderSpacing
	ctx            context.Context
	cancel         context.CancelFunc
	getServerState func() bool
}

func NewPluginBtn(owner vcl.IComponent, parent vcl.IWinControl, bs vcl.TControlBorderSpacing, info *config.Plugin, getServerState func() bool) *PluginBtn {
	btn := &PluginBtn{
		Btn:            vcl.NewBitBtn(owner),
		Info:           info,
		parent:         parent,
		bs:             bs,
		getServerState: getServerState,
	}

	btn.SetAttribute()
	btn.SetEvents()

	return btn
}

func (f *PluginBtn) SetAttribute() {
	f.Btn.SetCaption(f.Info.Name)
	f.Btn.SetAlign(types.AlLeft)
	f.Btn.SetParent(f.parent)
	f.Btn.SetVisible(true)

	f.Btn.SetBorderSpacing(&f.bs)
}

func (f *PluginBtn) SetEvents() {
	f.Btn.SetOnClick(func(sender vcl.IObject) {

		if f.getServerState != nil && f.getServerState() {
			vcl.ShowMessage("服务运行中，请先停止服务")
			return
		}

		f.OpenPlugin()
	})
}

func (f *PluginBtn) OpenPlugin() {
	// vcl.ShowMessage(f.Info.Name)
	f.ctx, f.cancel = context.WithCancel(context.Background())
	path := filepath.Join(global.ROOT_PATH, "plugins", f.Info.Process)
	msgChan, err := helper_cmd.Run(0)(f.ctx, false, path)
	if err != nil {
		vcl.ShowMessage(err.Error())
	}

	go func() {
		for msg := range msgChan {
			global.Sugared.Debugf("Plugin %s: %s", f.Info.Name, msg)
		}
	}()
}
