package main

import (
  "net/http"
  "html/template"
  "regexp"
)

func main() {
  c := new(color)
  re, _ := regexp.Compile("^#?([a-f0-9]{6}|[a-f0-9]{3})$")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    c.hex = r.URL.Path[1:]
    page := "error.html"

    if re.MatchString(c.hex) {
      page = "color.html"
    }

    t, _ := template.ParseFiles(page)
    t.Execute(w, c)
  })

  http.ListenAndServe(":8080", nil)
}

type color struct {
  hex string
}
