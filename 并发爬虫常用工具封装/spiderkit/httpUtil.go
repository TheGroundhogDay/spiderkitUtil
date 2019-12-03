package spiderkit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net"
	"time"
)

const (
	DIAL_TIMEOUT = 10 * time.Second
	RW_TIMEOUT   = 10 * time.Second
)

var (
	httpClient http.Client
)

func init() {
	httpClient = http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {

				//设置连接请求超时时间
				conn, err := net.DialTimeout(netw, addr, DIAL_TIMEOUT)
				if err != nil {
					return nil, err
				}

				//设置连接的读写超时时间
				deadline := time.Now().Add(RW_TIMEOUT)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}
}

/*获得页面html*/
func GetHtml(url string) string {
	resp, err := httpClient.Get(url)
	HandleError(err,`httpClient.Get(url)`)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}

/*下载文件*/
func DownloadFileWithClient(url string, filename string) {
	//fmt.Println("DownloadFileWithClient...")
	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(filename, "下载失败！")
		return
	}
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(filename, imgBytes, 0644)
	if err == nil {
		fmt.Println(filename, "下载成功！")
	} else {
		fmt.Println(filename, "下载失败！")
	}

}
