// http://www.w3schools.com/rss/default.asp
// http://www.tutorialspoint.com/rss/what-is-atom.htm
// http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter
package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"os"

	//	"code.google.com/p/go.text/encoding/charmap"
	//	"code.google.com/p/go.text/transform"
	//		"golang.org/x/text"
	"golang.org/x/net/html/charset"
)

// http://www.w3schools.com/rss/rss_syntax.asp
// http://www.w3schools.com/rss/rss_channel.asp
type Rss2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	// Required
	Title       string `xml:"channel>title"`
	Link        string `xml:"channel>link"`
	Description string `xml:"channel>description"`
	// Optional
	PubDate  string `xml:"channel>pubDate"`
	ItemList []Item `xml:"channel>item"`
}

// http://www.w3schools.com/rss/rss_item.asp
// http://stackoverflow.com/questions/7220670/difference-between-description-and-contentencoded-tags-in-rss2
// https://groups.google.com/d/topic/golang-nuts/uBMo1BpaQCM
type Item struct {
	// Required
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Description template.HTML `xml:"description"`
	// Optional
	Content  template.HTML `xml:"encoded"`
	PubDate  string        `xml:"pubDate"`
	Comments string        `xml:"comments"`
}

func main() {
	r := Rss2{}
	/*
		xmlContent, _ := ioutil.ReadFile("example-6.xml")

		sr := strings.NewReader(xmlContent)
		tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())

		buf, err := ioutil.ReadAll(tr)
	*/

	response, err := http.Get("http://minjust.ru/ru/unwanted/rss")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		decoder := xml.NewDecoder(response.Body)
		decoder.CharsetReader = charset.NewReaderLabel
		err = decoder.Decode(&r)

		//	err := xml.Unmarshal(buf, &r)
		if err != nil {
			panic(err)
		}
		for _, item := range r.ItemList {
			fmt.Println(item)
		}
	}
}
