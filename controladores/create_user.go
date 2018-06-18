package controladores

import (
	"../db"
	"../modelos"
	"github.com/kataras/iris"

)

func HandlerCreateUser (contexto iris.Context) {
	var user modelos.Usuario
	contexto.ReadJSON(&user)

	//Insertar documento
	db.GetSessionDB().DB("api").Col("usuarios").Save(&user)

	contexto.StatusCode(iris.StatusOK)
}