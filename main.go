package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
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
	config, err := NewConfig()
	if err != nil {
		log.Println(err)
		return
	}

	var addr = flag.String("addr", ":8080", "The address of the application")
	flag.Parse()

	githubProvider := config.Oauth["github"]

	gomniauth.SetSecurityKey("sgsdfghdtyjurye5rt434535trhggfh")
	gomniauth.WithProviders(
		github.New(githubProvider.ClientID, githubProvider.ClientSecret, githubProvider.RedirectURL),
	)

	r := newRoom()
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on addr: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ServeAndListen: ", err)
	}
}
