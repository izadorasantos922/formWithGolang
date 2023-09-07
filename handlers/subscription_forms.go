package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"webform/controllers"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	SucessMessage := ""
	ErrorMessage := ""
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "erro ao fazer parse do form: %v", err)
			return
		}
		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			ErrorMessage = "Ocorreu um erro"
			fmt.Fprintf(w, "erro ao salvar os dados do form: %v", err)
			return
		} else {
			SucessMessage = "Cadastro realizado"
		}
		SucessMessage = ""
	}
	data := struct {
		SucessMessage string
		ErrorMessage  string
	}{
		SucessMessage: SucessMessage,
		ErrorMessage:  ErrorMessage,
	}
	tmpl, err := template.ParseFiles("handlers/templates/subscription_form.html")
	if err != nil {
		fmt.Fprintf(w, "erro ao analisar o modelo: %v", err)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "erro ao executar o modelo: %v", err)
	}
	// http.ServeFile(w, r, "handlers/templates/subscription_form.html")
}
