package timefinder

import (
	"fmt"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/huichen/sego"
)

type TimeFinder struct {
	Segmenter sego.Segmenter
}

const timeFormat = "2006-01-02 15:04:05"
const dayFormat = "2006年01月02日"

// utilCnNum 中文数字
var utilCnNum = map[string]int{
	"零": 0, "一": 1, "二": 2, "两": 2, "三": 3, "四": 4,
	"五": 5, "六": 6, "七": 7, "八": 8, "九": 9,
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4,
	"5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

// utilCnUnit 中文单元
var utilCnUnit = map[string]int{
	"个": 1, "十": 10, "百": 100, "千": 1000, "万": 10000,
}

var keyDate = map[string]int{"今天": 0, "明天": 1, "明早": 1, "后天": 2, "大后天": 3, "昨天": -1, "前天": -2, "大前天": -3}
var keyYear = map[string]int{"今年": 0, "明年": 1, "后年": 2, "大后年": 3, "去年": -1, "前年": -2, "大前年": -2}
var keyMonth = map[string]int{"这个月": 0, "上个月": -1, "下个月": 1}
var keyWeekDay = map[string]int{
	"这个周": 0, "这周": 0, "本周": 0, "周": 0, "下周": 7, "下下周": 14, "上周": -7, "上上周": -14,
	"这个星期": 0, "这星期": 0, "本星期": 0, "星期": 0, "下个星期": 7, "下下个星期": 14, "上个星期": -7, "上上个星期": -14,
	"这个礼拜": 0, "这礼拜": 0, "本礼拜": 0, "礼拜": 0, "下个礼拜": 7, "下下个礼拜": 14, "上个礼拜": -7, "上上个礼拜": -14,
}

var jiebaTimeTag = []string{"m", "t", "f", "q"}

// cn2dig 中文单元转化为数字
func cn2dig(src string) (rsl int) {
	if src == "" {
		return -1
	}
	compile, err := regexp.Compile("\\d+")
	if err != nil {
		return
	}
	m := compile.FindString(src)
	if m == "0" {
		return
	}
	if m != "" {
		return stringToInt(m)
	}
	rsl = 0
	unit := 1
	for _, ele := range []rune(src) {
		_, exist := utilCnUnit[string(ele)]
		if !exist {
			num, exist := utilCnNum[string(ele)]
			if exist {
				rsl += num * unit
			} else {
				if string(ele) == "小" {
					continue
				}
				// 对礼拜、星期特殊单元的处理
				if string(ele) == "礼" || string(ele) == "星" {
					continue
				}
				return -1
			}
		} else {
			unit, _ = utilCnUnit[string(ele)]
		}
	}
	if rsl < unit {
		rsl += unit
	}
	return
}

// year2dig 年份转化为数字
func year2dig(year string) (rsl int) {
	var res string
	for _, ele := range []rune(year) {
		num, exist := utilCnNum[string(ele)]
		if exist {
			res = res + strconv.Itoa(num)
		} else {
			res = res + string(ele)
		}
	}

	rsl = -1
	compile, err := regexp.Compile("\\d+")
	if err != nil {
		return
	}
	m := compile.FindString(res)

	if m == "" {
		return
	}

	if len(m) == 2 {
		rsl = int(time.Now().Year()/100)*100 + stringToInt(m)
	} else {
		rsl = stringToInt(m)
	}

	return
}

// weekday2dig 周时转化为日差额数字
func weekday2dig(weekday string) (rsl int) {
	if weekday == "" {
		return
	}
	curWeekDay := int(time.Now().Weekday())
	if curWeekDay == 0 {
		curWeekDay = 7
	}
	weekdaySplit := []rune(weekday)
	numStr := weekdaySplit[len(weekdaySplit)-1]
	var num int
	if string(numStr) == "日" {
		num = 7
	} else {
		num = cn2dig(string(numStr))
	}
	weekdayPre := trimLastChar(weekday)
	for k, v := range keyWeekDay {
		if weekdayPre == k {
			rsl = v - (curWeekDay - num)
		}
	}

	return rsl
}

// parseDatetime 函数，用以将每个提取到的文本日期串进行时间转换。
// 其主要通过正则表达式将日期串进行切割，分为"年" "月" "日" "时" H分""秒"等具体维度，
// 然后针对每个子维度单独再进行识别。
func parseDatetime(msg string) (parseTime time.Time) {
	if len(msg) == 0 {
		return
	}

	var targetDate string

	compile, err := regexp.Compile("" +
		"([0-9零一二两三四五六七八九十]+年)?" +
		"([0-9一二两三四五六七八九十]+(?:个月|月))?" +
		"([0-9一二两三四五六七八九十]+[天号日])?" +
		"([0-9零一二三四五六七八九十百]+(?:星期|周|礼拜|个星期|个礼拜))?" +
		"((?:这周|这个周|本周|下周|下下周|上周|上上周|这星期|这个星期|下个星期|下下个星期|上个星期|上上个星期|这礼拜|这个礼拜|下个礼拜|下下个礼拜|上个礼拜|上上个礼拜|周|星期|礼拜)[1-7一二三四五六七日])?" +
		"(上午|中午|下午|晚上|晚|早上|早|凌晨｜傍晚)?" +
		"([0-9零一二两三四五六七八九十百]+(?:[点\\.时]|个小时|小时))?" +
		"([0-9零一二三四五六七八九十百]+分)?" +
		"([0-9零一二三四五六七八九十百]+秒)?" +
		//"([0-9零一二三四五六七八九十百]+(?:星期|周|礼拜|个星期|个礼拜))?" +
		//"((?:这周|这个周|本周|下周|下下周|上周|上上周|这星期|这个星期|下个星期|下下个星期|上个星期|上上个星期|这礼拜|这个礼拜|下个礼拜|下下个礼拜|上个礼拜|上上个礼拜|周|星期|礼拜)[1-7一二三四五六七日])?" +
		"([0-9零一二两三四五六七八九十百]+:[0-9零一二两三四五六七八九十百]+(?::[0-9零一二两三四五六七八九十百])?)?")

	if err != nil {
		return
	}

	allMatched := compile.FindStringSubmatch(msg)
	if len(allMatched) == 0 {
		return
	}
	emptyMatched := true
	for _, matched := range allMatched {
		if matched != "" {
			emptyMatched = false
		}
	}
	if emptyMatched {
		return
	}

	compileDirect, _ := regexp.Compile("[前后]")
	direction := compileDirect.FindString(msg)

	year := allMatched[1]
	month := allMatched[2]
	day := allMatched[3]
	hour := "00"
	if len(allMatched[7]) > 0 {
		hour = allMatched[7]
	}
	minute := "00"
	if len(allMatched[8]) > 0 {
		minute = allMatched[8]
	}
	// 增加对时半的支持
	if minute == "00" && strings.Contains(msg, hour+"半") {
		minute = "30"
	}

	second := "00"
	if len(allMatched[9]) > 0 {
		second = allMatched[9]
	}

	//0  1  2  3    4       5    6    7    8  9  10
	//空 年 月 日 weekday  week  上午  小时  分  秒 数

	week := allMatched[4]
	weekday := allMatched[5]

	if len(allMatched[10]) > 0 {
		matched10 := strings.Split(allMatched[10], ":")
		for i, s := range matched10 {
			if i == 0 {
				hour = s
			}
			if i == 1 {
				minute = s
			}
			if i == 2 {
				second = s
			}
		}
	}

	res := map[string]string{
		"year":    year,
		"month":   month,
		"day":     day,
		"week":    week,
		"weekday": weekday,
		"hour":    hour,
		"minute":  minute,
		"second":  second,
	}

	params := make(map[string]int)
	for k, v := range res {
		var tmp int
		if k == "year" {
			tmp = year2dig(trimLastChar(v))
		} else if k == "weekday" {
			tmp = weekday2dig(v)
		} else {
			_, err := strconv.Atoi(v)
			if err != nil {
				tmp = cn2dig(trimLastChar(v))
			} else {
				tmp = cn2dig(v)
			}
		}
		params[k] = tmp
	}
	// 获取今天的凌晨时间 (00:00:00)
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	targetDate = ternaryTime(params["year"], midnight.Year()) + "-" +
		ternaryTime(params["month"], int(midnight.Month())) + "-" +
		ternaryTime(params["day"], midnight.Day()) + " " +
		ternaryTime(params["hour"], 0) + ":" +
		ternaryTime(params["minute"], 0) + ":" +
		ternaryTime(params["second"], 0)

	nowUnix := now.Unix()
	midnightUnix := midnight.Unix()
	if len(direction) > 0 {
		for k, v := range params {
			if k == "year" && v > 0 {
				if direction == "前" {
					nowUnix = now.AddDate(-v, 0, 0).Unix()
				} else {
					nowUnix = now.AddDate(v, 0, 0).Unix()
				}
			}
			if k == "month" && v > 0 {
				if direction == "前" {
					nowUnix = now.AddDate(0, -v, 0).Unix()
				} else {
					nowUnix = now.AddDate(0, v, 0).Unix()
				}
			}
			if k == "day" && v > 0 {
				if direction == "前" {
					nowUnix = now.AddDate(0, 0, -v).Unix()
				} else {
					nowUnix = now.AddDate(0, 0, v).Unix()
				}
			}
			if k == "hour" && v > 0 {
				if direction == "前" {
					nowUnix = nowUnix - int64(v*60*60)
				} else {
					nowUnix = nowUnix + int64(v*60*60)
				}
			}
			if k == "minute" && v > 0 {
				if direction == "前" {
					nowUnix = nowUnix - int64(v*60)
				} else {
					nowUnix = nowUnix + int64(v*60)
				}
			}
			if k == "second" && v > 0 {
				if direction == "前" {
					nowUnix = nowUnix - int64(v)
				} else {
					nowUnix = nowUnix + int64(v)
				}
			}
			if k == "week" && v > 0 {
				if direction == "前" {
					nowUnix = now.AddDate(0, 0, -(v * 7)).Unix()
				} else {
					nowUnix = now.AddDate(0, 0, v*7).Unix()
				}
			}
		}
		targetDate = time.Unix(nowUnix, 0).Format(timeFormat)
	} else {
		// 创建一个基于midnight的时间变量，用于累积修改
		currentTime := midnight

		if params["year"] != -1 {
			currentTime = currentTime.AddDate(params["year"]-midnight.Year(), 0, 0)
			midnightUnix = currentTime.Unix()
			targetDate = time.Unix(midnightUnix, 0).Format(timeFormat)
		}

		if params["month"] != -1 {
			// 计算月份差值时，使用当前年份作为基准
			currentTime = currentTime.AddDate(0, params["month"]-int(currentTime.Month()), 0)
			midnightUnix = currentTime.Unix()
			targetDate = time.Unix(midnightUnix, 0).Format(timeFormat)
		}

		if params["day"] != -1 {
			// 计算日期差值时，使用当前月份作为基准
			currentTime = currentTime.AddDate(0, 0, params["day"]-currentTime.Day())
			midnightUnix = currentTime.Unix()
			targetDate = time.Unix(midnightUnix, 0).Format(timeFormat)
		}

		if params["weekday"] != -1 {
			currentTime = currentTime.AddDate(0, 0, params["weekday"])
			midnightUnix = currentTime.Unix()
			targetDate = time.Unix(midnightUnix, 0).Format(timeFormat)
		}

		if params["hour"] != 0 {
			currentTime = currentTime.Add(time.Duration(params["hour"]) * time.Hour)
			midnightUnix = currentTime.Unix()
			targetDate = time.Unix(midnightUnix, 0).Format(timeFormat)
		}

		if params["minute"] != 0 {
			currentTime = currentTime.Add(time.Duration(params["minute"]) * time.Minute)
			midnightUnix = currentTime.Unix()
			targetDate = time.Unix(midnightUnix, 0).Format(timeFormat)
		}

		if params["second"] != 0 {
			currentTime = currentTime.Add(time.Duration(params["second"]) * time.Second)
			midnightUnix = currentTime.Unix()
			targetDate = time.Unix(midnightUnix, 0).Format(timeFormat)
		}
	}

	//0  1  2  3  4   5   6   7  8  9  10
	//空 年 月 日 星期  周 上午 小时 分 秒 数

	isPm := allMatched[6]
	if len(isPm) > 0 {
		if isPm == "下午" || isPm == "晚上" || isPm == "中午" || isPm == "傍晚" {
			parse, err := time.Parse(timeFormat, targetDate)
			if err != nil {
				return
			}

			if parse.Hour() == 0 && parse.Minute() == 0 && parse.Second() == 0 {
				switch isPm {
				case "下午":
					parse = parse.Add(3 * 60 * 60 * time.Second)
				case "晚上":
					parse = parse.Add(8 * 60 * 60 * time.Second)
				case "傍晚":
					parse = parse.Add(6 * 60 * 60 * time.Second)
				}

				if parse.Hour() < 12 {
					parse = parse.Add(60 * 60 * 12 * time.Second)
				}
				targetDate = parse.Format(timeFormat)
			}

			if parse.Hour() == 0 && parse.Minute() == 0 && parse.Second() == 0 {
				parse = parse.Add(8 * 60 * 60 * time.Second)
			}

			if parse.Hour() < 12 {
				parse = parse.Add(60 * 60 * 12 * time.Second)
			}
			targetDate = parse.Format(timeFormat)
		}
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	parseTime, err = time.ParseInLocation(timeFormat, targetDate, loc)
	if err != nil {
		return
	}
	return
}

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func checkTimeValid(word string) (rsl string) {
	matched, err := regexp.MatchString("\\d+$", word)
	if err != nil {
		fmt.Printf("error : %v", err.Error())
		return
	}
	if matched && len(word) < 6 {
		return
	}

	compile, err := regexp.Compile("[号|日]\\d+$")
	if err != nil {
		fmt.Printf("error : %v", err.Error())
		return
	}
	replacedWord := compile.ReplaceAllString(word, "日")
	if replacedWord != word {
		return checkTimeValid(replacedWord)
	}
	return replacedWord
}

// New 初始化
func New(fileNames ...string) *TimeFinder {
	finder := new(TimeFinder)
	currentPath := ""
	if len(fileNames) < 1 {
		// 增加一些特殊词语的分词及词性
		currentPath = path.Join(path.Dir(getCurrentFilePath()), "./jieba_dict.txt") + "," + path.Join(path.Dir(getCurrentFilePath()), "./dictionary.txt")
	} else {
		currentPath = strings.Join(fileNames, ",")
	}

	finder.Segmenter.LoadDictionary(currentPath)

	return finder
}

// TimeExtract 通过Jieba分词将带有时间信息的词进行切分，然后记录连续时间信息的词
// time extract:对句子进行解析，提取其中所有能表示 日期时间的词，并进行上下文拼接
func (tf *TimeFinder) TimeExtract(text string) (finalRes []time.Time) {
	var (
		timeRes []string
		txtList []string
		word    string
		txt     string
		pegList []string
	)

	now := time.Now()

	// 分词
	segments := tf.Segmenter.Segment([]byte(text))

	//str := sego.SegmentsToString(segments, false)
	//fmt.Println(str)

	for _, tag := range segments {
		k := tag.Token().Text()
		v := tag.Token().Pos()
		isXNum := false
		if v == "x" {
			_, err := strconv.Atoi(k)
			if err != nil {
				isXNum = false
			} else {
				isXNum = true
			}
		}
		pegList = append(pegList, fmt.Sprintf("%v/%s", k, v))
		if cpDate, exist := keyDate[k]; exist {
			if word != "" {
				timeRes = append(timeRes, word)
				txtList = append(txtList, txt)
			}
			word = now.AddDate(0, 0, cpDate).Format(dayFormat)
			txt = k
		} else if cpYear, exist := keyYear[k]; exist {
			nYear := now.Year()
			word = strconv.Itoa(nYear) + strconv.Itoa(cpYear) + "年"
			txt += k
		} else if cpMonth, exist := keyMonth[k]; exist {
			nMonth := int(now.Month())
			word = strconv.Itoa(nMonth+cpMonth) + "月"
			txt += k
		} else if word != "" {
			if includes(jiebaTimeTag, v) || k == ":" || isXNum {
				word = word + k
				txt += k
			} else {
				timeRes = append(timeRes, word)
				txtList = append(txtList, txt)
				word = ""
				txt = ""
			}
		} else if includes(jiebaTimeTag, v) || k == ":" || isXNum {
			word = k
			txt = k
		}
	}

	if word != "" {
		timeRes = append(timeRes, word)
		txtList = append(txtList, txt)
	}

	var result []string
	for _, ele := range timeRes {
		valid := checkTimeValid(ele)
		if valid != "" {
			result = append(result, valid)
		}
	}

	fmt.Println(result)

	for _, ele := range result {
		if !parseDatetime(ele).IsZero() {
			finalRes = append(finalRes, parseDatetime(ele))
		}
	}

	return
}

func ternaryTime(ele int, defaultEle int) (re string) {
	eleStr := strconv.Itoa(ele)
	defaultEleStr := strconv.Itoa(defaultEle)
	if eleStr != "" && eleStr != "-1" {
		re = eleStr
	} else {
		re = defaultEleStr
	}
	if len(re) < 2 {
		re = "0" + re
	}
	return
}

func getCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}
