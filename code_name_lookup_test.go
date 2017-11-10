package dataaxe

import (
	"testing"
)

func TestOneCountryLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Country codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Country code</p>
    <p>AX</p>
    <p class="title">Country name</p>
    <p>Åland Islands</p>
  </div>
</body>
</html>
`

	codeLookup, _ := doCodeNameLookup(TYPE_COUNTRY, []string{"ax"}, nil)
	nameLookup, _ := doCodeNameLookup(TYPE_COUNTRY, nil, []string{"Åland"})
	if codeLookup != expected {
		t.Errorf("Search country by code failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
	if nameLookup != expected {
		t.Errorf("Search country by name failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
}

func TestTwoCountriesLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Country codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Country code</p>
    <p>AX</p>
    <p class="title">Country name</p>
    <p>Åland Islands</p>
  </div><div class="card">
    <p class="title">Country code</p>
    <p>SX</p>
    <p class="title">Country name</p>
    <p>Sint Maarten (Dutch part)</p>
  </div>
</body>
</html>
`

	codeLookup, _ := doCodeNameLookup(TYPE_COUNTRY, []string{"ax", "sx"}, nil)
	nameLookup, _ := doCodeNameLookup(TYPE_COUNTRY, nil, []string{"Åland", "Sint Maarten"})
	if codeLookup != expected {
		t.Errorf("Search country by code failed, expected\n" + expected + "\ngot\n" + codeLookup)
		/*for i, _ := range codeLookup {
			if codeLookup[i:i+1] != expected[i:i+1] {
				t.Errorf("missmatch at %d, '%s' vs '%s'", i, expected[i:i+1], codeLookup[i:i+1])
				break
			}
		}*/
	}
	if nameLookup != expected {
		t.Errorf("Search country by name failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
}

func TestOneLanguageLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Language codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Language code</p>
    <p>se</p>
    <p class="title">Language name</p>
    <p>Northern Sami</p>
  </div>
</body>
</html>
`

	codeLookup, _ := doCodeNameLookup(TYPE_LANGUAGE, []string{"se"}, nil)
	nameLookup, _ := doCodeNameLookup(TYPE_LANGUAGE, nil, []string{"Northern Sami"})
	if codeLookup != expected {
		t.Errorf("Search language by code failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
	if nameLookup != expected {
		t.Errorf("Search language by name failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
}

func TestTwoLanguagesLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Language codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Language code</p>
    <p>se</p>
    <p class="title">Language name</p>
    <p>Northern Sami</p>
  </div><div class="card">
    <p class="title">Language code</p>
    <p>sv</p>
    <p class="title">Language name</p>
    <p>Swedish</p>
  </div>
</body>
</html>
`

	codeLookup, _ := doCodeNameLookup(TYPE_LANGUAGE, []string{"se", "sv"}, nil)
	nameLookup, _ := doCodeNameLookup(TYPE_LANGUAGE, nil, []string{"Northern Sami", "Swedish"})
	if codeLookup != expected {
		t.Errorf("Search language by code failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
	if nameLookup != expected {
		t.Errorf("Search language by name failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
}

func TestTwoLanguagesCommaSeparatedLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Language codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Language code</p>
    <p>se</p>
    <p class="title">Language name</p>
    <p>Northern Sami</p>
  </div><div class="card">
    <p class="title">Language code</p>
    <p>sv</p>
    <p class="title">Language name</p>
    <p>Swedish</p>
  </div>
</body>
</html>
`

	codeLookup, _ := doCodeNameLookup(TYPE_LANGUAGE, []string{"se,sv"}, nil)
	if codeLookup != expected {
		t.Errorf("Search language by code failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
}

func TestThreeLanguagesLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Language codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Language code</p>
    <p>se</p>
    <p class="title">Language name</p>
    <p>Northern Sami</p>
  </div><div class="card">
    <p class="title">Language code</p>
    <p>sv</p>
    <p class="title">Language name</p>
    <p>Swedish</p>
  </div><div class="card">
    <p class="title">Language code</p>
    <p>sw</p>
    <p class="title">Language name</p>
    <p>Swahili</p>
  </div>
</body>
</html>
`

	codeLookup, _ := doCodeNameLookup(TYPE_LANGUAGE, []string{"se,sv", "sw"}, nil)
	if codeLookup != expected {
		t.Errorf("Search language by code failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
}

func TestOneCodeOneNameLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Language codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Language code</p>
    <p>se</p>
    <p class="title">Language name</p>
    <p>Northern Sami</p>
  </div><div class="card">
    <p class="title">Language code</p>
    <p>sv</p>
    <p class="title">Language name</p>
    <p>Swedish</p>
  </div>
</body>
</html>
`

	lookup, _ := doCodeNameLookup(TYPE_LANGUAGE, []string{"se"}, []string{"Swedish"})
	if lookup != expected {
		t.Errorf("Search language by code failed, expected\n" + expected + "\ngot\n" + lookup)
	}
}

func TestTwoCountriesOneSearchLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>Country codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Country code</p>
    <p>VG</p>
    <p class="title">Country name</p>
    <p>Virgin Islands, British</p>
  </div><div class="card">
    <p class="title">Country code</p>
    <p>VI</p>
    <p class="title">Country name</p>
    <p>Virgin Islands, U.S.</p>
  </div>
</body>
</html>
`
	codeLookup, _ := doCodeNameLookup(TYPE_COUNTRY, nil, []string{"virgin"})
	if codeLookup != expected {
		t.Errorf("Search country by code failed, expected\n" + expected + "\ngot\n" + codeLookup)
	}
}

func TestMccLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>MCC codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">MCC</p>
    <p>466</p>
    <p class="title">Country</p>
    <p>Taiwan</p>
  </div><div class="card">
    <p class="title">MCC</p>
    <p>240</p>
    <p class="title">Country</p>
    <p>Sweden</p>
  </div>
</body>
</html>
`

	lookup, _ := doCodeNameLookup(TYPE_MCC, []string{"466"}, []string{"Sweden"})
	if lookup != expected {
		t.Errorf("Search mcc by code failed, expected\n" + expected + "\ngot\n" + lookup)
	}
}

func TestHttpLookup(t *testing.T) {
	expected := `<!doctype html>
<html>
  <title>HTTP status codes</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">HTTP status code</p>
    <p><a href="https://httpstatuses.com/451">451</a></p>
    <p class="title">Name</p>
    <p>Unavailable For Legal Reasons</p>
  </div><div class="card">
    <p class="title">HTTP status code</p>
    <p><a href="https://httpstatuses.com/418">418</a></p>
    <p class="title">Name</p>
    <p>I'm a teapot</p>
  </div>
</body>
</html>
`

	lookup, _ := doCodeNameLookup(TYPE_HTTP_STATUS, []string{"451"}, []string{"teapot"})
	if lookup != expected {
		t.Errorf("Search http by code failed, expected\n" + expected + "\ngot\n" + lookup)
	}
}