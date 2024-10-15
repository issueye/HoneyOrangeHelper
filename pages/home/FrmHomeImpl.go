package home

import (
	"HoneyOrangeHelper/pages/about"
	"HoneyOrangeHelper/pages/server_mana"
	"HoneyOrangeHelper/pages/server_query"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrmHomeFields struct {
	IsServerRun    bool
	ShowLogCount   int32
	IsTrueClose    bool
	LineServerPath string // 排队叫号服务路径
	LineServerPort string // 排队叫号服务端口
}

func (f *TFrmHome) OnFormCreate(sender vcl.IObject) {
	f.IsServerRun = false

	f.TForm.SetPosition(types.PoScreenCenter)
	f.Tray_icon.SetVisible(true)
	f.Tray_icon.SetHint("蜜桔工具")

	f.SetEvents()

	f.Monitor()
}

func (f *TFrmHome) SetEvents() {
	f.Tray_icon.SetOnClick(f.OnTrayIconClick)

	f.Timer.SetOnTimer(f.OnTimer)
	f.Meu_item_about.SetOnClick(f.OnAboutClick)
	f.TForm.SetOnClose(f.OnFormClose)
	f.TForm.SetOnDestroy(f.OnFormDestroy)

	f.Meu_add_server.SetOnClick(f.OnServerAddClick)
	f.PM_close.SetOnClick(f.OnAppCloseClick)
	f.PM_show.SetOnClick(f.OnAppShowClick)
	f.Meu_server_list.SetOnClick(f.OnMeu_server_listClick)
}

func (f *TFrmHome) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	f.Hide()
	if !f.IsTrueClose {
		*action = types.CaHide
	}
}

func (f *TFrmHome) OnFormDestroy(sender vcl.IObject) {
}

func (f *TFrmHome) OnAppCloseClick(sender vcl.IObject) {
	f.IsTrueClose = true
	f.Close()
}

func (f *TFrmHome) OnAppShowClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnMeu_server_listClick(sender vcl.IObject) {
	page := server_query.NewFrm_server_qry(f)
	page.SetPosition(types.PoOwnerFormCenter)
	page.ShowModal()
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
	item := f.StatusBar.Panels().Items(1)
	now := time.Now().Format("2006-01-02 15:04:05")
	item.SetText(now)
}
