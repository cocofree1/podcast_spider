package common

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

func DescriptionRegexp(str string)string{
	str = strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str," ","")
	re, _ := regexp.Compile("<[^>]*>")
	description := re.ReplaceAllString(str, "")
	return description
}

func DateFormToTime(str string)time.Time{
	monthNumberMap := map[string]string{
		"Jan":"01",
		"Feb":"02",
		"Mar":"03",
		"Apr":"04",
		"May":"05",
		"Jun":"06",
		"Jul":"07",
		"Aug":"08",
		"Sep":"09",
		"Oct":"10",
		"Nov":"11",
		"Dec":"12",
	}

	dateArr := strings.Split(str," ")
	day := dateArr[1]
	month := monthNumberMap[dateArr[2]]
	year := dateArr[3]
	hour := dateArr[4]
	date := fmt.Sprintf("%s-%s-%s %s",year,month,day,hour)
	t, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil{
		log.Fatal(err)
	}
	return t
}