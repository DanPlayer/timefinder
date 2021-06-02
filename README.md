## TimeFinder

### 简介

- 分词基于SeGo
- 对自然语言（中文）提取时间

##### 使用

```
go get github.com/DanPlayer/timefinder
```

```
var msg string
var extract []time.Time

// 初始化timefinder
segmenter := timefinder.New()

msg = " 6月9日有一场show要去观看"
// 解析话语词汇
extract = segmenter.TimeExtract(msg)
fmt.Println(msg)
fmt.Println(extract[0].Format(timeFormat))
```

##### 示例

```
func main() {
    var msg string
    var extract []time.Time
    
    segmenter := timefinder.New()
    msg = " 6月9日有一场show要去观看"
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
}
```

```
6月9日有一场show要去观看
2021-06-09 00:00:00
后天早上10:35的会议，需要及时参与
2021-06-04 10:35:00
明天下午三点的飞机，提醒我坐车
2021-06-03 15:00:00
一个小时后提醒我喝水
2021-06-02 18:40:07
明天早上8:00喊我起床
2021-06-03 08:00:00
明天早上8点喊我起床
2021-06-03 08:00:00
明早十点喊我喝水
2021-06-03 10:00:00
明天早上十点喊我喝水
2021-06-03 10:00:00
明天下午三点提醒我喝水
2021-06-03 15:00:00
一天后提醒我喝水
2021-06-03 17:40:07
一年后提醒我喝水
2022-06-02 17:40:07
一个月后提醒我喝水
2021-07-02 17:40:07
一月后提醒我喝水
2021-07-02 17:40:07
我要住到大后天
2021-06-05 00:00:00
我要住到明天
2021-06-03 00:00:00
下个月到上个月再到这个月
2021-07-02 00:00:00
我要住到明天下午三点十分
2021-06-03 15:10:00
帮我预定明天凌晨3点的飞机
2021-06-03 03:00:00
今天13:00的飞机
2021-06-02 13:00:00
3月15号的飞机
2021-03-15 00:00:00
昨天凌晨2点
2021-06-01 02:00:00
十分钟后提醒我喝水
2021-06-02 17:50:07
```