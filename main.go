package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/otiai10/gosseract"
)

type ocrHolder struct {
	path      string
	ocrResult string
}

type ocrResults []ocrHolder

// in order to be able to sort we need to implement a three functions
// returns how many elements in the collection
func (slice ocrResults) Len() int {
	return len(slice)
}

// which element comes before the other
func (slice ocrResults) Less(i, j int) bool {
	return slice[i].path < slice[j].path
}

// shuffling of the elements
func (slice ocrResults) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// goroutine channel setup for single OCR
// takes an OC holder struct
func convertSingle(res ocrHolder, language string, resultChan chan<- ocrHolder) {
	parameters := gosseract.Params{Src: res.path, Languages: language}
	res.ocrResult = gosseract.Must(parameters)
	resultChan <- res
}

// processes a list of paths with OCR. puts file paths on the OCR convert pipeline
// reads processed items from the results pipeline and adds it to a container struct
func convertMultiple(fileList []string, language string, resultChannel chan ocrHolder) (Results ocrResults) {
	Results = make(ocrResults, len(fileList))

	for _, path := range fileList {
		oh := ocrHolder{path: path}
		go convertSingle(oh, language, resultChannel)
	}

	// consuming all elements in the channel
	for i := 0; i < len(fileList); i++ {
		Results[i] = <-resultChannel
	}

	return
}

func main() {
	resultChannel := make(chan ocrHolder, 1)
	fileList, _ := filepath.Glob("/home/tim/Pictures/schadeformulier2bw.png")
	// fileList = fileList[0:500]
	lang := "eng" // set language

	// convert multiple, writes to return channel
	Results := convertMultiple(fileList, lang, resultChannel)
	sort.Sort(Results)
	fmt.Println(Results)
}
