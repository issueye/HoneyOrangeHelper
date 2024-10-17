// Automatically generated by the res2go, do not edit.

package settings

import (
	"github.com/ying32/govcl/vcl"
)

type TFrmSettings struct {
	*vcl.TForm
	PageCtl            *vcl.TPageControl
	Page_log_settings  *vcl.TTabSheet
	Pnl_actions        *vcl.TPanel
	Btn_save           *vcl.TBitBtn
	Btn_cancel         *vcl.TBitBtn
	Panel1             *vcl.TPanel
	Sed_backups        *vcl.TSpinEdit
	Cmb_level          *vcl.TComboBox
	Sed_max_age        *vcl.TSpinEdit
	Sed_show_log_count *vcl.TSpinEdit
	Sed_max_back       *vcl.TSpinEdit
	Edt_log_path       *vcl.TLabeledEdit
	Label1             *vcl.TLabel
	Label2             *vcl.TLabel
	Label3             *vcl.TLabel
	Label4             *vcl.TLabel
	Label5             *vcl.TLabel
	Ckb_compress       *vcl.TCheckBox
	OD_select_server   *vcl.TOpenDialog
	Pmenu_params       *vcl.TPopupMenu
	MenuItem1          *vcl.TMenuItem
	MenuItem2          *vcl.TMenuItem
	ImageList1         *vcl.TImageList

	// ::private::
	TFrmSettingsFields
}

var FrmSettings *TFrmSettings

// Loaded in bytes.
// vcl.Application.CreateForm(&FrmSettings)

func NewFrmSettings(owner vcl.IComponent) (root *TFrmSettings) {
	vcl.CreateResForm(owner, &root)
	return
}

var FrmSettingsBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x0B\x46\x72\x6D\x53\x65\x74\x74\x69\x6E\x67\x73\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\x81\x01\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xF4\x01\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x00\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x44\x69\x61\x6C\x6F\x67\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x81\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xF4\x01\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0F\x70\x6F\x44\x65\x73\x6B\x74\x6F\x70\x43\x65\x6E\x74\x65\x72\x00\x0C\x54\x50\x61\x67\x65\x43\x6F\x6E\x74\x72\x6F\x6C\x07\x50\x61\x67\x65\x43\x74\x6C\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x81\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xF4\x01\x0A\x41\x63\x74\x69\x76\x65\x50\x61\x67\x65\x07\x11\x50\x61\x67\x65\x5F\x6C\x6F\x67\x5F\x73\x65\x74\x74\x69\x6E\x67\x73\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x08\x54\x61\x62\x49\x6E\x64\x65\x78\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x09\x54\x54\x61\x62\x53\x68\x65\x65\x74\x11\x50\x61\x67\x65\x5F\x6C\x6F\x67\x5F\x73\x65\x74\x74\x69\x6E\x67\x73\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x97\xA5\xE5\xBF\x97\xE8\xAE\xBE\xE7\xBD\xAE\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x63\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xEC\x01\x00\x06\x54\x50\x61\x6E\x65\x6C\x0B\x50\x6E\x6C\x5F\x61\x63\x74\x69\x6F\x6E\x73\x04\x4C\x65\x66\x74\x02\x03\x06\x48\x65\x69\x67\x68\x74\x02\x2D\x03\x54\x6F\x70\x03\x33\x01\x05\x57\x69\x64\x74\x68\x03\xE6\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x03\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x03\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x03\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x03\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x29\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE2\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x08\x42\x74\x6E\x5F\x73\x61\x76\x65\x04\x4C\x65\x66\x74\x03\x3D\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x4B\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x0A\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xA1\xAE\xE5\xAE\x9A\x06\x49\x6D\x61\x67\x65\x73\x07\x0A\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x31\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x0A\x42\x74\x6E\x5F\x63\x61\x6E\x63\x65\x6C\x04\x4C\x65\x66\x74\x03\x92\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x4B\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x05\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8F\x96\xE6\xB6\x88\x06\x49\x6D\x61\x67\x65\x73\x07\x0A\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x31\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x03\x06\x48\x65\x69\x67\x68\x74\x03\x2D\x01\x03\x54\x6F\x70\x02\x03\x05\x57\x69\x64\x74\x68\x03\xE6\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x03\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x03\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x03\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x03\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x29\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE2\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x0B\x53\x65\x64\x5F\x62\x61\x63\x6B\x75\x70\x73\x04\x4C\x65\x66\x74\x03\x4E\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x76\x05\x57\x69\x64\x74\x68\x03\x82\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x05\x56\x61\x6C\x75\x65\x02\x64\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x09\x43\x6D\x62\x5F\x6C\x65\x76\x65\x6C\x04\x4C\x65\x66\x74\x02\x67\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x34\x05\x57\x69\x64\x74\x68\x03\x82\x00\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x00\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x05\x44\x45\x42\x55\x47\x06\x04\x49\x4E\x46\x4F\x06\x04\x57\x41\x52\x4E\x06\x05\x45\x52\x52\x4F\x52\x00\x05\x53\x74\x79\x6C\x65\x07\x0E\x63\x73\x44\x72\x6F\x70\x44\x6F\x77\x6E\x4C\x69\x73\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x04\x54\x65\x78\x74\x06\x05\x44\x45\x42\x55\x47\x00\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x0B\x53\x65\x64\x5F\x6D\x61\x78\x5F\x61\x67\x65\x04\x4C\x65\x66\x74\x02\x67\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x76\x05\x57\x69\x64\x74\x68\x03\x82\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x05\x56\x61\x6C\x75\x65\x02\x1E\x00\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x12\x53\x65\x64\x5F\x73\x68\x6F\x77\x5F\x6C\x6F\x67\x5F\x63\x6F\x75\x6E\x74\x04\x4C\x65\x66\x74\x03\x4E\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xA0\x00\x05\x57\x69\x64\x74\x68\x03\x82\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x05\x56\x61\x6C\x75\x65\x03\x88\x13\x00\x00\x09\x54\x53\x70\x69\x6E\x45\x64\x69\x74\x0C\x53\x65\x64\x5F\x6D\x61\x78\x5F\x62\x61\x63\x6B\x04\x4C\x65\x66\x74\x02\x67\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xA0\x00\x05\x57\x69\x64\x74\x68\x03\x82\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x05\x56\x61\x6C\x75\x65\x02\x64\x00\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x0C\x45\x64\x74\x5F\x6C\x6F\x67\x5F\x70\x61\x74\x68\x04\x4C\x65\x66\x74\x02\x67\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x09\x05\x57\x69\x64\x74\x68\x03\x65\x01\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x11\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x02\x54\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE6\x97\xA5\xE5\xBF\x97\xE4\xBF\x9D\xE5\xAD\x98\xE8\xB7\xAF\xE5\xBE\x84\xEF\xBC\x9A\x15\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0D\x4C\x61\x62\x65\x6C\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x06\x6C\x70\x4C\x65\x66\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x1B\xE8\xAF\xB7\xE8\xBE\x93\xE5\x85\xA5\xE6\x97\xA5\xE5\xBF\x97\xE4\xBF\x9D\xE5\xAD\x98\xE8\xB7\xAF\xE5\xBE\x84\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x27\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x35\x05\x57\x69\x64\x74\x68\x02\x3C\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE6\x97\xA5\xE5\xBF\x97\xE7\xAD\x89\xE7\xBA\xA7\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x7A\x05\x57\x69\x64\x74\x68\x02\x54\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE6\x97\xA5\xE5\xBF\x97\xE4\xBF\x9D\xE5\xAD\x98\xE5\xA4\xA9\xE6\x95\xB0\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x03\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x78\x05\x57\x69\x64\x74\x68\x02\x48\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE6\x97\xA5\xE5\xBF\x97\xE5\xA4\x87\xE4\xBB\xBD\xE6\x95\xB0\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x02\x10\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\xA4\x00\x05\x57\x69\x64\x74\x68\x02\x54\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE6\x97\xA5\xE5\xBF\x97\xE5\x88\x86\xE5\x89\xB2\xE5\xA4\xA7\xE5\xB0\x8F\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x03\xF7\x00\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x03\xA4\x00\x05\x57\x69\x64\x74\x68\x02\x54\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE6\x9C\x80\xE5\xA4\xA7\xE6\x97\xA5\xE5\xBF\x97\xE8\xA1\x8C\xE6\x95\xB0\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x09\x54\x43\x68\x65\x63\x6B\x42\x6F\x78\x0C\x43\x6B\x62\x5F\x63\x6F\x6D\x70\x72\x65\x73\x73\x04\x4C\x65\x66\x74\x03\xF4\x00\x06\x48\x65\x69\x67\x68\x74\x02\x15\x03\x54\x6F\x70\x02\x35\x05\x57\x69\x64\x74\x68\x02\x5E\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x12\xE6\x98\xAF\xE5\x90\xA6\xE5\x8E\x8B\xE7\xBC\xA9\xE6\x97\xA5\xE5\xBF\x97\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x06\x00\x00\x00\x00\x00\x0B\x54\x4F\x70\x65\x6E\x44\x69\x61\x6C\x6F\x67\x10\x4F\x44\x5F\x73\x65\x6C\x65\x63\x74\x5F\x73\x65\x72\x76\x65\x72\x04\x4C\x65\x66\x74\x03\x24\x01\x03\x54\x6F\x70\x03\xFF\x00\x00\x00\x0A\x54\x50\x6F\x70\x75\x70\x4D\x65\x6E\x75\x0C\x50\x6D\x65\x6E\x75\x5F\x70\x61\x72\x61\x6D\x73\x04\x4C\x65\x66\x74\x03\x8B\x01\x03\x54\x6F\x70\x03\xFE\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x4D\x65\x6E\x75\x49\x74\x65\x6D\x31\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x96\xB0\xE5\xA2\x9E\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x09\x4D\x65\x6E\x75\x49\x74\x65\x6D\x32\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x88\xA0\xE9\x99\xA4\x00\x00\x00\x0A\x54\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x0A\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x31\x04\x4C\x65\x66\x74\x03\xCD\x00\x03\x54\x6F\x70\x03\x00\x01\x06\x42\x69\x74\x6D\x61\x70\x0A\xCC\x01\x00\x00\x4C\x7A\x02\x00\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\xB6\x01\x00\x00\x00\x00\x00\x00\x78\xDA\xDD\x53\x3D\x4F\x02\x41\x14\xA4\xC4\x8E\x80\x11\x0B\x4D\xAE\x01\xEC\xEE\x0A\x0B\x0B\x13\xE8\xB8\x92\x0E\x3A\xEE\x1F\x1C\xB1\xE1\x1A\xCD\xC6\xCA\xCE\x96\xC2\x82\xC6\x40\x07\x05\x09\x96\x96\x56\x4A\xA7\xA5\x1D\x15\x8A\x95\xA5\xCE\x84\x77\xE6\xB2\xD9\xC3\x25\xC6\x98\x78\xC9\x64\xBF\xDE\xEC\x9B\x9D\xF7\x2E\x93\xF9\xF9\x77\xEC\x6C\x55\xA2\x5A\x21\x04\x3A\x82\x30\x05\x27\x32\xFA\x49\x3E\xD6\x6A\x12\xEC\x8D\x80\x6B\x22\xAA\xE5\xDB\x3A\xB0\x7F\xBA\x54\xA5\x7B\x8C\x6A\xA9\xCA\xB7\x3A\x7F\x15\xB7\xCA\x93\x3C\xEB\x56\xF3\x87\x2F\x67\xE5\x73\x8C\x1E\xF8\xE3\x6E\x6D\xDB\xD9\x8C\x5F\xA8\xBE\xAA\xD2\xD5\x5F\xF1\xE1\xED\x4E\x54\xCD\x1F\xD9\xF2\x27\xC1\xBE\x42\x4C\x4E\x07\x7C\xF3\x2D\xF8\xAD\x37\x55\x7E\x48\x03\xBD\x5F\xC7\xB7\xE9\x95\xFF\xCA\x87\x37\x03\x53\xDF\x1A\xFA\xF8\x42\xE7\xA3\xC6\xEE\x4A\x83\x35\x1A\x99\x5F\xFC\x76\xF1\xB9\xAE\xDB\x04\x42\x41\x03\x5B\x39\x0B\x5E\xAE\x5E\xAF\xF7\xDA\xED\xF6\x3B\x70\x07\x0C\x81\x3E\x30\x03\xE6\x40\x2F\xED\x9E\x62\xB1\xE8\xE2\xFC\x19\x98\x7A\x9E\xE7\x98\x34\xE1\x6C\xC0\x18\x9C\x1F\xE8\x79\xB9\x8F\xDC\x97\xDF\x69\xF4\x7D\x5F\x31\x96\xF7\xC5\x7B\xA2\x73\x6A\xEB\x8F\xBC\xA9\x2F\xB9\xB3\x7C\x2F\xC6\x2F\xCD\xF4\x0B\x79\x3A\x89\x9C\x5C\x47\x09\xBD\x8E\x70\xB2\x8C\xC5\xFC\x49\x7B\x0F\xCF\xE7\xD4\x0A\xB4\x30\xFF\xC0\x18\x68\x1A\x66\xE4\x02\x5D\xCC\x6F\x0C\x7E\x56\xB0\xBF\x30\x71\xE3\x37\x90\xCB\xFA\x62\x3E\x36\xF8\x14\x48\x1D\x17\xD4\x91\xC2\x0F\x45\xFF\xA3\x89\x4B\xED\xA2\x63\xAE\xDF\xC1\xFE\x88\xFB\x8A\xB1\xC9\x9A\x06\x41\x10\xE1\xCC\xD7\xDE\xD2\x8B\xD7\x8C\x8D\xFD\xE3\x5A\x7A\x6E\xB8\x41\xFD\xA6\xE4\x24\xFD\xA6\x16\x1B\x6E\xA2\x7F\x52\xFF\x07\x53\xFF\xE2\x7E\x17\x39\x47\xF4\x82\xEF\x59\xC3\xF5\xC4\xF7\x99\xF4\xE5\x98\xFE\xB2\x8E\xD4\x6C\xFB\x1F\x4A\x6F\x84\x52\xE3\x66\x1A\xEF\x13\xB3\x0D\x56\x7E\x00\x00\x00")

// 注册窗口资源
var _ = vcl.RegisterFormResource(FrmSettings, &FrmSettingsBytes)
