package main

import (
	"testing"
	"strings"
	"fmt"
	"go_interface/tool"
	"strconv"
	"time"
	"github.com/tidwall/gjson"
)

// 测试字符串数组转字符串
func TestStrArray2Str(t *testing.T){
	array := []string{"katy", "key", "test"}
	tmp := strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
	fmt.Println(tmp)
}

// 测试过滤数组里面的重复元素
func filter(array []string) []string {
	result := make(map[string]int)

	for i, data := range array {
		result[data] = i
	}
	var tmpArr []string
	for k, _ := range result {
		tmpArr = append(tmpArr, k)
	}
	return tmpArr
}
// 测试字符串数组去重
// 测试去除数组中重复元素
func TestFilterDuplicateElem(t *testing.T){
	array := []string {"hello", "katy", "key", "hello world", "key", "katy"}
	r := filter(array)
	fmt.Println(tool.StringifyJson(r))
}

// 从map[string]string中找到最大的value对应的key值
func getMaxKeyValue(hkjlMap map[string]string) (key, value int) {
	for k, v := range hkjlMap {
		// 将string类型的k转为int,这个绝对不会有err,因为key都是数字1-24
		intKey, _ := strconv.Atoi(k)
		// 将string类型的value转int，如果没有错误，就说明该value是字符串数字
		intValue, err := strconv.Atoi(v)
		if err == nil {
			// 如果value相等，取最大的key
			if value == intValue && key < intKey{
				value = intValue
				key = intKey
			}else if key < intKey && value <= intValue{
				value = intValue
				key = intKey
			}
		}
	}
	return
}

// 查找数组中最小的元素
func getMinValueInArray(array []int) int {
	min := array[0]
	for i := 1; i < len(array)-1; i++ {
		if min > array[i] {
			min = array[i]
		}
	}
	return min
}

// 计算两个日期之间相差多少个月份
func GetDiffMonth(dateOne time.Time, dateTwo time.Time) int {

	// 思路：(年 - 年) * 12 + (月 - 月)
	count := (dateOne.Year() - dateTwo.Year()) * 12 + (int(dateOne.Month()) - int(dateTwo.Month()))
	return count
}

func TestGetJsonData(t *testing.T){
	str := `{"status":200,"total":0,"data":[{"uuid":"45ffeb65-a3cd-49f8-b8a1-c43d8823827c","roleUuid":"1","userName":"root","password":"ED45BEF8","phone":"18689841696","phoneState":1,"email":"463121419@qq.com","emailState":1,"updateTime":"2018-05-29T10:30:01+08:00","createTime":"2018-05-16T12:07:55+08:00","describe":"-测试环境初始化账号","roleName":"Super_admin","alias":"超级管理员","status":1},{"uuid":"7e3cd09a-0ae3-4c1a-b65f-64f86d66d5fe","roleUuid":"9f6afc27-d1e8-4084-ba30-1b27aae94a3f","userName":"qeq","password":"ED45BEF8","phone":"","phoneState":0,"email":"","emailState":0,"updateTime":"2018-05-24T08:47:11+08:00","createTime":"2018-05-24T15:19:00+08:00","describe":"eqeq","roleName":"q","alias":"","status":0},{"uuid":"c6c44e04-8ea8-444b-b726-4c790f159cb6","roleUuid":"dc49b101-2bb8-4316-b551-eef6a24ef1a7","userName":"fanxing","password":"019A9BF2","phone":"15764336108","phoneState":1,"email":"642664360@qq.com","emailState":1,"updateTime":"2018-05-29T06:59:44+08:00","createTime":"2018-05-24T17:43:31+08:00","describe":"fanxing","roleName":"fanxing","alias":"fanxing","status":1},{"uuid":"e4281d2d-ac27-4a61-91a6-4c19f5ffcf70","roleUuid":"9f6afc27-d1e8-4084-ba30-1b27aae94a3f","userName":"qqqqq","password":"8CE74A94","phone":"18689841694","phoneState":0,"email":"323344@qq.com","emailState":1,"updateTime":"2018-05-25T10:11:24+08:00","createTime":"2018-05-23T18:44:00+08:00","describe":"qqq1111","roleName":"q","alias":"root","status":0}],"reason":""}`

	status := gjson.Get(str, "status").Raw
	fmt.Println(status)
}

