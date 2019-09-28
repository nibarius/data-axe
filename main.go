package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	CodeParameter      = "code"
	NameParameter      = "name"
	TimestampParameter = "t"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/country", countryHandler)
	http.HandleFunc("/language", languageHandler)
	http.HandleFunc("/mcc", mccHandler)
	http.HandleFunc("/http", httpHandler)
	http.HandleFunc("/ts", timestampHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/ascii", asciiHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, getInstructionsDocument())
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	codeNameHandler(w, r, TYPE_LANGUAGE)
}

func countryHandler(w http.ResponseWriter, r *http.Request) {
	codeNameHandler(w, r, TYPE_COUNTRY)
}

func mccHandler(w http.ResponseWriter, r *http.Request) {
	codeNameHandler(w, r, TYPE_MCC)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	codeNameHandler(w, r, TYPE_HTTP_STATUS)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	codeNameHandler(w, r, TYPE_ANDROID_API_LEVEL)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	codeNameHandler(w, r, TYPE_ASCII)
}

func codeNameHandler(w http.ResponseWriter, r *http.Request, theType string) {
	code := sanitize(r.URL.Query()[CodeParameter])
	name := sanitize(r.URL.Query()[NameParameter])
	html, err := doCodeNameLookup(theType, code, name)
	if err != nil {
		_, _ = fmt.Fprint(w, getInstructionsDocument())
	} else {
		_, _ = fmt.Fprint(w, html)
	}
}

func timestampHandler(w http.ResponseWriter, r *http.Request) {
	ts := sanitize(r.URL.Query()[TimestampParameter])
	if ts == nil {
		_, _ = fmt.Fprint(w, getInstructionsDocument())
		return
	}
	html, err := timestampToDate(strings.TrimSpace(ts[0]))

	if err != nil {
		_, _ = fmt.Fprint(w, getInstructionsDocument())
		return
	}

	_, _ = fmt.Fprint(w, html)
}

func sanitize(in []string) []string {
	if in == nil {
		return nil
	}
	out := make([]string, 0, len(in))
	sanitizer := bluemonday.StrictPolicy()
	for _, v := range in {
		out = append(out, sanitizer.Sanitize(v))
	}
	return out
}
