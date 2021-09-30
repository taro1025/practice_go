package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"errors"
)

type Page struct {
	Title string
	Body []byte
}
//template.Must validates
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
			http.NotFound(w, r)
			return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) //Third parameter is permission
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
			return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//view & controller
func viewHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/view/"):] no validation
	title, err := getTitle(w, r)       // validated
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//t, _ := template.ParseFiles("view.html")    |move to render func
	//t.Execute(w, p) // p is used in view.html   |
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
			return
	}
	p, err := loadPage(title)
	if err != nil {
			p = &Page{Title: title}
	}
	//t, _ := template.ParseFiles("edit.html")   |move to render func
	//t.Execute(w, p) // p is used in edit.html  |
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/save/"):] no validation
	title, err := getTitle(w, r)        // validated
	if err != nil {
			return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	//t, err := template.ParseFiles(tmpl + ".html")     //Parse template
	err := templates.ExecuteTemplate(w, tmpl+".html", p) //Execute Tmplate already parsed
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
}

func main() {
	//route
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	//litsten 8000port
	log.Fatal(http.ListenAndServe(":8080", nil))
}