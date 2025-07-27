package main

import (
	"control_horario/database" // Conexión a base de datos
	"control_horario/tablas"   // Modelo de Empleado

	"github.com/gin-gonic/gin" // Framework web
)

func main() {
	//  Conectar a PostgreSQL
	database.Conectar()

	//  Crear el router
	router := gin.Default()

	// Ruta para guardar empleados en la base de datos
	router.POST("/empleado", func(c *gin.Context) {
		var nuevo tablas.Empleado

		// Leer JSON recibido
		if err := c.ShouldBindJSON(&nuevo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Guardar en la base de datos
		result := database.DB.Create(&nuevo)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, nuevo)
	})

	//  Iniciar el servidor en el puerto 8081
	router.Run(":8081")
}
