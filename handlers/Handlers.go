package handlers

import (
	"Practice/Exercise1/parser"
	"Practice/Exercise1/rendoring"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	strip "github.com/grokify/html-strip-tags-go"
)

var viewTemplate = filepath.Join("view")

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	rendoring.RenderTemplate(w, viewTemplate, nil)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("urlToCrawl")
	if len(url) == 0 {
		rendoring.RenderTemplate(w, viewTemplate, &rendoring.Page{Error: "Please enter a valid URl."})
	}
	resp, err := http.Get(url)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	if err != nil {
		rendoring.RenderTemplate(w, viewTemplate, &rendoring.Page{Error: "No data found for the entered URL. Try other URL."})
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rendoring.RenderTemplate(w, viewTemplate, &rendoring.Page{Error: "No data found for the entered URL. Try other URL."})
	}
	contentWithoutHTMLTags := strip.StripTags(string(content))
	wordsWithCount := parser.WordCountFromText(contentWithoutHTMLTags)

	rendoring.RenderTemplate(w, viewTemplate, &rendoring.Page{Words: wordsWithCount, CrawlNotification:fmt.Sprintf("URL to be crwaled: %s", url)})
}
