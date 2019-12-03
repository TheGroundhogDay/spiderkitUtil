package spiderkit


import (
	"sync"
	"time"
	"math/rand"
)

var(
	randomMT   sync.Mutex
)

/*生成[start,end)之间的随机数*/
func GetRandomInt(start, end int) int {
	randomMT.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := start + r.Intn(end-start)
	randomMT.Unlock()
	return ret
}