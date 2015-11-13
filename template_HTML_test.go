package main

import (
    "testing"
    "html/template"
    "os"
    "log"
    "fmt")

//define 的使用
func TestTextTemplate(t *testing.T){
//    tmpl, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
    tmpl, _ := template.ParseFiles("temp/temp*.html")
    tmpl, _ = tmpl.Parse(`{{define "W"}} nihaoya, {{.}}!!! {{end}}`)
//    tmpl, _ = tmpl.Parse(`default , {{.}}!!!`)
    tmpl, _ = tmpl.ParseFiles("temp/temp3.html")
//    tmpl = tmpl.Lookup("W")
    for k, tm := range tmpl.Templates(){
        log.Printf("第%d个名字是%s", k, tm.Name())
    }
//    tmpl.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
    if err := tmpl.Execute(os.Stdout, template.HTML("<script>alert('you have been pwned')</script>")); err != nil{
        log.Printf("execute err : %v", err)
    }else{
        log.Println("success")
    }
}

func TestHTMLEscape(t *testing.T){
    const s = `"Fran & Freddie's Diner" <tasty@example.com>`
    v := []interface{}{`"Fran & Freddie's Diner"`, ' ', `<tasty@example.com>`}

    fmt.Println(template.HTMLEscapeString(s))
    template.HTMLEscape(os.Stdout, []byte(s))
    fmt.Fprint(os.Stdout, "")

    fmt.Println(template.JSEscapeString(s))
    fmt.Println(template.JSEscaper(v...))
    fmt.Println(template.URLQueryEscaper(v...))
}