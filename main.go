package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar Gin
	r := gin.Default()

	// Definir la ruta que quieres (como la de tu captura)
	r.GET("/usuario", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"nombre": "Guillermo",
			"rol":    "estudiante",
		})
	})

	// Ejecutar el servidor
	r.Run(":8080")
}
