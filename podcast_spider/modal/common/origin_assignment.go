package common

import (
	"podcast_spider/dao"
	"strings"
)

// todo
func GetOriginAlbumDataByPodcastList(list dao.PodcastList) dao.OriginAlbumData {
	podcastList := dao.OriginAlbumData{
		WrapperType:            list.WrapperType,
		Kind:                   list.Kind,
		ArtistId:               list.ArtistId,
		CollectionId:           list.CollectionId,
		TrackId:                list.TrackId,
		ArtistName:       	    list.ArtistName,
		CollectionName:		    list.CollectionName,
		TrackName:              list.TrackName,
		CollectionCensoredName: list.CollectionCensoredName,
		TrackCensoredName:      list.TrackCensoredName,
		ArtistViewUrl:          list.ArtistViewUrl,
		CollectionViewUrl:      list.CollectionViewUrl,
		FeedUrl:                list.FeedUrl,
		TrackViewUrl:           list.TrackViewUrl,
		ArtworkUrl30:           list.ArtworkUrl30,
		ArtworkUrl60:           list.ArtworkUrl60,
		ArtworkUrl100:          list.ArtworkUrl100,
		ReleaseDate:            list.ReleaseDate,
		CollectionExplicitness: list.CollectionExplicitness,
		TrackExplicitness:      list.TrackExplicitness,
		TrackCount:             list.TrackCount,
		Country:                list.Country,
		Currency:               list.Currency,
		PrimaryGenreName:       list.PrimaryGenreName,
		ContentAdvisoryRating:  list.ContentAdvisoryRating,
		ArtworkUrl600:          list.ArtworkUrl600,
		GenreIds:               strings.Join(list.GenreIds,","),
		Genres:                 strings.Join(list.Genres,","),
	}
	return podcastList
}

func GetOriginProgramDataByXmlItem(content dao.XmlItem)dao.OriginProgramData {
	description := DescriptionRegexp(content.Description)
	podcastContent := dao.OriginProgramData{
		Title: content.Title,
		Author: content.Author,
		Subtitle: content.Subtitle,
		Image: content.Image.Href,
		Guid: content.Guid,
		PubDate: content.PubDate,
		Duration: content.Duration,
		Link: content.Link,
		Description: description,
		Order: content.Order,
	}
	return podcastContent
}
