package handlers

import (
	"bibliotecadigital/database"
	"bibliotecadigital/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PrestarLibro(c *gin.Context) {
	var nuevoPrestamo models.Prestamo
	if err := c.ShouldBindJSON(&nuevoPrestamo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nuevoPrestamo.ID = len(database.Prestamos) + 1
	database.Prestamos = append(database.Prestamos, nuevoPrestamo)
	c.JSON(http.StatusCreated, nuevoPrestamo)
}

func ListarPrestamos(c *gin.Context) {
	c.JSON(http.StatusOK, database.Prestamos)
}
