## TimeFinder

### 简介

- 分词基于GoJieba
- 对自然语言（中文）提取时间

##### 示例
```
func main() {
    var msg string
	msg = "我要住到大后天"
	extract := timefinder.TimeExtract(msg)
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

	msg = "帮我我预定明天凌晨3点的飞机"
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
}
```
```
我要住到大后天
[2021-05-23 15:02:00]
我要住到明天
[2021-05-21 15:02:15]
下个月到上个月再到这个月
[20211-05-20 15:02:30 2021-05-20 15:202:30 2021-05-20 15:02:30]
我要住到明天下午三点十分
[2021-05-21 01:01:51]
帮我我预定明天凌晨3点的飞机
[2021-05-21 15:03:11 2021-05-20 15:03:11]
今天13:00的飞机
[2021-05-20 13:03:32]
3月15号的飞机
[2021-05-20 15:03:53]
昨天凌晨2点
[2021-05-19 15:04:12 2021-05-20 15:04:12]
```