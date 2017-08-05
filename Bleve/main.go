package main

import (
	"encoding/json"
	"fmt"
	"google-search-place/datamodel"
	"io/ioutil"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/yanyiwu/gojieba"
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
	if err != nil {
		fmt.Println("Open index Error!!", err)
	}
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

func jieba(INDEX_DIR stringg) error {
	indexMapping := bleve.NewIndexMapping()
	os.RemoveAll(INDEX_DIR)
	// clean index when example finished
	defer os.RemoveAll(INDEX_DIR)

	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"type":         "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		panic(err)
	}
	indexMapping.DefaultAnalyzer = "gojieba"

	index, err := bleve.New(INDEX_DIR, indexMapping)
	if err != nil {
		panic(err)
	}
	for _, msg := range messages {
		if err := index.Index(msg.Id, msg); err != nil {
			panic(err)
		}
	}

	querys := []string{
		"咖啡",
	}

	for _, q := range querys {
		req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
		req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}
