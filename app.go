package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

const base = "http://www.marlex.org"

var (
	reIsBlogpost    = regexp.MustCompile(`^\/\d+\-(\w|\-)+\/?$`)
	reIsNewBlogpost = regexp.MustCompile(`^\/(\w|\-)+\/\d+\/?$`)
	reIsKeyword     = regexp.MustCompile(`^\/(clave|tags)\/(\w|\-)+\/?$`)
	reIsSearch      = regexp.MustCompile(`^\/q\/(\w|\-|\+)+\/?$`)
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	switch {
	case reIsBlogpost.MatchString(p):
		redirToBlogpost(w, strings.Trim(p, "/"), false)

	case reIsNewBlogpost.MatchString(p):
		redirToBlogpost(w, strings.Trim(p, "/"), true)

	case reIsKeyword.MatchString(p):
		redirToTag(w, strings.Trim(p, "/"))

	case reIsSearch.MatchString(p):
		redirToSearch(w, strings.Trim(p, "/"))

	default:
		permanent(w, fmt.Sprintf("%s/", base))
	}
}

// this will contain the redirection on path, to parse it
// use the end boolean. If true, the ID is at the end
// separated by "/:path/:id". If false, the ID is
// at the beginning, separated by "/:id-:path/"
func redirToBlogpost(w http.ResponseWriter, path string, end bool) {
	var id, name string

	// if it's at the end, reverse the name and id, then
	// find where the slash location is, otherwise just
	// find the initial slash and save it
	if end {
		pos := strings.Index(path, "/")
		name, id = path[:pos], path[pos+1:]
	} else {
		pos := strings.Index(path, "-")
		id, name = path[:pos], path[pos+1:]
	}

	// construct the URL
	permanent(w, fmt.Sprintf("%s/%s/%s/", base, name, id))
}

// this function will allow to redirect to the following
// locations: "/clave/:word" and "/tags/:word"
func redirToTag(w http.ResponseWriter, path string) {
	// remove the prefixes that are common from these
	// old paths in the blog
	path = strings.TrimPrefix(path, "clave/")
	path = strings.TrimSuffix(path, "tags/")

	// construct the URL
	permanent(w, fmt.Sprintf("%s/tema/%s/", base, path))
}

// handle search endpoints with the format "/q/:word", where
// the most important thing is that words can be either with
// words separated by plus (+) or (-)
func redirToSearch(w http.ResponseWriter, path string) {
	// remove the "q/" part initially set on old searches
	// also remove all the plusses in the querystring
	// and remove any dash too
	path = strings.TrimPrefix(path, "q/")
	path = strings.Replace(path, "+", " ", -1)

	// construct the URL
	temporary(w, fmt.Sprintf("https://www.google.com/search?q=site:www.marlex.org %s", path))
}

// create an http redirection
func permanent(w http.ResponseWriter, location string) {
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusMovedPermanently)
}

// create an http redirection
func temporary(w http.ResponseWriter, location string) {
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusFound)
}
