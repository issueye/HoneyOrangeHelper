package home

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/pages/about"
	"HoneyOrangeHelper/pages/plugin"
	"HoneyOrangeHelper/pages/server"
	"HoneyOrangeHelper/pages/settings"
	"context"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type SrvObject struct {
	ctx    context.Context
	cancel context.CancelFunc
	Id     int64
	Info   *config.Server
	Frm    *server.TFrm_server_item
}

type ItemMenuObject struct {
	Menu *vcl.TMenuItem
	Info *config.ToolPlugin
}

// ::private::
type TFrmHomeFields struct {
	IsServerRun  bool
	ShowLogCount int32
	IsTrueClose  bool
	ActiveItem   int64
	ctx          context.Context
	cancel       context.CancelFunc

	SrvList []*SrvObject
	IMList  []*ItemMenuObject
}

func (f *TFrmHome) OnFormCreate(sender vcl.IObject) {
	f.IsServerRun = false

	f.TForm.SetPosition(types.PoScreenCenter)
	f.Tray_icon.SetVisible(true)
	f.Tray_icon.SetHint("蜜桔工具")

	f.ctx, f.cancel = context.WithCancel(context.Background())

	f.InitData()
	f.SetEvents()

	f.Monitor()
	f.EventMonitor(f.ctx)
}

func (f *TFrmHome) SetEvents() {
	f.TForm.SetOnClose(f.OnFormClose)
	f.TForm.SetOnDestroy(f.OnFormDestroy)

	f.Tray_icon.SetOnClick(f.OnTrayIconClick)

	f.Timer.SetOnTimer(f.OnTimer)
	f.Meu_item_about.SetOnClick(f.OnAboutClick)
	f.Meu_log_settings.SetOnClick(f.OnMeu_log_settingsClick)
	f.Meu_server_mana.SetOnClick(f.OnMeu_server_manaClick)
	f.Meu_plugin_mana.SetOnClick(f.OnMeu_plugin_manaClick)

	f.PM_close.SetOnClick(f.OnAppCloseClick)
	f.PM_show.SetOnClick(f.OnAppShowClick)
	f.Btn_server_add.SetOnClick(f.OnBtn_server_addClick)
}

func (f *TFrmHome) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	f.Hide()
	if !f.IsTrueClose {
		*action = types.CaHide
	}
}

func (f *TFrmHome) OnFormDestroy(sender vcl.IObject) {
	f.cancel()
}

func (f *TFrmHome) OnAppCloseClick(sender vcl.IObject) {
	f.IsTrueClose = true
	for _, srv := range f.SrvList {
		srv.Frm.Free()
		srv.cancel()
	}
	f.Close()
}

func (f *TFrmHome) OnAppShowClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnBtn_server_addClick(sender vcl.IObject) {
	data := server.GetSelectData(f)
	if data != nil {
		f.AppendServer(data)

		if len(f.SrvList) > 0 {
			f.SrvList[0].Frm.OnBtn_itemClick(nil)
		}
	}
}

func (f *TFrmHome) OnMeu_log_settingsClick(sender vcl.IObject) {
	settings.NewForm(f).ShowModal()
}

func (f *TFrmHome) OnMeu_server_manaClick(sender vcl.IObject) {
	page := server.NewFrm_server_qry(f)
	page.SetPosition(types.PoOwnerFormCenter)
	page.ShowModal()
}

func (f *TFrmHome) OnMeu_plugin_manaClick(sender vcl.IObject) {
	page := plugin.NewPlugin(f)
	page.ShowModal()
}

func (f *TFrmHome) OnAboutClick(sender vcl.IObject) {
	about := about.NewFrmAbout(f)
	about.SetPosition(types.PoOwnerFormCenter)
	about.ShowModal()
}

func (f *TFrmHome) OnServerAddClick(sender vcl.IObject) {
	frm := server.NewFrm_server_mana(f)
	frm.SetPosition(types.PoOwnerFormCenter)
	frm.ShowModal()
}

func (f *TFrmHome) OnTrayIconClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnTimer(sender vcl.IObject) {
	item := f.StatusBar.Panels().Items(1)
	now := time.Now().Format("2006-01-02 15:04:05")
	item.SetText(now)
}
