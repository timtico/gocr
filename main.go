package main

import "github.com/otiai10/gosseract"

type ocrHolder struct {
	path      string
	ocrResult string
}

// goroutine channel setup for single OCR
func convertSingle(res ocrHolder, language string, resultChan chan<- ocrHolder) {
	parameters := gosseract.Params{Src: res.path, Languages: language}
	res.ocrResult = gosseract.Must(parameters)
	resultChan <- res
}

func main() {
	resultChannel := make(chan ocrHolder, 1)
	fileList := []string{"testimage1.png", "testimage2.png"}
	// client, _ := gosseract.NewClient()
	lang := "eng"
	Results := make([]ocrHolder, len(fileList))

	for _, path := range fileList {
		oh := ocrHolder{path: path}
		go convertSingle(oh, lang, resultChannel)
	}

	for i := 0; i < len(fileList); i++ {
		elem := <-resultChannel
		Results[i] = elem
	}
	//fmt.Println(Results)
}
