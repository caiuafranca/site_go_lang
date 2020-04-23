package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("./pages/index.html")
	data := map[string]string{
		"Title":    "SonSon Doces",
		"Wellcome": "Bem Vindos, aproveitem nossas ofertas!",
	}

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Servidor esta UP em: http://localhost:8888/")
	http.ListenAndServe(":8888", nil)
}
