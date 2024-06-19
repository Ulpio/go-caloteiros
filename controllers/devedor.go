package controllers

import (
	"caloteiros/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	devedores := models.GetAllDevedores()
	temp.ExecuteTemplate(w, "Index", devedores)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		telefone := r.FormValue("telefone")
		valorstr := r.FormValue("valorDevido")
		valorDevido, err := strconv.ParseFloat(valorstr, 64)
		if err != nil {
			panic(err)
		}
		models.CreateDevedor(nome, telefone, valorDevido)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteDevedor(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	devedor := models.EditarDevedor(id)
	temp.ExecuteTemplate(w, "Edit", devedor)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		telefone := r.FormValue("telefone")
		valorDevido := r.FormValue("valorDevido")
		valor := strings.Replace(valorDevido, ",", ".", -1)
		valorDevidoFloat, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			fmt.Println("Erro ao converter valor devido")
		}
		id := r.FormValue("id")
		idInt, err := strconv.Atoi(id)
		//It is getting a terror
		if err != nil {
			fmt.Println("Erro ao converter id")
		}
		models.AtualizarDevedor(idInt, nome, telefone, valorDevidoFloat)
	}
	http.Redirect(w, r, "/", 301)
}
