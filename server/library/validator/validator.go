// Package validator
package validator

import (
	"github.com/gogf/gf/text/gregex"
	"regexp"
	"time"
)

// IsMobile 验证手机号格式
func IsMobile(mobile string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

// IsDateTime 是否为日期+时间 ,格式为2006-01-02 15:04:05
func IsDateTime(value string) bool {
	layout := "2006-01-02 15:04:05"
	p, err := time.Parse(layout, value)
	if err != nil {
		return false
	}
	if p.Format(layout) == value {
		return true
	}
	return false
}

// IsDate 是否为日期，格式为2006-01-02
func IsDate(value string) bool {
	if !gregex.IsMatchString("\\d{4}-\\d{2}-\\d{2}", value) {
		return false
	}
	return true
}

// IsMonth 是否为日期，格式为2006-01
func IsMonth(value string) bool {
	if !gregex.IsMatchString("\\d{4}-\\d{2}", value) {
		return false
	}
	return true
}

// IsAlphaNumber 是否为数字和字母组成
func IsAlphaNumber(value string) bool {
	if !gregex.IsMatchString("[a-zA-Z0-9]", value) {
		return false
	}
	return true
}
