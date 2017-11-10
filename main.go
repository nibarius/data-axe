package dataaxe

import (
	"fmt"
	"net/http"
)

const (
	CODE_PARAMETER      = "code"
	NAME_PARAMETER      = "name"
	TIMESTAMP_PARAMETER = "t"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/country", countryHandler)
	http.HandleFunc("/language", languageHandler)
	http.HandleFunc("/mcc", mccHandler)
	http.HandleFunc("/http", httpHandler)
	http.HandleFunc("/ts", timestampHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getInstructionsDocument())
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

func codeNameHandler(w http.ResponseWriter, r *http.Request, theType string) {
	code := r.URL.Query()[CODE_PARAMETER]
	name := r.URL.Query()[NAME_PARAMETER]
	html, err := doCodeNameLookup(theType, code, name)
	if err != nil {
		fmt.Fprint(w, getInstructionsDocument())
	} else {
		fmt.Fprint(w, html)
	}
}

func timestampHandler(w http.ResponseWriter, r *http.Request) {
	ts := r.URL.Query()[TIMESTAMP_PARAMETER]
	if ts == nil {
		fmt.Fprint(w, getInstructionsDocument())
		return
	}
	html, err := timestampToDate(ts[0])

	if err != nil {
		fmt.Fprint(w, getInstructionsDocument())
		return
	}

	fmt.Fprint(w, html)
}