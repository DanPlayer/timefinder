package tests

import (
	"fmt"
	"github.com/516396859/timefinder/v2"
	"testing"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func TestAny(t *testing.T) {
	segmenter := timefinder.New("../jieba_dict.txt", "../dictionary.txt")

	//var msg = "下个月8日"
	//nlpTransferDate(msg, segmenter)

	var msg = "下周三晚上9点"
	nlpTransferDate(msg, segmenter)

	msg = "一个小时后提醒我喝水"
	nlpTransferDate(msg, segmenter)

	msg = "周日19点"
	nlpTransferDate(msg, segmenter)

	msg = "周日下午2点半"
	nlpTransferDate(msg, segmenter)

	msg = "周日下午18点"
	nlpTransferDate(msg, segmenter)

	msg = "19号下午9点"
	nlpTransferDate(msg, segmenter)

	msg = "下个月16号上午9点"
	nlpTransferDate(msg, segmenter)

	msg = "下周三晚上9点"
	nlpTransferDate(msg, segmenter)

	msg = "下个月上午17点"
	nlpTransferDate(msg, segmenter)

	msg = "下个月8号"
	nlpTransferDate(msg, segmenter)

}

func nlpTransferDate(msg string, segmenter *timefinder.TimeFinder) {
	var extract []time.Time
	extract = segmenter.TimeExtract(msg)
	fmt.Println()
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))
}

func TestTimeFinder(t *testing.T) {
	var msg string
	var extract []time.Time

	segmenter := timefinder.New("../jieba_dict.txt", "../dictionary.txt")
	msg = "今天晚上8点"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "今天下午三点"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "9点半去跑步"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "6月9日有一场show要去观看"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "后天早上10:35的会议，需要及时参与"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "明天下午三点的飞机，提醒我坐车"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "一个小时后提醒我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "明天早上8:00喊我起床"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "明天早上8点喊我起床"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "明早十点喊我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "明天早上十点喊我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "明天下午三点提醒我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "一天后提醒我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "一年后提醒我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "一个月后提醒我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "一月后提醒我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "我要住到大后天"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "我要住到明天"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "下个月到上个月再到这个月"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "我要住到明天下午三点十分"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "帮我预定明天凌晨3点的飞机"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "今天13:00的飞机"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "3月15号的飞机"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "昨天凌晨2点"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "十分钟后提醒我喝水"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	fmt.Println("end tests")
}

func TestSpecialWeek(t *testing.T) {
	var msg string
	var extract []time.Time

	segmenter := timefinder.New()

	msg = "一个礼拜后提醒我开心"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "一个星期后提醒我开心"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "一周后提醒我开心"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	fmt.Println("end tests")
}

func TestWeekDay(t *testing.T) {
	var msg string
	var extract []time.Time

	segmenter := timefinder.New()

	msg = "下下个礼拜六提醒我做事"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "下下个星期六提醒我做事"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "这周日提醒我做事"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "这周二提醒我做事"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))

	msg = "下下周六提醒我做事"
	extract = segmenter.TimeExtract(msg)
	fmt.Println(msg)
	fmt.Println(extract[0].Format(timeFormat))
}
