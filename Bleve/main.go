package main

import (
	"encoding/json"
	"fmt"
	"google-search-place/datamodel"
	"io/ioutil"

	"github.com/blevesearch/bleve"
	"github.com/yanyiwu/gojieba"
)

type JiebaTokenizer struct {
	handle *gojieba.Jieba
}

var (
	filename = "CoffeeComment.json"
)

func main() {

	/*		com, err := Read(filename)
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
	/*
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
	*/

	err := jiebatest()
	if err != nil {
		fmt.Println("jieba Error!!", err)
	}
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

func jiebatest() error {
	indexMapping := bleve.NewIndexMapping()
	err := indexMapping.AddCustomTokenizer("gojieba",
		map[string]interface{}{
			"dictpath":   "jieba/dict.txt",
			"hmmpath":    "jieba/hmm_model.utf8",
			"idf":        "idf.utf8",
			"stop_words": "stop_word.utf8",
			"type":       "unicode",
		},
	)
	if err != nil {
		fmt.Println("Tokenizer Error!!", err)
	}

	indexMapping.DefaultAnalyzer = "gojieba"

	querys := []string{
		"環境舒服",
		"不錯",
		"咖啡好喝",
		"好喝",
		"好",
	}

	index, err := bleve.Open("coffee.bleve")
	if err != nil {
		fmt.Println("Open index Error!!", err)
	}

	for _, q := range querys {
		req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
		req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			panic(err)
		}
		fmt.Println(prettify(res))
	}

	return nil
}

func prettify(res *bleve.SearchResult) string {
	type Result struct {
		Id    string
		Score float64
	}
	results := []Result{}
	for _, item := range res.Hits {
		results = append(results, Result{item.ID, item.Score})
	}

	b, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	//counter
	type conditional struct {
		cond []int
	}
	dataCounter := make(map[string]conditional)
	for i := 0; i < len(results); i++ {
		for j := 0; i < len(results[i].Id); i++ {
			dataCount[results[i].id[j]].cond[i]++
		}
	}

	for k, v := range dataCounter {
		total := 0
		for i := 0; i < len(v.cond); i++ {
			total += v.cond[i]
		}
	}

	return nil
}
