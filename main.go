package main

import (
	"encoding/json"
	"html/template"
	"time"

	"io/ioutil"
	"log"
	"net/http"
)

var t *template.Template
var c config

type data struct {
	Weekday  time.Weekday
	Greeting string
}

type config struct {
	Greeting string
}

func main() {
	loadConfig()
	initializeTemplates()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func loadConfig() {
	contentBytes, _ := ioutil.ReadFile("config.json")
	json.Unmarshal(contentBytes, &c)
}

func initializeTemplates() {
	t, _ = template.ParseFiles("index.html")
}
func handler(w http.ResponseWriter, r *http.Request) {
	data := &data{time.Now().Weekday(), c.Greeting}

	t.Execute(w, data)
}
