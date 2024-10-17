package server

import (
	"HoneyOrangeHelper/internal/config"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func NewServerQryForm(owner vcl.IComponent) *TFrm_server_qry {
	if Frm_server_qry == nil {
		Frm_server_qry = NewFrm_server_qry(owner)
	}
	Frm_server_qry.SetPosition(types.PoOwnerFormCenter)
	return Frm_server_qry
}

// ::private::
type TFrm_server_qryFields struct {
	list []*config.Server
}

func (f *TFrm_server_qry) OnFormCreate(sender vcl.IObject) {
	f.list = make([]*config.Server, 0)

	f.SetEvents()
}

func (f *TFrm_server_qry) SetEvents() {
	f.SetOnShow(f.OnShow)
	f.Btn_query.SetOnClick(f.OnBtn_queryClick)
	f.Btn_add.SetOnClick(f.OnBtn_addClick)
	f.Table_data.SetOnButtonClick(f.OnTable_dataButtonClick)
}

func (f *TFrm_server_qry) OnBtn_addClick(sender vcl.IObject) {
	page := NewServerManaForm(f, 0, nil)
	page.ShowModal()
	f.GetData()
}

func (f *TFrm_server_qry) OnBtn_queryClick(sender vcl.IObject) {
	f.GetData()
}

func (f *TFrm_server_qry) OnShow(sender vcl.IObject) {
	f.GetData()
}

func (f *TFrm_server_qry) OnTable_dataButtonClick(sender vcl.IObject, col, row int32) {
	// 如果是删除按钮
	if col == 3 {
		rt := vcl.MessageDlg("确定要删除吗？", types.MtWarning, types.MbYes, types.MbNo)
		if rt == 6 {
			config.DeleteServer(f.list[row-1].ID)
			f.GetData()
			return
		}
	}

	// 如果是修改按钮，则弹出修改窗口
	if col == 2 {
		NewServerManaForm(f, 1, f.list[row-1]).ShowModal()
		f.GetData()
	}
}

func (f *TFrm_server_qry) GetData() {

	// 清空表格数据
	// f.Table_data.ClearRows()
	f.Table_data.Clear()
	f.Table_data.SetRowCount(2)

	list, err := config.GetServerList(f.Edt_condition.Text())
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
		f.Table_data.SetCells(3, row, "删除")

		// 如果是最后一行，则不再添加新行
		if row == length {
			break
		}

		f.Table_data.SetRowCount(f.Table_data.RowCount() + 1)
	}
}