type Resp struct {
	Status int `json:"status"`
	Total int `json:"total"`
	Data []Data `json:"data"`
	
}
type Data struct {
	Uuid string `json:"uuid"`
	RoleUuid string `json:"roleUuid"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	PhoneState int `json:"phoneState"`
	Email string `json:"email"`
	EmailState int `json:"emailState"`
	UpdateTime string `json:"updateTime"`
	CreateTime string `json:"createTime"`
	Describe string `json:"describe"`
	RoleName string `json:"roleName"`
	Alias string `json:"alias"`
	Status int `json:"status"`
}

func TestShow(t *testing.T){
	fmt.Println(Resp{})
}

// 得到今天的时间范围
func GetToday()(start time.Time, end time.Time){
	now := time.Now()
	start = now.Truncate(time.Hour * 24).Add(time.Hour * -8 )
	next := now.Add(time.Hour * 24 )
	end = time.Date(next.Year(), next.Month(), next.Day(), 0,0,0,0,next.Location())
	return
}

// 得到月初和月末的日期
func GetMonthStartEndDate(checkTime time.Time)(time.Time, time.Time){
	year, month , _ := checkTime.Date()
	start := time.Date(year, month, 1, 0, 0,0, 0, time.Local)
	endTmp := checkTime.AddDate(0, 1, 0)
	yearEndTmp, monthEndTmp , _ := endTmp.Date()
	end := time.Date(yearEndTmp, monthEndTmp, 0,0,0,0,0,  time.Local)
	return tool.GetDate(start), tool.GetDate(end)
}

// 测试往后多少个月
func TestMonthMove(t *testing.T){
	/**
	往后一个月的意思是     月份数+1    日期数-1
	往前N个月的意思是    月份数-N   日期数+1
	 */
	now := time.Now()
	// 当前时间往后1个月
	date1 := now.AddDate(0, 1, -1)
	fmt.Println(date1)
	date2 := now.AddDate(0, -1, 1)
	fmt.Println(date2)
}

// 给定的时间往后month个月:往后一个月的意思是     月份数+1    日期数-1
func GetBackMonth(checkTime time.Time, month int) (time.Time) {
	date := checkTime.AddDate(0, month, -1)
	return date
}

// 当前时间往前month个月:往前N个月的意思是    月份数-N   日期数+1
func GetBeforeMonth(checkTime time.Time, month int) (time.Time) {
	date := checkTime.AddDate(0, month, 1)
	return date
}

// 两个时间相差多少天
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}

// 测试过滤数组里面的重复元素
func FilterDuplicateElem(array []string) []string {
	// key为数组元素，value可以任意，这里value为数组下标
	result := make(map[string]int)
	// 循环数组元素，将每个元素存入result，map会过滤掉重复元素
	for i, data := range array {
		result[data] = i
	}
	// 将map里面的key存入数组中
	var tmpArr []string
	for k, _ := range result {
		tmpArr = append(tmpArr, k)
	}
	return tmpArr
}

// 将string类型的日期转为time.Time类型的日期
func ChangeStrTime2Time(strTime string) time.Time {
	date, err := time.Parse("2006.01.02", strTime)
	if err != nil {
		date, err = time.Parse("2006年01月02日", strTime)
		if err != nil {
			date, err = time.Parse("2006.01 .02", strTime)
			if err != nil {
				date, err = time.Parse("2006年01月", strTime)
				if err != nil {
					date, err = time.Parse("2006 01 02", strTime)
					if err != nil {
						date, err = time.Parse("2006-1-2", strTime)
						if err != nil {
							date, err = time.Parse("2006.01", strTime)
							if err != nil {
								logger.Warn("warn", "从征信json里面取出来的查询日期转换格式出错", err.Error())
							}
						}

					}
				}
			}
		}
	}
	return date
}