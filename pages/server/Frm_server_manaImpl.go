package server

import (
	"HoneyOrangeHelper/internal/config"
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func NewServerManaForm(owner vcl.IComponent, operationType int, data *config.Server) *TFrm_server_mana {
	if Frm_server_mana == nil {
		Frm_server_mana = NewFrm_server_mana(owner)
	}

	Frm_server_mana.OperationType = operationType
	Frm_server_mana.SetPosition(types.PoOwnerFormCenter)
	Frm_server_mana.PageCtl.SetActivePageIndex(0)

	if operationType == 1 && data != nil {
		Frm_server_mana.SetData(data)
	} else {
		Frm_server_mana.ResetData()
	}

	return Frm_server_mana
}

// ::private::
type TFrm_server_manaFields struct {
	data          *config.Server
	OperationType int // 0 新增 1 修改
}

func (f *TFrm_server_mana) ResetData() {
	f.Edt_server_name.Clear()
	f.Edt_server_path.Clear()
	f.Table_params.Clear()
	f.Table_params.SetRowCount(2)
	f.Table_params.SetCells(4, f.Table_params.RowCount()-1, "清空")

	f.Tb_plugin.Clear()
	f.Tb_plugin.SetRowCount(2)
	f.Tb_plugin.SetCells(3, f.Tb_plugin.RowCount()-1, "清空")
}

func (f *TFrm_server_mana) SetData(data *config.Server) {
	f.OperationType = 1
	f.data = data

	f.Table_params.Clear()
	f.Table_params.SetRowCount(2)
	f.Tb_plugin.Clear()
	f.Tb_plugin.SetRowCount(2)

	f.Edt_server_name.SetText(data.Name)
	f.Edt_server_path.SetText(data.Path)
	f.Edt_close_path.SetText(data.CloseScript)

	for index, v := range f.data.Params {
		row := int32(index + 1)
		f.Table_params.SetCells(0, row, v.Key)
		f.Table_params.SetCells(1, row, v.Type)
		f.Table_params.SetCells(2, row, v.Value)
		f.Table_params.SetCells(3, row, v.Mark)
		f.Table_params.SetCells(4, row, "删除")

		f.Table_params.SetRowCount(f.Table_params.RowCount() + 1)

	}

	for index, v := range f.data.Plugins {
		row := int32(index + 1)
		f.Tb_plugin.SetCells(0, row, v.Name)
		f.Tb_plugin.SetCells(1, row, v.Process)
		f.Tb_plugin.SetCells(2, row, v.Mark)
		f.Tb_plugin.SetCells(3, row, "删除")

		f.Tb_plugin.SetRowCount(f.Tb_plugin.RowCount() + 1)
	}

	// 最后一行设置为清空
	f.Table_params.SetCells(4, f.Table_params.RowCount()-1, "清空")
	f.Tb_plugin.SetCells(3, f.Tb_plugin.RowCount()-1, "清空")
}

func (f *TFrm_server_mana) OnFormCreate(sender vcl.IObject) {
	f.InitPage()
	f.SetEvents()
}

func (f *TFrm_server_mana) SetEvents() {
	f.Btn_cancel.SetOnClick(f.OnBtn_cancelClick)
	f.Btn_save.SetOnClick(f.OnBtn_saveClick)
	f.Btn_select_server.SetOnClick(f.OnBtn_select_serverClick)
	f.Table_params.SetOnButtonClick(f.OnTable_paramsOnButtonClick)
	f.Tb_plugin.SetOnButtonClick(f.OnTb_pluginOnButtonClick)
	f.Mi_param_add.SetOnClick(f.OnMi_param_addClick)
	f.Btn_select_close.SetOnClick(f.OnBtn_select_closeClick)
}

func (f *TFrm_server_mana) InitPage() {
	f.Table_params.SetCells(4, 1, "清空")
}

func (f *TFrm_server_mana) OnBtn_select_serverClick(sender vcl.IObject) {
	if f.Od_select_path.Execute() {
		f.Edt_server_path.SetText(f.Od_select_path.FileName())
	}
}

func (f *TFrm_server_mana) OnBtn_select_closeClick(sender vcl.IObject) {
	if f.Od_select_path.Execute() {
		f.Edt_close_path.SetText(f.Od_select_path.FileName())
	}
}

func (f *TFrm_server_mana) OnTb_pluginOnButtonClick(sender vcl.IObject, aCol, aRow int32) {
	// 如果是最后一行，则是清空最后一行的数据
	if aRow == (f.Tb_plugin.RowCount() - 1) {
		f.Tb_plugin.SetCells(0, aRow, "")
		f.Tb_plugin.SetCells(1, aRow, "")
		f.Tb_plugin.SetCells(2, aRow, "")
		return
	}

	rtn := vcl.MessageDlg("确定删除吗？", types.MtWarning)
	fmt.Println(rtn)

	if rtn == 1 {
		f.Tb_plugin.DeleteRow(aRow)
	}
}

func (f *TFrm_server_mana) OnTable_paramsOnButtonClick(sender vcl.IObject, aCol, aRow int32) {
	// 如果是最后一行，则是清空最后一行的数据
	if aRow == (f.Table_params.RowCount() - 1) {
		f.Table_params.SetCells(0, aRow, "")
		f.Table_params.SetCells(1, aRow, "")
		f.Table_params.SetCells(2, aRow, "")
		f.Table_params.SetCells(3, aRow, "")
		return
	}

	rtn := vcl.MessageDlg("确定删除吗？", types.MtWarning)
	fmt.Println(rtn)

	if rtn == 1 {
		f.Table_params.DeleteRow(aRow)
	}
}

func (f *TFrm_server_mana) OnBtn_cancelClick(sender vcl.IObject) {
	f.Close()
}

func (f *TFrm_server_mana) SaveData() {
	params := make([]*config.ServerArg, 0)

	for i := int32(1); i < f.Table_params.RowCount(); i++ {

		if f.Table_params.Cells(0, i) == "" {
			continue
		}

		params = append(params, &config.ServerArg{
			Key:   f.Table_params.Cells(0, i),
			Type:  f.Table_params.Cells(1, i),
			Value: f.Table_params.Cells(2, i),
			Mark:  f.Table_params.Cells(3, i),
		})
	}

	plugins := make([]*config.Plugin, 0)

	for i := int32(1); i < f.Tb_plugin.RowCount(); i++ {

		if f.Tb_plugin.Cells(0, i) == "" {
			continue
		}

		plugins = append(plugins, &config.Plugin{
			Name:    f.Tb_plugin.Cells(0, i),
			Process: f.Tb_plugin.Cells(1, i),
			Mark:    f.Tb_plugin.Cells(2, i),
		})
	}

	srv := &config.Server{
		Name:        f.Edt_server_name.Text(),
		Path:        f.Edt_server_path.Text(),
		CloseScript: f.Edt_close_path.Text(),
		Params:      params,
		Plugins:     plugins,
	}

	if f.OperationType == 0 {
		config.AddServer(srv)
	} else {
		srv.ID = f.data.ID
		config.UpdateServer(srv)
	}
}

func (f *TFrm_server_mana) OnBtn_saveClick(sender vcl.IObject) {
	// 将数据保存到数据库中
	f.SaveData()
	f.Close()
}

func (f *TFrm_server_mana) OnMi_param_addClick(sender vcl.IObject) {
	// 如果最后一行的参数名称没有值，则不允许添加
	if f.Table_params.Cells(0, f.Table_params.RowCount()-1) == "" {
		vcl.ShowMessage("请先填写参数名称")
		return
	}

	f.Table_params.SetRowCount(f.Table_params.RowCount() + 1)
	f.Table_params.SetCells(4, f.Table_params.RowCount()-2, "删除")
	f.Table_params.SetCells(4, f.Table_params.RowCount()-1, "清空")
}
