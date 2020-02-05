package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"flag"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The address of the application")
	flag.Parse()

	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on addr: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ServeAndListen: ", err)
	}
}
