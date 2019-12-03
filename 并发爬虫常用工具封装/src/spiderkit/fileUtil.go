package spiderkit

import (
	"strconv"
	"time"
)

/*生成时间戳_随机数文件名*/
func GetRandomFileName() string {
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	randomNum := strconv.Itoa(GetRandomInt(1000, 10000))
	return timestamp + "_" + randomNum
}
