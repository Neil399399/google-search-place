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
	filename  = "CoffeeComment.json"
	index_dir = "coffee.bleve"
)

func main() {

	/*	com, err := Read(filename)
		if err != nil {
			fmt.Println("Read Error!!", err)
		}

			err = CreateIndex(com, index_dir)
			if err != nil {
				fmt.Println("CreateIndex Error!!", err)
			}
	*/
	query, err := jiebatest()
	if err != nil {
		fmt.Println("jieba Error!!", err)
	}
	dataCounter, err := CountResult(index_dir, query)
	if err != nil {
		fmt.Println("CountTesult Error!!", err)
	}
	error := SortTotal(dataCounter)
	if err != nil {
		fmt.Println("Sort Total Error!!", err)
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

func CreateIndex(com []datamodel.Comment, index_dir string) error {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(index_dir, mapping)
	if err != nil {
		fmt.Println(err)
	}

	// index some data
	for i := 0; i < len(com); i++ {
		err = index.Index(com[i].ID, com[i].Comment)
		//fmt.Println(com[i].Comment)
	}

	return nil
}

func jiebatest() ([]string, error) {
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

	return querys, nil
}

func prettify(res *bleve.SearchResult) (string, error) {
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

	return string(b), nil
}

func CountResult(index_dir string, querys []string) (map[string]int, error) {
	type Result struct {
		Id    string
		Score float64
	}

	index, err := bleve.Open(index_dir)
	if err != nil {
		fmt.Println("Open index Error!!", err)
	}
	dataCounter := make(map[string]int)
	for _, q := range querys {
		req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
		req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			panic(err)
		}
		results := []Result{}
		for _, item := range res.Hits {
			results = append(results, Result{item.ID, item.Score})
		}

		for i := 0; i < len(results); i++ {
			dataCounter[results[i].Id]++
		}
	}
	for k, v := range dataCounter {
		fmt.Println("id:", k)
		fmt.Println("total:", v)

	}
	return dataCounter, nil
}

func SortTotal(data map[string]int) error {

	fmt.Println(data)
	fmt.Println(data[2])
	fmt.Println(len(data))

	return nil
}
