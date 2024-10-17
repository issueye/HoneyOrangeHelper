// Automatically generated by the res2go, do not edit.

package server

import (
	"github.com/ying32/govcl/vcl"
)

type TFrm_server struct {
	*vcl.TForm
	Pnl_header          *vcl.TPanel
	Pnl_name            *vcl.TPanel
	Lbl_id              *vcl.TLabel
	Lbl_name            *vcl.TLabel
	Lbl_path            *vcl.TLabel
	Pnl_actions         *vcl.TPanel
	Btn_server_run      *vcl.TBitBtn
	Btn_server_settings *vcl.TBitBtn
	Btn_remove_server   *vcl.TBitBtn
	Btn_clear_log       *vcl.TBitBtn
	Mmo_run_log         *vcl.TMemo
	IList               *vcl.TImageList

	// ::private::
	TFrm_serverFields
}

var Frm_server *TFrm_server

// Loaded in bytes.
// vcl.Application.CreateForm(&Frm_server)

func NewFrm_server(owner vcl.IComponent) (root *TFrm_server) {
	vcl.CreateResForm(owner, &root)
	return
}

var Frm_serverBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x0A\x46\x72\x6D\x5F\x73\x65\x72\x76\x65\x72\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\xD7\x01\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xB2\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x9C\x8D\xE5\x8A\xA1\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xD7\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xB2\x02\x00\x06\x54\x50\x61\x6E\x65\x6C\x0A\x50\x6E\x6C\x5F\x68\x65\x61\x64\x65\x72\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x02\x7B\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xB2\x02\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x7B\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xB2\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x08\x50\x6E\x6C\x5F\x6E\x61\x6D\x65\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x45\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xB0\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x01\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x01\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x41\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xAC\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x62\x6C\x5F\x69\x64\x04\x4C\x65\x66\x74\x02\x18\x06\x48\x65\x69\x67\x68\x74\x02\x10\x03\x54\x6F\x70\x02\x07\x05\x57\x69\x64\x74\x68\x02\x16\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x05\x49\x44\xEF\xBC\x9A\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF5\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x0C\xE5\xBE\xAE\xE8\xBD\xAF\xE9\x9B\x85\xE9\xBB\x91\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x57\x6F\x72\x64\x57\x72\x61\x70\x09\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x08\x4C\x62\x6C\x5F\x6E\x61\x6D\x65\x04\x4C\x65\x66\x74\x02\x0D\x06\x48\x65\x69\x67\x68\x74\x02\x10\x03\x54\x6F\x70\x02\x1A\x05\x57\x69\x64\x74\x68\x02\x21\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE5\x90\x8D\xE7\xA7\xB0\xEF\xBC\x9A\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF5\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x0C\xE5\xBE\xAE\xE8\xBD\xAF\xE9\x9B\x85\xE9\xBB\x91\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x57\x6F\x72\x64\x57\x72\x61\x70\x09\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x08\x4C\x62\x6C\x5F\x70\x61\x74\x68\x04\x4C\x65\x66\x74\x02\x0D\x06\x48\x65\x69\x67\x68\x74\x02\x10\x03\x54\x6F\x70\x02\x2D\x05\x57\x69\x64\x74\x68\x02\x21\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x09\xE8\xB7\xAF\xE5\xBE\x84\xEF\xBC\x9A\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF5\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x0C\xE5\xBE\xAE\xE8\xBD\xAF\xE9\x9B\x85\xE9\xBB\x91\x06\x4C\x61\x79\x6F\x75\x74\x07\x08\x74\x6C\x43\x65\x6E\x74\x65\x72\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x57\x6F\x72\x64\x57\x72\x61\x70\x09\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x0B\x50\x6E\x6C\x5F\x61\x63\x74\x69\x6F\x6E\x73\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x31\x03\x54\x6F\x70\x02\x4A\x05\x57\x69\x64\x74\x68\x03\xB0\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x42\x6F\x74\x74\x6F\x6D\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x01\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x01\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x2D\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xAC\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x0E\x42\x74\x6E\x5F\x73\x65\x72\x76\x65\x72\x5F\x72\x75\x6E\x04\x4C\x65\x66\x74\x03\x71\x01\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x55\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x05\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x05\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xBF\x90\xE8\xA1\x8C\xE6\x9C\x8D\xE5\x8A\xA1\x06\x49\x6D\x61\x67\x65\x73\x07\x05\x49\x4C\x69\x73\x74\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x13\x42\x74\x6E\x5F\x73\x65\x72\x76\x65\x72\x5F\x73\x65\x74\x74\x69\x6E\x67\x73\x04\x4C\x65\x66\x74\x03\xCB\x01\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x46\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x05\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x05\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE8\xAE\xBE\xE7\xBD\xAE\x06\x49\x6D\x61\x67\x65\x73\x07\x05\x49\x4C\x69\x73\x74\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x03\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x11\x42\x74\x6E\x5F\x72\x65\x6D\x6F\x76\x65\x5F\x73\x65\x72\x76\x65\x72\x04\x4C\x65\x66\x74\x03\x61\x02\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x46\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x05\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x05\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE7\xA7\xBB\xE9\x99\xA4\x06\x49\x6D\x61\x67\x65\x73\x07\x05\x49\x4C\x69\x73\x74\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x0D\x42\x74\x6E\x5F\x63\x6C\x65\x61\x72\x5F\x6C\x6F\x67\x04\x4C\x65\x66\x74\x03\x16\x02\x06\x48\x65\x69\x67\x68\x74\x02\x23\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x46\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x05\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x05\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\xB8\x85\xE7\xA9\xBA\x06\x49\x6D\x61\x67\x65\x73\x07\x05\x49\x4C\x69\x73\x74\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x04\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x00\x00\x05\x54\x4D\x65\x6D\x6F\x0B\x4D\x6D\x6F\x5F\x72\x75\x6E\x5F\x6C\x6F\x67\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\x59\x01\x03\x54\x6F\x70\x02\x7E\x05\x57\x69\x64\x74\x68\x03\xB0\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x01\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x03\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x01\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x06\x62\x73\x4E\x6F\x6E\x65\x05\x43\x6F\x6C\x6F\x72\x07\x0A\x63\x6C\x4D\x65\x6E\x75\x54\x65\x78\x74\x0C\x46\x6F\x6E\x74\x2E\x43\x68\x61\x72\x53\x65\x74\x07\x0E\x47\x42\x32\x33\x31\x32\x5F\x43\x48\x41\x52\x53\x45\x54\x0A\x46\x6F\x6E\x74\x2E\x43\x6F\x6C\x6F\x72\x07\x07\x63\x6C\x43\x72\x65\x61\x6D\x0B\x46\x6F\x6E\x74\x2E\x48\x65\x69\x67\x68\x74\x02\xF3\x09\x46\x6F\x6E\x74\x2E\x4E\x61\x6D\x65\x06\x0C\xE5\xBE\xAE\xE8\xBD\xAF\xE9\x9B\x85\xE9\xBB\x91\x0A\x46\x6F\x6E\x74\x2E\x50\x69\x74\x63\x68\x07\x0A\x66\x70\x56\x61\x72\x69\x61\x62\x6C\x65\x0C\x46\x6F\x6E\x74\x2E\x51\x75\x61\x6C\x69\x74\x79\x07\x07\x66\x71\x44\x72\x61\x66\x74\x0A\x50\x61\x72\x65\x6E\x74\x46\x6F\x6E\x74\x08\x08\x52\x65\x61\x64\x4F\x6E\x6C\x79\x09\x0A\x53\x63\x72\x6F\x6C\x6C\x42\x61\x72\x73\x07\x0A\x73\x73\x41\x75\x74\x6F\x42\x6F\x74\x68\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x08\x57\x6F\x72\x64\x57\x72\x61\x70\x08\x00\x00\x0A\x54\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x05\x49\x4C\x69\x73\x74\x04\x4C\x65\x66\x74\x03\xDF\x01\x03\x54\x6F\x70\x03\x8B\x00\x06\x42\x69\x74\x6D\x61\x70\x0A\x21\x06\x00\x00\x4C\x7A\x05\x00\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\x0B\x06\x00\x00\x00\x00\x00\x00\x78\xDA\xD5\x57\xED\x4F\x9B\x55\x14\xEF\xCB\xD3\x37\xED\x3B\xAC\x2D\x2F\xED\xC8\x32\x10\x30\x2E\x15\x68\xD9\x92\x25\xA2\x33\x13\x01\xA7\x8B\x83\xB0\xE8\xB2\xC6\x97\x80\x6B\x0B\x05\xB7\x31\xDD\x12\x21\xC6\xA5\xC6\x0F\x63\x89\xBC\x8C\x68\x58\x4B\xAB\xC9\x40\x5E\x62\x36\xD8\x37\x62\x5C\x82\x9B\x71\xB8\x98\xD0\x0E\xB6\xCE\x2F\xE3\x9F\x30\xC1\xDF\x29\xCF\xD3\x3E\xAD\xCF\xD3\x52\x75\x31\xDE\xE4\xE6\xB9\xF7\x9E\xF3\xBB\xF7\x9C\xDF\x39\x07\x4E\x25\x92\xCC\x61\xAB\xD5\x1B\x9D\xED\xA5\xBD\x9E\x69\xF7\xBC\x67\xDA\xB5\xEA\x99\x71\xDF\xA5\xB5\xB3\xA3\xB4\x8B\x64\x92\x1C\xA3\xF9\x93\x9A\x6E\x60\x36\x9B\x07\xAB\x87\x9D\xED\x65\x4D\xCE\x63\xE5\xB6\xE4\xC4\x1A\xB2\x60\x52\x36\x54\xD3\x2D\x88\x1D\xAC\x9E\xC4\x3B\x2B\xB6\x5A\x9D\x4D\xEC\x7E\x92\x41\x67\xB1\x79\xA8\x3A\x98\x89\xAD\x09\x10\x56\xB2\x83\x61\xAB\xD1\xAB\x3D\xD7\xDC\x73\x9C\x1D\xE4\x13\xEC\x4A\x64\xFB\x46\x76\x8B\xDB\x91\xC6\x38\x3B\xCA\xBA\xF0\xFE\x70\xB6\x0E\xE4\x6B\xB0\xF3\x9C\x38\x57\xD5\x03\xC4\x29\xF4\x96\x88\x1F\x01\x7C\x02\x3E\x6D\x81\x97\x41\x21\x3C\x30\xCE\x64\x8C\x10\x1F\x21\x5B\x39\xBC\xD8\x1D\xCE\x37\xCB\xE1\x83\x7B\x8D\x66\x3E\x3C\xDE\x98\x13\xE4\x31\x95\x1B\x65\x8D\xA2\xF6\x0F\x55\x7F\x2B\x62\x7F\x13\xF9\x4E\x3C\x20\x16\x41\x21\x3C\xE5\x84\x28\x7F\xC0\x10\xF7\xA2\xF1\x13\xE0\x54\x28\x7E\xBC\xFC\x59\x94\xEC\x70\x20\x7F\x16\x61\x5B\x77\xD6\xD9\x30\xDD\x81\x1C\xAD\xC8\xF9\xEE\xB5\x64\xFE\x8E\xE5\xAC\x1F\xF2\x2D\xA3\x7E\x4A\xF7\xB3\xF5\x93\x00\x5F\x9E\x9C\xF9\x8D\x1A\x21\x4E\x89\x5B\x36\x3E\x2B\x6C\x8C\x3A\xF3\xD5\x2F\x0D\x99\xDA\x65\x94\xE9\x7D\xBD\x4A\x47\x6C\x5E\x61\x8F\xAD\x2A\x1D\xF1\xBB\xB4\xC6\x59\x17\xC9\x72\x61\x15\xD6\x70\x37\x30\x9B\x8C\x25\x3C\x2C\x37\x78\x9B\x18\xA3\xD7\x46\x13\xD8\x26\xC8\x82\x24\x23\x1D\x21\x2C\x30\x93\x78\x67\x45\xAA\x72\x89\xD6\x1D\xC9\xA0\xB3\x08\xDD\x60\xD6\xBB\x01\xC2\xEE\x24\x76\xF0\x41\x0D\xDD\x39\xCE\x0E\xF2\x09\x76\x25\xF2\xF9\x96\xCD\x11\x87\x91\x1B\x7C\x5D\xB8\x2B\xA3\xFE\x15\x96\xD0\x51\xB9\x7D\xED\x73\xFE\x94\x3A\x66\x0E\x65\xF9\x3B\x40\x9C\xE2\x9E\x25\xE2\x27\x03\x6F\x8F\x4D\xC0\xC6\x2D\xFE\xA4\x33\xBE\x0E\xF8\x75\x52\x5C\x28\x3E\xC4\x71\xA1\x78\x60\xE0\x43\x7C\x8D\xE6\xDF\xC1\x6F\xF3\xB8\x9D\x1B\xE0\xA0\xB1\x60\x3C\xE5\x04\x7C\x27\x1E\x28\x37\x0A\xE5\x8F\x30\xC4\xFD\x3F\x8D\x1F\x2F\x7F\x76\x5C\xFF\x6C\x0E\x76\x67\x9D\x0D\xD3\x39\x72\xB4\x22\xD7\xBB\x2C\x76\x2C\x57\xFD\x90\x6F\xC4\x4F\xBA\x7E\xBC\xFB\xD9\xFA\x49\x00\x9B\xB3\xFE\xA9\x46\x92\x9C\x82\x5B\x36\x3E\x2B\x6C\x8C\x3A\x0B\xE1\x88\x46\x69\x69\x69\x05\x4D\x49\x81\x43\xAB\xD5\xD6\xD6\xD7\xD7\x3F\xC2\xDC\x62\xE7\x23\x3A\xDB\x29\xBE\xAE\xAE\xEE\x0E\xE6\xED\x92\x92\x92\xBD\x34\x69\x8D\x3B\xEE\x89\xE9\x57\x56\x56\xFA\x20\x4F\xF0\xE6\x1F\x59\xFB\x04\x6B\x47\x6A\x4F\x18\x9E\x9F\xF5\x55\x55\x55\xCB\x3C\x7B\x73\x4E\xD2\x25\x0C\xDF\x06\xD8\x39\xB8\x53\x3C\xE9\x66\xFB\xF0\x2F\xE3\x37\xE1\x9F\x8B\xDB\xB3\xEB\xCD\x02\xF0\x09\x8A\x3B\xB7\x67\xD7\x89\xFF\x11\xFE\x31\x62\xE4\xE2\xC5\x8B\xD6\x8F\x9F\x30\xFF\x27\x0A\xC0\x9F\x10\xA8\x99\x22\xE4\xF9\xA3\x7C\x58\xD2\x21\x5D\x91\xBA\xD3\xE1\xEE\x93\x64\x9F\xC8\x3C\x49\x3A\x92\x27\x38\xF0\xD7\x4F\xEA\x6F\x55\x37\xF7\xBC\xAA\x7E\x86\x3B\xC3\x7A\x5F\x4F\xAB\xFA\xA5\x7C\x58\x77\x25\xE3\xD8\x18\x35\xAF\xAE\x8F\x99\x97\x23\x7D\xBA\xCB\x1B\x63\xE6\x89\x8D\xB1\xA2\x2B\xEB\xA3\xE6\x1F\x23\x01\xDD\x38\xBE\x2B\xD0\xB1\x8B\xE1\xD7\x47\x4C\x5F\x43\x2F\xD5\x73\xF7\xB4\x68\x76\xF7\xB4\x6A\x6A\x5D\x95\x0A\x39\xED\x23\xFD\xC6\x06\xD2\x11\xC3\x3F\x1C\x37\x5D\xF6\xB7\x6A\x0E\x8A\xC9\x71\x57\xD1\x83\x71\xD3\x25\x11\xDB\x2B\x60\xDF\xCF\xF8\xEA\xD9\xBD\x05\x77\xBD\x8F\x79\xCA\x5D\xC5\x3C\x95\xB2\x71\xD4\x14\x69\x84\x8C\x8F\x0D\xF5\x19\x8F\x00\xFB\x4B\xB4\xCF\xF8\x42\x12\xBB\x97\x31\xD2\x5D\x91\x80\xDE\x0F\x2E\x3B\x5C\x95\x8C\x9A\xD3\x05\xD6\x0E\xD9\x0D\xDC\xFB\x36\x77\xB6\x31\x6A\x0A\xC3\x36\x27\xB7\x07\xE6\xDD\x68\x40\xDF\x9B\xDE\x6B\x9A\x70\xD6\xCB\x7F\x73\x63\xC4\x3C\x9C\x5E\x9B\x66\x80\xAF\xE1\xE9\xE3\x5D\x4D\x7B\x6A\xDF\xA2\x79\x05\xFB\x8C\x9C\x47\x9C\xC2\xDC\x3A\xE2\xD7\x7D\xB8\x3E\x62\xF9\x8A\x38\x48\x73\x61\xFA\x0D\x77\xB6\xF8\x5B\xD4\x07\x60\xDF\x1D\xE0\x1B\x58\x59\x71\x34\xA0\xFB\xEC\xC1\x97\xC5\x17\xF9\xF7\x45\xFB\x74\xC7\x89\x03\x8E\x2B\x60\x9F\x7B\x38\x6E\x1C\xC1\xD9\x9C\xBF\x4D\x73\x20\xC5\xDF\x98\x79\x0E\xF8\x77\x84\x62\x40\x39\x82\x77\x8E\x88\xC5\xCF\x77\x98\x31\xC0\xD7\x79\xD1\xFC\x19\x35\x4D\x00\x7F\x58\x4C\xDE\xDB\xA6\x29\x46\xFE\x44\xC5\xE4\xA1\x80\xA6\x11\x77\xDC\x8F\xF4\x6A\xCF\xC0\x9F\x00\xFC\x9E\xC5\x7E\x09\xEB\xE0\xEC\x79\xCD\x71\x70\xF6\x6B\xB4\x5F\xDB\x96\xA7\x06\x88\x9F\xC1\xB0\x5F\xDF\x91\xE6\x57\xFF\x56\x24\xA0\xBD\x00\x99\x49\xF2\x1F\x0C\xF4\x13\xF4\x7B\xE0\x1B\x59\x71\x28\x80\xFE\xE2\xE9\x82\xC0\x4A\xB7\x02\xD8\xEF\x25\x0A\x97\x9C\x7A\x26\xF4\x21\xDE\x42\xE0\x72\xBD\xEF\xA0\xD2\x3A\xF5\x51\x72\x5D\x7C\xB5\x5F\x66\xF0\x1D\xCA\x8D\x68\x42\x6F\xEC\x7B\x03\x7A\x47\xA5\xBA\x0F\x8A\x19\xCB\xD4\x10\xF6\xF5\x6C\x8F\xF6\x9D\x4C\xD9\xA8\xCA\xE3\x6B\x10\xF3\x02\xF0\x9D\x4C\x79\x1C\x7D\x6C\xFC\x96\x44\xE3\x96\xCA\xD5\x8D\xF0\x23\x3E\x9B\x1B\x3B\x15\x00\xF6\x63\xC1\xBE\xCC\xEC\xDD\xA5\xA4\x9E\xCC\x1E\x5B\x66\xAC\xA1\x01\x89\x6A\xBF\x2C\xA3\x37\x34\xF8\xF6\x28\x1C\xF1\xDB\x52\xD5\x59\x43\x9E\xFE\x4E\x8B\x77\x2E\xC8\x77\x85\x4F\x65\x08\x60\x23\x53\x16\x8B\x2A\x1C\xB1\x65\xF8\xB9\xC4\x58\xC3\x17\x89\x2B\xF4\x7C\x7F\xF1\x37\xE9\x8B\x3D\x36\x23\xD0\xD7\xDE\xA4\xBE\x5E\xA2\xAC\x97\x81\x33\xB7\xC2\x1A\x3A\xCF\xF6\x91\xD7\xE5\x96\xD0\x39\xC4\xA2\x81\x64\xAC\xEE\x02\x71\x9D\xF2\x7D\xD7\xD4\xEB\x8C\xE5\xDE\x25\x31\x9B\x61\x4B\x9B\x62\xF7\xFD\xCB\xC0\xDD\x82\x9F\xB3\xF8\xAE\xCA\x8C\x7E\x2B\xCB\xF9\x39\x95\x23\x36\x29\x64\xAB\x60\x4E\x58\x42\xA7\xC1\x81\x7F\x3B\x3F\xBC\x65\x2A\xC7\x5A\x88\x17\xBF\xF7\xF0\x5F\xE4\x53\xA9\xCA\x6D\x95\x1A\x7A\x5E\x14\x8C\xD3\x9E\xD8\x17\xB0\xA7\x9A\xED\xC9\x95\xE0\xE2\x26\xFC\x6D\x07\x67\x1E\xD8\xB5\x2E\xB7\x85\x9F\x05\x8F\x3F\xC8\x0D\x33\xAF\xE1\x9E\x72\xF8\x9F\xFC\x9F\x21\x35\xF9\xF6\x50\xDC\x94\xF6\xF8\xB4\x4C\x75\x56\xC5\xF7\x8F\xF0\xE8\xD9\x8F\x01\x7F\x03\x9C\x5D\xC7\xF7\x27\x59\xF9\xD5\xBD\xF4\x95\x97\x84\x3D\x52\xD5\x99\x0A\xAC\x63\xD0\xD5\xE0\x37\xD7\x42\xAE\xFC\x85\xAE\x0E\x36\x04\x60\xC3\x0C\x74\xA7\x15\xD6\xDF\x07\x61\xE3\x02\xEC\x3B\x2D\x55\xBB\xB5\x42\xB1\x13\x1D\x2A\x17\x03\xDB\x9E\x47\xED\xBC\x0C\x3F\x42\xF4\x5B\x40\x6A\x9E\xDE\xC7\x89\xFF\x04\x31\x25\x3B\x82\x00\x00\x00")

// 注册窗口资源
var _ = vcl.RegisterFormResource(Frm_server, &Frm_serverBytes)
