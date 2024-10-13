package config

type Server struct {
	ID     int64  `gorm:"column:id;primaryKey;autoIncrement:false;type:int" json:"id" form:"id"` // 编码
	Name   string `gorm:"column:name;size:255;not null;comment:名称" json:"name" form:"name"`      // 名称
	Path   string `gorm:"column:path;size:255;not null;comment:路径" json:"path" form:"path"`      // 路径
	Params Arr    `gorm:"column:params;type:text;comment:参数" json:"params" form:"params"`        // 参数
}

func (Server) TableName() string {
	return "server_info"
}

func AddServer(server *Server) error {
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
