package main

import (
	"fmt"

	"google_place/datamodel"
	"google_place/storage"

	"github.com/wangbin/jiebago"
)

var (
	seg      jiebago.Segmenter
	filename = "d://golang/src/google_place/coffee_comment.json"
	coff     datamodel.Coffee
	a        []string
	b        string
)

func init() {
	seg.LoadDictionary("d://golang/src/google_place/jieba/dict.txt")

}

func print(ch <-chan string) {
	for word := range ch {
		fmt.Printf(" %s /", word)
	}
	fmt.Println()
}

func main() {
	fmt.Print("全模式：")
	var err error
	b = "花"
	Store, err := filestore.NewWriteInFile(filename)
	if err != nil {
		fmt.Println("Load File Error!!", err)
	}
	coffees, err := Store.Read()
	if err != nil {
		fmt.Println("File Read Error!!", err)
	}
	//fmt.Println(coffees[0].TEXT)
	//fmt.Println(coffees)
	//for i := 0; i < len(coffees[0].TEXT); i++ {

	print(seg.Cut(coffees[0].TEXT[0], true))
	for str := range seg.CutAll(coffees[0].TEXT[0]) {
		seg.AddWord(str, 0)
		fmt.Println(seg.Frequency(b))
	}
	jiebago.Segmenter.CutForSearch

	//}

}
