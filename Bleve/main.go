package main

import (
	"encoding/json"
	"fmt"
	"google-search-place/datamodel"
	"io/ioutil"

	"github.com/blevesearch/bleve"
)

var (
	filename = "CoffeeComment.json"
)

func main() {
	/*
		com, err := Read(filename)
		if err != nil {
			fmt.Println("Read Error!!", err)
		}
		//fmt.Println(com[0].Comment)

		// open a new index
		mapping := bleve.NewIndexMapping()
		index, err := bleve.New("coffee.bleve", mapping)
		if err != nil {
			fmt.Println(err)
		}

		// index some data
		for i := 0; i < len(com); i++ {
			err = index.Index(com[i].ID, com[i].Comment)
			fmt.Println(com[i].Comment)
		}
	*/

	// search for some text
	index, err := bleve.Open("coffee.bleve")
	query := bleve.NewMatchQuery("咖啡")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(searchResults)

}

func Read(filename string) ([]datamodel.Comment, error) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read error: ", err)
	}
	//Create new List and append
	com := []datamodel.Comment{}
	// unmarshal each list

	//unmarshal and Change String to byte
	err = json.Unmarshal(b, &com)
	if err != nil {
		fmt.Println("json err:", err)
	}

	return com, nil
}
