// Automatically generated by the res2go, do not edit.

package server

import (
	_ "embed"

	"github.com/ying32/govcl/vcl"
)

type TFrm_server_item struct {
	*vcl.TForm
	Btn_item *vcl.TBitBtn
	IList    *vcl.TImageList

	// ::private::
	TFrm_server_itemFields
}

var Frm_server_item *TFrm_server_item

// Loaded in bytes.
// vcl.Application.CreateForm(&Frm_server_item)

func NewFrm_server_item(owner vcl.IComponent) (root *TFrm_server_item) {
	vcl.CreateResForm(owner, &root)
	return
}

//go:embed Frm_server_item.gfm
var Frm_server_itemBytes []byte

// 注册窗口资源
var _ = vcl.RegisterFormResource(Frm_server_item, &Frm_server_itemBytes)
