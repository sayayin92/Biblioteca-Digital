package routers

import (
	"bibliotecadigital/handlers"

	"github.com/gin-gonic/gin"
)

func Rutas(r *gin.Engine) {
	r.POST("/usuarios", handlers.RegistrarUsuario)
	r.GET("/usuarios", handlers.ListarUsuarios)

	r.POST("/libros", handlers.RegistrarLibro)
	r.GET("/libros", handlers.ListarLibros)
	r.GET("/libros/:id", handlers.BuscarLibro)

	r.POST("/prestamos", handlers.PrestarLibro)
	r.GET("/prestamos", handlers.ListarPrestamos)
}
