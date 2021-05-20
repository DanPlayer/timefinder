package timefinder

import "strconv"

// stringToInt 字符串转INT
func stringToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func includes(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}