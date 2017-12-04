package dataaxe

import (
	"fmt"
	"net/http"
	"github.com/microcosm-cc/bluemonday"
	"strings"
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
	http.HandleFunc("/api", apiHandler)
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

func apiHandler(w http.ResponseWriter, r *http.Request) {
	codeNameHandler(w, r, TYPE_ANDROID_API_LEVEL)
}

func codeNameHandler(w http.ResponseWriter, r *http.Request, theType string) {
	code := sanitize(r.URL.Query()[CODE_PARAMETER])
	name := sanitize(r.URL.Query()[NAME_PARAMETER])
	html, err := doCodeNameLookup(theType, code, name)
	if err != nil {
		fmt.Fprint(w, getInstructionsDocument())
	} else {
		fmt.Fprint(w, html)
	}
}

func timestampHandler(w http.ResponseWriter, r *http.Request) {
	ts := sanitize(r.URL.Query()[TIMESTAMP_PARAMETER])
	if ts == nil {
		fmt.Fprint(w, getInstructionsDocument())
		return
	}
	html, err := timestampToDate(strings.TrimSpace(ts[0]))

	if err != nil {
		fmt.Fprint(w, getInstructionsDocument())
		return
	}

	fmt.Fprint(w, html)
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
