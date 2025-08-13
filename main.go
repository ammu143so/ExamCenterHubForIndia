package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"sort"

	handlerpkg "exam-center-assignment/internal/handler"
)

//go:embed templates/*.html static/*
var content embed.FS

type Server struct {
	h *handlerpkg.ExamCenterHandler
	t *template.Template
}

func newServer() *Server {
	// Parse templates from embedded FS
	tmpl := template.Must(template.ParseFS(content, "templates/*.html"))
	return &Server{
		h: handlerpkg.NewExamCenterHandler(),
		t: tmpl,
	}
}

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	// Static files
	staticFS, _ := fs.Sub(content, "static")
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	mux.HandleFunc("/", s.handleHome)
	mux.HandleFunc("/search", s.handleSearch)
	return mux
}

type HomePageData struct {
	Title       string
	Cities      []string
	Error       string
	SelectedCity string
}

type ResultCity struct {
	Name     string
	Distance string
	Centers  []string
}

type ResultsPageData struct {
	Title     string
	HomeCity  string
	Results   []ResultCity
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	cities := s.h.GetAvailableCities()
	sort.Strings(cities)
	data := HomePageData{Title: "ExamCenterHub — Find Nearest Exam Centers", Cities: cities}
	if q := r.URL.Query().Get("error"); q != "" {
		data.Error = q
	}
	_ = s.t.ExecuteTemplate(w, "index.html", data)
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/?error="+urlQueryEscape("Invalid form submission"), http.StatusSeeOther)
		return
	}
	homeCityInput := r.FormValue("home_city")
	homeCity, err := s.h.ValidateCity(homeCityInput)
	if err != nil {
		http.Redirect(w, r, "/?error="+urlQueryEscape(err.Error()), http.StatusSeeOther)
		return
	}
	nearest, err := s.h.FindNearestCities(homeCity, 3)
	if err != nil {
		http.Redirect(w, r, "/?error="+urlQueryEscape(err.Error()), http.StatusSeeOther)
		return
	}
	var results []ResultCity
	for _, cd := range nearest {
		var centers []string
		for _, c := range cd.Centers {
			centers = append(centers, c.Name)
		}
		results = append(results, ResultCity{
			Name:     cd.City.Name,
			Distance: fmt.Sprintf("%.1f km", cd.Distance),
			Centers:  centers,
		})
	}
	data := ResultsPageData{Title: "Results — ExamCenterHub", HomeCity: homeCity, Results: results}
	_ = s.t.ExecuteTemplate(w, "results.html", data)
}

func urlQueryEscape(s string) string {
	// Simple replacement to avoid importing net/url just for this
	replacer := map[string]string{
		" ": "+",
		"\n": "",
	}
	out := s
	for k, v := range replacer {
		out = stringReplaceAll(out, k, v)
	}
	return out
}

func stringReplaceAll(s, old, new string) string {
	for {
		idx := indexOf(s, old)
		if idx < 0 { return s }
		s = s[:idx] + new + s[idx+len(old):]
	}
}

func indexOf(s, sub string) int {
	// naive search
	ls, lsub := len(s), len(sub)
	if lsub == 0 { return 0 }
	for i := 0; i+ lsub <= ls; i++ {
		if s[i:i+lsub] == sub { return i }
	}
	return -1
}

func main() {
	srv := newServer()
	addr := ":8080"
	log.Printf("ExamCenterHub web UI listening on %s", addr)
	if err := http.ListenAndServe(addr, srv.routes()); err != nil {
		log.Fatal(err)
	}
} 