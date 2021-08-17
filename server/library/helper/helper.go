// Package helper
package helper

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"math"
	netURL "net/url"
	"regexp"
	"sort"
	"strings"
)

func NowTime(args ...int) uint {
	var (
		nowTime uint
		timeAdd int
	)
	if len(args) > 0 {
		timeAdd = args[0]
		nowTime = uint(gtime.Now().Unix()) + uint(timeAdd)
	} else {
		nowTime = uint(gtime.Now().Unix())
	}
	return nowTime
}
func StringInArray(str string, array []string) bool {
	var res bool
	for _, v := range array {
		if v == str {
			res = true
			break
		}
	}
	return res
}

// Round 四舍五入，ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// FormatSelectColumn 第一个prefix作为as 的值，后面的作为忽略的field
// 如:helper.FormatSelectColumn(userFields, "u", "user_name", "content")
func FormatSelectColumn(data g.MapStrStr, prefix ...string) string {
	prefixLen := len(prefix)
	var prefixStr string
	ignore := g.SliceStr{}
	if prefixLen > 0 {
		prefixStr = prefix[0]
		if prefixLen > 1 {
			ignore = prefix[1:prefixLen]
		}
	} else {
		prefixStr = ""
	}
	columns := make([]string, 0)
	for _, v := range data {
		if len(ignore) > 0 && StringInArray(v, ignore) {
			continue
		}
		var field string
		if len(prefixStr) > 0 {
			field = prefixStr + ".`" + v + "`"
		} else {
			field = v
		}
		columns = append(columns, field)
	}
	return gstr.Join(columns, ",")
}

// GetCurrentUrl 获取当前页面的URL
func GetCurrentUrl(r *ghttp.Request) string {
	var scheme string
	if r.TLS == nil {
		scheme = "http://"
	} else {
		scheme = "https://"
	}
	url := scheme + r.Host + r.RequestURI
	return url
}
func GetUploadBaseUrl() string {
	var uploadUrl string
	uploadUrl = gstr.Trim(g.Cfg().GetString("upload.url"), "/")
	if uploadUrl == "" {
		appBaseUrl := gstr.TrimRight(g.Cfg().GetString("app.baseUrl"), "/")
		uploadUrl = appBaseUrl + "/" + g.Cfg().GetString("upload.dir") + "/"
	} else {
		uploadUrl += "/"
	}
	return uploadUrl
}
func ConvertToUploadUrl(fileName string) string {
	if len(fileName) == 0 {
		return ""
	}
	if strings.HasPrefix(fileName, ":/") {
		return fileName
	}
	localSaveDir := g.Cfg().GetString("upload.localSaveDir") + "/"
	baseUrl := GetUploadBaseUrl()
	fileName = gstr.Replace(fileName, localSaveDir, baseUrl)
	return fileName
}
func ConvertToSaveAs(url string) string {
	if len(url) == 0 {
		return ""
	}
	if !gstr.Contains(url, ":/") {
		return url
	}
	localSaveDir := g.Cfg().GetString("upload.localSaveDir") + "/"
	baseUrl := GetUploadBaseUrl()
	url = gstr.Replace(url, baseUrl, localSaveDir)
	return url
}
func TrimUploadUrl(url string) string {
	if len(url) == 0 {
		return ""
	}
	if !gstr.Contains(url, ":/") {
		return url
	}
	baseUrl := GetUploadBaseUrl()
	if !gstr.Contains(url, baseUrl) {
		return url
	}
	return gstr.Replace(url, baseUrl, "")
}

func GetQrcodeUploadDir(qrcodePath ...string) string {
	fileDir := g.Cfg().GetString("upload.localSaveDir") + "/"
	if len(qrcodePath) > 0 {
		fileDir += qrcodePath[0] + "/"
	}
	dateDir := gtime.Now().Format("Y/md") + "/"
	saveDir := fileDir + dateDir
	if !gfile.Exists(saveDir) {
		_ = gfile.Mkdir(saveDir)
	}
	return saveDir
}

func MergeMap(maps ...map[string]interface{}) map[string]interface{} {
	if len(maps) <= 0 {
		return nil
	}
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
func GetMapKeys(m map[string]interface{}) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

// MapKeySort 获得map排序后的key
func MapKeySort(data map[string]interface{}) []string {
	var keys []string
	if len(data) <= 0 {
		return nil
	}
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

// ListJsonDecode 解析标注格式的list json
// str @example [{"sort":"1", "title":"", "image":"", "link":""}]
func ListJsonDecode(str string) ([]interface{}, error) {
	buf := []byte(str)
	var res []interface{}
	err := json.Unmarshal(buf, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
func UrlEncode(url string) string {
	return netURL.QueryEscape(url)
}

func UrlDecode(str string) (string, error) {
	return netURL.QueryUnescape(str)
}

// RightPad2Len 右补齐字符
func RightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// LeftPad2Len 左补齐字符
func LeftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
