package filereader

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

const FileName string = "TestPage"

func (p *Page) Save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	body, err := os.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func FileWriteAndRead() {
	p1 := &Page{Title: FileName, Body: []byte("This is test body")}
	p1.Save()

	p2, _ := LoadPage(FileName + ".txt")
	fmt.Println(string(p2.Body))
}
