package main

import (
	"net/http"

	"github.com/alwindoss/prism"
)

// Handlers for different pages
func homeHandler(renderer prism.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "home.html", "default", map[string]interface{}{
			"Title": "Home Page",
		})
	}
}

func aboutHandler(renderer prism.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "about.html", "default", map[string]interface{}{
			"Title": "About Us",
		})
	}
}

func contactHandler(renderer prism.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "contact.html", "default", map[string]interface{}{
			"Title": "Contact Us",
		})
	}
}

func main() {
	cfg := prism.Config{
		LayoutPath:   "web/layouts/*.html",
		PagesPath:    "web/pages/*.html",
		PartialsPath: "web/partials/*.html",
	}
	renderer := prism.New(&cfg)
	http.HandleFunc("/", homeHandler(renderer))
	http.HandleFunc("/about", aboutHandler(renderer))
	http.HandleFunc("/contact", contactHandler(renderer))

	// Serve on port 8080
	http.ListenAndServe(":8080", nil)
}
