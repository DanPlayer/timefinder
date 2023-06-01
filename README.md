## TimeFinder

### 简介

TimeFinder 是一个用于在文本中查找和提取时间信息的 Golang 库。它提供了一种简单的方式来解析文本并识别其中的时间表达。

- 分词基于SeGo
- 对自然语言（中文）提取时间

## 特性

- 快速准确的时间信息提取。
- 支持多种日期和时间格式。
- 提供灵活的时间范围识别功能。
- 易于集成和使用。

## 安装

要使用 TimeFinder，您需要先安装 Golang。然后，可以使用以下命令从 GitHub 下载和安装 TimeFinder：

```
$ go get github.com/DanPlayer/timefinder

```

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
上述代码会在文本中查找时间信息，并将结果打印输出。您可以根据需要自定义输出格式和进一步处理提取的时间信息。

##### 支持的时间格式
TimeFinder 支持多种日期和时间格式的识别和提取，包括但不限于以下格式：

- 年月日：2023年6月1日、2023-06-01、6/1/2023
- 时间：15:30、15点30分、下午3点30分
- 相对时间：明天、下周五、三天后
- 时间范围：6月1日至6月5日、上午9点到下午5点

##### 更多用法和定制化
TimeFinder 提供了更多功能和选项，以满足不同的时间信息提取需求。您可以查阅 TimeFinder 的文档以获取更详细的用法说明和定制化选项。

##### 贡献和反馈
TimeFinder 是一个开源项目，欢迎您的贡献和反馈。如果您发现问题、有改进建议或者想要贡献代码，请参阅项目的贡献指南并提交 Issue 或 Pull Request。

##### 许可证
TimeFinder 使用 MIT 许可证。有关详细信息，请参阅项目的许可证文件。

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
