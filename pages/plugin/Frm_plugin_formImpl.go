package plugin

import (
	"HoneyOrangeHelper/internal/config"
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func NewPluginForm(owner vcl.IComponent, operationType int, data *config.ToolPlugin) *TFrm_plugin_form {
	if Frm_plugin_form == nil {
		Frm_plugin_form = NewFrm_plugin_form(owner)
	}

	Frm_plugin_form.OperationType = operationType
	Frm_plugin_form.SetPosition(types.PoOwnerFormCenter)
	Frm_plugin_form.PageCtl.SetActivePageIndex(0)

	if operationType == 1 && data != nil {
		Frm_plugin_form.SetData(data)
	} else {
		Frm_plugin_form.ResetData()
	}

	return Frm_plugin_form
}

// ::private::
type TFrm_plugin_formFields struct {
	data          *config.ToolPlugin
	OperationType int // 0 新增 1 修改
}

func (f *TFrm_plugin_form) ResetData() {
	f.Edt_name.Clear()
	f.Edt_process_path.Clear()

	f.Table_event.SetRowCount(6)

	f.Table_event.SetCells(0, 1, "0")
	f.Table_event.SetCells(1, 1, "server_start")
	f.Table_event.SetCells(2, 1, "服务启动")
	f.Table_event.SetCells(3, 1, "启动服务时进行推送")

	f.Table_event.SetCells(0, 2, "0")
	f.Table_event.SetCells(1, 2, "server_stop")
	f.Table_event.SetCells(2, 2, "服务停止")
	f.Table_event.SetCells(3, 2, "停止服务时进行推送")

	f.Table_event.SetCells(0, 3, "0")
	f.Table_event.SetCells(1, 3, "server_setting_ok")
	f.Table_event.SetCells(2, 3, "服务设置完成")
	f.Table_event.SetCells(3, 3, "设置服务信息完成时进行推送")

	f.Table_event.SetCells(0, 4, "0")
	f.Table_event.SetCells(1, 4, "server_remove")
	f.Table_event.SetCells(2, 4, "服务移除")
	f.Table_event.SetCells(3, 4, "移除服务时进行推送")

	f.Table_event.SetCells(0, 5, "0")
	f.Table_event.SetCells(1, 5, "server_add")
	f.Table_event.SetCells(2, 5, "服务添加")
	f.Table_event.SetCells(3, 5, "添加服务时进行推送")
}

func (f *TFrm_plugin_form) SetData(data *config.ToolPlugin) {
	f.OperationType = 1
	f.data = data

	f.Table_event.Clear()
	f.Table_event.SetRowCount(2)

	f.Edt_name.SetText(data.Name)
	f.Edt_process_path.SetText(data.Path)

	for index, v := range f.data.Events {
		row := int32(index + 1)

		check := "0"
		if v.IsCheck {
			check = "1"
		}

		f.Table_event.SetCells(0, row, check)
		f.Table_event.SetCells(1, row, v.Name)
		f.Table_event.SetCells(2, row, v.Title)
		f.Table_event.SetCells(3, row, v.Mark)

		if !(index == len(f.data.Events)-1) {
			f.Table_event.SetRowCount(f.Table_event.RowCount() + 1)
		}
	}
}

func (f *TFrm_plugin_form) OnFormCreate(sender vcl.IObject) {
	f.InitPage()
	f.SetEvents()
}

func (f *TFrm_plugin_form) SetEvents() {
	f.Btn_cancel.SetOnClick(f.OnBtn_cancelClick)
	f.Btn_save.SetOnClick(f.OnBtn_saveClick)
	f.Btn_select.SetOnClick(f.OnBtn_selectClick)
	f.Table_event.SetOnButtonClick(f.OnTable_eventOnButtonClick)
}

func (f *TFrm_plugin_form) InitPage() {
}

func (f *TFrm_plugin_form) OnBtn_selectClick(sender vcl.IObject) {
	if f.Od_select_path.Execute() {
		f.Edt_process_path.SetText(f.Od_select_path.FileName())
	}
}

func (f *TFrm_plugin_form) OnBtn_select_closeClick(sender vcl.IObject) {
	if f.Od_select_path.Execute() {
		f.Edt_process_path.SetText(f.Od_select_path.FileName())
	}
}

func (f *TFrm_plugin_form) OnTable_eventOnButtonClick(sender vcl.IObject, aCol, aRow int32) {
	// 如果是最后一行，则是清空最后一行的数据
	if aRow == (f.Table_event.RowCount() - 1) {
		f.Table_event.SetCells(0, aRow, "")
		f.Table_event.SetCells(1, aRow, "")
		f.Table_event.SetCells(2, aRow, "")
		f.Table_event.SetCells(3, aRow, "")
		return
	}

	rtn := vcl.MessageDlg("确定删除吗？", types.MtWarning)
	fmt.Println(rtn)

	if rtn == 1 {
		f.Table_event.DeleteRow(aRow)
	}
}

func (f *TFrm_plugin_form) OnBtn_cancelClick(sender vcl.IObject) {
	f.Close()
}

func (f *TFrm_plugin_form) SaveData() {
	events := make([]*config.Event, 0)
	for i := int32(1); i < f.Table_event.RowCount(); i++ {

		if f.Table_event.Cells(0, i) == "" {
			continue
		}

		event := &config.Event{
			Name:  f.Table_event.Cells(1, i),
			Title: f.Table_event.Cells(2, i),
			Mark:  f.Table_event.Cells(3, i),
		}

		if f.Table_event.Cells(0, i) == "1" {
			event.IsCheck = true
		}

		events = append(events, event)
	}

	tool := &config.ToolPlugin{
		Name:        f.Edt_name.Text(),
		Path:        f.Edt_process_path.Text(),
		CookieKey:   f.Edt_cookie_key.Text(),
		CookieValue: f.Edt_cookie_value.Text(),
		Events:      events,
	}

	if f.OperationType == 0 {
		config.AddToolPlugin(tool)
	} else {
		tool.ID = f.data.ID
		config.UpdateToolPlugin(tool)
	}
}

func (f *TFrm_plugin_form) OnBtn_saveClick(sender vcl.IObject) {
	// 将数据保存到数据库中
	f.SaveData()
	f.SetModalResult(types.MrOk)
	f.Close()
}
