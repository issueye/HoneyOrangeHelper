package settings

import (
	"HoneyOrangeHelper/internal/config"
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func NewForm(owner vcl.IComponent) *TFrmSettings {
	if FrmSettings == nil {
		FrmSettings = NewFrmSettings(owner)
	}

	FrmSettings.SetPosition(types.PoOwnerFormCenter)
	return FrmSettings
}

// ::private::
type TFrmSettingsFields struct {
}

func (f *TFrmSettings) OnFormCreate(sender vcl.IObject) {
	f.Sed_show_log_count.SetMinValue(20)
	f.SetCaption("系统设置")

	f.InitParams()

	// 事件
	f.Btn_cancel.SetOnClick(f.OnCancelClick)
	f.Btn_save.SetOnClick(f.OnSaveClick)
	f.Sed_show_log_count.SetOnChange(f.OnEdtShowLogCountChange)
}

func (f *TFrmSettings) OnCancelClick(sender vcl.IObject) {
	// 事件
	f.Close()
}

func (f *TFrmSettings) OnSaveClick(sender vcl.IObject) {
	// 事件

	config.SetParam(config.LOG, "level", "日志等级", f.Cmb_level.Text())
	config.SetParam(config.LOG, "path", "日志存放路径", f.Edt_log_path.Text())
	config.SetParam(config.LOG, "max-size", "日志文件最大大小", f.Sed_max_back.Text())
	config.SetParam(config.LOG, "max-backups", "最大备份日志文件数", f.Sed_backups.Text())
	config.SetParam(config.LOG, "max-age", "日志文件最大保存时间", f.Sed_max_age.Text())
	config.SetParam(config.LOG, "compress", "是否压缩日志文件", f.Ckb_compress.Checked())
	config.SetParam(config.LOG, "show-log-count", "显示日志条数", f.Sed_show_log_count.Text())
	vcl.ShowMessage("配置保存成功!")
	f.SetModalResult(1)
}

func (f *TFrmSettings) OnEdtShowLogCountChange(sender vcl.IObject) {
	minValue := f.Sed_show_log_count.MinValue()
	// 事件
	if f.Sed_show_log_count.Value() < minValue {
		f.Sed_show_log_count.SetValue(minValue)
	}
}

func (f *TFrmSettings) InitParams() {
	// 日志
	showLogCount := config.GetParam(config.LOG, "show-log-count", 100).Int()
	logPath := config.GetParam(config.LOG, "path", "logs").String()
	compress := config.GetParam(config.LOG, "compress", true).Bool()
	level := config.GetParam(config.LOG, "level", 0).Int()
	MaxSize := config.GetParam(config.LOG, "max-size", 100).Int()
	maxBackups := config.GetParam(config.LOG, "max-backups", 100).Int()
	maxAge := config.GetParam(config.LOG, "max-age", 100).Int()

	f.Sed_show_log_count.SetValue(int32(showLogCount))
	f.Edt_log_path.SetText(logPath)
	f.Ckb_compress.SetChecked(compress)
	f.Cmb_level.SetItemIndex(int32(level + 1))
	f.Sed_backups.SetValue(int32(maxBackups))
	f.Sed_max_back.SetValue(int32(MaxSize))
	f.Sed_max_age.SetValue(int32(maxAge))
}

func (f *TFrmSettings) OnTable_paramsButtonClick(sender vcl.IObject, aCol int32, aRow int32) {
	vcl.ShowMessage(fmt.Sprintf("点击了第%d列，第%d行", aCol, aRow))
}
