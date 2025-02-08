package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Dados struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.POST("/dados", func(c *gin.Context) {
		var dados Dados

		// Recebe os dados do POST e os deserializa para a struct Dados
		if err := c.ShouldBindJSON(&dados); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Imprime os dados recebidos no console
		println("Nome:", dados.Nome)
		println("Email:", dados.Email)

		// Retorna uma resposta de sucesso
		c.JSON(http.StatusOK, gin.H{
			"message": "Dados recebidos com sucesso!",
			"dados":   dados,
		})
	})

	r.Run(":8080")
}
