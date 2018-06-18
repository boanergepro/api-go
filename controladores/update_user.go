package controladores

import (
	"../db"
	"../modelos"
	arango "github.com/diegogub/aranGO"
	"github.com/kataras/iris"

)

func HandlerUpdateUser (contexto iris.Context) {
	key_params := contexto.Params().Get("key")

	query := arango.NewQuery(`
			FOR usuario in usuarios
			FILTER usuario._key == @key
			RETURN usuario
		`)
	query.BindVars = map[string]interface{}{
		"key": key_params,
	}

	cursor, err := db.GetSessionDB().DB("api").Execute(query)

	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	var user_actualiar modelos.Usuario

	cursor.FetchOne(&user_actualiar)
	
	contexto.ReadJSON(&user_actualiar)

	err = db.GetSessionDB().DB("api").Col("usuarios").Replace(user_actualiar.Key, user_actualiar)
	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	contexto.StatusCode(iris.StatusOK)
}