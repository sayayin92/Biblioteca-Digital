package handlers

import (
	"bibliotecadigital/database"
	"bibliotecadigital/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegistrarLibro(c *gin.Context) {
	var nuevoLibro models.Libro
	if err := c.ShouldBindJSON(&nuevoLibro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nuevoLibro.ID = len(database.Libros) + 1
	database.Libros = append(database.Libros, nuevoLibro)
	c.JSON(http.StatusCreated, nuevoLibro)
}

func ListarLibros(c *gin.Context) {
	c.JSON(http.StatusOK, database.Libros)
}

func BuscarLibro(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, l := range database.Libros {
		if l.ID == id {
			c.JSON(http.StatusOK, l)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Libro no encontrado"})
}
