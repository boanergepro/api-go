package controladores

import (
	"github.com/kataras/iris"
)

func HandleIndex (contexto iris.Context) {
	contexto.HTML(`
		<h1>Api corriendo</h1><br>
		<h4>Recursos:</h4><br>
		<ul>
			<li>
				<a href="http://localhost:8080/api/usuarios"> usuarios
			</li>
		</ul>
    `)
}