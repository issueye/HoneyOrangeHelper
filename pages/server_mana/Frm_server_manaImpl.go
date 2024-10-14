package server_mana

import (
	"HoneyOrangeHelper/internal/config"
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
	f.Btn_cancel.SetOnClick(f.OnBtn_cancelClick)
	f.Btn_save.SetOnClick(f.OnBtn_saveClick)
	f.Btn_select_server.SetOnClick(f.OnBtn_select_serverClick)
	f.Table_params.SetOnButtonClick(f.OnTable_paramsOnButtonClick)
	f.Mi_param_add.SetOnClick(f.OnMi_param_addClick)
}

func (f *TFrm_server_mana) InitPage() {
	f.Table_params.SetCells(4, 1, "清空")
}

func (f *TFrm_server_mana) OnBtn_select_serverClick(sender vcl.IObject) {
	if f.Od_select_path.Execute() {
		f.Edt_server_path.SetText(f.Od_select_path.FileName())
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

func (f *TFrm_server_mana) OnBtn_saveClick(sender vcl.IObject) {
	// 将数据保存到数据库中

	params := make([]*config.ServerArg, 0)

	for i := int32(1); i < f.Table_params.RowCount(); i++ {

		if f.Table_params.Cells(0, i) == "" {
			continue
		}

		params = append(params, &config.ServerArg{
			Key:   f.Table_params.Cells(0, i),
			Value: f.Table_params.Cells(1, i),
			Type:  f.Table_params.Cells(2, i),
			Mark:  f.Table_params.Cells(3, i),
		})
	}

	config.AddServer(&config.Server{
		Name:   f.Edt_server_name.Text(),
		Path:   f.Edt_server_path.Text(),
		Params: params,
	})

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
