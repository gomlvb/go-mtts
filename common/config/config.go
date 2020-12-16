/*****************************************************************************************************
copyright (C),2020-2060,wondershare .Co.,Ltd.

FileName     : config.go
Author       : Shijh      Version : 1.0    Date: 2020年11月27日
Description  : 公共配置接口
Version      : 1.0
Function List:
			// Package config is used to parse config
			// Usage:
			//
			//  cnf, err := config.NewConfig("ini", "config.conf")
			//
			//  cnf APIS:
			//
			//  cnf.Set(key, val string) error
			//  cnf.String(key string) string
			//  cnf.Strings(key string) []string
			//  cnf.Int(key string) (int, error)
			//  cnf.Int64(key string) (int64, error)
			//  cnf.Bool(key string) (bool, error)
			//  cnf.Float(key string) (float64, error)
			//  cnf.DefaultString(key string, defaultVal string) string
			//  cnf.DefaultStrings(key string, defaultVal []string) []string
			//  cnf.DefaultInt(key string, defaultVal int) int
			//  cnf.DefaultInt64(key string, defaultVal int64) int64
			//  cnf.DefaultBool(key string, defaultVal bool) bool
			//  cnf.DefaultFloat(key string, defaultVal float64) float64
			//  cnf.DIY(key string) (interface{}, error)
			//  cnf.GetSection(section string) (map[string]string, error)
			//  cnf.SaveConfigFile(filename string) error
History      :
<author>       <time>             <version>            <desc>
Shijh       2020年11月27日          1.0          build this moudle
******************************************************************************************************/

package config

import "fmt"

// 定义外部调用接口
type Configer interface {
	Set(key, val string) error
	String(key string) string
	Strings(key string) []string
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	DefaultString(key string, defaultVal string) string
	DefaultStrings(key string, defaultVal []string) []string //get string slice
	DefaultInt(key string, defaultVal int) int
	DefaultInt64(key string, defaultVal int64) int64
	DefaultBool(key string, defaultVal bool) bool
	DefaultFloat(key string, defaultVal float64) float64
	DIY(key string) (interface{}, error)
	GetSection(section string) (map[string]string, error)
	SaveConfigFile(filename string) error
}

// 定义格式化接口
type Config interface {
	Parse(key string) (Configer, error)
	ParseData(data []byte) (Configer, error)
}

// 配置缓存
var adapters = make(map[string]Config)

// 通过name 注册对应的不同类型的配置文件 ini/json/xml/yaml
func Register(name string, adapter Config) {
	if adapter == nil {
		panic("config: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("config: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// adapterName is ini/json/xml/yaml
// fileName 文件名路径
func NewConfig(adapterName, fileName string) (Configer, error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		return nil, fmt.Errorf("config: unknown adaptername %q (forgotten import?)", adapterName)
	}
	return adapter.Parse(fileName)
}

// adapterName is ini/json/xml/yaml
// data 配置数据
func NewConfigData(adapterName string, data []byte) (Configer, error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		return nil, fmt.Errorf("config: unknown adaptername %q (forgotten import?)", adapterName)
	}
	return adapter.ParseData(data)
}

// 解析是否成功
func ParseBool(val interface{}) (value bool, err error) {
	if val != nil {
		switch v := val.(type) {
		case bool:
			return v, nil
		case string:
			switch v {
			case "1", "t", "T", "true", "TRUE", "True", "YES", "yes", "Yes", "Y", "y", "ON", "on", "On":
				return true, nil
			case "0", "f", "F", "false", "FALSE", "False", "NO", "no", "No", "N", "n", "OFF", "off", "Off":
				return false, nil
			}
		case int8, int32, int64:
			strV := fmt.Sprintf("%s", v)
			if strV == "1" {
				return true, nil
			} else if strV == "0" {
				return false, nil
			}
		case float64:
			if v == 1 {
				return true, nil
			} else if v == 0 {
				return false, nil
			}
		}
		return false, fmt.Errorf("parsing %q: invalid syntax", val)
	}
	return false, fmt.Errorf("parsing <nil>: invalid syntax")
}
