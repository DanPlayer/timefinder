package tests

import (
	"fmt"
	"github.com/DanPlayer/timefinder"
	"testing"
)

func TestTimeFinder(t *testing.T) {
	var msg string
	var extract []string

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
