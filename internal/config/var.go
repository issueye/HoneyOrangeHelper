package config

import (
	"database/sql/driver"
	"encoding/json"
)

type ServerArg struct {
	Key   string `json:"key"`   // key
	Value string `json:"value"` // value
	Type  string `json:"type"`  // type
	Mark  string `json:"mark"`  // mark
}

type ServerArgs []*ServerArg

// 存入数据库
func (args ServerArgs) Value() (driver.Value, error) {
	return json.Marshal(args)
}

// 从数据库取数据
func (args *ServerArgs) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), args)
}

func (arr *ServerArgs) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]*ServerArg)(arr))
}

type Plugin struct {
	Name    string `json:"name"`    // name
	Process string `json:"process"` // process
	Mark    string `json:"mark"`    // mark
}

type Plugins []*Plugin

// 存入数据库
func (args Plugins) Value() (driver.Value, error) {
	return json.Marshal(args)
}

// 从数据库取数据
func (args *Plugins) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), args)
}

func (arr *Plugins) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]*Plugin)(arr))
}

type Event struct {
	IsCheck bool   `json:"is_check"` // is_check
	Name    string `json:"name"`     // name
	Title   string `json:"title"`    // title
	Mark    string `json:"mark"`     // mark
}

type Events []*Event

// 存入数据库
func (args Events) Value() (driver.Value, error) {
	return json.Marshal(args)
}

// 从数据库取数据
func (args *Events) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), args)
}

func (arr *Events) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*[]*Event)(arr))
}
