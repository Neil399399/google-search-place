package main

import (
	"fmt"
	"google-search-place/datamodel"
	"google-search-place/storage/sqlstore"
	"log"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var (
	APIKey    = "AIzaSyAFictx33AgxsMkYF-fHCkeakTlBiIZIV4"
	Radius    uint
	filename  = "coffee_comment.json"
	fileStore Storage
	sqlStore  Storage
	cof       datamodel.Coffee
)

func main() {
	//fileStore := fileStore.NewWriteInfile()

	var err error
	Location := &maps.LatLng{Lat: 25.054989, Lng: 121.533359}
	Radius = 50
	Keyword := "coffee"
	Language := "zh-TW"

	err = PlaceSearch(Location, Radius, Keyword, Language)
	if err != nil {
		fmt.Println("google Place Search Error!!", err)
	}

}

func PlaceSearch(location *maps.LatLng, radius uint, keyword string, language string) error {
	sqlStore := sqlstore.NewWriteToSQL("root", "123456", "localhost", "hello")
	c, err := maps.NewClient(maps.WithAPIKey(APIKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	//Store, err := filestore.NewWriteInFile(filename)
	if err != nil {
		fmt.Println("Load File Error!!", err)
	}
	request := &maps.NearbySearchRequest{}
	request.Location = location
	request.Radius = radius
	request.Keyword = keyword
	request.Language = language
	for {
		resp, err := c.NearbySearch(context.Background(), request)

		if resp.NextPageToken != "" {
			request.PageToken = resp.NextPageToken
		}

		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}

		//	text := []string{}
		for i := 0; i < len(resp.Results); i++ {
			id := resp.Results[i].PlaceID
			Name := resp.Results[i].Name
			Rate := resp.Results[i].Rating

			cof.Id = id
			cof.Name = Name
			cof.Rate = Rate
			cof.Reviews = []datamodel.Review{}

			req := &maps.PlaceDetailsRequest{}
			req.PlaceID = id
			req.Language = language

			respd, err := c.PlaceDetails(context.Background(), req)
			if err != nil {
				log.Fatalf("fatal error: %s", err)
			}

			for j := 0; j < len(respd.Reviews); j++ {
				review := datamodel.Review{cof.Id, respd.Reviews[j].Text}
				cof.Reviews = append(cof.Reviews, review)
			}
			err = sqlStore.Write(cof)
			if err != nil {
				fmt.Println("Write In Sql Error!!", err)
			}

		}
	}
	return nil
}
