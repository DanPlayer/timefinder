## TimeFinder

### 简介

- 分词基于SeGo
- 对自然语言（中文）提取时间

##### 使用
```
go get github.com/DanPlayer/timefinder
```

##### 示例
```
func main() {
    var msg string
    var extract []string
    
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
}
```
```
一天后提醒我喝水
[2021-05-22 10:44:11]
一年后提醒我喝水
[2022-05-21 10:44:26]
一个月后提醒我喝水
[2021-06-21 10:44:41]
一月后提醒我喝水
[2021-06-21 10:45:00]
一个小时后提醒我喝水
[2021-05-21 11:45:19]
我要住到大后天
[2021-05-24 00:00:00]
我要住到明天
[2021-05-22 00:00:00]
下个月到上个月再到这个月
[2021-06-21 00:00:00 2021-04-21 00:00:00 2021-05-21 00:00:00]
我要住到明天下午三点十分
[2021-05-22 03:10:00]
帮我预定明天凌晨3点的飞机
[2021-05-22 03:00:00]
今天13:00的飞机
[2021-05-21 13:00:00]
3月15号的飞机
[2021-03-15 00:00:00]
昨天凌晨2点
[2021-05-20 02:00:00]
十分钟后提醒我喝水
[2021-05-21 10:58:21]
```