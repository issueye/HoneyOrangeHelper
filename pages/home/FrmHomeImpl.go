package home

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/pages/about"
	"HoneyOrangeHelper/pages/server"
	"HoneyOrangeHelper/pages/server_mana"
	"HoneyOrangeHelper/pages/server_query"
	"HoneyOrangeHelper/pages/settings"
	"HoneyOrangeHelper/pkg/utils"
	"context"
	"slices"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type SrvObject struct {
	Id   int64
	Info *config.Server
	Frm  *server.TFrm_server_item
}

// ::private::
type TFrmHomeFields struct {
	IsServerRun  bool
	ShowLogCount int32
	IsTrueClose  bool
	ActiveItem   int64

	SrvList []*SrvObject
}

func (f *TFrmHome) OnFormCreate(sender vcl.IObject) {
	f.IsServerRun = false

	f.TForm.SetPosition(types.PoScreenCenter)
	f.Tray_icon.SetVisible(true)
	f.Tray_icon.SetHint("蜜桔工具")

	f.InitData()
	f.SetEvents()

	f.Monitor()

	f.EventMonitor()
}

func (f *TFrmHome) EventMonitor() {
	messages, err := global.PubSub.Subscribe(context.Background(), global.TOPIC_SERVER_REMOVE)
	if err != nil {
		return
	}
	go func() {
		for msgInfo := range messages {
			err := f.closeItem(msgInfo)
			if err != nil {
				global.Sugared.Errorf("关闭项目失败 %s", err.Error())
				continue
			}
		}
	}()
}

func (f *TFrmHome) closeItem(msgInfo *message.Message) error {
	defer msgInfo.Ack()

	data := string(msgInfo.Payload)
	cfg, err := config.FromToJson(data)
	if err != nil {
		return err
	}

	f.SrvList = slices.DeleteFunc(f.SrvList, func(v *SrvObject) bool {
		if v.Id == cfg.Code {
			v.Frm.CloseWindow()
			return true
		}

		return false
	})

	return nil
}

func (f *TFrmHome) InitData() {
	f.ActiveItem = 0
	f.SrvList = make([]*SrvObject, 0)

	list, err := config.GetServerList("")
	if err != nil {
		return
	}

	// 加载到主页面中，添加 PageControl 页面
	for _, v := range list {
		f.AppendServer(v)
	}

	if len(f.SrvList) > 0 {
		f.ActiveItem = f.SrvList[0].Info.Code
		f.SrvList[0].Frm.OnBtn_itemClick(nil)
	}
}

func (f *TFrmHome) AppendServer(cfg *config.Server) {
	Id := utils.GenID()
	cfg.Code = Id
	frm := server.NewItemForm(f, Id, f.Sb_box, f.Pnl_body, cfg)
	obj := &SrvObject{Info: cfg, Frm: frm}
	obj.Id = Id
	f.SrvList = append(f.SrvList, obj)
	frm.Show()
}

func (f *TFrmHome) SetEvents() {
	f.Tray_icon.SetOnClick(f.OnTrayIconClick)

	f.Timer.SetOnTimer(f.OnTimer)
	f.Meu_item_about.SetOnClick(f.OnAboutClick)
	f.Meu_log_settings.SetOnClick(f.OnMeu_log_settingsClick)

	f.TForm.SetOnClose(f.OnFormClose)
	f.TForm.SetOnDestroy(f.OnFormDestroy)

	f.PM_close.SetOnClick(f.OnAppCloseClick)
	f.PM_show.SetOnClick(f.OnAppShowClick)
	f.Meu_server_mana.SetOnClick(f.OnMeu_server_manaClick)
	f.Btn_server_add.SetOnClick(f.OnBtn_server_addClick)
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

func (f *TFrmHome) OnBtn_server_addClick(sender vcl.IObject) {
	data := GetSelectData(f)
	if data != nil {
		f.AppendServer(data)

		if len(f.SrvList) > 0 {
			f.SrvList[0].Frm.OnBtn_itemClick(nil)
		}
	}
}

func (f *TFrmHome) OnMeu_log_settingsClick(sender vcl.IObject) {
	page := settings.NewForm(f)
	page.ShowModal()
}

func (f *TFrmHome) OnMeu_server_manaClick(sender vcl.IObject) {
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
