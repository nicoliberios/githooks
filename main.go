package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin" // Importamos Gin
)

func main() {
	// Creamos una nueva instancia de Gin
	r := gin.Default()

	// Definimos una ruta simple en la raíz
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "¡Hola, Mundo desde Gin!",
		})
	})

	// Definimos una ruta para probar la salud de la aplicación
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Iniciamos el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en puerto 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
