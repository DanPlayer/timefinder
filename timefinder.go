package timefinder

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

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
	"个": 1,"十": 10, "百": 100, "千": 1000, "万": 10000,
}

var keyDate = map[string]int{"今天": 0, "明天": 1, "后天": 2, "大后天": 3, "昨天": -1, "前天": -2, "大前天": -3}
var keyYear = map[string]int{"今年": 0, "明年": 1, "后年": 2, "大后年": 3, "去年": -1, "前年": -2, "大前年": -2}
var keyMonth = map[string]int{"这个月": 0, "上个月": -1, "下个月": 1}

var jiebaTimeTag = []string{"m", "t", "f"}

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

// parseDatetime 函数，用以将每个提取到的文本日期串进行时间转换。
// 其主要通过正则表达式将日期串进行切割，分为"年" "月" "日" "时" H分""秒"等具体维度，
// 然后针对每个子维度单独再进行识别。
func parseDatetime(msg string) (targetDate string) {
	if len(msg) == 0 {
		return
	}

	compile, err := regexp.Compile("" +
		"([0-9零一二两三四五六七八九十]+年)?" +
		"([0-9一二两三四五六七八九十]+(?:个月|月))?" +
		"([0-9一二两三四五六七八九十]+[天号日])?" +
		"([上中下午晚早凌晨]+)?" +
		"([0-9零一二两三四五六七八九十百]+(?:[点:\\.时]|个小时|小时))?" +
		"([0-9零一二三四五六七八九十百]+分?)?" +
		"([0-9零一二三四五六七八九十百]+秒)?")
	if err != nil {
		return
	}

	allMatched := compile.FindStringSubmatch(msg)
	if len(allMatched) == 0 {
		return
	}

	compileDirect, _ := regexp.Compile("[前后]")
	direction := compileDirect.FindString(msg)

	year := allMatched[1]
	month := allMatched[2]
	day := allMatched[3]
	hour := "00"
	if len(allMatched[5]) > 0 {
		hour = allMatched[5]
	}
	minute := "00"
	if len(allMatched[6]) > 0 {
		minute = allMatched[6]
	}
	second := "00"
	if len(allMatched[7]) > 0 {
		second = allMatched[7]
	}

	res := map[string]string{
		"year":   year,
		"month":  month,
		"day":    day,
		"hour":   hour,
		"minute": minute,
		"second": second,
	}

	params := make(map[string]int)
	for k, v := range res {
		var tmp int
		if k == "year" {
			tmp = year2dig(trimLastChar(v))
		} else {
			tmp = cn2dig(trimLastChar(v))
		}
		params[k] = tmp
	}
	now := time.Now()
	if len(direction) > 0 {
		nowUnix := now.Unix()
		for k, v := range params {
			if k == "year" && v > 0 {
				if direction == "前" {
					nowUnix = now.AddDate( -v, 0, 0).Unix()
				} else {
					nowUnix = now.AddDate( v, 0, 0).Unix()
				}
			}
			if k == "month" && v > 0 {
				if direction == "前" {
					nowUnix = now.AddDate( 0, -v, 0).Unix()
				} else {
					nowUnix = now.AddDate( 0, v, 0).Unix()
				}
			}
			if k == "day" && v > 0 {
				if direction == "前" {
					nowUnix = now.AddDate( 0, 0, -v).Unix()
				} else {
					nowUnix = now.AddDate( 0, 0, v).Unix()
				}
			}
			if k == "hour" && v > 0 {
				if direction == "前" {
					nowUnix = nowUnix - int64(v * 60 * 60)
				} else {
					nowUnix = nowUnix + int64(v * 60 * 60)
				}
			}
			if k == "minute" && v > 0 {
				if direction == "前" {
					nowUnix = nowUnix - int64(v * 60)
				} else {
					nowUnix = nowUnix + int64(v * 60)
				}
			}
			if k == "second" && v > 0 {
				if direction == "前" {
					nowUnix = nowUnix - int64(v)
				} else {
					nowUnix = nowUnix + int64(v)
				}
			}
		}
		targetDate = time.Unix(nowUnix, 0).Format(timeFormat)
	} else {
		// 需要在today的基础上修正replace
		targetDate = ternaryTime(params["year"], now.Year()) + "-" +
			ternaryTime(params["month"], int(now.Month())) + "-" +
			ternaryTime(params["day"], now.Day()) + " " +
			ternaryTime(params["hour"], now.Hour()) + ":" +
			ternaryTime(params["minute"], now.Minute()) + ":" +
			ternaryTime(params["second"], now.Second())
	}

	isPm := allMatched[4]
	if len(isPm) > 0 {
		if isPm == "下午" || isPm == "晚上" || isPm == "中午" {
			parse, err := time.Parse(timeFormat, targetDate)
			if err != nil {
				return
			}

			if parse.Hour() < 12 {
				parse = parse.Add(60 * 60 * 12)
			}
			targetDate = parse.Format(timeFormat)
		}
	}
	return targetDate
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

// TimeExtract 通过Jieba分词将带有时间信息的词进行切分，然后记录连续时间信息的词
// time extract:对句子进行解析，提取其中所有能表示 日期时间的词，并进行上下文拼接
func TimeExtract(text string) (finalRes []string) {
	var (
		timeRes []string
		txtList []string
		word    string
		txt     string
		pegList []string
	)

	now := time.Now()

	// 增加一些特殊词语的分词及词性
	currentPath := path.Join(path.Dir(getCurrentFilePath()), "./jieba_dict.txt")
	gojieba.USER_DICT_PATH = currentPath
	psg := gojieba.NewJieba()
	defer psg.Free()

	cutList := psg.Tag(text)
	for _, tag := range cutList {
		tagSplit := strings.Split(tag, "/")
		k := tagSplit[0]
		v := tagSplit[1]
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
			if includes(jiebaTimeTag, v) || k == ":" {
				word = word + k
				txt += k
			} else {
				timeRes = append(timeRes, word)
				txtList = append(txtList, txt)
				word = ""
				txt = ""
			}
		} else if includes(jiebaTimeTag, v) || k == ":" {
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

	for _, ele := range result {
		finalRes = append(finalRes, parseDatetime(ele))
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
