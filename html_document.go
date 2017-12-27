package dataaxe

import "github.com/osteele/liquid"

type cardParameters struct {
	Title1, Body1, Title2, Body2 string
}

func getHtmlDocument(pageTitle string, params []cardParameters) (string, error) {
	engine := liquid.NewEngine()
	bindings := map[string]interface{}{
		"pageTitle": pageTitle,
		"cards":     params,
	}
	return engine.ParseAndRenderString(resultDocument, bindings)
}

func getInstructionsDocument() string {
	return instructionsDocument
}

const resultDocument = `<!doctype html>
<html>
  <title>{{ pageTitle }}</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  {% for card in cards %}<div class="card">
    <p class="title">{{ card.Title1 }}</p>
    <p>{{ card.Body1 }}</p>
    <p class="title">{{ card.Title2 }}</p>
    <p>{{ card.Body2 }}</p>
  </div>{% endfor %}
</body>
</html>
`

const instructionsDocument = `<!doctype html>
<html>
  <title>Data axe</title>
  <link rel="stylesheet" type="text/css" href="/static/main.css">
</head>
<body>
  <div class="card">
    <p class="title">Try one of these</p>
    <p class="example"><a href="/api?code=21">/api?code=21</a></p>
    <p class="example"><a href="/mcc?code=410">/mcc?code=410</a></p>
    <p class="example"><a href="/http?code=451">/http?code=451</a></p>
    <p class="example"><a href="/ascii?code=127">/ascii?code=127</a></p>
    <p class="example"><a href="/ts?t=1323782116">/ts?t=1323782116</a></p>
    <p class="example"><a href="/country?code=se">/country?code=se</a></p>
    <p class="example"><a href="/language?code=se">/language?code=se</a></p>
    <p class="example"><a href="/country?name=virgin">/country?name=virgin</a></p>
    <p class="example"><a href="/language?name=swedish">/language?name=swedish</a></p>
  </div>
</body>
</html>
`
