package spiderkit

import (
	"fmt"
	"strings"
)

/*获得页面上的全部文件信息*/
func GetPageImginfos2Chan(url, imgDir string, chImgMaps chan<- map[string]string) {
	imginfos := GetPageImginfos(url, imgDir)
	fmt.Println("imginfos=", imginfos)

	for _, infoMap := range imginfos {
		chImgMaps <- infoMap
	}
}

/*获得页面上的全部图片信息（链接+文件名）*/
func GetPageImginfos(url, imgDir string) []map[string]string {

	html := GetHtml(url)
	//html = string(ConvertToByte(html, "gbk", "utf8"))
	//fmt.Println(html)

	//re := regexp.MustCompile(reImgTagStr)
	rets := reImgTag.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))

	imginfos := make([]map[string]string, 0)
	for _, ret := range rets {
		imgInfo := make(map[string]string)
		imgUrl := ret[1]
		imgInfo["url"] = imgUrl
		imgInfo["filename"] = GetImgNameFromTag(ret[0], imgUrl, imgDir)

		imginfos = append(imginfos, imgInfo)
	}

	return imginfos
}

/*
从<img>标签中提取文件名（含地址）：
优先级：alt+链接文件名 > 链接文件名 > 时间戳+随机数
参数：
imgTag 图片<img>标签
imgUrl 图片链接，可以为空字符串，不传时会自动从<img>标签中检索
imgDir 目录位置
suffix 文件名后缀
*/
func GetImgNameFromTag(imgTag, imgUrl, imgDir string) (filename string) {
	//var filename string

	if imgUrl == "" {
		//re := regexp.MustCompile(reTagSrcStr)
		rets := reTagSrc.FindAllStringSubmatch(imgTag, -1)
		if len(rets) > 0 {
			imgUrl = rets[0][1]
		}
	}

	//获得图片格式
	imgName := GetImgNameFromImgurl(imgUrl)
	suffix := ".jpg"
	if imgName != "" {
		suffix = imgName[strings.LastIndex(imgName, "."):]
	}

	//尝试从imgTag中提取alt
	//re := regexp.MustCompile(reTagAltStr)
	rets := reTagAlt.FindAllStringSubmatch(imgTag, 1)

	if len(rets) > 0 && imgName != "" {
		//首选alt
		alt := rets[0][1]
		alt = strings.Replace(alt, ":", "_", -1)
		filename = alt + imgName
	} else if imgName != "" {
		//次选链接中的文件名
		filename = imgName
	} else {
		//最末时间戳+随机数
		filename = GetRandomFileName() + suffix
	}
	filename = imgDir + filename
	return filename
}

/*从imgUrl中摘取图片名称*/
func GetImgNameFromImgurl(imgUrl string) string {
	//re := regexp.MustCompile(reImgNameStr)
	rets := reImgName.FindAllStringSubmatch(imgUrl, -1)
	if len(rets) > 0 {
		return rets[0][1]
	} else {
		return ""
	}
}
