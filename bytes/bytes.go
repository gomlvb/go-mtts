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
	"fmt"
	"regexp"
	"strconv"
)

type Bytes struct {
}

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

var (
	pattern = regexp.MustCompile(`(?i)^(-?\d+)([KMGTP]B?|B)$`)
	global  = New()
)

// 工厂模式构造初始化字节实例
func New() *Bytes {
	return &Bytes{}
}

// 将整数[b字节]格式化为对应的单位字符串
// eg: 31323 bytes -> "30.59KB".
func (*Bytes) Format(b uint64) string {
	multiple := ""
	value := float64(b)

	switch {
	case b < KB:
		return strconv.FormatUint(b, 10) + "B"
	case b < MB:
		value /= KB
		multiple = "KB"
	case b < GB:
		value /= MB
		multiple = "MB"
	case b < TB:
		value /= GB
		multiple = "GB"
	case b < PB:
		value /= TB
		multiple = "TB"
	case b < EB:
		value /= PB
		multiple = "PB"
	}

	return fmt.Sprintf("%.02f%s", value, multiple)
}

// 将单位字符串格式化为对应的整数[b字节]
// eg: 6GB (6G is also valid) -> 6442450944.
func (*Bytes) Parse(value string) (i uint64, err error) {
	parts := pattern.FindStringSubmatch(value)
	if len(parts) < 3 {
		return 0, fmt.Errorf("error parsing value=%s", value)
	}
	bytesString := parts[1]
	multiple := parts[2]
	bytes, err := strconv.ParseUint(bytesString, 10, 64)
	if err != nil {
		return
	}

	switch multiple {
	case "B":
		return bytes * B, nil
	case "K", "KB":
		return bytes * KB, nil
	case "M", "MB":
		return bytes * MB, nil
	case "G", "GB":
		return bytes * GB, nil
	case "T", "TB":
		return bytes * TB, nil
	case "P", "PB":
		return bytes * PB, nil
	}

	return
}

// 全局字节序列化的格式化函数
func Format(b uint64) string {
	return global.Format(b)
}

// 全局字节序列的解析函数
func Parse(val string) (uint64, error) {
	return global.Parse(val)
}
