package dao

import (
	"encoding/xml"
)

type XmlItem struct{
	Item        xml.Name      `xml:"item"`
	Title       string        `xml:"title"`
	Author      string        `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd author"`
	Subtitle    string        `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd subtitle"`
	Image       ImageAttr     `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
	Enclosure   EnclosureAttr `xml:"enclosure"`
	Guid        string        `xml:"guid"`
	PubDate     string        `xml:"pubDate"`
	Duration    string        `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd duration"`
	Link        string        `xml:"link"`
	Description string        `xml:"description"`
	Order       string        `xml:"order"`
}

type XmlItemCont struct {
	XmlItem          XmlItem
	CollectionId     int
}

type ImageAttr struct {
	Href       string     `xml:"href,attr"`
}

type EnclosureAttr struct {
	Url       string     `xml:"url,attr"`
	Type      string     `xml:"type,attr"`
	Length    string     `xml:"length,attr"`
}

type XmlChannel struct {
	Channel      xml.Name   `xml:"channel"`
	Copyright    string     `xml:"copyright"`
	Language     string     `xml:"language"`
	Link         string     `xml:"link"`
	Title        string     `xml:"title"`
	Author       string     `xml:"author"`
	Subtitle     string     `xml:"subtitle"`
	Summary      string     `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd summary"`
	Owner      	 string     `xml:"owner"`
	Description  string     `xml:"description"`
	Image        string     `xml:"image"`
	Explicit     string     `xml:"explicit"`
	Item         []XmlItem  `xml:"item"`
}

type XmlRss struct {
	Rss      xml.Name     `xml:"rss"`
	Channel  []XmlChannel `xml:"channel"`
}
