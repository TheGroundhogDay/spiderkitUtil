package spiderkit

import "regexp"

var(
	/*图片：
	<img...src="(http[\s\S]+?)"...>
	<img...alt="中国军团“备战”维密 大秀前先来一波颜值霸屏！"...src="http://cms-bucket.nosdn.127.net/2018/11/09/0e08464f5c9f4448ad5db7683283f571.jpeg?imageView&amp;thumbnail=200y125&amp;quality=85">
	*/
	//reImgTagStr = `<img[\s\S]+?src="(http[\s\S]+?)"[\s\S]*?>`
	//reImgTagStr = `<img[\s\S]+?src="(http[\s\S]+?)"`
	reImgTagStr = `<img.+?src="(http.+?)".*?>`
	reImgAltStr = `<img.+?alt="(.+?)"`

	/*img标签中的alt属性*/
	reTagAltStr = `alt="([\s\S]+?)"`
	/*img标签中的src属性*/
	reTagSrcStr = `src="(http.+?)"`

	/*
	图片链接中的图片名称
	http://cms-bucket.nosdn.127.net/2018/11/09/7e88b8526ff141129809d8ae7c718e51.jpeg?imageView&thumbnail=185y116&quality=85
	http://img2.money.126.net/chart/hs/time/180x120/0000001.png
	http://cms-bucket.nosdn.127.net/2018/05/31/bc7d30ff42194c35a4743834a77ec97b.png?imageView&thumbnail=90y90&quality=85
	*/
	reImgNameStr = `/(\w+\.((jpg)|(jpeg)|(png)|(gif)|(bmp)|(webp)|(swf)|(ico)))`


	/*预编译正则对象*/
	reImgTag,reImgAlt,reTagAlt,reTagSrc,reImgName *regexp.Regexp
)

func init() {
	reImgTag = regexp.MustCompile(reImgTagStr)
	reImgAlt = regexp.MustCompile(reImgAltStr)
	reTagAlt = regexp.MustCompile(reTagAltStr)
	reTagSrc = regexp.MustCompile(reTagSrcStr)
	reImgName = regexp.MustCompile(reImgNameStr)
}
