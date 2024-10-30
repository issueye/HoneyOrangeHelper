package plugin

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
	"HoneyOrangeHelper/internal/utils"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func NewPlugin(owner vcl.IComponent) *TFrm_plugin {
	if Frm_plugin == nil {
		Frm_plugin = NewFrm_plugin(owner)
	}
	Frm_plugin.SetPosition(types.PoOwnerFormCenter)
	return Frm_plugin
}

// ::private::
type TFrm_pluginFields struct {
	list []*config.ToolPlugin
}

func (f *TFrm_plugin) OnFormCreate(sender vcl.IObject) {
	f.list = make([]*config.ToolPlugin, 0)

	f.SetEvents()
}

func (f *TFrm_plugin) SetEvents() {
	f.SetOnShow(f.OnShow)
	f.Btn_query.SetOnClick(f.OnBtn_queryClick)
	f.Btn_add.SetOnClick(f.OnBtn_addClick)
	f.Table_data.SetOnButtonClick(f.OnTable_dataButtonClick)
	f.Btn_refresh_home_menu.SetOnClick(f.OnBtn_refresh_home_menuClick)
}

func (f *TFrm_plugin) OnBtn_addClick(sender vcl.IObject) {
	page := NewPluginForm(f, 0, nil)
	page.ShowModal()
	f.GetData()
}

func (f *TFrm_plugin) OnBtn_refresh_home_menuClick(sender vcl.IObject) {
	msg := message.NewMessage(watermill.NewUUID(), message.Payload([]byte{}))
	global.PubSub.Publish(global.TOPIC_SERVER_REFRESH_TOOL_PLUGIN, msg)
}

func (f *TFrm_plugin) OnBtn_queryClick(sender vcl.IObject) {
	f.GetData()
}

func (f *TFrm_plugin) OnShow(sender vcl.IObject) {
	f.GetData()
}

func (f *TFrm_plugin) OnTable_dataButtonClick(sender vcl.IObject, col, row int32) {
	// 如果是删除按钮
	if col == 5 {
		rt := vcl.MessageDlg("确定要删除吗？", types.MtWarning, types.MbYes, types.MbNo)
		if rt == 6 {
			config.DeleteToolPlugin(f.list[row-1].ID)
			f.GetData()
			return
		}
	}

	// 如果是加载按钮， 则发送加载事件
	if col == 3 {
		msg := message.NewMessage(watermill.NewUUID(), message.Payload(utils.Int64ToBytes(f.list[row-1].ID)))
		global.PubSub.Publish(global.TOPIC_SERVER_LOAD_TOOL_PLUGIN, msg)
	}

	// 如果是卸载按钮， 则发送卸载事件
	if col == 4 {
		msg := message.NewMessage(watermill.NewUUID(), message.Payload(utils.Int64ToBytes(f.list[row-1].ID)))
		global.PubSub.Publish(global.TOPIC_SERVER_UNLOAD_TOOL_PLUGIN, msg)
	}

	// 如果是修改按钮，则弹出修改窗口
	if col == 2 {
		NewPluginForm(f, 1, f.list[row-1]).ShowModal()
		f.GetData()
	}
}

func (f *TFrm_plugin) GetData() {

	// 清空表格数据
	// f.Table_data.ClearRows()
	f.Table_data.Clear()
	f.Table_data.SetRowCount(2)

	list, err := config.GetToolPluginList(f.Edt_condition.Text())
	if err != nil {
		return
	}

	f.list = list

	length := int32(len(list))
	for i := int32(0); i < length; i++ {

		if list[i].Name == "" {
			continue
		}

		row := i + 1

		f.Table_data.SetCells(0, row, list[i].Name)
		f.Table_data.SetCells(1, row, list[i].Path)
		f.Table_data.SetCells(2, row, "修改")
		f.Table_data.SetCells(3, row, "加载")
		f.Table_data.SetCells(4, row, "卸载")
		f.Table_data.SetCells(5, row, "删除")

		// 如果是最后一行，则不再添加新行
		if row == length {
			break
		}

		f.Table_data.SetRowCount(f.Table_data.RowCount() + 1)
	}
}
