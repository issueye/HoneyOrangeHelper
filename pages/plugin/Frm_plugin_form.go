// Automatically generated by the res2go, do not edit.

package plugin

import (
	_ "embed"
	"github.com/ying32/govcl/vcl"
)

type TFrm_plugin_form struct {
	*vcl.TForm
	GB_base          *vcl.TGroupBox
	Btn_select       *vcl.TBitBtn
	Edt_name         *vcl.TLabeledEdit
	Edt_process_path *vcl.TLabeledEdit
	PageCtl          *vcl.TPageControl
	Page_event       *vcl.TTabSheet
	Table_event      *vcl.TStringGrid
	Pnl_actions      *vcl.TPanel
	Btn_save         *vcl.TBitBtn
	Btn_cancel       *vcl.TBitBtn
	Od_select_path   *vcl.TOpenDialog
	IList            *vcl.TImageList
	Pm_menu_param    *vcl.TPopupMenu
	Mi_param_add     *vcl.TMenuItem
	Edt_cookie_key   *vcl.TLabeledEdit
	Edt_cookie_value *vcl.TLabeledEdit

	// ::private::
	TFrm_plugin_formFields
}

var Frm_plugin_form *TFrm_plugin_form

// Loaded in bytes.
// vcl.Application.CreateForm(&Frm_plugin_form)

func NewFrm_plugin_form(owner vcl.IComponent) (root *TFrm_plugin_form) {
	vcl.CreateResForm(owner, &root)
	return
}

//go:embed Frm_plugin_form.gfm
var Frm_plugin_formBytes []byte

// 注册窗口资源
var _ = vcl.RegisterFormResource(Frm_plugin_form, &Frm_plugin_formBytes)
