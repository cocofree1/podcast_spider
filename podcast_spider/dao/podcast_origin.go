package dao

import (
	"time"
)

type OriginAlbumData struct {
	Id                		 int      `orm:"column(id);auto;pk"`
	WrapperType      	 	 string
	Kind             	 	 string
	ArtistId         		 int
	CollectionId     	 	 int
	TrackId           		 int
	ArtistName       		 string
	CollectionName			 string
	TrackName                string
	CollectionCensoredName   string
	TrackCensoredName        string
	ArtistViewUrl            string
	CollectionViewUrl        string
	FeedUrl                  string
	TrackViewUrl             string
	ArtworkUrl30             string
	ArtworkUrl60             string
	ArtworkUrl100            string
	ReleaseDate              time.Time
	CollectionExplicitness   string
	TrackExplicitness        string
	TrackCount               int
	Country                  string
	Currency                 string
	PrimaryGenreName         string
	ContentAdvisoryRating    string
	ArtworkUrl600            string
	GenreIds                 string
	Genres                   string
}

type OriginProgramData struct {
	Id           int     	     `orm:"column(id);auto;pk"`
	Title        string
	Author       string
	Subtitle     string
	Image        string
	Guid         string
	PubDate      string
	Duration     string
	Link         string
	Description  string            `orm:"type(text)"`
	Order        string
}
