package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/filereader"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := filereader.LoadPage(title)
	if err != nil {
		p = &filereader.Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\"/>"+
		"</form>", p.Title, p.Title, p.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := filereader.LoadPage("TestPage.txt")
	fmt.Fprintf(w, "<head><title>%s</title></head><h1>%s</h1>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
