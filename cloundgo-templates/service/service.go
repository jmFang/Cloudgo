package service

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			fmt.Println(root)
		}
	}

	//中间件的顺序很重要

	//
	mx.HandleFunc("/api/data", apiFetchDataHandler(formatter)).Methods("GET")
	mx.PathPrefix("/api").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if strings.Index(url, "/api") == 0 {
			if strings.Index(url, "/api/data") != 0 {
				fmt.Fprintln(w, "501 Not Implemented ")
			}
		}
	})
	//支持静态文件访问
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	mx.HandleFunc("/result", submit(formatter)).Methods("POST")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}
