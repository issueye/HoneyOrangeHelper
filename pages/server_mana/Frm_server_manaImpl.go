package server_mana

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrm_server_manaFields struct {
}

func (f *TFrm_server_mana) OnFormCreate(sender vcl.IObject) {
	f.InitPage()
	f.SetEvents()
}

func (f *TFrm_server_mana) SetEvents() {
	f.Btn_select_server.SetOnClick(f.OnBtn_select_serverClick)
	f.Btn_param_add.SetOnClick(f.OnBtn_param_addClick)
	f.Table_params.SetOnButtonClick(f.OnTable_paramsOnButtonClick)
}

func (f *TFrm_server_mana) InitPage() {
	f.Table_params.SetCells(3, 1, "")
}

func (f *TFrm_server_mana) OnBtn_select_serverClick(sender vcl.IObject) {
	if f.Od_select_path.Execute() {
		f.Edt_server_path.SetText(f.Od_select_path.FileName())
	}
}

func (f *TFrm_server_mana) OnBtn_param_addClick(sender vcl.IObject) {
	f.Table_params.SetCells(0, f.Table_params.RowCount()-1, f.Edt_param_key.Text())
	f.Table_params.SetCells(1, f.Table_params.RowCount()-1, fmt.Sprintf("%d", f.Rg_param_type.ItemIndex()))
	f.Table_params.SetCells(2, f.Table_params.RowCount()-1, f.Edt_param_value.Text())
	f.Table_params.SetCells(3, f.Table_params.RowCount()-1, f.Edt_param_remark.Text())
	f.Table_params.SetCells(4, f.Table_params.RowCount()-1, "删除")

	f.Edt_param_key.SetText("")
	f.Edt_param_value.SetText("")
	f.Edt_param_remark.SetText("")

	f.Table_params.SetRowCount(f.Table_params.RowCount() + 1)
}

func (f *TFrm_server_mana) OnTable_paramsOnButtonClick(sender vcl.IObject, aCol, aRow int32) {
	if aRow == (f.Table_params.RowCount() - 1) {
		return
	}

	rtn := vcl.MessageDlg("确定删除吗？", types.MtWarning)
	fmt.Println(rtn)

	if rtn == 1 {
		f.Table_params.DeleteRow(aRow)
	}
}
