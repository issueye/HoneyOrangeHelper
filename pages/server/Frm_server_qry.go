// Automatically generated by the res2go, do not edit.

package server

import (
	"github.com/ying32/govcl/vcl"
)

type TFrm_server_qry struct {
	*vcl.TForm
	Pnl_query_form *vcl.TPanel
	Btn_query      *vcl.TBitBtn
	Edt_condition  *vcl.TLabeledEdit
	Btn_add        *vcl.TBitBtn
	Table_data     *vcl.TStringGrid
	Imgs           *vcl.TImageList

	// ::private::
	TFrm_server_qryFields
}

var Frm_server_qry *TFrm_server_qry

// Loaded in bytes.
// vcl.Application.CreateForm(&Frm_server_qry)

func NewFrm_server_qry(owner vcl.IComponent) (root *TFrm_server_qry) {
	vcl.CreateResForm(owner, &root)
	return
}

var Frm_server_qryBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x0E\x46\x72\x6D\x5F\x73\x65\x72\x76\x65\x72\x5F\x71\x72\x79\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\x9D\x01\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xEC\x02\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\x8D\xE5\x8A\xA1\xE7\xAE\xA1\xE7\x90\x86\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x9D\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xEC\x02\x00\x06\x54\x50\x61\x6E\x65\x6C\x0E\x50\x6E\x6C\x5F\x71\x75\x65\x72\x79\x5F\x66\x6F\x72\x6D\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x02\x2D\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xEA\x02\x05\x41\x6C\x69\x67\x6E\x07\x05\x61\x6C\x54\x6F\x70\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x01\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x01\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x02\x29\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xE6\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x09\x42\x74\x6E\x5F\x71\x75\x65\x72\x79\x04\x4C\x65\x66\x74\x03\x4A\x02\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x49\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x05\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x05\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x9F\xA5\xE8\xAF\xA2\x06\x49\x6D\x61\x67\x65\x73\x07\x04\x49\x6D\x67\x73\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x0D\x45\x64\x74\x5F\x63\x6F\x6E\x64\x69\x74\x69\x6F\x6E\x04\x4C\x65\x66\x74\x02\x48\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x09\x05\x57\x69\x64\x74\x68\x03\xF2\x00\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x11\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x02\x3C\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x0F\xE6\x9C\x8D\xE5\x8A\xA1\xE5\x90\x8D\xE7\xA7\xB0\xEF\xBC\x9A\x15\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0D\x4C\x61\x62\x65\x6C\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x06\x6C\x70\x4C\x65\x66\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x08\x54\x65\x78\x74\x48\x69\x6E\x74\x06\x15\xE8\xAF\xB7\xE8\xBE\x93\xE5\x85\xA5\xE6\x9C\x8D\xE5\x8A\xA1\xE5\x90\x8D\xE7\xA7\xB0\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x07\x42\x74\x6E\x5F\x61\x64\x64\x04\x4C\x65\x66\x74\x03\x98\x02\x06\x48\x65\x69\x67\x68\x74\x02\x1F\x03\x54\x6F\x70\x02\x05\x05\x57\x69\x64\x74\x68\x02\x49\x05\x41\x6C\x69\x67\x6E\x07\x07\x61\x6C\x52\x69\x67\x68\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x05\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x05\x14\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x42\x6F\x74\x74\x6F\x6D\x02\x05\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\xB7\xBB\xE5\x8A\xA0\x06\x49\x6D\x61\x67\x65\x73\x07\x04\x49\x6D\x67\x73\x0A\x49\x6D\x61\x67\x65\x49\x6E\x64\x65\x78\x02\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x00\x0B\x54\x53\x74\x72\x69\x6E\x67\x47\x72\x69\x64\x0A\x54\x61\x62\x6C\x65\x5F\x64\x61\x74\x61\x04\x4C\x65\x66\x74\x02\x01\x06\x48\x65\x69\x67\x68\x74\x03\x6B\x01\x03\x54\x6F\x70\x02\x32\x05\x57\x69\x64\x74\x68\x03\xEA\x02\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x12\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x4C\x65\x66\x74\x02\x01\x11\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x54\x6F\x70\x02\x05\x13\x42\x6F\x72\x64\x65\x72\x53\x70\x61\x63\x69\x6E\x67\x2E\x52\x69\x67\x68\x74\x02\x01\x08\x43\x6F\x6C\x43\x6F\x75\x6E\x74\x02\x04\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x0D\x54\x69\x74\x6C\x65\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\x8D\xE5\x8A\xA1\xE5\x90\x8D\xE7\xA7\xB0\x05\x57\x69\x64\x74\x68\x03\xC8\x00\x00\x01\x0D\x54\x69\x74\x6C\x65\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE6\x9C\x8D\xE5\x8A\xA1\xE8\xB7\xAF\xE5\xBE\x84\x05\x57\x69\x64\x74\x68\x03\x9A\x01\x00\x01\x0B\x42\x75\x74\x74\x6F\x6E\x53\x74\x79\x6C\x65\x07\x0F\x63\x62\x73\x42\x75\x74\x74\x6F\x6E\x43\x6F\x6C\x75\x6D\x6E\x0F\x54\x69\x74\x6C\x65\x2E\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x0D\x54\x69\x74\x6C\x65\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE4\xBF\xAE\xE6\x94\xB9\x00\x01\x0B\x42\x75\x74\x74\x6F\x6E\x53\x74\x79\x6C\x65\x07\x0F\x63\x62\x73\x42\x75\x74\x74\x6F\x6E\x43\x6F\x6C\x75\x6D\x6E\x0F\x54\x69\x74\x6C\x65\x2E\x41\x6C\x69\x67\x6E\x6D\x65\x6E\x74\x07\x08\x74\x61\x43\x65\x6E\x74\x65\x72\x0D\x54\x69\x74\x6C\x65\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x88\xA0\xE9\x99\xA4\x00\x00\x09\x46\x69\x78\x65\x64\x43\x6F\x6C\x73\x02\x00\x04\x46\x6C\x61\x74\x09\x08\x52\x6F\x77\x43\x6F\x75\x6E\x74\x02\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x0A\x54\x49\x6D\x61\x67\x65\x4C\x69\x73\x74\x04\x49\x6D\x67\x73\x04\x4C\x65\x66\x74\x03\xAC\x01\x03\x54\x6F\x70\x02\x0F\x06\x42\x69\x74\x6D\x61\x70\x0A\x00\x03\x00\x00\x4C\x7A\x02\x00\x00\x00\x10\x00\x00\x00\x10\x00\x00\x00\xEA\x02\x00\x00\x00\x00\x00\x00\x78\xDA\xA5\x54\xDF\x4B\x53\x61\x18\x3E\x98\x13\x17\x3B\xAB\xB3\x9B\xBA\x08\x12\x4D\xC1\x90\x73\x44\x48\x21\x08\xBA\xC8\x5D\x6C\x43\xFA\x03\x3C\x5B\x48\x17\x3B\xE0\xDC\x46\xF4\x03\xA1\x2E\x23\xF2\x6C\x4B\x68\xAB\xBC\x28\x68\x9B\x57\x5D\x78\x11\x48\xBA\xCD\x84\xBA\xCA\xC2\x3B\x45\x17\x45\xC5\xA8\x20\x6F\x22\xE9\x07\x68\xCF\xA3\x67\x78\x3C\x6E\x3A\xEB\x83\x87\x7D\xDF\xFB\x3E\xCF\xFB\xBE\xDF\xFB\x7E\x67\x82\xB0\xB5\xC2\xBE\xC6\x8E\xA5\x94\xF4\x70\x39\xE5\x7A\x07\xFC\x06\x56\x80\x17\x99\xA8\x78\xB9\xBB\xB5\xFE\x90\xB0\xCB\xCA\x44\xC4\x61\x6A\x8A\x29\xA9\x90\x8D\x8A\xD7\x42\x5E\xBB\x07\xF0\x2F\x27\x45\x1D\xF6\x12\xF0\x01\x31\xE4\x8A\xDA\xA8\x38\x08\xFF\x7A\x36\xEA\xBC\x54\xC9\x0F\xDD\xB1\xA5\x51\xF1\x29\x38\x5F\xB0\x6F\x12\x2C\x35\x33\xEF\x78\xD4\x31\x20\xEC\xB1\xC0\x7B\x05\xBC\x34\xDB\x8A\x49\xE9\x31\x6A\x7E\x26\xD4\xB0\x7A\x5A\xEB\x8F\x33\x57\xC8\xDB\x78\xC6\x14\xF3\x7D\x36\x2C\x0E\x0B\x35\x2E\xF0\x17\xD3\x21\xE7\xB0\xE9\xCC\x78\xBE\x9A\xF5\x49\xE9\x49\x31\xE9\x7A\x60\xD2\xAF\xB0\xCF\xB5\xE7\x97\x0A\xE9\x88\xF3\x8A\xE9\x3C\xCB\x19\xED\xA3\xFE\xD2\x90\xD7\x7E\xB6\x7C\x4E\x47\xC4\x21\xDA\x38\xA3\xBD\xB4\xE9\xA8\x78\xA1\x98\x72\x7D\xEC\x6E\xB3\x39\xB6\x7A\x6A\x73\xF0\x6D\x70\xBE\xBB\x69\x39\x77\x68\xBF\x65\x42\xCE\xFE\x0A\x3E\x19\xF7\xF8\xC4\xF9\x72\x46\x95\xF4\x61\x9F\x1D\x1C\xD7\x5A\x36\x22\xF6\x57\x99\xED\x91\xB7\xF7\xA5\x71\xE3\xDD\x2F\xB2\xCF\xF8\x9D\xC6\xDB\x58\x0E\x79\xEC\xB2\x91\xE7\x34\xF2\xAC\xE2\xAD\xAB\xD5\xEA\x1C\xF2\x36\xCA\xE3\x11\xC7\x0D\xE8\xC7\x32\x61\xC7\x75\xF4\xEA\x7C\x4F\x9B\xCD\x66\xCA\xD3\x85\x7B\xAC\x20\xC6\x80\xF0\x8F\x0B\x75\x9C\x44\x6D\x5F\x33\x11\x67\xF0\x7F\x62\xA0\x8E\xD7\x56\x7B\xC0\x1D\x68\xD3\xB5\xD8\xCD\x69\x3D\xBF\x00\xFC\x30\xB0\x40\x1B\x7D\xBB\xC5\xD4\xB5\xF8\x6D\x70\x7F\x01\x9F\xB1\xBF\xEB\x77\xFB\x07\x09\xEE\x69\xA3\x8F\x9C\x2A\xDA\x51\xF8\xD7\xF5\x60\x6C\x44\x69\x56\x76\xFC\xD7\xC8\x2D\x8A\x2B\xA6\xC5\xEF\x6D\x70\xC0\x35\xFB\x46\xB4\xF8\x20\xED\x31\x2D\xB1\xAD\x1F\x72\xB3\x7C\x98\x30\xDB\xC0\xD1\xC8\xA5\x86\x67\xE5\x84\x62\xC7\xB9\x04\x4C\x5A\x73\x4E\xEB\xB9\x09\x62\xA7\x3D\x3F\x49\x0D\xB5\xAA\xDB\xDF\xC7\x78\xAA\x3B\xE0\xB5\xF2\xA6\xF4\xDC\x73\xC2\x6A\x27\x77\x53\xE3\xEF\xCB\xDF\x99\xBD\x85\xFD\x1F\xDC\xB9\x81\xBE\x8E\xA6\x8E\x3A\xB5\x57\x6D\x22\xCA\xFA\xF2\x99\xBE\x8D\x9A\xC1\xA5\x66\x53\x9B\x1B\x63\x2D\xA6\xDA\x26\x18\xBB\x0A\x26\x4C\xBC\x12\xB5\x39\x7D\x26\x86\xFD\x6A\xD9\x8E\x79\xB5\xA3\xBE\x00\x01\xFB\x3C\x51\x3E\xD3\x67\xD2\xAF\x52\x8B\x3B\xA8\xC6\xFD\xBB\xF6\x71\xFF\x2E\xE3\xFE\x2A\xEE\x22\x61\xFF\x3D\x17\x2B\x3C\xAA\xD0\xE7\x29\xC2\x6A\x27\x97\x1A\x6A\x37\x67\x1A\xD7\x71\x5E\x43\x3C\x8F\xE5\xCD\xB4\x13\xDB\x73\xFB\x3D\xE4\x52\x63\x7E\x27\xB0\xBD\x61\x4C\xD8\x2F\xE2\x5C\x67\xCD\xA9\xB4\x2A\x07\xF4\x60\x22\x08\xCE\x4F\x72\xAD\xEF\x4A\x69\x51\x0E\xE6\xF4\xFC\x8C\xD1\xE7\xF9\x91\x60\xE2\x2A\xFA\xE5\x23\xB8\x37\x7A\xB9\x4E\x0E\xB9\x95\xBE\x81\xCE\x96\xCE\x86\xB8\x96\x40\xDF\x0B\x73\x3B\x67\x57\x98\xA3\x8F\x9C\x5A\xBE\x6F\xF5\x9C\x7A\x54\xED\x0D\x9C\xDA\x00\xF6\xD5\x78\x7F\x01\xEE\xE0\xD2\xBB\x00\x00\x00")

// 注册窗口资源
var _ = vcl.RegisterFormResource(Frm_server_qry, &Frm_server_qryBytes)