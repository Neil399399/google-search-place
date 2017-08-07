package main

import (
	"encoding/json"
	"fmt"
	"google-search-place/datamodel"
	"io/ioutil"
	"log"

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

/*	err := jiebatest()
	if err != nil {
		fmt.Println("jieba Error!!", err)
*/	}
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

func jieba(INDEX_DIR string) error {
	indexMapping := bleve.NewIndexMapping()

	err := indexMapping.AddCustomTokenizer("jieba",
		map[string]interface{}{
			"file": "jieba/dict.txt",
			"type": "咖啡",
		})
	if err != nil {
		log.Fatal(err)
	}

	err = indexMapping.AddCustomAnalyzer("jieba",
		map[string]interface{}{
			"type":      "custom",
			"tokenizer": "咖啡",
			"token_filters": []string{
				"possessive_en",
				"to_lower",
				"stop_en",
			},
		})

	if err != nil {
		log.Fatal(err)
	}

	indexMapping.DefaultAnalyzer = "jieba"

	index, err := bleve.Open(INDEX_DIR)
	if err != nil {
		fmt.Println("Open index Error!!", err)
	}
	for _, keyword := range []string{"咖啡"} {
		query := bleve.NewMatchQuery(keyword)
		search := bleve.NewSearchRequest(query)
		search.Highlight = bleve.NewHighlight()
		searchResults, err := index.Search(search)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Result of %s: %s\n", keyword, searchResults)
	}
	return nil
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
	err = indexMapping.AddCustomAnalyzer("gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
indexMapping.DefaultType()

	if err != nil {
		fmt.Println("Analyzer Error!!", err)
	}
	indexMapping.DefaultAnalyzer = "gojieba"

	querys := []string{
		"舒服",
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
	return string(b)
}
