package models

import (
	"caloteiros/database"
	"fmt"
)

type Devedor struct {
	Id          int
	Nome        string
	Telefone    string
	ValorDevido float64
}

func GetAllDevedores() []Devedor {
	db := database.ConnectDB()
	selectAll, err := db.Query("SELECT * FROM devedores")
	if err != nil {
		panic(err)
	}
	d := Devedor{}
	devedores := []Devedor{}
	for selectAll.Next() {
		var id int
		var nome, telefone string
		var valorDevido float64

		err = selectAll.Scan(&id, &nome, &telefone, &valorDevido)

		if err != nil {
			panic(err)
		}

		d.Id = id
		d.Nome = nome
		d.Telefone = telefone
		d.ValorDevido = valorDevido

		devedores = append(devedores, d)
	}
	defer db.Close()
	return devedores
}

func CreateDevedor(nome, telefone string, preco float64) {
	db := database.ConnectDB()
	InsertData, err := db.Prepare("INSERT INTO devedores(nome, telefone, valorDevido) VALUES($1,$2,$3)")
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserido com sucesso: ", nome, telefone, preco)
	InsertData.Exec(nome, telefone, preco)

	defer db.Close()
}

func DeleteDevedor(id string) {
	db := database.ConnectDB()
	deleteData, err := db.Prepare("DELETE FROM devedores WHERE id=$1")
	if err != nil {
		panic(err)
	}
	deleteData.Exec(id)
	defer db.Close()
}

func EditarDevedor(id string) Devedor {
	db := database.ConnectDB()
	selectData, err := db.Query("SELECT * FROM devedores WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	d := Devedor{}
	for selectData.Next() {
		var id int
		var nome, telefone string
		var valorDevido float64

		err = selectData.Scan(&id, &nome, &telefone, &valorDevido)

		if err != nil {
			panic(err)
		}

		d.Id = id
		d.Nome = nome
		d.Telefone = telefone
		d.ValorDevido = valorDevido
	}

	defer db.Close()
	return d
}

func AtualizarDevedor(id int, nome, telefone string, preco float64) {
	db := database.ConnectDB()
	updateData, err := db.Prepare("UPDATE devedores SET nome=$1, telefone=$2, valorDevido=$3 WHERE id=$4")
	if err != nil {
		panic(err)
	}
	updateData.Exec(nome, telefone, preco, id)
	defer db.Close()
}
