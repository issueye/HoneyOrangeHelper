package config

import (
	"HoneyOrangeHelper/pkg/utils"
	"encoding/json"
	"fmt"
	"strings"
)

type ToolPlugin struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:false;type:int" json:"id" form:"id"`                              // 编码
	Name        string `gorm:"column:name;size:255;not null;comment:名称" json:"name" form:"name"`                                   // 名称
	Path        string `gorm:"column:path;size:255;not null;comment:路径" json:"path" form:"path"`                                   // 路径
	CookieKey   string `gorm:"column:cookie_key;size:255;not null;comment:cookie_key" json:"cookie_key" form:"cookie_key"`         // cookie_key
	CookieValue string `gorm:"column:cookie_value;size:255;not null;comment:cookie_value" json:"cookie_value" form:"cookie_value"` // cookie_value
	Events      Events `gorm:"column:events;type:text;comment:事件" json:"events" form:"events"`                                     // 事件
}

func (srv *ToolPlugin) ToJson() string {
	data, _ := json.Marshal(srv)
	return string(data)
}

func (ToolPlugin) FromToJson(jsonStr string) (*ToolPlugin, error) {
	var ToolPlugin ToolPlugin
	err := json.Unmarshal([]byte(jsonStr), &ToolPlugin)
	return &ToolPlugin, err
}

func (ToolPlugin) TableName() string {
	return "tool_plugin"
}

func AddToolPlugin(ToolPlugin *ToolPlugin) error {
	ToolPlugin.ID = utils.GenID()
	return configDB.Create(ToolPlugin).Error
}

func UpdateToolPlugin(ToolPlugin *ToolPlugin) error {
	return configDB.Where("id = ?", ToolPlugin.ID).Save(ToolPlugin).Error
}

func DeleteToolPlugin(id int64) error {
	return configDB.Where("id = ?", id).Delete(&ToolPlugin{}).Error
}

func GetToolPlugin(id int64) (*ToolPlugin, error) {
	var ToolPlugin ToolPlugin
	err := configDB.Where("id = ?", id).First(&ToolPlugin).Error
	return &ToolPlugin, err
}

func GetToolPluginName(name string) (*ToolPlugin, error) {
	var ToolPlugin ToolPlugin
	err := configDB.Where("name = ?", name).First(&ToolPlugin).Error
	return &ToolPlugin, err
}

func GetToolPluginList(condition string) ([]*ToolPlugin, error) {
	ToolPluginList := make([]*ToolPlugin, 0)

	db := configDB.Model(&ToolPlugin{})

	qCondition := strings.TrimSpace(condition)

	if qCondition != "" {
		db = db.Where(fmt.Sprintf(" name like '%%%s%%' or path like '%%%s%%' ", qCondition, qCondition))
	}

	err := db.Find(&ToolPluginList).Error
	return ToolPluginList, err
}
