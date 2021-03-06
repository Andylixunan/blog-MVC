package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
)

// 定义查询文章，每页显示的文章量
const ArticleDisplayNum = 5

// 日志工具
var Logger *log.Logger

func init() {
	Logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
}

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

// 将传入的时间戳转为时间
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

func SwitchMarkdownToHTML(content string) template.HTML {
	// transform md to html
	markdown := blackfriday.MarkdownCommon([]byte(content))
	// find element with <code></code> tags
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	if err != nil {
		log.Fatal(err)
	}
	// highlight the code elements
	doc.Find("code").Each(func(i int, s *goquery.Selection) {
		hightlighted, err := syntaxhighlight.AsHTML([]byte(s.Text()))
		if err != nil {
			log.Fatal(err)
		}
		s.SetHtml(string(hightlighted))
	})
	htmlString, err := doc.Html()
	if err != nil {
		log.Fatal(err)
	}
	return template.HTML(htmlString)
}
