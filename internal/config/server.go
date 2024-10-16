package config

import (
	"HoneyOrangeHelper/pkg/utils"
	"encoding/json"
	"fmt"
	"strings"
)

type Server struct {
	Code   int64      `gorm:"-" json:"code"`                                                         // 编码
	ID     int64      `gorm:"column:id;primaryKey;autoIncrement:false;type:int" json:"id" form:"id"` // 编码
	Name   string     `gorm:"column:name;size:255;not null;comment:名称" json:"name" form:"name"`      // 名称
	Path   string     `gorm:"column:path;size:255;not null;comment:路径" json:"path" form:"path"`      // 路径
	Params ServerArgs `gorm:"column:params;type:text;comment:参数" json:"params" form:"params"`        // 参数
}

func (srv *Server) ToJson() string {
	data, _ := json.Marshal(srv)
	return string(data)
}

func FromToJson(jsonStr string) (*Server, error) {
	var server Server
	err := json.Unmarshal([]byte(jsonStr), &server)
	return &server, err
}

func (Server) TableName() string {
	return "server_info"
}

func AddServer(server *Server) error {
	server.ID = utils.GenID()
	return configDB.Create(server).Error
}

func UpdateServer(server *Server) error {
	return configDB.Where("id = ?", server.ID).Save(server).Error
}

func DeleteServer(id int64) error {
	return configDB.Where("id = ?", id).Delete(&Server{}).Error
}

func GetServer(id int64) (*Server, error) {
	var server Server
	err := configDB.Where("id = ?", id).First(&server).Error
	return &server, err
}

func GetServerName(name string) (*Server, error) {
	var server Server
	err := configDB.Where("name = ?", name).First(&server).Error
	return &server, err
}

func GetServerList(condition string) ([]*Server, error) {
	serverList := make([]*Server, 0)

	db := configDB.Model(&Server{})

	qCondition := strings.TrimSpace(condition)

	if qCondition != "" {
		db = db.Where(fmt.Sprintf(" name like '%%%s%%' or path like '%%%s%%' ", qCondition, qCondition))
	}

	err := db.Find(&serverList).Error
	return serverList, err
}
