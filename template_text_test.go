package main
import (
    "testing"
    "text/template"
    "os"
    "log")

func TestVariables(t *testing.T){
    tmpl, _ := template.ParseFiles("temp/temp_text")
    if err := tmpl.Execute(os.Stdout, ""); err != nil{
        log.Printf("execute err : %v", err)
    }
}
