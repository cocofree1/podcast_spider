package dao

import "time"

type PodcastList struct {
	WrapperType      	 	 string     `json:"wrapperType"`
	Kind             	 	 string     `json:"kind"`
	ArtistId         		 int        `json:"artistId"`
	CollectionId     	 	 int        `json:"collectionId"`
	TrackId           		 int		`json:"trackId"`
	ArtistName       		 string     `json:"artistName"`
	CollectionName			 string		`json:"collectionName"`
	TrackName                string		`json:"trackName"`
	CollectionCensoredName   string		`json:"collectionCensoredName"`
	TrackCensoredName        string		`json:"trackCensoredName"`
	ArtistViewUrl            string		`json:"artistViewUrl"`
	CollectionViewUrl        string		`json:"collectionViewUrl"`
	FeedUrl                  string		`json:"feedUrl"`
	TrackViewUrl             string		`json:"trackViewUrl"`
	ArtworkUrl30             string		`json:"artworkUrl30"`
	ArtworkUrl60             string		`json:"artworkUrl60"`
	ArtworkUrl100            string		`json:"artworkUrl100"`
	CollectionPrice          float64	`json:"collectionPrice"`
	TrackPrice               float64	`json:"trackPrice"`
	TrackRentalPrice         float64	`json:"trackRentalPrice"`
	CollectionHdPrice        float64	`json:"collectionHdPrice"`
	TrackHdPrice             float64	`json:"trackHdPrice"`
	TrackHdRentalPrice       float64	`json:"trackHdRental_price"`
	ReleaseDate              time.Time	`json:"releaseDate"`
	CollectionExplicitness   string		`json:"collectionExplicitness"`
	TrackExplicitness        string		`json:"trackExplicitness"`
	TrackCount               int		`json:"trackCount"`
	Country                  string		`json:"country"`
	Currency                 string		`json:"currency"`
	PrimaryGenreName         string		`json:"primaryGenreName"`
	ContentAdvisoryRating    string		`json:"contentAdvisoryRating"`
	ArtworkUrl600            string		`json:"artworkUrl600"`
	GenreIds                 []string	`json:"genreIds"`
	Genres                   []string	`json:"genres"`
}

type PodcastResult struct {
	ResultCount   int            `json:"resultCount"`
	Results       []PodcastList  `json:"results"`
}

type PodcastTask struct {
	PodcastList     PodcastList
	AlbumTag        AlbumTag
}