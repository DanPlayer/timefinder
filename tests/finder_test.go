package tests

import (
	"fmt"
	"github.com/DanPlayer/timefinder"
	"testing"
)

func TestTimeFinder(t *testing.T) {
	var msg string
	var extract []string

	msg = "明天下午三点提醒我喝水"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "一天后提醒我喝水"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "一年后提醒我喝水"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "一个月后提醒我喝水"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "一月后提醒我喝水"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "一个小时后提醒我喝水"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "我要住到大后天"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "我要住到明天"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "下个月到上个月再到这个月"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "我要住到明天下午三点十分"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "帮我预定明天凌晨3点的飞机"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "今天13:00的飞机"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "3月15号的飞机"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "昨天凌晨2点"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "十分钟后提醒我喝水"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	fmt.Println("end tests")
}

func TestSpecialWeek(t *testing.T) {
	var msg string
	var extract []string

	msg = "一个礼拜后提醒我开心"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "一个星期后提醒我开心"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "一周后提醒我开心"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	fmt.Println("end tests")
}

func TestWeekDay(t *testing.T)  {
	var msg string
	var extract []string

	msg = "这周日提醒我做事"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "这周二提醒我做事"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "下下周六提醒我做事"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "下下个礼拜六提醒我做事"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)

	msg = "下下个星期六提醒我做事"
	extract = timefinder.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract)
}