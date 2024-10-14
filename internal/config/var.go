package config

import (
	"database/sql/driver"
	"encoding/json"
)

type ServerArg struct {
	Key   string `json:"key"`   // key
	Value string `json:"value"` // value
	Mark  string `json:"mark"`  // mark
}

type ServerArgs []ServerArg

// 存入数据库
func (args ServerArgs) Value() (driver.Value, error) {
	return json.Marshal(args)
}

// 从数据库取数据
func (args *ServerArgs) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), args)
}

func (arr *ServerArgs) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]ServerArg)(arr))
}
