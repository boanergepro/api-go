package controladores

import (
	"../db"
	"github.com/kataras/iris"

)

func HandlerDeleteUser (contexto iris.Context) {
	key_params := contexto.Params().Get("key")

	err := db.GetSessionDB().DB("api").Col("usuarios").Delete(key_params)

	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	contexto.StatusCode(iris.StatusOK)
}