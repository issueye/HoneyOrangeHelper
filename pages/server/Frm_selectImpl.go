package server

import (
	"HoneyOrangeHelper/internal/config"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func NewSelectForm(owner vcl.IComponent) *TFrm_select {
	if Frm_select == nil {
		Frm_select = NewFrm_select(owner)
	}

	Frm_select.SetPosition(types.PoOwnerFormCenter)
	return Frm_select
}

func GetSelectData(owner vcl.IComponent) *config.Server {
	NewSelectForm(owner).ShowModal()
	return Frm_select.IndexData
}

// ::private::
type TFrm_selectFields struct {
	list      []*config.Server
	IndexData *config.Server
}

func (f *TFrm_select) OnFormCreate(sender vcl.IObject) {
	f.list = make([]*config.Server, 0)
	f.SetCaption("选择服务")
	f.SetEvents()
}

func (f *TFrm_select) SetEvents() {
	f.SetOnShow(f.OnShow)
	f.Btn_query.SetOnClick(f.OnBtn_queryClick)
	f.Table_data.SetOnButtonClick(f.OnTable_dataButtonClick)
}

func (f *TFrm_select) OnBtn_queryClick(sender vcl.IObject) {
	f.GetData()
}

func (f *TFrm_select) OnShow(sender vcl.IObject) {
	f.IndexData = nil
	f.GetData()
}

func (f *TFrm_select) OnTable_dataButtonClick(sender vcl.IObject, col, row int32) {
	// 如果是选择按钮，则关闭按钮
	if col == 2 {
		f.IndexData = f.list[row-1]
		f.Close()
	}
}

func (f *TFrm_select) GetData() {

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
		f.Table_data.SetCells(2, row, "选择")

		// 如果是最后一行，则不再添加新行
		if row == length {
			break
		}

		f.Table_data.SetRowCount(f.Table_data.RowCount() + 1)
	}
}
