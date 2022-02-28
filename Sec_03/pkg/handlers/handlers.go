package handlers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page")
	render.renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	// sum, _ := addValues(2, 2)
	// // fmt.Fprintf(w, "This is the about page")
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("About page. 2 + 2 =%d", sum))
	render.renderTemplate(w, "about.page.html")
}
