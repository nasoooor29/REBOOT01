package main

import "net/http"

// var IMPORTANT_FILES = []string{
// 	"templates/400.tmpl.html",
// 	"templates/404.tmpl.html",
// 	"templates/500.tmpl.html",
// 	"templates/index.html",
// }

var IMPORTANT_FILES = map[int]string{
	400: "templates/400.tmpl.html",
	404: "templates/404.tmpl.html",
	406: "templates/406.tmpl.html",
	500: "templates/500.tmpl.html",
	200: "templates/index.html",

}

var TEMPLATE_FS = http.Dir("static")
