package server

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func NewItemForm(owner vcl.IComponent, Id int64, parent vcl.IWinControl, body vcl.IWinControl, data *config.Server) *TFrm_server_item {
	page := NewFrm_server_item(owner)
	page.data = data
	page.Id = Id
	page.BodyPnl = body
	page.SetParent(parent)
	page.SetAlign(types.AlLeft)
	page.SetBorderStyle(types.BsNone)
	page.Btn_remove.SetVisible(false)
	page.ServerForm = NewServerForm(owner, body, data)
	return page
}

// ::private::
type TFrm_server_itemFields struct {
	Id         int64
	data       *config.Server
	ServerForm *TFrm_server
	BodyPnl    vcl.IWinControl
}

func (f *TFrm_server_item) CloseWindow() {
	f.ServerForm.CloseWindow()
	f.Close()
}

func (f *TFrm_server_item) OnFormCreate(sender vcl.IObject) {
	f.SetEvents()
}

func (f *TFrm_server_item) OnFormShow(sender vcl.IObject) {
	f.InitData()
}

func (f *TFrm_server_item) InitData() {
	f.Btn_item.SetCaption(f.data.Name)
}

func (f *TFrm_server_item) SetEvents() {
	f.SetOnShow(f.OnFormShow)
	f.Btn_item.SetOnClick(f.OnBtn_itemClick)
	f.Btn_remove.SetOnClick(f.OnImg_removeClick)
	f.SetOnDestroy(f.OnFormDestroy)
}

func (f *TFrm_server_item) OnBtn_itemClick(sender vcl.IObject) {
	f.ServerForm.Show()
}

func (f *TFrm_server_item) OnFormDestroy(sender vcl.IObject) {
	f.ServerForm.Close()
}

func (f *TFrm_server_item) OnImg_removeClick(sender vcl.IObject) {
	msg := message.NewMessage(watermill.NewUUID(), []byte(f.data.ToJson()))
	global.PubSub.Publish(global.TOPIC_SERVER_REMOVE, msg)
}
