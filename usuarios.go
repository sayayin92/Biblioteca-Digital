package handlers

import (
	"bibliotecadigital/database"
	"bibliotecadigital/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistrarUsuario(c *gin.Context) {
	var nuevoUsuario models.Usuario
	if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nuevoUsuario.ID = len(database.Usuarios) + 1
	database.Usuarios = append(database.Usuarios, nuevoUsuario)
	c.JSON(http.StatusCreated, nuevoUsuario)
}

func ListarUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, database.Usuarios)
}
