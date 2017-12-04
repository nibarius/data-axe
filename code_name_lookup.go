package dataaxe

import (
	"strings"
	"errors"
)

type codeNamePair struct {
	code, name string
}

const (
	TYPE_LANGUAGE          = "Language"
	TYPE_COUNTRY           = "Country"
	TYPE_MCC               = "MCC"
	TYPE_HTTP_STATUS       = "HTTP"
	TYPE_ANDROID_API_LEVEL = "Android API level"
	LOOKUP_BY_CODE         = iota
	LOOKUP_BY_NAME
)

func doCodeNameLookup(theType string, code []string, name []string) (string, error) {
	if code == nil && name == nil {
		return getInstructionsDocument(), errors.New("nothing to search for")
	}

	cards := make([]cardParameters, 0, len(code)+len(name))
	cards = addCards(theType, LOOKUP_BY_CODE, code, cards)
	cards = addCards(theType, LOOKUP_BY_NAME, name, cards)

	pageTitle := getPageTitle(theType)
	return getHtmlDocument(pageTitle, cards)
}

func getPageTitle(theType string) (title string) {
	switch theType {
	case TYPE_HTTP_STATUS:
		title = "HTTP status codes"
	case TYPE_ANDROID_API_LEVEL:
		title = "Android API levels"
	default:
		title = theType + " codes"
	}
	return
}

func addCards(theType string, method int, from []string, to []cardParameters) []cardParameters {
	var haystack []codeNamePair
	var search func([]codeNamePair, string) []codeNamePair
	codeTitle := theType + " code"
	nameTitle := theType + " name"

	switch method {
	case LOOKUP_BY_NAME:
		search = getPairByName
	case LOOKUP_BY_CODE:
		search = getPairByCode
	}

	switch theType {
	case TYPE_LANGUAGE:
		haystack = languages
	case TYPE_COUNTRY:
		haystack = countries
	case TYPE_MCC:
		haystack = mccs
		codeTitle = theType
		nameTitle = "Country"
	case TYPE_HTTP_STATUS:
		haystack = httpStatusCodes
		codeTitle = "HTTP status code"
		nameTitle = "Name"
	case TYPE_ANDROID_API_LEVEL:
		haystack = androidApiLevels
		codeTitle = theType
		nameTitle = "Android version"
	}

	for _, value := range from {
		var values []string
		switch method {
		case LOOKUP_BY_NAME:
			values = []string{value}
		case LOOKUP_BY_CODE:
			values = strings.Split(value, ",")
		}
		for _, expandedValue := range values {
			expandedValue = strings.TrimSpace(expandedValue)
			matches := search(haystack, expandedValue)
			for _, codeName := range matches {
				var body1 string
				switch theType {
				case TYPE_HTTP_STATUS:
					body1 = "<a href=\"https://httpstatuses.com/" + codeName.code + "\">" + codeName.code + "</a>"
				default:
					body1 = codeName.code

				}
				params := cardParameters{
					codeTitle,
					body1,
					nameTitle,
					codeName.name,
				}
				to = append(to, params)
			}
		}
	}
	return to
}

func getPairByCode(haystack []codeNamePair, languageCode string) []codeNamePair {
	for _, pair := range haystack {
		if strings.ToLower(pair.code) == strings.ToLower(languageCode) {
			return []codeNamePair{pair}
		}
	}
	return []codeNamePair{{languageCode, "Unknown code"}}
}

func getPairByName(haystack []codeNamePair, languageName string) []codeNamePair {
	ret := []codeNamePair{}
	for _, pair := range haystack {
		if strings.Contains(strings.ToLower(pair.name), strings.ToLower(languageName)) {
			ret = append(ret, pair)
		}
	}
	if len(ret) == 0 {
		ret = append(ret, codeNamePair{"Unknown name", languageName})
	}
	return ret
}
