package main

import (
	"fmt"  
	"io/ioutil"
	"net/http"
	"path/filepath"
	"os"
	"encoding/json"
//	"html/template"
)
type Glite struct{
	Html []byte
	Css []byte
	Js []byte
	Path []byte
}
type CError struct {
	Message string
}
func (e CError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}
func (g *Glite)saveProject() error{
	path := string(g.Path)
	pe, _ := pathExists(path)
	if pe != true {
		os.MkdirAll(string(filepath.Separator) + path,0777)	
		fmt.Println("created")
	}
	herr := ioutil.WriteFile(path+"index1.html", g.Html, 0777)
	cerr := ioutil.WriteFile(path+"style.css", g.Css, 0777)
	jerr := ioutil.WriteFile(path+"script.js", g.Js, 0777)
	if herr != nil || cerr != nil || jerr != nil{
		return CError{Message:"Write error"}
	}
	return nil
}
func importProject(path string) (string, error){
	html, herr := ioutil.ReadFile(path+"index1.html")
	css, cerr := ioutil.ReadFile(path+"style.css")
	js, jerr := ioutil.ReadFile(path+"script.js")
	if herr != nil || cerr != nil || jerr != nil {
		return "error",CError{Message:"Reading Error"}
	}
	g := &Glite{Html:[]byte(html), Css:[]byte(css), Js:[]byte(js), Path:[]byte(path)}
	gjson, err := json.Marshal(g)
	if err != nil {
		return "error",CError{Message:"json creation Error"}	
	}
	return string(gjson), nil	
}
func saveHandler(w http.ResponseWriter, r *http.Request){
	html := r.FormValue("html")
	css := r.FormValue("css")
	js := r.FormValue("js")
	path := r.FormValue("path")
	fmt.Println(html)
	g := &Glite{Html: []byte(html) , Css: []byte(css), Js: []byte(js), Path:[]byte(path)}
	err := g.saveProject()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return	
	}
	fmt.Fprintf(w, "saved")
}
func importHandler(w http.ResponseWriter, r *http.Request){
	path := r.FormValue("path")
	//fmt.Fprintf(w, path)
	s,err := importProject(path)
	if err != nil {
		fmt.Fprintf(w, err.Error())	
	}
	fmt.Fprintf(w, s)
}
func main(){
	fs := http.FileServer(http.Dir("."))
	http.Handle("/",fs)
	http.HandleFunc("/save/",saveHandler)
	http.HandleFunc("/import/",importHandler)
	http.ListenAndServe(":80",nil)
	
}
