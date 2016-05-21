package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
//	"html/template"
)
type Glite struct{
	Html []byte
	Css []byte
	Js []byte
}
func (g *Glite)save() error{
	ioutil.WriteFile("./test/index.html", g.Html, 0600)
	ioutil.WriteFile("./test/style.css", g.Css, 0600)
	return ioutil.WriteFile("./test/script.js", g.Js, 0600)
}
func saveHandler(w http.ResponseWriter, r *http.Request){
	html := r.FormValue("html")
	css := r.FormValue("css")
	js := r.FormValue("js")
	fmt.Println(html)
	fmt.Println(css)
	g := &Glite{Html: []byte(html) , Css: []byte(css), Js: []byte(js)}
	err := g.save()
	if err != nil {
		http.Redirect(w, r, "/",http.StatusFound)
		return	
	}
	http.Redirect(w, r, "/sample.html",http.StatusFound)
}
func main(){
	fs := http.FileServer(http.Dir("."))
	http.Handle("/",fs)
	http.HandleFunc("/save/",saveHandler)
	http.ListenAndServe(":80",nil)
	
}
