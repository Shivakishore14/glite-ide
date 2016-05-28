package main

import (
	"fmt"  
	"io/ioutil"
	"net/http"
	"path/filepath"
	"os"
	"encoding/json"
	"runtime"
	"strings"
)
var home string
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
	if err == nil { 
		return true, nil 
	}
	if os.IsNotExist(err) { 
		return false, nil 
	}
	return false, err
}
func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	fmt.Println(path)
	if err != nil{
		return false,err	
	}
	return fileInfo.IsDir(), err
}

func (g *Glite)saveProject() error{
	path := string(g.Path)
	pe, _ := pathExists(path)
	if pe != true {
		os.MkdirAll(string(filepath.Separator) + path,0777)	
		fmt.Println(path + " created")
	}
	herr := ioutil.WriteFile(path+"index1.html", g.Html, 0777)
	cerr := ioutil.WriteFile(path+"style.css", g.Css, 0777)
	jerr := ioutil.WriteFile(path+"script.js", g.Js, 0777)
	cferr := ioutil.WriteFile(path+"proj.cnf", []byte("configuration"), 0777)
	if herr != nil || cerr != nil || jerr != nil || cferr != nil{
		return CError{Message:"Write error"}
	}
	return nil
}
func importProject(path string) (string, error){
	a,_ := pathExists(path+"proj.cnf")
	if a != true {
		return "error", CError{Message: "not project"}	
	}
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
	l := len(path)
	if path[l-1] != '/' {  //linux
		path = path + "/"
		fmt.Println("added /")
	}
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
	} else {
		fmt.Fprintf(w, s)
	}
}
func ftHandler(w http.ResponseWriter, r *http.Request){
	path := r.FormValue("dir")
	fmt.Println("path = 	"+path)
	head := "<ul class=\"jqueryFileTree\" style=\"display: none;\"> \n"
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if b,_ := pathExists(path+f.Name()); b{
			//fmt.Println(path)
			a,_ := isDir(path+f.Name())
			
			if a {
				c,_ := pathExists(path+f.Name()+"/proj.cnf")
				s := ""
				if c != true {
					s = fmt.Sprintf("<li class=\"directory collapsed\"><a href=\"#\" rel=\"%s/\"> %s </a></li> \n" ,path+f.Name(), f.Name() )
				}else{
					s = fmt.Sprintf("<li class=\"file ext_%s\"><a href=\"#\" rel=\"%s\"> %s </a></li> \n" ,"gproj", path+f.Name(), f.Name() )
				}
				head = head + s

			} else  {
				//gproj ext for the folder
				ext := strings.Split(f.Name(), ".")
				l := len(ext)
				if l > 1 {
					ext[0] = ext[l-1]				
				} else {
					ext[0] = "unknown"
				}
				s := fmt.Sprintf("<li class=\"file ext_%s\"><a href=\"#\" rel=\"%s\"> %s </a></li> \n" ,ext[0], path+f.Name(), f.Name() )
				head = head + s
			}
		}
	}
	head = head + "</ul>"
	fmt.Println(head)
	fmt.Fprintf(w, head)
}
func main(){
	if runtime.GOOS == "windows" {
		home = "C:\\glite"
		fmt.Println("Hello from Windows")
	}
	if runtime.GOOS == "linux" {
		home = "/glite"
		fmt.Println("Hello from linux")
	}
	pe, _ := pathExists(home)
	if pe != true {
		e := os.Mkdir(string(filepath.Separator) + home, 0777)
		if e != nil {
			fmt.Println("error")			
		}	
		fmt.Println("created")
	}
	fs := http.FileServer(http.Dir("."))
	http.Handle("/",fs)
	http.HandleFunc("/save/",saveHandler)
	http.HandleFunc("/import/",importHandler)
	http.HandleFunc("/filetree/",ftHandler)
	http.ListenAndServe(":80",nil)
	
}
