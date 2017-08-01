package main

import (
	"fmt"
	"google_place/datamodel"
	"google_place/storage"
	"log"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var (
	APIKey   = "AIzaSyAFictx33AgxsMkYF-fHCkeakTlBiIZIV4"
	Radius   uint
	cof      datamodel.Coffee
	filename = "coffee_comment.json"
	Store    Storage
)

func main() {
	var err error

	Location := &maps.LatLng{Lat: 25.054989, Lng: 121.533359}
	Radius = 500
	Keyword := "coffee"
	Language := "zh-TW"

	err = PlaceSearch(Location, Radius, Keyword, Language)
	if err != nil {
		fmt.Println("google Place Search Error!!", err)
	}

}

func PlaceSearch(location *maps.LatLng, radius uint, keyword string, language string) error {

	c, err := maps.NewClient(maps.WithAPIKey(APIKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	Store, err := filestore.NewWriteInFile(filename)
	if err != nil {
		fmt.Println("Load File Error!!", err)
	}
	request := &maps.NearbySearchRequest{}
	request.Location = location
	request.Radius = radius
	request.Keyword = keyword
	request.Language = language

	resp, err := c.NearbySearch(context.Background(), request)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	text := []string{}
	for i := 0; i < len(resp.Results); i++ {
		id := resp.Results[i].PlaceID
		Name := resp.Results[i].Name
		Rate := resp.Results[i].Rating
		cof.Name = Name
		cof.Rate = Rate

		req := &maps.PlaceDetailsRequest{}
		req.PlaceID = id
		req.Language = language
		respd, err := c.PlaceDetails(context.Background(), req)
		if err != nil {
			log.Fatalf("fatal error: %s", err)
		}

		for j := 0; j < len(respd.Reviews); j++ {
			Text := respd.Reviews[j].Text
			text = append(text, Text)
		}
		cof.TEXT = text

	}
	Store.Write(cof)

	return nil
}
