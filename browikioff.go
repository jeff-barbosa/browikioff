package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "net/http"
  "html/template"
  "io/ioutil"
)

func main() {
  fmt.Println("========================================================")
  fmt.Println("bROWikiOff - Conteúdo do site bROWiki disponível offline")
  fmt.Println("========================================================")
  fmt.Println("Acesse a wiki no link: http://localhost:8080")

  http.Handle("/w/", http.StripPrefix("/w/", http.FileServer(http.Dir("resources"))))

  http.HandleFunc("/ajax", handleAjax)
  http.HandleFunc("/wiki/", handleWiki)
  http.HandleFunc("/", handle)

  http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
  pageName := r.URL.Path[1:]

  if pageName == "favicon.ico" {
    return
  }

  if pageName == "" {
    pageName = "Página principal"
  }
  
  pageContent, err := ioutil.ReadFile("resources/pages/"+ pageName + ".html")
  if err != nil {
    fmt.Println(err)
    return
  }
  
  tmpl, err := template.ParseFiles("resources/template/base.html")
  if err != nil {
    fmt.Println(err)
    return
  }

  tmpl.Execute(
    w, 
    &map[string]interface{}{
      "Title": strings.Replace(pageName,"_"," ",-1),
      "Content": template.HTML(string(pageContent)),
  })
}

func handleWiki(w http.ResponseWriter, r *http.Request) {
  r.URL.Path = r.URL.Path[len("/wiki"):]
  handle(w,r)
}

func handleAjax(w http.ResponseWriter, r *http.Request) {
  vals := r.URL.Query()
  key := vals["q"][0]

  filehandler, err := os.Open("resources/browiki_links.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  scanner := bufio.NewScanner(filehandler)
  reply := ""
  count := 0

  for scanner.Scan() {
    if strings.Contains(strings.ToLower(scanner.Text()), strings.ToLower(key)) {
      reply += "<a href=\""+ scanner.Text() +"\"> "+ scanner.Text() +"</a><br >"
      count++
    }

    if count > 9 {
      fmt.Fprintf(w, reply)
      return
    }
  }

  if reply == "" {
    reply = "Nenhum resultado"
  }

  fmt.Fprintf(w, reply)
}