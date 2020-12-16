/*****************************************************************************************************
copyright (C),2020-2060,wondershare .Co.,Ltd.

FileName     : bytes.go
Author       : Shijh      Version : 1.0    Date: 2020年11月27日
Description  : 字节处理类
Version      : 1.0
Function List:

History      :
<author>       <time>             <version>            <desc>
Shijh       2020年11月27日          1.0          build this moudle
******************************************************************************************************/

package bytes

import (
	"testing"
)

func TestBytes_Format(t *testing.T) {
	type args struct {
		b uint64
	}
	tests := []struct {
		name string
		b    *Bytes
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Format(tt.args.b); got != tt.want {
				t.Errorf("Bytes.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes_Parse(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		b       *Bytes
		args    args
		wantI   uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, err := tt.b.Parse(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bytes.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotI != tt.wantI {
				t.Errorf("Bytes.Parse() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	type args struct {
		b uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.b); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
