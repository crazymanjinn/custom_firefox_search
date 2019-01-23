package main

const indexTemplate = `<head>
{{range .}}<link rel="search" type="application/opensearchdescription+xml" title="{{.Title}}" href="/xml/{{.File}}">{{end}}
</head>
<h1>Search Engines</h1>
<ul>
{{range .}}<li><a href="/xml/{{.File}}">{{.Title}}</a><xmp>{{.Raw}}</xmp></li>{{end}}
</ul>`
