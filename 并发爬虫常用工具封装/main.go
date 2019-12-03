package main

import (
	"strconv"
	"sync"
	"spiderkit"
)

var (
	chSem      = make(chan int, 10)
	chImgMaps   = make(chan map[string]string,1000)

	wg4Imginfo  sync.WaitGroup
	wg4Download sync.WaitGroup

	//图片存储地址
	imgDir = `C:\Users\75515\Desktop\爬虫演示\`
)

func main() {
	baseUrl := "https://www.duotoo.com/zt/rbmn/index"
	for i := 1; i < 5; i++ {
		var url string
		if i != 1 {
			url = baseUrl + "_" + strconv.Itoa(i) + ".html"
		} else {
			url = baseUrl + ".html"
		}

		wg4Imginfo.Add(1)
		go func(theUrl string) {
			spiderkit.GetPageImginfos2Chan(theUrl,imgDir,chImgMaps)
			wg4Imginfo.Done()
		}(url)
	}

	go func() {
		wg4Imginfo.Wait()
		close(chImgMaps)
		//fmt.Println("chImgMaps closed!")
	}()

	for imgMap := range chImgMaps {
		//fmt.Println("imgMap got:",imgMap)
		wg4Download.Add(1)
		go func(im map[string]string) {
			chSem <- 123
			spiderkit.DownloadFileWithClient(im["url"], im["filename"])
			<-chSem
			wg4Download.Done()
		}(imgMap)
	}

	wg4Download.Wait()
}
