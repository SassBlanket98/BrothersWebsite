package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	ListenAddr = ":3000"
)

func homeRoute(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		tmpl *template.Template
	)
	tmpl, err = template.ParseFiles("pages/home.html", "pages/navbar.html", "pages/head.html")
	err = tmpl.Execute(writer, nil)
	if err != nil {
		return
	}
}
func calendarRoute(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		tmpl *template.Template
	)
	tmpl, err = template.ParseFiles("pages/calendar.html", "pages/navbar.html", "pages/head.html")
	err = tmpl.Execute(writer, "meme")
	if err != nil {
		return
	}
}
func pricingRoute(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		tmpl *template.Template
	)
	tmpl, err = template.ParseFiles("pages/pricing.html", "pages/navbar.html", "pages/head.html")
	err = tmpl.Execute(writer, "meme")
	if err != nil {
		return
	}
}
func aboutUsRoute(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		tmpl *template.Template
	)
	tmpl, err = template.ParseFiles("pages/aboutUs.html", "pages/navbar.html", "pages/head.html")
	err = tmpl.Execute(writer, "meme")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/home", homeRoute)
	http.HandleFunc("/calendar", calendarRoute)
	http.HandleFunc("/pricing", pricingRoute)
	http.HandleFunc("/about-us", aboutUsRoute)

	err := http.ListenAndServe(ListenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
