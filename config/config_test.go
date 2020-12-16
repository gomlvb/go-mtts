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

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		adapterName string
		fileName    string
	}
	tests := []struct {
		name    string
		args    args
		want    Configer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.adapterName, tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
