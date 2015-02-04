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
    c.Color = r.URL.Path[1:]
    page := "error.html"

    if re.MatchString(c.Color) {
      page = "color.html"
    }

    t, _ := template.ParseFiles(page)
    t.Execute(w, c)
  })

  http.ListenAndServe(":8080", nil)
}

type color struct {
  Color string
}
