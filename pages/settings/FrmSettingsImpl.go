package settings

import (
	"HoneyOrangeHelper/internal/config"
	"HoneyOrangeHelper/internal/global"
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
	// initialize.WriteToIniFile("System", "Port", f.EdtPort.Text())
	// initialize.WriteToIniFile("System", "FormatDateTime", f.EdtCron.Text())
	// initialize.WriteToIniFile("System", "AutoRestart", strconv.FormatBool(f.CheckBox1.Checked()))
	// initialize.WriteToIniFile("System", "AutoRestartCron", f.EdtCron.Text())
	// initialize.WriteToIniFile("Loginfo", "ShowLogCount", f.EdtShowLogCount.Text())
	// initialize.WriteToIniFile("Loginfo", "LogPath", f.EdtLogPath.Text())
	// initialize.WriteToIniFile("Loginfo", "Compress", strconv.FormatBool(f.CkbCompress.Checked()))
	// initialize.WriteToIniFile("Loginfo", "Level", fmt.Sprintf("%d", f.CmbLevel.ItemIndex()))
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
	showLogCount := config.GetParam(config.LOG, "showLogCount", 100).Int()
	logPath := config.GetParam(config.LOG, "path", fmt.Sprintf("%s/logs", global.ROOT_PATH)).String()
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
