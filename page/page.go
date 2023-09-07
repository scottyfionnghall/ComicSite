package page

import (
	"html/template"
	"log"
	"os"
	"slices"

	"github.com/gomarkdown/markdown"
)

type Page struct {
	Title    string
	Body     template.HTML
	Next     template.HTML
	Previous template.HTML
}

func checkIndex(index int, pages []string) bool {
	if index >= 0 && index < len(pages) {
		return true
	} else {
		return false
	}
}

func getAllPages() []string {
	files, err := os.ReadDir("data/")
	if err != nil {
		log.Fatal(err)
	}
	var return_files []string
	for _, file := range files {
		return_files = append(return_files, file.Name())
	}
	return return_files
}

// func (p *Page) Save() error {
// 	filename := "data/" + p.Title + ".md"
// 	return os.WriteFile(filename, p.Body, 0600)
// }

func getPrevOrNext(index int, pages []string) template.HTML {
	if checkIndex(index, pages) {
		temp := pages[index]
		temp = temp[:len(temp)-3]
		return template.HTML(temp)
	} else {
		return template.HTML("HOME")
	}
}

func LoadPage(title string) (*Page, error) {
	filename := "data/" + title + ".md"
	context, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	context = markdown.ToHTML(context, nil, nil)
	body := template.HTML(context)
	all_flies := getAllPages()
	next_index := slices.Index(all_flies, title+".md") + 1
	next := getPrevOrNext(next_index, all_flies)
	previous_index := slices.Index(all_flies, title+".md") - 1
	previous := getPrevOrNext(previous_index, all_flies)
	return &Page{Title: title, Body: body, Next: next, Previous: previous}, nil
}
